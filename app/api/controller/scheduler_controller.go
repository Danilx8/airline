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

// GetSchedules godoc
// @Summary	List of schedule
// @Description get schedule
// @Tags Schedule
// @Accept json
// @Produce json
// @Param        order    query     string true  "order of routes [date, ticketPrice, ticketStatus]"
// @Param        from    query     string true  "from of routes"
// @Param        to    query     string true  "to of routes"
// @Param        outbound    query     string true  "outbound of routes [1970-10-24]"
// @Param        flightNumber    query     string true  "flight number of routes [000]"
// @Success 200 {array} domain.Schedules
// @Failure 500 {object} domain.ErrorMessage
// @Router /schedules [get]
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

// UpdateFlightByNum godoc
// @Summary	Update flight by num
// @Description update flight
// @Tags Schedule
// @Accept json
// @Produce json
// @Param        data    body     domain.Schedules true  "scheme of schedules"
// @Success 200 {object} domain.Schedules
// @Failure 500 {object} domain.ErrorMessage
// @Router /schedules/flight [put]
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

// LoadFile godoc
// @Summary	Load file
// @Description Load file with schedule
// @Tags Schedule
// @Accept multipart/form-data
// @Param  file formData file true "account image"
// @Produce json
// @Success 200 {object} domain.FormResponse
// @Failure 500 {object} domain.ErrorMessage
// @Router /schedules/file [post]
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

	formResponse, err := s.ScheduleUsecase.FormHandler(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{
			Header:      "can't open file",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, formResponse)
}
