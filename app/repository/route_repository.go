package repository

import (
	"app/app/domain"
	"fmt"

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

func (r routeRepository) GetByDepartureID(depId int) (*domain.Route, error) {
	var route domain.Route
	if err := r.db.Table("Routes").
		Where("DepartureAirportID = ?", depId).
		Find(&route).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(route)

	return &route, nil
}

func (r routeRepository) GetByArrivalID(arrId int) (*domain.Route, error) {
	var route domain.Route
	if err := r.db.Table("Routes").
		Where("ArrivalAirportID = ?", arrId).
		Find(&route).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &route, nil
}

func (r routeRepository) GetByDepartureIDAndArrivalID(depId, arrId int) (*domain.Route, error) {
	var route domain.Route
	if err := r.db.Table("Routes").
		Where("ArrivalAirportID = ?", arrId).
		Where("DepartureAirportID = ?", depId).
		Find(&route).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &route, nil
}
