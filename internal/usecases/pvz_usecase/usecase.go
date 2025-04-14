package pvz_usecase

import (
	"avito-spring-test/internal/models/dao"
	"avito-spring-test/internal/models/dto"
	"context"
	"fmt"
	"strconv"
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

func (uc *UseCase) CreatePvz(ctx context.Context, request dto.CreatePvzRequest) (dto.PvzResponse, error) {
	const op = "pvz_usecase.CreatePvz"

	switch request.City {
	case "Москва", "Санкт-Петербург", "Казань":
		pvz, err := uc.repository.CreatePVZ(ctx, request.City)
		if err != nil {
			return dto.PvzResponse{}, fmt.Errorf(op, err)
		}
		return daoPvzToDto(pvz), nil

	default:
		return dto.PvzResponse{}, fmt.Errorf(op, "invalid city")
	}
}

func (uc *UseCase) GetPvzList(ctx context.Context, startDate, endDate, page, limit string) ([]dto.PVZWithReceptions, error) {
	const op = "pvz_usecase.GetPvzList"
	//парсим данные
	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf(op, "err parsing start date")
	}
	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf(op, "err parsing end date")
	}
	pageParsed, err := strconv.Atoi(page)
	if err != nil && page == "" {
		pageParsed = 1
	} else if err != nil {
		return nil, fmt.Errorf(op, "err parsing page")
	}
	limitParsed, err := strconv.Atoi(limit)
	if err != nil && limit == "" {
		limitParsed = 10
	} else if err != nil {
		return nil, fmt.Errorf(op, "err parsing limit")
	}
	offset := (pageParsed - 1) * limitParsed

	pvzList, err := uc.repository.ListPVZWithReceptions(ctx, startDateParsed, endDateParsed, limitParsed, offset)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}
	return pvzList, nil
}

func daoPvzToDto(pvz dao.Pvz) dto.PvzResponse {
	return dto.PvzResponse{ID: pvz.ID, RegistrationDate: pvz.RegistrationDate, City: pvz.City}
}
