package route

import (
	"app/app/bootstrap"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRouter(env *bootstrap.Env, group *gin.RouterGroup) {
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
