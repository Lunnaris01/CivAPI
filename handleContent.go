package main

import (
	"encoding/json"
	"github.com/Lunnaris01/CivAPI/internal/auth"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (cfg apiConfig) handlerDashboard(w http.ResponseWriter, r *http.Request) {
	cfg.displayFileserverContent(w, "/content")
}

func (cfg apiConfig) handlerGetGames(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)

	if err != nil {
		log.Printf("Unable to read Authentification header")
		respondWithError(w, http.StatusUnauthorized, "Unable to read Authentification header", err)
		return
	}

	userID, err := auth.ValidateJWT(bearerToken, cfg.secretKey)
	if err != nil {
		log.Printf("Token missmatch")
		respondWithError(w, http.StatusUnauthorized, "Unable to verify token", err)
		return
	}
	games, err := cfg.db.GetGamesByUserID(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failes to load games", err)
	}
	log.Printf("Authentification successfull for user: %v with token %v", userID, bearerToken)
	respondWithJSON(w, http.StatusOK, games)

}

func (cfg apiConfig) handlerAddGame(w http.ResponseWriter, r *http.Request) {

	type reqBody struct {
		Country      string `json:"country"`
		GameWon      bool   `json:"game_won"`
		WinCondition string `json:"win_condition"`
	}

	bearerToken, err := auth.GetBearerToken(r.Header)

	if err != nil {
		log.Printf("Unable to read Authentification header")
		respondWithError(w, http.StatusUnauthorized, "Unable to read Authentification header", err)
		return
	}
	var userIDstr string
	userIDstr, err = auth.ValidateJWT(bearerToken, cfg.secretKey)
	if err != nil {
		log.Printf("Token missmatch")
		respondWithError(w, http.StatusUnauthorized, "Unable to verify", err)
		return
	}

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		log.Printf("Unable to fetch UserID")
		respondWithError(w, http.StatusBadRequest, "Unable to verify", err)
		return
	}

	rBody := reqBody{}
	rData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read Body")
		respondWithError(w, http.StatusBadRequest, "Error adding the game", err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(rData, &rBody)
	if err != nil {
		log.Printf("Failed to Unmarshal JSON")
		respondWithError(w, http.StatusBadRequest, "Error adding the game", err)
		return
	}

	err = cfg.db.AddGame(r.Context(), userID, rBody.Country, rBody.GameWon, rBody.WinCondition)
	if err != nil {
		log.Printf("Failed to add new Game to the Database")
		respondWithError(w, http.StatusBadRequest, "Error adding the game", err)
		return
	}
	log.Printf("Added new Game to the Database!")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Added Successful"))

}
