package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

var db *sql.DB

func initDB() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []Task

    rows, err := db.Query("SELECT id, description, completed FROM tasks")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Description, &task.Completed); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskID := params["id"]

    var task Task
    sqlStatement := `SELECT id, description, completed FROM tasks WHERE id = $1;`
    row := db.QueryRow(sql: sqlStatement, taskID)

    err := row.Scan(&task.ID, &task.Description, &task.Completed)
    if err == sql.ErrNoRows {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sqlStatement := `
        INSERT INTO tasks (description, completed)
        VALUES ($1, $2)
        RETURNING id`
    id := 0
    err := db.QueryRow(sqlStatement, task.Description, task.Completed).Scan(&id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(struct{ ID int `json:"id"` }{ID: id})
}

func updateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var task Task
    taskId := params["id"]
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sqlStatement := `
        UPDATE tasks
        SET description = $2, completed = $3
        WHERE id = $1;`
    _, err := db.Exec(sqlStatement, taskId, task.Description, task.Completed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskId := params["id"]

    sqlStatement := `DELETE FROM tasks WHERE id = $1;`
    _, err := db.Exec(sqlStatement, taskId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func main() {
    initDB()

    r := mux.New;

    r.HandleFunc("/tasks", getTasks).Methods("GET")
    r.HandleFunc("/tasks/{id}", getTaskByID).Methods("GET") // Added route for getting a single task by ID
    r.HandleFunc("/tasks", createTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
    r.HandleFunc("tasks/{id}", deletePhotograph).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}