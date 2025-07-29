package handler

import (
	"corelyst-api/middleware"
	"corelyst-api/model"
	"corelyst-api/service"
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := service.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("user berhasil di daftarkan"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	u, err := service.LoginUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, _ := middleware.GenerateJWT(u.Username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := service.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}
