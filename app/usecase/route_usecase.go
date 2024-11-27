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
