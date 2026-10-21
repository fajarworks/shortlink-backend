package router

import (
	_ "github.com/fajarworks/shortlink-backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	AuthRouter(r, db)
}
