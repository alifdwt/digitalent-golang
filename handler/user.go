package handler

import (
	"mygram-api/domain/requests/user"
	"mygram-api/pkg/token"
	"net/http"
	"strconv"

	_ "mygram-api/domain/responses/user"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserGroup(api *gin.Engine) {
	user := api.Group("/users")

	user.Use(authMiddleware(h.tokenMaker))
	user.PUT("/", h.handlerUpdateUser)
	user.DELETE("/", h.handlerDeleteUser)
}

// handlerUpdateUser function
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param data body user.UpdateUserRequest true "user data"
// @Security BearerAuth
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /users [put]
func (h *Handler) handlerUpdateUser(c *gin.Context) {
	var body user.UpdateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updatedUser, err := h.services.User.UpdateUserById(uint(userID), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// handlerDeleteUser function
// @Summary Delete user
// @Description Delete user from the application
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} responses.ErrorMessage
// @Router /users [delete]
func (h *Handler) handlerDeleteUser(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = h.services.User.DeleteUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your account has been successfully deleted"})
}
