package main

import (
	"fmt"
	"net/http"

	"github.com/utkarshverm4/rssagg/internal/auth"
	"github.com/utkarshverm4/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithErr(w, 403, fmt.Sprintf("Auth err: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithErr(w, 400, fmt.Sprintf("Couldn't find user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
