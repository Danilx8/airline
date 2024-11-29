package repository

import (
	"app/app/domain"

	"gorm.io/gorm"
)

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) domain.RouteRepository {
	return &routeRepository{
		db: db,
	}
}

func (r routeRepository) GetByDepartureID(route *domain.Route, depId int) error {
	if err := r.db.Table("Routes").
		Where("DepartureAirportID = ?", depId).
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("DepartureAirport.Country").
		Preload("ArrivalAirport.Country").
		First(route).Error; err != nil {
		return err
	}
	return nil
}

func (r routeRepository) GetByArrivalID(route *domain.Route, arrId int) error {
	if err := r.db.Table("Routes").
		Where("ArrivalAirportID = ?", arrId).
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("DepartureAirport.Country").
		Preload("ArrivalAirport.Country").
		First(route).Error; err != nil {
		return err
	}

	return nil
}

func (r routeRepository) GetByDepartureIDAndArrivalID(route *domain.Route, depId, arrId int) error {
	if err := r.db.Table("Routes").
		Where("ArrivalAirportID = ?", arrId).
		Where("DepartureAirportID = ?", depId).
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("DepartureAirport.Country").
		Preload("ArrivalAirport.Country").
		First(route).Error; err != nil {
		return err
	}
	return nil
}

func (r routeRepository) GetAllRoutes(routes *[]domain.Route) error {
	if err := r.db.Table("Routes").
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("DepartureAirport.Country").
		Preload("ArrivalAirport.Country").
		Find(&routes).Error; err != nil {
		return err
	}
	return nil
}

func (r routeRepository) GetRouteIDByFromAndTo(from string, to string) (int, error) {
	var route domain.Route
	if err := r.db.Table("Routes").
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Joins("JOIN Airports AS DepartureAirport ON DepartureAirport.ID = Routes.DepartureAirportID").
		Where("DepartureAirport.IATACode = ?", from).
		Joins("JOIN Airports AS ArrivalAirport ON ArrivalAirport.ID = Routes.ArrivalAirportID").
		Where("ArrivalAirport.IATACode = ?", to).
		Find(&route).Error; err != nil {
		return 0, err
	}
	return route.ID, nil
}
