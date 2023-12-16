package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/islamyakin/jwt/auth"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", auth.ProtectedHandler).Methods("GET")

	fmt.Println("Starting the server")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
	fmt.Println("Server started. Listening on port 4000")
}