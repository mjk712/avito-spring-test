package receptions

import "log/slog"

type Handler struct {
	receptionsUsecase receptionsUsecase
	log               *slog.Logger
}

func New(receptionsUsecase receptionsUsecase, log *slog.Logger) *Handler {
	return &Handler{
		receptionsUsecase: receptionsUsecase,
		log:               log,
	}
}
