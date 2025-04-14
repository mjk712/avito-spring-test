package register

import (
	"avito-spring-test/internal/models/dto"
	"avito-spring-test/internal/tools"
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Handler struct {
	authUsecase authUsecase
	log         *slog.Logger
}

func New(authUsecase authUsecase, log *slog.Logger) *Handler {
	return &Handler{
		authUsecase: authUsecase,
		log:         log,
	}
}

func (h *Handler) Register(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.register"

		w.Header().Set("Content-Type", "application/json")
		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req dto.RegisterRequest

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("error register", tools.ErrAttr(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, dto.ErrorResponse{Message: "error register"})
			return
		}
		
		user, err := h.authUsecase.RegisterUser(ctx, req)
		if err != nil {
			log.Error("error register", tools.ErrAttr(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, dto.ErrorResponse{Message: "error register"})
			return
		}
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, user)
	}
}
