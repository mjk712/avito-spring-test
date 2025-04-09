package product_repository

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) ProductRepo {
	return &repository{
		db: db,
	}
}
