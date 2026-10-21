package router

import (
	"github.com/fajarworks/shortlink-backend/internal/controller"
	"github.com/fajarworks/shortlink-backend/internal/repository"
	"github.com/fajarworks/shortlink-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRouter(r *gin.Engine, db *pgxpool.Pool) {
	authRouter := r.Group("auth")
	authRepo := repository.NewAuthRepo(db)
	authService := service.NewAuthService(authRepo)
	AuthController := controller.NewAuthController(authService)
	authRouter.POST("register", AuthController.CreateUser)

}
