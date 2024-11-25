package controller

import (
	"app/app/domain"
	"app/app/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RouteController struct {
	RouteUsecase usecase.RouteUsecase
}

func (r *RouteController) GetFlight(c *gin.Context) {
	sortBy := "date"
	queryParams := c.Request.URL.Query()
	departure, _ := strconv.Atoi(queryParams.Get("departure"))
	arrival, _ := strconv.Atoi(queryParams.Get("arrival"))

	if checkSortBy := queryParams.Get("sortBy"); checkSortBy != "" {
		sortBy = checkSortBy
	}

	if departure == 0 && arrival == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{
			Header:      "Not found in query params departure or arrival or they have zero value",
			Description: fmt.Sprintf("Result found departure: %d and arrival: %d", departure, arrival),
		})
		return
	}

	route, err := r.RouteUsecase.GetByDepatureIDAndArrivalID(c, departure, arrival, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "Get error from RouteUsecase",
			Description: err.Error(),
		})
		return
	}
	if route == nil {
		c.JSON(http.StatusNotFound, domain.ErrorMessage{
			Header:      "Not found",
			Description: "Not found route by query",
		})
		return
	}
	c.JSON(http.StatusOK, route)
}

func (r *RouteController) GetAllRoutes(c *gin.Context) {
	routes, err := r.RouteUsecase.GetAllRoutes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "Get error from RouteUsecase",
			Description: err.Error(),
		})
		return
	}
	fmt.Println("asdasd")
	c.JSON(http.StatusOK, routes)
}
