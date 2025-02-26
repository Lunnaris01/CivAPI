package main

import(
	"log"
	"net/http"
	"github.com/Lunnaris01/CivAPI/internal/auth"

)


func (cfg apiConfig) handlerDashboard(w http.ResponseWriter, r *http.Request) {
	cfg.displayFileserverContent(w,"/content")
}

func (cfg apiConfig) handlerGetGames(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	
	if err != nil {
		log.Printf("Unable to read Authentification header")
		respondWithError(w,401,"Unable to read Authentification header",err)
		return
	}

	userID, err := auth.ValidateJWT(bearerToken, cfg.secretKey)
	if err != nil {
		log.Printf("Token missmatch")
		respondWithError(w,401,"Unable to verify token",err)
		return
	}
	games, err := cfg.db.GetGamesByUserID(r.Context(),userID)
	if err != nil{
		respondWithError(w,401,"Failes to load games",err)
	}
	log.Printf("Authentification successfull for user: %v with token %v", userID, bearerToken)
	respondWithJSON(w,200,games)

}