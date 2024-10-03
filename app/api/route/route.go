package route

import (
	"app/app/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewUserRouter(env, timeout, db, publicRouter)
	NewAuthRouter(env, timeout, db, publicRouter)
}
