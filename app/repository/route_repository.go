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

func (r *routeRepository) GetByDepatureID(depId int) *domain.Route {
	return nil
}

func (r *routeRepository) GetByArrivalID(arrId int) *domain.Route {
	return nil
}

func (r *routeRepository) GetByDepatureIDAndArrivalID(depId, arrId int) *domain.Route {
	return nil
}
