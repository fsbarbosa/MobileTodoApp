package main

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

const (
    CreateTodoTableSQL = `
    CREATE TABLE IF NOT EXISTS todos (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "task" TEXT NOT NULL,
        "status" TEXT NOT NULL
    );
    `
    DbFileName = "todo.db"
)

func InitializeDB() *sql.DB {
    db, err := sql.Open("sqlite3", DbFileName)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    _, err = db.Exec(CreateTodoTableSQL)
    if err != nil {
        log.Fatalf("Error executing database creation SQL: %v", err)
        db.Close()
    }

    return db
}

func InsertTodo(db *sql.DB, task, status string) error {
    query := `INSERT INTO todos(task, status) VALUES (?, ?)`
    _, err := db.Exec(query, task, status)
    if err != nil {
        log.Printf("Error executing insert todo: %v", err)
        return err
    }
    return nil
}

func GetTodos(db *sql.DB) ([]Todo, error) {
    var todos []Todo
    query := `SELECT id, task, status FROM todos`
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("Error querying todos: %v", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var t Todo
        if err := rows.Scan(&t.ID, &t.Task, &t.Status); err != nil {
            log.Printf("Error scanning todo: %v", err)
            return nil, err
        }
        todos = append(todos, t)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error iterating through todos: %v", err)
        return todos, err
    }

    return todos, nil
}

type Todo struct {
    ID     int
    Task   string
    Status string
}

func main() {
    db := InitializeRepositoryDB()
    defer db.Close()

    if err := InsertTodo(db, "Learn Go database interaction", "pending"); err != nil {
        log.Printf("Error inserting todo: %v", err)
    }

    todos, err := GetTodos(db)
    if err != nil {
        log.Printf("Error getting todos: %v", err)
        return
    }

    for _, todo := range todos {
        log.Println(todo)
    }
}