package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/fajarworks/shortlink-backend/internal/dto"
	"github.com/fajarworks/shortlink-backend/internal/service"
	"github.com/fajarworks/shortlink-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

// CreateUser godoc
// @Summary      Register new user
// @Description  Create a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload body dto.CreateUser true "User registration data"
// @Success      201 {object} dto.Response "User successfully registered"
// @Failure      400 {object} dto.Response "Invalid request body or empty field"
// @Failure      409 {object} dto.ErrorResponse "User already exists"
// @Failure      500 {object} dto.Response "Internal server error"
// @Router       /auth/register [post]
func (c *AuthController) CreateUser(ctx *gin.Context) {
	var payload dto.CreateUser
	if err := ctx.ShouldBindWith(&payload, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}
	if err := c.service.CreateUser(ctx.Request.Context(), payload); err != nil {
		log.Println("[registrastion] err :", err.Error())
		switch {
		case errors.Is(err, utils.ErrFieldEmpty):
			ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		case errors.Is(err, utils.ErrFieldEmpty):
			ctx.JSON(http.StatusConflict, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: err.Error(),
			})

		}
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "user created",
	})
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and return JWT token with user information
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload body dto.LoginRequest true "Login credentials"
// @Success      200 {object} dto.Response "Login successfully"
// @Failure      400 {object} dto.ErrorResponse "Invalid request or empty field"
// @Failure      401 {object} dto.ErrorResponse "Wrong email or password"
// @Failure      500 {object} dto.ErrorResponse "Internal server error"
// @Router       /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var payload dto.LoginRequest

	if err := ctx.ShouldBindWith(&payload, binding.JSON); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "server error occured",
		})
		return
	}
	user, token, err := c.service.Login(ctx, payload)
	if err != nil {
		switch {

		case errors.Is(err, utils.ErrFieldEmpty):
			ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		case errors.Is(err, utils.ErrWrongEmailPass):
			ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		default:
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Message: "server error occured",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "login successfully",
		Result: gin.H{
			"token": token,
			"user":  user,
		},
	})
}
