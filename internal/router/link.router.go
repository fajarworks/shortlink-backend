package router

import (
	"github.com/fajarworks/shortlink-backend/internal/controller"
	"github.com/fajarworks/shortlink-backend/internal/middleware"
	"github.com/fajarworks/shortlink-backend/internal/repository"
	"github.com/fajarworks/shortlink-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func LinkRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	linkRouter := r.Group("links")
	linkRepo := repository.NewLinkRepo(db)
	linkService := service.NewLinkService(linkRepo)
	linkCOntroller := controller.NewLinkController(linkService)
	linkRouter.POST("", middleware.AuthMiddleware(rdb), linkCOntroller.CreateLink)

}
