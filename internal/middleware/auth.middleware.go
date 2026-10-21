package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fajarworks/shortlink-backend/internal/dto"
	"github.com/fajarworks/shortlink-backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(rc *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.GetHeader("Authorization")

		if bearer == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "please login first",
			})
			return
		}

		result := strings.Split(bearer, " ")
		if len(result) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}
		if result[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "invalid bearer token",
			})
			return
		}
		var token pkg.JWTClaims
		err := token.VerifyToken(result[1])
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
					Success: false,
					Message: "invalid token",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "an server error occured",
			})
			return

		}
		key := fmt.Sprintf("shortlink:blacklist:%s", token.ID)

		blacklisted, err := rc.Exists(ctx.Request.Context(), key).Result()
		if err != nil {
			log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "server error occured",
			})
			return
		}
		if blacklisted > 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Message: "token has expired",
			})
			return
		}
		ctx.Set("token", token)
		ctx.Next()
	}

}
