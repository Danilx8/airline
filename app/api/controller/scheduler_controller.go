package controller

import (
	"app/app/domain"
	"app/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	ScheduleUsecase usecase.ScheduleUsecase
}

func (s *ScheduleController) GetSchedules(c *gin.Context) {
	defaultQuery := map[string]string{
		"order":        "date",
		"from":         "",
		"to":           "",
		"outbound":     "",
		"flightNumber": "",
	}
	queryParams := c.Request.URL.Query()
	if orderInQuery := queryParams.Get("order"); orderInQuery != "" {
		defaultQuery["order"] = orderInQuery
	}
	defaultQuery["from"] = queryParams.Get("from")
	defaultQuery["to"] = queryParams.Get("to")
	defaultQuery["outbound"] = queryParams.Get("outbound")
	defaultQuery["flightNumber"] = queryParams.Get("flightNumber")

	schedules, err := s.ScheduleUsecase.GetSchedules(c, defaultQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "Get error from ScheduleUsecase",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (s *ScheduleController) UpdateFlightByNum(c *gin.Context) {
	var flight domain.Schedules

	err := c.BindJSON(&flight)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{
			Header:      "Get error while parse flight",
			Description: err.Error(),
		})
		return
	}

	err = s.ScheduleUsecase.UpdateFlightByNum(c, &flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "Get error from ScheduleUsecase",
			Description: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, flight)
}

func (s *ScheduleController) LoadFile(c *gin.Context) {
	rFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{
			Header:      "file format error",
			Description: err.Error(),
		})
		return
	}
	file, err := rFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "can't open file",
			Description: err.Error(),
		})
		return
	}

	s.ScheduleUsecase.FormHandler(c, file)

	c.JSON(http.StatusOK, gin.H{"OK": "OK"})
}
