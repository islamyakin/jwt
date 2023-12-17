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
		_, err := fmt.Fprint(w, "Hello world")
		if err != nil {
			return
		}
	})
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/whoiam", auth.ProtectedHandler).Methods("GET")

	fmt.Println("Starting the server")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
	fmt.Println("Server started. Listening on port 4000")
}
