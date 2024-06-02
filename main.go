package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "sync"

    _ "github.com/joho/godotenv/autoload"
    _ "github.com/lib/pq"
)

var (
    db *sql.DB
    cache = struct {
        sync.RWMutex
        m map[string][]Todo 
    }{m: make(map[string][]Todo)}
)

type Todo struct {
    ID    int    `json:"id"`
    Task  string `json:"task"`
    Done  bool   `json:"done"`
}

func main() {
    initDB()
    setupRoutes()
    startHTTPServer()
}

func initDB() {
    var err error
    db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
}

func setupRoutes() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/todos", todosHandler)
    http.HandleFunc("/todo", todoHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintf(w, "Welcome to the Go Todo Application!")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        user, _, _ := r.BasicAuth()
        todos, found := getCachedTodos(user)
        if !found {
            rows, err := db.Query("SELECT id, task, done FROM todos WHERE user_id = $1", user)
            if err != nil {
                http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
                return
            }
            defer rows.Close()
            for rows.Next() {
                var todo Todo
                if err := rows.Scan(&todo.ID, &todo.Task, &todo.Done); err != nil {
                    http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
                    return
                }
                todos = append(todos, todo)
            }
            setCachedTodos(user, todos)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todos)

    case http.MethodPost:
        var todo Todo
        if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
            http.Error(w, "Failed to decode todo", http.StatusBadRequest)
            return
        }
        user, _, _ := r.BasicAuth()

        err := db.QueryRow("INSERT INTO todos (task, done, user_id) VALUES ($1, $2, $3) RETURNING id", todo.Task, todo.Done, user).Scan(&todo.ID)
        if err != nil {
            http.Error(w, "Failed to add todo", http.StatusInternalServerError)
            return
        }

        invalidateCache(user)

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todo)

    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
}

func getCachedTodos(user string) ([]Todo, bool) {
    cache.RLock()
    defer cache.RUnlock()
    todos, found := cache.m[user]
    return todos, found
}

func setCachedTd(r todos []Todo) {
    cache.Lock()
    defer cache.Unlock()
    cache.m[user] = todos
}

func invalidateCache(user string) {
    cache.Lock()
    defer cache.Unlock()
    delete(cache.m, user)
}

func startHTTPServer() {
    log.Println("Starting HTTP server on port :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}