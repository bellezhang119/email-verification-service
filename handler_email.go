package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bellezhang119/email-verification-service/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateEmail(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	email, err := apiCfg.DB.CreateEmail(r.Context(), database.CreateEmailParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     params.Email,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create email: %v", err))
		return
	}

	respondWithJSON(w, 201, email)
}

func (apiCfg *apiConfig) handlerGetEmail(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("id")

	if param == "" {
		respondWithError(w, 400, "Invalid parameter")
		return
	}

	id, err := uuid.Parse(param)

	// if param is unable to be parsed as UUID
	if err != nil {
		// search param as string
		email, err := apiCfg.DB.GetEmail(r.Context(), param)
		if err != nil {
			respondWithError(w, 404, "Invalid email")
			return
		}
		respondWithJSON(w, 200, email)
	} else {
		// if param is able to be parsed as UUID, search param as UUID
		email, err := apiCfg.DB.GetEmailByID(r.Context(), id)
		if err != nil {
			respondWithError(w, 404, "Invalid id")
			return
		}
		respondWithJSON(w, 200, email)
	}
}
