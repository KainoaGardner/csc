package api

import (
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerTestRoutes(r chi.Router) {
	r.Get("/test/moveConvert", h.moveTest)
}

func (h *Handler) moveTest(w http.ResponseWriter, r *http.Request) {
	engine.RunTests()
	w.Write([]byte("Test Finished"))
}
