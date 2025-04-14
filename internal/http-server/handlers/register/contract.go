package register

import (
	"avito-spring-test/internal/models/dto"
	"context"
)

type authUsecase interface {
	RegisterUser(ctx context.Context, reqData dto.RegisterRequest) (dto.UserResponse, error)
}
