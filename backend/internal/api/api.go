package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/KainoaGardner/csc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/websockets"
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run(client *mongo.Client, config config.Config) error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "127.0.0.1", "localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/", func(r chi.Router) {
		handHandler := NewHandler(client, config)
		handHandler.RegisterRoutes(r)

	})

	engine.StartGlobalTimeCheck(5*time.Second, client, config, websockets.GameOver)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, r)

}
