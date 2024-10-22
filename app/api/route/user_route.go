package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(env *bootstrap.Env, userController controller.UserController, group *gin.RouterGroup) {
	group.GET("users", userController.GetEmployeeUsers)
	group.POST("users/create", userController.CreateUser)
	group.PUT("users/update", userController.UpdateUser)
	group.DELETE("users/delete", userController.DeleteUser)
	group.POST("users/ban", userController.BanUser)
	//TODO: вынести публичные действия в отдельный роутер
	group.GET("user/sessions", userController.GetUsersSessions)
}
