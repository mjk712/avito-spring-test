package pvz

import (
	"avito-spring-test/internal/models/dto"
	"context"
)

type pvzUsecase interface {
	CreatePvz(ctx context.Context, request dto.CreatePvzRequest) (dto.PvzResponse, error)
	GetPvzList(ctx context.Context, startDate, endDate, page, limit string) ([]dto.PVZWithReceptions, error)
}
