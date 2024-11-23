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
	queryParams := c.Request.URL.Query()
	departure, _ := strconv.Atoi(queryParams.Get("departure"))
	arrival, _ := strconv.Atoi(queryParams.Get("arrival"))

	if departure == 0 && arrival == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{
			Header:      "Not found in query params departure or arrival or they have zero value",
			Description: fmt.Sprintf("Result found departure: %d and arrival: %d", departure, arrival),
		})
		return
	}

	route, err := r.RouteUsecase.GetByDepatureIDAndArrivalID(c, departure, arrival)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "Get error from RouteUsecase",
			Description: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, route)
}
