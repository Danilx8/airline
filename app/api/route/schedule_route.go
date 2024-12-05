package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewScheduleRouter(env *bootstrap.Env, scheduleController controller.ScheduleController, group *gin.RouterGroup) {
	group.GET("schedules/", scheduleController.GetSchedules)
	group.PUT("schedules/flight", scheduleController.UpdateFlightByNum)
	group.POST("schedules/file", scheduleController.LoadFile)
}

