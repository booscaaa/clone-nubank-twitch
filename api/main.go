package main

import (
	"api/provider"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Não consegui carregar as variaveis de ambiente")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	provider := provider.GetProvider()

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CompressHandler(provider)))
}
