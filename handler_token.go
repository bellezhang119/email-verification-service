package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/bellezhang119/email-verification-service/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateToken(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		respondWithError(w, 400, "Invalid parameter")
		return
	}

	email, err := apiCfg.DB.GetEmailByID(r.Context(), id)

	if err != nil {
		respondWithError(w, 400, "Invalid parameter")
	}

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenString := make([]byte, 60)

	for i := range tokenString {
		var num *big.Int
		var err error
		// Retry loop for generating a random index
		for {
			num, err = rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			if err == nil {
				break // Exit the retry loop if no error
			}
		}
		tokenString[i] = letters[num.Int64()]
	}

	token, err := apiCfg.DB.CreateToken(r.Context(), database.CreateTokenParams{
		ID:        uuid.New(),
		EmailID:   email.ID,
		Token:     string(tokenString),
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().Add(2 * time.Hour),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating token: %v", err))
	}

	respondWithJSON(w, 201, token)
}

func (apiCfg *apiConfig) handlerVerify(w http.ResponseWriter, r *http.Request) {
	tokenParam := r.URL.Query().Get("token")

	if tokenParam == "" {
		respondWithError(w, 400, "Invalid verification token")
	}

	dbToken, err := apiCfg.DB.GetToken(r.Context(), tokenParam)

	if err != nil {
		respondWithError(w, 404, "Invalid verification token")
	}

	if dbToken.ExpiresAt.After(time.Now()) {
		err = apiCfg.DB.UpdateTokenIsUsed(r.Context(), database.UpdateTokenIsUsedParams{
			Token:  tokenParam,
			IsUsed: true,
		})

		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Error updating token: %v", err))
		}

		err = apiCfg.DB.UpdateEmailIsVerified(r.Context(), database.UpdateEmailIsVerifiedParams{
			ID:         dbToken.EmailID,
			IsVerified: true,
		})

		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Error updating email: %v", err))
		}

		respondWithJSON(w, 200, "Email successfully verified")
		return
	}

	respondWithError(w, 400, "Token is expired, please request another one")
}
