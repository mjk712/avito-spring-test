package auth_usecase

import (
	"avito-spring-test/internal/models/dao"
	"avito-spring-test/internal/models/dto"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type UseCase struct {
	repository repository
}

func New(repository repository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (uc *UseCase) DummyLogin(ctx context.Context, req dto.DummyLoginRequest) (string, error) {
	const op = "usecase.DummyLogin"

	switch req.Role {
	case "moderator":
		token, err := GenerateJWT(req.Role)
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		return token, nil

	case "employee":
		token, err := GenerateJWT(req.Role)
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		return token, nil

	default:
		return "", fmt.Errorf("%s: error validate request data", op)
	}
}

func (uc *UseCase) RegisterUser(ctx context.Context, reqData dto.RegisterRequest) (dto.UserResponse, error) {
	const op = "usecase.RegisterUser"
	//check user exists
	_, err := uc.repository.GetUserByEmail(ctx, reqData.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Создаём нового пользователя если он не существует
			user, err := uc.repository.RegisterUser(ctx, reqData.Email, reqData.Password, reqData.Role)
			if err != nil {
				return dto.UserResponse{}, fmt.Errorf("%s: %w", op, err)
			}
			res := userDaoToDto(user)

			return res, nil
		} else {
			return dto.UserResponse{}, fmt.Errorf("%s: %w", op, err)
		}
	} else {
		// Нашли пользователя - сообщаем об этом
		return dto.UserResponse{}, fmt.Errorf("%s: user exists", op)
	}
}

func (uc *UseCase) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	const op = "usecase.Login"
	user, err := uc.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: user not found", op)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	switch user.Role {
	case "moderator":
		token, err := GenerateJWT(user.Role)
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		return token, nil

	case "employee":
		token, err := GenerateJWT(user.Role)
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
		return token, nil

	default:
		return "", fmt.Errorf("%s: error invalid role", op)
	}
}

func userDaoToDto(user dao.User) dto.UserResponse {
	var res dto.UserResponse
	res.ID = user.ID
	res.Email = user.Email
	res.Role = user.Role
	return res
}

func GenerateJWT(role string) (string, error) {
	const op = "service.GenerateJWT"
	claims := jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return tokenString, nil
}
