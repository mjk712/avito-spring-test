package pvz

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

		w.Header().Set("Content-Type", "application/json")
		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req dto.CreatePvzRequest

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("error create pvz: req data pass", tools.ErrAttr(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, dto.ErrorResponse{Message: "error pvz"})
			return
		}

		pvz, err := h.pvzUsecase.CreatePvz(r.Context(), req)
		if err != nil {
			log.Error("error create pvz: ", tools.ErrAttr(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, dto.ErrorResponse{Message: "error pvz"})
			return
		}
		render.JSON(w, r, pvz)
	}
}

func (h *Handler) GetPvzList(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetPvzList"

		w.Header().Set("Content-Type", "application/json")
		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		start := r.URL.Query().Get("startDate")
		end := r.URL.Query().Get("endDate")
		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")

		pvzList, err := h.pvzUsecase.GetPvzList(ctx, start, end, page, limit)
		if err != nil {
			log.Error("error get pvz list: ", tools.ErrAttr(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, dto.ErrorResponse{Message: "error get pvz list"})
			return
		}
		render.JSON(w, r, pvzList)
	}
}
