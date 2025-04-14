package dummy_login

import (
	"avito-spring-test/internal/models/dto"
	"context"
)

type authUsecase interface {
	DummyLogin(ctx context.Context, req dto.DummyLoginRequest) (string, error)
}
