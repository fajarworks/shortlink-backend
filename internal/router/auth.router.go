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

func AuthRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	authRouter := r.Group("auth")
	authRepo := repository.NewAuthRepo(db)
	authService := service.NewAuthService(authRepo, rdb)
	AuthController := controller.NewAuthController(authService)
	authRouter.POST("register", AuthController.CreateUser)
	authRouter.POST("login", AuthController.Login)
	authRouter.POST("logout", middleware.AuthMiddleware(rdb), AuthController.Logout)
}
