package main

import (
	"log"
	"net/http"

	"github.com/Lunnaris01/CivAPI/internal/auth"
	//"github.com/Lunnaris01/CivAPI/internal/database"

)






func (cfg apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	user,err := cfg.db.GetUserByUsername(r.Context(),username)
	if err != nil {
		log.Printf("Unable to find User: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Trying to log in User %s",user.Username)
	// auth user
	cfg.displayFileserverContent(w, "/login")

}

func (cfg apiConfig) handlerSignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	hashed_password, err := auth.HashPassword(password)
	if err != nil {
		log.Printf("Error hashin password: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cfg.db.AddUser(r.Context(), username, hashed_password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Successfully added user")
	cfg.displayFileserverContent(w, "/login")

}
