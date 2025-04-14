package http_server

import (
	"avito-spring-test/internal/http-server/handlers/dummy_login"
	"avito-spring-test/internal/http-server/handlers/login"
	"avito-spring-test/internal/http-server/handlers/pvz"
	"avito-spring-test/internal/http-server/handlers/register"
	"avito-spring-test/internal/http-server/middleware/role_middleware"
	"avito-spring-test/internal/repositories/pvz_repository"
	"avito-spring-test/internal/repositories/user_repository"
	"avito-spring-test/internal/usecases/auth_usecase"
	"avito-spring-test/internal/usecases/pvz_usecase"
	"context"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"avito-spring-test/internal/config"
)

func NewServer(ctx context.Context, log *slog.Logger, cfg *config.Config, storage *sqlx.DB) *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api", func(r chi.Router) {
		r.Post("/dummyLogin", dummy_login.New(auth_usecase.New(user_repository.New(storage)), log).DummyLogin(ctx))
		r.Post("/register", register.New(auth_usecase.New(user_repository.New(storage)), log).Register(ctx))
		r.Post("/login", login.New(auth_usecase.New(user_repository.New(storage)), log).Login(ctx))
		r.With(role_middleware.RequireRole("moderator")).Post("/pvz", pvz.New(pvz_usecase.New(pvz_repository.New(storage)), log).CreatePvz(ctx))
		r.With(role_middleware.RequireRole("moderator", "employee")).Get("/pvz", pvz.New(pvz_usecase.New(pvz_repository.New(storage)), log).GetPvzList(ctx))
		/*
			/dummyLogin POST pass
			/register POST pass
			/login POST pass
			/pvz POST pass
			/pvz GET
			/pvz/{pvzId}/close_last_reception POST
			/pvz/{pvzId}/delete_last_product POST
			/receptions POST
			/products POST
		*/
	})

	return &http.Server{
		Addr:              cfg.Address,
		BaseContext:       func(net.Listener) context.Context { return ctx },
		Handler:           router,
		ReadHeaderTimeout: 1 * time.Second,
	}
}
