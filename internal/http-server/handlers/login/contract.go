package login

import (
	"avito-spring-test/internal/models/dto"
	"context"
)

type authUsecase interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}
