package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return
	}
	fmt.Printf("The user request value %v", u)

	if u.Username == "kanaya" && u.Password == "rainbowdrinker" {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = fmt.Errorf("no username found")
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

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprint(w, "Invalid token")
		if err != nil {
			return
		}
		return
	}
	_, err = fmt.Fprint(w, "Hello Mr. Kanaya")
	if err != nil {
		return
	}

}
