package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Fetching all users...")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintf(w, "Fetching user with ID %s...\n", userId)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating new user...")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintf(w, "Updating user with ID %s...\n", userId)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintf(w, "Deleting user with ID %s...\n", userId)
}