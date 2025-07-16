package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/KainoaGardner/csc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	client *mongo.Client

	config config.Config
}

func NewHandler(client *mongo.Client, config config.Config) *Handler {
	var result Handler
	result.client = client
	result.config = config

	return &result
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	h.registerUserRoutes(r)
	h.registerUserStatRoutes(r)
	h.registerGameRoutes(r)
	h.registerGameLogRoutes(r)
	h.registerTestRoutes(r)
}
