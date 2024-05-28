package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/joho/godotenv/autoload"
    _ "github.com/lib/pq"
)

var db *sql.DB

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
    // Add more routes here
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    _, _ = fmt.Fprintf(w, "Welcome to the Go Application!")
}

func startHTTPServer() {
    log.Println("Starting HTTP server on port :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}