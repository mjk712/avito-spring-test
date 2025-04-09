package pvz_repository

import (
	"avito-spring-test/internal/models/dao"
	"context"
)

type PvzRepo interface {
	CreatePVZ(ctx context.Context, city string) (dao.Pvz, error)
}
