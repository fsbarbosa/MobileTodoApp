package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := setupRouter()
	enhancedRouter := loggingMiddleware(router)
	port := getServerPort()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), enhancedRouter))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{userId}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/apiusers/{userId}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{userId}", DeleteUser).Methods("DELETE")
	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("[%s] %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
		}()
		next.ServeHTTP(w, r)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, "Welcome to the Home Page!", http.StatusOK)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, "Fetching all users...", http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	writeResponse(w, fmt.Sprintf("Fetching user with ID %s...\n", userId), http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, "Creating new user...", http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	writeResponse(w, fmt.Sprintf("Updating user with ID %s...\n", userId), http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	writeResponse(w, fmt.Sprintf("Deleting user with ID %s...\n", userId), http.StatusOK)
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Print("Port not set in .env file, defaulting to 8080")
		port = "8080"
	}
	return port
}

func writeResponse(w http.ResponseWriter, response string, statusCode int) {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprintln(w, response)
	if err != nil {
		log.Printf("Error sending response: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}