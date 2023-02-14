package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port string
}

type Broker struct {
	config *Config
	router *mux.Router
}

type Server interface {
	Config() *Config
}

func (b *Broker) Config() *Config {
	return b.config
}

var S Server

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	S = broker

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)

	log.Println("Starting server on port ", b.Config().Port)

	err := http.ListenAndServe(b.Config().Port, handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
