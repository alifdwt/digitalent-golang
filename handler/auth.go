package handler

import (
	"mygram-api/domain/requests/auth"
	_ "mygram-api/domain/responses/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthGroup(api *gin.Engine) {
	auth := api.Group("/users")

	auth.POST("/register", h.register)
	auth.POST("/login", h.login)
}

// register function
// @Summary Register to the application
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param data body auth.RegisterRequest true "user data"
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /users/register [post]
func (h *Handler) register(c *gin.Context) {
	var body auth.RegisterRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Auth.Register(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// login function
// @Summary Login to the application
// @Description Login with email and password to get a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param data body auth.LoginRequest true "user data"
// @Success 200 {object} responses.Token
// @Failure 400 {object} responses.ErrorMessage
// @Router /users/login [post]
func (h *Handler) login(c *gin.Context) {
	var body auth.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Auth.Login(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}
