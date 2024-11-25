package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"

	"github.com/gin-gonic/gin"
)

func NewRoutesRouter(env *bootstrap.Env, routeController controller.RouteController, group *gin.RouterGroup) {
	group.GET("routes/", routeController.GetAllRoutes)
	group.GET("routes/query", routeController.GetFlight)
}
