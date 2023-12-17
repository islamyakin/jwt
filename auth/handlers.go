package auth

import (
	"encoding/json"
	"fmt"
	"github.com/islamyakin/jwt/models"
	"net/http"
	"strings"
)

//type User struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return
	}
	fmt.Printf("The user request value %v", u)

	if AuthenticateUser(u.Username, u.Password) {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprint(w, "Error creating Token")
		}
		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprint(w, tokenString)
		if err != nil {
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_, err = fmt.Fprint(w, "Invalid credentials")
		if err != nil {
			return
		}
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprint(w, "Missing authorization header")
		if err != nil {
			return
		}
		return
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := getUsernameFromToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprint(w, "Invalid token")
		if err != nil {
			return
		}
		return
	}

	userRoles, err := GetUserRoles(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, "Error retrieving user roles")
		if err != nil {
			return
		}
		return
	}

	requiredRoles := "admin"
	if !hasRequiredRoles(userRoles, requiredRoles) {
		w.WriteHeader(http.StatusForbidden)
		message := fmt.Sprintf("Insufficient roles to access this endpoint, %s", username)
		_, err := fmt.Fprint(w, message)
		if err != nil {
			return
		}
		return
	}
	message := fmt.Sprintf("Access granted, hello %s", username)
	_, err = fmt.Fprint(w, message)
	if err != nil {
		return
	}
}
func hasRequiredRoles(userRoles, requiredRoles string) bool {
	userRolesList := strings.Split(userRoles, ",")
	requiredRolesList := strings.Split(requiredRoles, ",")

	for _, role := range requiredRolesList {
		if !contains(userRolesList, role) {
			return false
		}
	}

	return true
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
