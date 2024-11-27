package usecase

import (
	"app/app/domain"
	"app/app/internal"
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/gin-gonic/gin"
)

type ScheduleUsecase struct {
	scheduleRepository domain.ScheduleRepository
	routeUsecase       RouteUsecase
}

func NewScheduleUsecase(scheduleRepo domain.ScheduleRepository, routeUsecase RouteUsecase) ScheduleUsecase {
	return ScheduleUsecase{
		scheduleRepository: scheduleRepo,
		routeUsecase:       routeUsecase,
	}
}

func (s *ScheduleUsecase) GetSchedules(c *gin.Context, defaultQuery map[string]string) (*[]domain.Schedules, error) {
	var schedules []domain.Schedules
	if err := s.scheduleRepository.GetSchedules(&schedules, defaultQuery); err != nil {
		return nil, err
	}
	for i := range schedules {
		route := &schedules[i]
		route.BussinesPrice = math.Floor(route.EconomyPrice * 1.35)
		route.FirstClassPrice = math.Floor(route.BussinesPrice * 1.3)
	}

	return &schedules, nil
}
func (s *ScheduleUsecase) UpdateFlightByNum(c *gin.Context, flight *domain.Schedules) error {
	if err := s.scheduleRepository.UpdateFlightByNum(flight); err != nil {
		return err
	}
	return nil
}

func (s *ScheduleUsecase) FormHandler(c *gin.Context, file io.Reader) (*domain.FormResponse, error) {
	var formResponse domain.FormResponse
	records, err := internal.ReadRecords(file, &formResponse)
	if !errors.Is(err, io.EOF) {
		return nil, err
	}
	mask := internal.CheckDuplicate(records, &formResponse)

	var recordModel domain.Record
	for i := range records {
		if mask[i] {
			continue
		}
		internal.ParsingRecord(records[i], &recordModel)
		switch recordModel.Type {
		case "ADD":

		case "EDIT":

		default:
		}
	}

	formResponse.SuccessfulChanges -= (formResponse.RecordWithMissingFields + formResponse.DuplicateRecords)
	if formResponse.SuccessfulChanges < 0 {
		return nil, fmt.Errorf("Invalid data in form")
	}
	return nil, nil
}
