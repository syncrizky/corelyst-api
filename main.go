package main

import (
	"corelyst-api/config"
	"corelyst-api/handler"
	"corelyst-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()

	r := mux.NewRouter()

	// Handler Public
	r.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)
	api.HandleFunc("/users", handler.GetUsersHandler).Methods("GET")
	api.HandleFunc("/nasabah", handler.GetAllNasabahHandler).Methods("GET")
	api.HandleFunc("/add_nasabah", handler.AddNasabahHandler).Methods("POST")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
