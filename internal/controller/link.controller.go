package controller

import (
	"log"
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

// GetMyLinks godoc
// @Summary      Get my shortlinks
// @Description  Get all active shortlinks owned by the authenticated user
// @Tags         Shortlinks
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} dto.ShortlinkResponse
// @Failure      401 {object} dto.Response
// @Failure      500 {object} dto.Response
// @Router       /links [get]
func (c *LinkController) GetMyLinks(ctx *gin.Context) {
	value, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}

	claims, ok := value.(pkg.JWTClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "invalid token",
		})
		return
	}

	links, err := c.service.GetLinksByUserID(
		ctx.Request.Context(),
		claims.Id,
	)
	if err != nil {
		log.Println("failed to get user links:", err)

		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "failed to get shortlinks",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "shortlinks retrieved successfully",
		"data":    links,
	})
}
