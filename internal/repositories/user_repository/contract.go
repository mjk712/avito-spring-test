package user_repository

import (
	"avito-spring-test/internal/models/dao"
	"context"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, email string, password string, role string) (dao.User, error)
	GetUserByEmail(ctx context.Context, email string) (dao.User, error)
}
