package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "sync"

    _ "github.com/joho/godotenv/autoload"
    _ "github.com/lib/pq"
)

var (
    db    *sql.DB
    cache = struct {
        sync.RWMutex
        m map[string]string
    }{m: make(map[string]string)}

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
```

```go
func getCachedValue(key string) (string, bool) {
    cache.RLock()
    defer cache.RUnlock()
    value, found := cache.m[key]
    return value, found
}

func setCachedValue(key string, value string) {
    cache.Lock()
    defer cache.Unlock()
    cache.m[key] = value
}