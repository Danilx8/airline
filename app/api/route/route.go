package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"
	"app/app/middlewares"
	"app/app/repository"
	"app/app/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	routeRepository := repository.NewRouteRepository(db)
	scheduleRepository := repository.NewScheduleRepository(db)

	userController := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(userRepository, timeout),
	}
	authMiddleware := middlewares.AuthMiddleware{
		AuthUsecase: usecase.NewAuthUsecase(userRepository, sessionRepository, timeout),
		Env:         *env,
	}
	routeUsecase := usecase.NewRouteUsecase(routeRepository)
	routeController := controller.RouteController{
		RouteUsecase: routeUsecase,
	}
	scheduleController := controller.ScheduleController{
		ScheduleUsecase: usecase.NewScheduleUsecase(scheduleRepository, routeUsecase),
	}

	publicRouter := gin.Group("")
	NewAuthRouter(env, timeout, db, publicRouter)
	NewRoutesRouter(env, routeController, publicRouter)
	NewScheduleRouter(env, scheduleController, publicRouter)
	SwaggerRouter(env, publicRouter)

	privateRouter := gin.Group("")
	privateRouter.Use(authMiddleware.CheckAuth)
	NewUserRouter(env, userController, privateRouter)
}
