package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"
	"app/app/repository"
	"app/app/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	userController := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(userRepository, timeout),
	}
	group.GET("users", userController.GetEmployeeUsers)
	group.POST("users/create", userController.CreateUser)
	group.PUT("users/update", userController.UpdateUser)
	group.DELETE("users/delete", userController.DeleteUser)
}
