package service

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"github.com/nathakusuma/bcc-be-freepass-2024/repository"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/errortypes"
	"net/http"
	"time"
)

type ElectionPeriodService struct {
	PeriodRepo *repository.ElectionPeriodRepository
}

func NewElectionPeriodService(periodRepo *repository.ElectionPeriodRepository) *ElectionPeriodService {
	return &ElectionPeriodService{periodRepo}
}

func (service *ElectionPeriodService) GetPeriod() (*model.GetElectionPeriodResponse, *errortypes.ApiError) {
	start, end, err := service.PeriodRepo.GetPeriod()
	if err != nil {
		return nil, &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to get election period",
			Data:    err,
		}
	}
	return &model.GetElectionPeriodResponse{
		Start: start,
		End:   end,
	}, nil
}

func (service *ElectionPeriodService) SetPeriod(start, end time.Time) *errortypes.ApiError {
	if err := service.PeriodRepo.SetPeriod(start, end); err != nil {
		return &errortypes.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "fail to save election period",
			Data:    err,
		}
	}
	return nil
}

func (service *ElectionPeriodService) IsInPeriod() (bool, error) {
	start, end, err := service.PeriodRepo.GetPeriod()
	if err != nil {
		return false, err
	}
	now := time.Now()
	if now.Before(start) || now.After(end) {
		return false, nil
	}
	return true, nil
}
