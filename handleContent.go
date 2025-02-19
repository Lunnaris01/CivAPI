package main

import(
	//"log"
	"net/http"
	//"github.com/Lunnaris01/CivAPI/internal/auth"

)


func (cfg apiConfig) handlerDashboard(w http.ResponseWriter, r *http.Request) {
	/*token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.secretKey)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	} */

	cfg.displayFileserverContent(w,"/dashboard")
	
}
