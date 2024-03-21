package handler

import (
	"errors"
	socialmedia "mygram-api/domain/requests/social_media"
	_ "mygram-api/domain/responses/social_media"
	"mygram-api/pkg/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initSocialMediaGroup(api *gin.Engine) {
	socialMedia := api.Group("/socialmedias")

	socialMedia.Use(authMiddleware(h.tokenMaker))
	socialMedia.GET("/", h.handlerGetSocialMediaAll)
	socialMedia.GET("/:socialMediaId", h.handlerGetSocialMediaById)
	socialMedia.POST("/", h.handlerCreateSocialMedia)
	socialMedia.PUT("/:socialMediaId", h.handlerUpdateSocialMedia)
	socialMedia.DELETE("/:socialMediaId", h.handlerDeleteSocialMedia)
}

// handlerGetSocialMediaAll function
// @Summary Get all social medias
// @Description Get all social medias from the application
// @Tags social medias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []socialmedia.SocialMediaResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /socialmedias [get]
func (h *Handler) handlerGetSocialMediaAll(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.SocialMedia.GetSocialMediaAll(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetSocialMediaById function
// @Summary Get social media
// @Description Get social media by id
// @Tags social medias
// @Accept json
// @Produce json
// @Param socialMediaId path int true "Social Media ID"
// @Security BearerAuth
// @Success 200 {object} socialmedia.SocialMediaResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /socialmedias/{socialMediaId} [get]
func (h *Handler) handlerGetSocialMediaById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userId, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.SocialMedia.GetSocialMediaById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if res.UserID != uint(userId) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to access this social media")))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerCreateSocialMedia function
// @Summary Create social media
// @Description Create social media
// @Tags social medias
// @Accept json
// @Produce json
// @Param data body socialmedia.CreateSocialMediaRequest true "social media data"
// @Security BearerAuth
// @Success 201 {object} socialmedia.SocialMediaResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /socialmedias [post]
func (h *Handler) handlerCreateSocialMedia(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body socialmedia.CreateSocialMediaRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.SocialMedia.CreateSocialMedia(uint(userID), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerUpdateSocialMedia function
// @Summary Update social media
// @Description Update social media
// @Tags social medias
// @Accept json
// @Produce json
// @Param socialMediaId path int true "Social Media ID"
// @Param data body socialmedia.UpdateSocialMediaRequest true "social media data"
// @Security BearerAuth
// @Success 200 {object} socialmedia.SocialMediaResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /socialmedias/{socialMediaId} [put]
func (h *Handler) handlerUpdateSocialMedia(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body socialmedia.UpdateSocialMediaRequest
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

	socialMedia, err := h.services.SocialMedia.GetSocialMediaById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if socialMedia.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to update this social media")))
		return
	}

	res, err := h.services.SocialMedia.UpdateSocialMedia(uint(id), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteSocialMedia function
// @Summary Delete social media
// @Description Delete social media
// @Tags social medias
// @Accept json
// @Produce json
// @Param socialMediaId path int true "Social Media ID"
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} responses.ErrorMessage
// @Router /socialmedias/{socialMediaId} [delete]
func (h *Handler) handlerDeleteSocialMedia(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)
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

	socialMedia, err := h.services.SocialMedia.GetSocialMediaById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if socialMedia.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to delete this social media")))
		return
	}

	_, err = h.services.SocialMedia.DeleteSocialMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "social media deleted successfully"})
}
