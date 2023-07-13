package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	log.Println("Server is running on http://localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Perform your login authentication logic here
		// For simplicity, we'll just print the entered username and password
		fmt.Printf("Login: username=%s, password=%s\n", username, password)
	}

	http.ServeFile(w, r, "public/login.html") // Path to your login.html file
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Perform your registration logic here
		// For simplicity, we'll just print the entered username and password
		fmt.Printf("Registration: username=%s, password=%s\n", username, password)
	}

	http.ServeFile(w, r, "public/register.html") // Path to your register.html file
}
