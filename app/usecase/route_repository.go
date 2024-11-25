package usecase

import (
	"app/app/domain"

	"github.com/gin-gonic/gin"
)

var orderMapper map[string]string = map[string]string{
	"date": "FlightTime",
}

type RouteUsecase struct {
	routeRepository domain.RouteRepository
}

func NewRouteUsecase(routeRepo domain.RouteRepository) RouteUsecase {
	return RouteUsecase{
		routeRepository: routeRepo,
	}
}

func (r *RouteUsecase) GetByDepatureIDAndArrivalID(c *gin.Context, depId, arrId int, sortBy string) (*domain.Route, error) {
	var route domain.Route
	var err error

	if depId == 0 && arrId != 0 {
		err = r.routeRepository.GetByArrivalID(&route, arrId)
		if err != nil {
			return nil, err
		}
	} else if depId != 0 && arrId == 0 {
		err = r.routeRepository.GetByDepartureID(&route, depId)
		if err != nil {
			return nil, err
		}
	} else {
		err = r.routeRepository.GetByDepartureIDAndArrivalID(&route, depId, arrId)
		if err != nil {
			return nil, err
		}
	}
	return &route, nil
}

func (r *RouteUsecase) GetAllRoutes(c *gin.Context) (*[]domain.Route, error) {
	var routes []domain.Route

	if err := r.routeRepository.GetAllRoutes(&routes); err != nil {
		return nil, err
	}

	return &routes, nil
}
