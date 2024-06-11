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

var database *sql.DB

func initializeDatabase() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    database, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    if err = database.Ping(); err != nil {
    panic(err)
    }
}

func retrieveTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []Task

    rows, err := database.Query("SELECT id, description, completed FROM tasks")
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

func retrieveTaskByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskID := params["id"]

    var task Task
    sqlStatement := `SELECT id, description, completed FROM tasks WHERE id = $1;`
    row := database.QueryRow(sqlStatement, taskID)

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

func createNewTask(w http.ResponseWriter, r *http.Request) {
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
    err := database.QueryRow(sqlStatement, task.Description, task.Completed).Scan(&id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(struct{ ID int `json:"id"` }{ID: id})
}

func modifyTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var task Task
    taskID := params["id"]
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sqlStatement := `
        UPDATE tasks
        SET description = $2, completed = $3
        WHERE id = $1;`
    _, err := database.Exec(sqlStatement, taskID, task.Description, task.Completed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func removeTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskID := params["id"]

    sqlStatement := `DELETE FROM tasks WHERE id = $1;`
    _, err := database.Exec(sqlStatement, taskID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func main() {
    initializeDatabase()

    router := mux.NewRouter()

    router.HandleFunc("/tasks", retrieveTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}", retrieveTaskByID).Methods("GET")
    router.HandleFunc("/tasks", createNewTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", modifyTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", removeTask).Methods("DELETE")

    http.ListenAndServe(":8080", router)
}