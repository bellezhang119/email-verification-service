package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bellezhang119/email-verification-service/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port: ", portString)

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// TODO: create token schema, queries and handler
	router := http.NewServeMux()

	router.HandleFunc("GET /ready", handlerReadiness)
	router.HandleFunc("GET /err", handleErr)

	router.HandleFunc("POST /email", apiCfg.handlerCreateEmail)
	router.HandleFunc("GET /email/{id}", apiCfg.handlerGetEmail)

	server := http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
