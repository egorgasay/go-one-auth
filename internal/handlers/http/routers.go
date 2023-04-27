package httphandler

import "github.com/go-chi/chi/v5"

func (h *Handler) PublicRoutes(r chi.Router) {
	r.Post("/v1/http/register", h.CreatePair)
	r.Post("/v1/http/verify", h.Verify)
	r.Head("/v1/http/ping", h.Ping)
}
