package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port: ", portString)

	router := http.NewServeMux()

	router.HandleFunc("GET /ready", handlerReadiness)
	router.HandleFunc("GET /err", handleErr)

	router.HandleFunc("POST /hello", handler1)
	router.HandleFunc("POST /bye", handler2)

	server := http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
