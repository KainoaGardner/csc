package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/KainoaGardner/csc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	client   *mongo.Client
	dbConfig config.DB
}

func NewHandler(client *mongo.Client, dbConfig config.DB) *Handler {
	var result Handler
	result.client = client
	result.dbConfig = dbConfig

	return &result
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	h.registerGameRoutes(r)
	h.registerGameLogRoutes(r)
	h.registerTestRoutes(r)
}
