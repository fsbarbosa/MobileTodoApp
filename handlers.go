package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "sync"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

type taskCache struct {
    sync.Mutex
    tasks map[string]Task // Using the task ID as the key
}

func (c *taskCache) Get(key string) (Task, bool) {
    c.Lock()
    defer c.Unlock()
    task, exists := c.tasks[key]
    return task, exists
}

func (c *taskCache) Set(key string, task Task) {
    c.Lock()
    defer c.Unlock()
    c.tasks[key] = task
}

var (
    database *sql.DB
    cache    taskCache
)

func initializeDatabase() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    database, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    if err = database.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }
}

func init() {
    cache.tasks = make(map[string]Task)
}

func main() {
    initializeDatabase()

    router := mux.NewRouter()

    router.HandleFunc("/tasks", retrieveTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}", retrieveTaskByID).Methods("GET")
    router.HandleFunc("/tasks", createNewTask).Methods("POST")
    router.HandleFunc("/tasks/{ua}", modifyTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", removeTask).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))
}