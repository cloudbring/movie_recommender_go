package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(100)"`
}

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := &User{
		Username: username,
		Password: password,
	}

	if err := uh.DB.Create(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := &User{}
	if err := uh.DB.Where("username = ? AND password = ?", username, password).First(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// In a real-world application, you would handle user sessions here.
	// For simplicity, we're just sending a 200 OK status.
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/register", uh.Register).Methods("POST")
	router.HandleFunc("/login", uh.Login).Methods("POST")
	router.HandleFunc("/logout", uh.Logout).Methods("POST")
}
