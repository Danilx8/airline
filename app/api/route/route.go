package route

import (
	"app/app/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewUserRouter(env, timeout, db, publicRouter)
	NewAuthRouter(env, timeout, db, publicRouter)
	SwaggerRouter(env, publicRouter)
}
