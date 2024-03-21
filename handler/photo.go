package handler

import (
	"errors"
	"mygram-api/domain/requests/photo"
	_ "mygram-api/domain/responses/photo"
	"mygram-api/pkg/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initPhotoGroup(api *gin.Engine) {
	photo := api.Group("/photos")

	photo.Use(authMiddleware(h.tokenMaker))
	photo.GET("/", h.handlerGetPhotoAll)
	photo.GET("/:photoId", h.handlerGetPhotoById)

	photo.POST("/", h.handlerCreatePhoto)
	photo.PUT("/:photoId", h.handlerUpdatePhoto)
	photo.DELETE("/:photoId", h.handlerDeletePhoto)
}

// handlerCreatePhoto function
// @Summary Create photo
// @Description Create photo
// @Tags photos
// @Accept json
// @Produce json
// @Param data body photo.CreatePhotoRequest true "photo data"
// @Security BearerAuth
// @Success 201 {object} photo.PhotoResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos [post]
func (h *Handler) handlerCreatePhoto(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body photo.CreatePhotoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Photo.CreatePhoto(uint(userID), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerGetPhotoAll function
// @Summary Get all photos
// @Description Get all photos from the application
// @Tags photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []photo.PhotoWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos [get]
func (h *Handler) handlerGetPhotoAll(c *gin.Context) {
	res, err := h.services.Photo.GetPhotoAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetPhotoById function
// @Summary Get photo
// @Description Get photo by id
// @Tags photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param photoId path int true "Photo ID"
// @Success 200 {object} photo.PhotoWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos/{photoId} [get]
func (h *Handler) handlerGetPhotoById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Photo.GetPhotoById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerUpdatePhoto function
// @Summary Update photo
// @Description Update photo by id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoId path int true "Photo ID"
// @Param data body photo.UpdatePhotoRequest true "photo data"
// @Security BearerAuth
// @Success 200 {object} photo.PhotoResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos/{photoId} [put]
func (h *Handler) handlerUpdatePhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body photo.UpdatePhotoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	photo, err := h.services.Photo.GetPhotoById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if photo.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to update this photo")))
		return
	}

	res, err := h.services.Photo.UpdatePhoto(uint(userID), uint(id), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeletePhoto function
// @Summary Delete photo
// @Description Delete photo
// @Tags photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param photoId path int true "Photo ID"
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} responses.ErrorMessage
// @Router /photos/{photoId} [delete]
func (h *Handler) handlerDeletePhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	photo, err := h.services.Photo.GetPhotoById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if photo.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to delete this photo")))
		return
	}

	_, err = h.services.Photo.DeletePhoto(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your photo has been successfully deleted"})
}
