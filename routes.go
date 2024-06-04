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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{userId}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{userId}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{userId}", DeleteUser).Methods("DELETE")

	enhancedRouter := loggingMiddleware(router)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not set in .env file")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), enhancedRouter))
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
	_, err := fmt.Fprintln(w, "Welcome to the Home Page!")
	if err != nil {
		log.Printf("Error sending response in HomeHandler: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Fetching all users...")
	if err != nil {
		log.Printf("Error sending response in GetUsers: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	_, err := fmt.Fprintf(w, "Fetching user with ID %s...\n", userId)
	if err != nil {
		log.Printf("Error sending response in GetUser: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Creating new user...")
	if err != nil {
		log.Printf("Error sending response in CreateUser: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	_, err := fmt.Fprintf(w, "Updating user with ID %s...\n", userId)
	if err != nil {
		log.Printf("Error sending response in UpdateUser: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	_, err := fmt.Fprintf(w, "Deleting user with ID %s...\n", userId)
	if err != nil {
		log.Printf("Error sending response in DeleteUser: %v", err)
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}
}