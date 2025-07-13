package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/KainoaGardner/csc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	client   *mongo.Client
	dbConfig config.DB
	jwt      config.JWT
}

func NewHandler(client *mongo.Client, dbConfig config.DB, jwt config.JWT) *Handler {
	var result Handler
	result.client = client
	result.dbConfig = dbConfig
	result.jwt = jwt

	return &result
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	h.registerUserRoutes(r)
	h.registerUserStatRoutes(r)
	h.registerGameRoutes(r)
	h.registerGameLogRoutes(r)
	h.registerTestRoutes(r)
}
