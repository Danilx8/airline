package usecase

import (
	"app/app/domain"

	"github.com/gin-gonic/gin"
)

type RouteUsecase struct {
	routeRepository domain.RouteRepository
}

func NewRouteUsecase(routeRepo domain.RouteRepository) RouteUsecase {
	return RouteUsecase{
		routeRepository: routeRepo,
	}
}

func (r *RouteUsecase) GetByDepatureIDAndArrivalID(c *gin.Context, depId, arrId int) (*domain.Route, error) {
	var route *domain.Route
	var err error

	if depId == 0 && arrId != 0 {
		route, err = r.routeRepository.GetByArrivalID(arrId)
		if err != nil {
			return nil, err
		}
	} else if depId != 0 && arrId == 0 {
		route, err = r.routeRepository.GetByDepartureID(depId)
		if err != nil {
			return nil, err
		}
	} else {
		route, err = r.routeRepository.GetByDepartureIDAndArrivalID(depId, arrId)
		if err != nil {
			return nil, err
		}
	}
	return route, nil
}
