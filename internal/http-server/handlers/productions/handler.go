package productions

import "log/slog"

type Handler struct {
	productionUsecase productionUsecase
	log               *slog.Logger
}

func New(productionUsecase productionUsecase, log *slog.Logger) *Handler {
	return &Handler{
		productionUsecase: productionUsecase,
		log:               log,
	}
}
