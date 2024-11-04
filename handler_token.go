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

	respondWithJSON(w, 200, token)
}
