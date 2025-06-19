package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var user UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Cannot decode body to json", http.StatusBadRequest)
	}

	fmt.Println(user)
}
