package pvz_repository

import (
	"avito-spring-test/internal/models/dao"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) PvzRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) CreatePVZ(ctx context.Context, city string) (dao.Pvz, error) {
	query, args, err := squirrel.Insert("pvz").
		Columns("registration_date", "city").
		Values(squirrel.Expr("now()"), city).
		Suffix("RETURNING id, registration_date, city").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return dao.Pvz{}, err
	}

	var pvz dao.Pvz
	err = r.db.GetContext(ctx, &pvz, query, args...)
	return pvz, err
}

func (r *repository) ListPVZWithReceptions(ctx context.Context, start, end string, limit, offset int) ([]dao.Pvz, error) {
	builder := squirrel.Select("DISTINCT pvz.*").
		From("pvz").
		Join("receptions ON pvz.id = receptions.pvz_id").
		Where(squirrel.And{
			squirrel.GtOrEq{"receptions.date_time": start},
			squirrel.LtOrEq{"receptions.date_time": end},
		}).
		OrderBy("pvz.registration_date").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var result []dao.Pvz
	err = r.db.SelectContext(ctx, &result, query, args...)
	return result, err
}
