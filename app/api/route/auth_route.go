package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"
	"app/app/repository"
	"app/app/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	authController := controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(userRepository, sessionRepository, timeout),
		Env:         *env,
	}
	group.POST("login", authController.Login)
}
