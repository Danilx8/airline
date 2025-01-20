package repository

import (
	"app/app/domain"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"time"
)

var orderMapper map[string]string = map[string]string{
	"date":         "Date, Time",
	"ticketPrice":  "EconomyPrice",
	"ticketStatus": "Confirmed",
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) domain.ScheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (s *scheduleRepository) GetSchedules(schedules *[]domain.Schedules, defaultQuery map[string]string) error {
	query := s.db.Table("Schedules").
		Preload("Aircraft").
		Preload("Route").
		Preload("Route.DepartureAirport").
		Preload("Route.DepartureAirport.Country").
		Preload("Route.ArrivalAirport").
		Preload("Route.ArrivalAirport.Country").
		Joins("JOIN Routes ON Routes.ID = Schedules.RouteID").
		Order(orderMapper[defaultQuery["order"]])

	if defaultQuery["from"] != "" {
		query = query.
			Joins("JOIN Airports AS DepartureAirport ON DepartureAirport.ID = Routes.DepartureAirportID").
			Where("DepartureAirport.IATACode = ?", defaultQuery["from"])
	}
	if defaultQuery["to"] != "" {
		query = query.
			Joins("JOIN Airports AS ArrivalAirport ON ArrivalAirport.ID = Routes.ArrivalAirportID").
			Where("ArrivalAirport.IATACode = ?", defaultQuery["to"])
	}
	if defaultQuery["outbound"] != "" {
		query = query.Where("Date = ?", defaultQuery["outbound"])
	}
	if defaultQuery["flightNumber"] != "" {
		query = query.Where("FlightNumber = ?", defaultQuery["flightNumber"])
	}

	if err := query.Find(&schedules).Error; err != nil {
		return err
	}
	return nil
}

func (s *scheduleRepository) UpdateFlightByNum(schedule *domain.Schedules) error {
	scheduleOld := &domain.Schedules{}
	result := s.db.Table("Schedules").
		Where("ID = ?", schedule.ID).
		First(scheduleOld)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch schedule with id %d: %w", schedule.ID, result.Error)
	}
	scheduleVal := reflect.ValueOf(schedule).Elem()
	scheduleOldVal := reflect.ValueOf(scheduleOld).Elem()

	date, _ := time.Parse(time.DateOnly, scheduleOldVal.FieldByName("Date").String())
	scheduleOldVal.FieldByName("Date").Set(reflect.ValueOf(date.Format(time.DateOnly)))

	for i := 0; i < scheduleVal.NumField(); i++ {
		value := scheduleVal.Field(i)
		if !value.IsValid() || scheduleVal.Type().Field(i).Name == "ID" || value.IsZero() {
			continue
		}

		if value.IsValid() && scheduleVal.Type().Field(i).Name == "Date" && !value.IsZero() {
			date, _ := time.Parse(time.DateOnly, value.String())
			scheduleOldVal.Field(i).Set(reflect.ValueOf(date.Format(time.DateOnly)))
		} else if value.IsValid() && scheduleVal.Type().Field(i).Name == "Time" && !value.IsZero() {
			timeOnly, _ := time.Parse(time.TimeOnly, value.String())
			scheduleOldVal.Field(i).Set(reflect.ValueOf(timeOnly.Format(time.TimeOnly)))
		} else if scheduleVal.Type().Field(i).Name == "confirmed" {
			scheduleOldVal.Field(i).Set(value)
		} else {
			scheduleOldVal.Field(i).Set(value)
		}
	}

	result = s.db.Table("Schedules").Save(scheduleOld)
	if result.Error != nil {
		return fmt.Errorf("failed to update schedule with id %d: %w", schedule.ID, result.Error)
	}

	return nil
}

func (s *scheduleRepository) AddRoute(schedule *domain.Schedules) error {
	if err := s.db.Table("Schedules").
		Create(schedule).Error; err != nil {
		return err
	}
	return nil
}

func (s *scheduleRepository) EditRoute(routeId int, schedule *domain.Schedules) error {
	schedId := s.GetIDByFields(routeId, schedule)
	if schedId == 0 {
		return fmt.Errorf("Not found record")
	}
	scheduleOld := &domain.Schedules{}
	result := s.db.Table("Schedules").
		Where("ID = ?", schedId).
		First(scheduleOld)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch schedule with id %d: %w", schedule.ID, result.Error)
	}
	scheduleVal := reflect.ValueOf(schedule).Elem()
	scheduleOldVal := reflect.ValueOf(scheduleOld).Elem()

	for i := 0; i < scheduleVal.NumField(); i++ {
		value := scheduleVal.Field(i)
		if !value.IsValid() || scheduleVal.Type().Field(i).Name == "ID" {
			continue
		}
		scheduleOldVal.Field(i).Set(value)
	}

	result = s.db.Table("Schedules").Save(scheduleOld)
	if result.Error != nil {
		return fmt.Errorf("failed to update schedule with id %d: %w", schedule.ID, result.Error)
	}

	return nil
}

func (s *scheduleRepository) GetIDByFields(routeId int, schedule *domain.Schedules) int {
	var resultSchedule domain.Schedules
	if err := s.db.Table("Schedules").
		Where("Date = ?", schedule.Date).
		Where("Time = ?", schedule.Time).
		Where("AircraftID = ?", schedule.AircraftID).
		Where("RouteID = ?", routeId).
		Where("EconomyPrice = ?", schedule.EconomyPrice).
		Where("FlightNumber = ?", schedule.FlightNumber).
		Find(&resultSchedule).Error; err != nil {
		return 0
	}
	return resultSchedule.ID
}
