package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := chi.NewRouter()
	fmt.Println("hello world")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in ENV")
	}
	fmt.Println("Port", portString)
}
