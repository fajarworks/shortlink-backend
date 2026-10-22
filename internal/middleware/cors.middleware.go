package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:8080"}
	if slices.Contains(allowedOrigins, ctx.GetHeader("Origin")) {
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
	}

	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS, PATCH, PUT, POST, DELETE")

	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}
