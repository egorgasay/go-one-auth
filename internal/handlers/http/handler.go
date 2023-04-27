package httphandler

import (
	"encoding/json"
	"errors"
	"go-one-auth/internal/schema"
	"go-one-auth/internal/storage/service"
	"go-one-auth/internal/usecase"
	"go-one-auth/pkg/logger"
	"io"
	"net/http"
)

type Handler struct {
	logic  *usecase.UseCase
	logger logger.ILogger
}

func New(logic *usecase.UseCase, logger logger.ILogger) *Handler {
	return &Handler{
		logic:  logic,
		logger: logger,
	}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	return
}

func bindJSON(body []byte, v interface{}) error {
	return json.Unmarshal(body, v)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	var kv schema.Pair

	res, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	if err = bindJSON(res, &kv); err != nil {
		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	if ok, err := h.logic.VerifyPair(r.Context(), kv.Key, kv.Value); err != nil || !ok {
		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreatePair(w http.ResponseWriter, r *http.Request) {
	var kv schema.Pair

	res, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	if err = bindJSON(res, &kv); err != nil {
		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err = h.logic.CreatePair(r.Context(), kv.Key, kv.Value); err != nil {
		if errors.Is(err, service.ErrPairExists) {
			w.WriteHeader(http.StatusConflict)
			return
		}

		h.logger.Warn(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
