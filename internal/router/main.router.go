package router

import (
	_ "github.com/fajarworks/shortlink-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(r *gin.Engine) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
