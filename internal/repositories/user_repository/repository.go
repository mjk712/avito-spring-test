package user_repository

import (
	"avito-spring-test/internal/models/dao"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) RegisterUser(ctx context.Context, email string, password string, role string) (dao.User, error) {
	const op = "user_repository.RegisterUser"
	query, args, err := squirrel.Insert("users").
		Columns("email", "password", "role").
		Values(email, password, role).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return dao.User{}, fmt.Errorf("%w: %s", err, op)
	}
	query += " RETURNING id, email, role"

	var user dao.User
	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&user)
	if err != nil {
		return dao.User{}, fmt.Errorf("%w: %s", err, op)
	}
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (dao.User, error) {
	const op = "user_repository.GetUserByEmail"
	var user dao.User
	query := `
	SELECT * FROM users WHERE email = $1;
`
	err := r.db.QueryRowxContext(ctx, query, email).StructScan(&user)
	if err != nil {
		return dao.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}
