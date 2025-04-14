package pvz

import (
	"context"
	"log/slog"
	"net/http"
)

type Handler struct {
	pvzUsecase pvzUsecase
	log        *slog.Logger
}

func New(pvzUsecase pvzUsecase, log *slog.Logger) *Handler {
	return &Handler{
		pvzUsecase: pvzUsecase,
		log:        log,
	}
}

func (h *Handler) CreatePvz(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.CreatePvz"
	}
}
