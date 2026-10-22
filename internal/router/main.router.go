package router

import (
	_ "github.com/fajarworks/shortlink-backend/docs"
	"github.com/fajarworks/shortlink-backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	r.Use(middleware.Cors)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	AuthRouter(r, db, rdb)
	LinkRouter(r, db, rdb)
}
