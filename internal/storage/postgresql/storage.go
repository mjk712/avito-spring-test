package postgresql

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	// Подключаем драйвер PostgreSQL для работы с миграциями
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// Подключаем поддержку миграций из файловой системы
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func New(connectionString string) (*sqlx.DB, error) {
	const op = "storage.postgres.new"
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	m, err := migrate.New("file:///app/internal/migrations", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return db, nil
}
