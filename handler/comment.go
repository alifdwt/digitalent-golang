package handler

import (
	"errors"
	"mygram-api/domain/requests/comment"
	_ "mygram-api/domain/responses/comment"
	"mygram-api/pkg/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initCommentGroup(api *gin.Engine) {
	comment := api.Group("/comments")

	comment.Use(authMiddleware(h.tokenMaker))
	comment.GET("/", h.handlerGetCommentAll)
	comment.GET("/:commentId", h.handlerGetCommentById)

	comment.POST("/", h.handlerCreateComment)
	comment.PUT("/:commentId", h.handlerUpdateComment)
	comment.DELETE("/:commentId", h.handlerDeleteComment)
}

// handlerGetCommentAll function
// @Summary Get all comments
// @Description Get all comments from the application
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []comment.CommentWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /comments [get]
func (h *Handler) handlerGetCommentAll(c *gin.Context) {
	res, err := h.services.Comment.GetCommentAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerGetCommentById function
// @Summary Get comment
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param commentId path int true "Comment ID"
// @Success 200 {object} comment.CommentWithRelationResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /comments/{commentId} [get]
func (h *Handler) handlerGetCommentById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("commentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Comment.GetCommentById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerCreateComment function
// @Summary Create comment
// @Description Create comment
// @Tags comments
// @Accept json
// @Produce json
// @Param data body comment.CreateCommentRequest true "comment data"
// @Security BearerAuth
// @Success 201 {object} comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /comments [post]
func (h *Handler) handlerCreateComment(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseUint(authPayload.Subject, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body comment.CreateCommentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := body.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.services.Comment.CreateComment(uint(userID), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// handlerUpdateComment function
// @Summary Update comment
// @Description Update comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Param data body comment.UpdateCommentRequest true "comment data"
// @Security BearerAuth
// @Success 200 {object} comment.CommentResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /comments/{commentId} [put]
func (h *Handler) handlerUpdateComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("commentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body comment.UpdateCommentRequest
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

	comment, err := h.services.Comment.GetCommentById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if comment.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to update this comment")))
		return
	}

	res, err := h.services.Comment.UpdateComment(uint(id), uint(id), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// handlerDeleteComment function
// @Summary Delete comment
// @Description Delete comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} responses.ErrorMessage
// @Router /comments/{commentId} [delete]
func (h *Handler) handlerDeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("commentId"), 10, 64)
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

	comment, err := h.services.Comment.GetCommentById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if comment.UserID != uint(userID) {
		c.JSON(http.StatusUnauthorized, errorResponse(errors.New("you are not allowed to delete this comment")))
		return
	}

	_, err = h.services.Comment.DeleteComment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your comment has been successfully deleted"})
}
