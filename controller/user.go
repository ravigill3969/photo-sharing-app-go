package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ravigill3969/backend/database"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(w http.ResponseWriter, r *http.Request) {

	db, err := database.ConnectDB()

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (username, email , password) VALUES ($1, $2, $3)")

	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var user UserRegister
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Cannot decode body to json", http.StatusBadRequest)
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		http.Error(w, "Unable to hash password", http.StatusInternalServerError)
	}

	cUser, err := stmt.Exec(user.Username, user.Email, hashedPassword)

	if err != nil {
		http.Error(w, "Error executing insert", http.StatusInternalServerError)
		return
	}

	fmt.Println(cUser)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "User registered successfully",
		"username": user.Username,
	})
}
