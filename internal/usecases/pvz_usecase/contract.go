package pvz_usecase

import (
	"avito-spring-test/internal/models/dao"
	"avito-spring-test/internal/models/dto"
	"context"
	"time"
)

type repository interface {
	CreatePVZ(ctx context.Context, city string) (dao.Pvz, error)
	ListPVZWithReceptions(ctx context.Context, start, end time.Time, limit, offset int) ([]dto.PVZWithReceptions, error)
}
