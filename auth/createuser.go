package auth

import (
	"encoding/json"
	"fmt"
	"github.com/islamyakin/jwt/models"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Invalid request body")
		return
	}
	fmt.Printf("The user request value %v", u)

	// Validate input
	if u.Username == "" || u.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Username and password are required")
		return
	}

	// Check if the username already exists
	if UserExists(u.Username) {
		w.WriteHeader(http.StatusConflict)
		_, _ = fmt.Fprint(w, "Username already exists")
		return
	}

	// Create the user in the database
	err = CreateUser(u.Username, u.Password, u.Roles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "Error creating user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprint(w, "User created successfully")
}

// UserExists checks if a user with the given username already exists in the database
func UserExists(username string) bool {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	return result.Error == nil
}
