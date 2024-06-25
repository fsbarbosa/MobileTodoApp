package main

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

const (
    CreateTodoTableSQL = `CREATE TABLE IF NOT EXISTS todos (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "task" TEXT NOT NULL,
        "status" TEXT NOT NULL
    );`
    DbFileName = "todo.db"
)

func InitializeDB() *sql.DB {
    db, err := sql.Open("sqlite3", DbFileName)
    if err != nil {
        log.Fatal(err)
    }

    statement, err := db.Prepare(CreateTodoTableSQL)
    if err != nil {
        log.Fatal(err)
    }
    statement.Exec()

    return db
}

func InsertTodo(db *sql.DB, task, status string) error {
    query := `INSERT INTO todos(task, status) VALUES (?, ?)`
    statement, err := db.Prepare(query)
    if err != nil {
        return err
    }
    _, err = statement.Exec(task, status)
    return err
}

func GetTodos(db *sql.DB) ([]Todo, error) {
    query := `SELECT id, task, status FROM todos`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []Todo
    for rows.Next() {
        var t Todo
        if err := rows.Scan(&t.ID, &t.Task, &t.Status); err != nil {
            return nil, err
        }
        todos = append(todos, t)
    }
    return todos, nil
}

type Todo struct {
    ID     int
    Task   string
    Status string
}

func main() {
    db := InitializeDB()
    defer db.Close()

    err := InsertTodo(db, "Learn Go database interaction", "pending")
    if err != nil {
        log.Fatal(err)
    }

    todos, err := GetTodos(db)
    if err != nil {
        log.Fatal(err)
    }
    for _, todo := range todos {
        log.Println(todo)
    }
}