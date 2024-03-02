package main

import (
	"fmt"
	"github.com/derrosilee/rssagg/internal/auth"
	"github.com/derrosilee/rssagg/internal/database"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, response *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication error %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't Retrieve User: %v", err))
			return
		}
		handler(w, r, user)
	}
}
