package controller

import (
	"net/http"

	"github.com/fajarworks/shortlink-backend/internal/dto"
	"github.com/fajarworks/shortlink-backend/internal/service"
	"github.com/fajarworks/shortlink-backend/pkg"
	"github.com/gin-gonic/gin"
)

type LinkController struct {
	service *service.LinkService
}

func NewLinkController(service *service.LinkService) *LinkController {
	return &LinkController{
		service: service,
	}
}

// CreateLink godoc
// @Summary      Create shortlink
// @Description  Membuat shortlink baru dengan custom slug atau slug otomatis
// @Tags         Shortlinks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateShortlinkRequest true "Data shortlink"
// @Success      201 {object} dto.Response{result=dto.ShortlinkResponse}
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /links [post]
func (c *LinkController) CreateLink(ctx *gin.Context) {
	var payload dto.CreateShortlinkRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	claims, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	token, ok := claims.(pkg.JWTClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Message: "Invalid token",
		})
		return
	}

	result, err := c.service.CreateLink(ctx.Request.Context(), token.Id, payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Shortlink created successfully",
		Result:  result,
	})
}
