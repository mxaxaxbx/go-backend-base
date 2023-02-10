package main

import (
	"context"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mxaxaxbx/go-backend-base/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {

}
