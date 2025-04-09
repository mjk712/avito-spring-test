package user_repository

import (
	"avito-spring-test/internal/models/dao"
	"context"
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

func (r *repository) RegisterUser(ctx context.Context, user dao.User) error {
	query, args, err := squirrel.Insert("users").
		Columns("email", "password_hash", "role").
		Values(user.Email, user.PasswordHash, user.Role).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (dao.User, error) {
	query, args, err := squirrel.Select("id", "email", "password_hash", "role").
		From("users").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return dao.User{}, err
	}

	var user dao.User
	err = r.db.GetContext(ctx, &user, query, args...)
	return user, err
}
