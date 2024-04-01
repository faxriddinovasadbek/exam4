package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"

	models "api-gateway/api/handlers/models"
	l "api-gateway/pkg/logger"
	"api-gateway/pkg/utils"

	pbc "api-gateway/protos/comment-service"
)

// CreateComment ...
// @Summary CreateComment
// @Security ApiKeyAuth
// @Description Api for creating a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body models.CommentCreate true "create Comment Model"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/ [post]
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		body        models.CommentCreate
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().CreateComment(ctx, &pbc.Comment{
		Content: body.Content,
		PostId:  body.PostId,
		OwnerId: body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetComment gets comment by id
// @Summary GetComment
// @Security ApiKeyAuth
// @Description Api for getting comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/{id} [get]
func (h *handlerV1) GetComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetComment(
		ctx, &pbc.IdRequst{
			Id: id,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListUsers returns list of comments
// @Summary GetComment
// @Security ApiKeyAuth
// @Description Api for getting comment by page and limit
// @Tags comment
// @Accept json
// @Produce json
// @Param page path string true "PAGE"
// @Param limit path string true "LIMIT"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/ [get]
func (h *handlerV1) ListComment(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)

	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetAllComment(
		ctx, &pbc.GetAllCommentsRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by id
// @Summary UpdateComment
// @Security ApiKeyAuth
// @Description Api for getting comment by body
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body models.Comment true "updateCommentModel"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/update [put]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	var (
		body        pbc.Comment
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().UpdateComment(ctx, &pbc.Comment{
		Id:      body.Id,
		Content: body.Content,
		PostId:  body.PostId,
		OwnerId: body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes user by id
// @Summary DeleteComment
// @Security ApiKeyAuth
// @Description Api for getting comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments [delete]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().DeleteComment(ctx, &pbc.IdRequst{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}


// GetAllData data
// @Summary GetAllData
// @Security ApiKeyAuth
// @Description Api for get all data
// @Tags comment
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/all [get]
func (h *handlerV1) GetAllData(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Query("page")
	limit := c.Query("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetAllUsers(
		ctx, &pbc.GetAllCommentsRequest{
			Page:  cast.ToInt64(page),
			Limit: cast.ToInt64(limit),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all data", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetPostById gets post by id
// @Summary GetPostById
// @Security ApiKeyAuth
// @Description Api for getting post by id
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/post [get]
func (h *handlerV1) GetPostById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetPostById(
		ctx, &pbc.GetPostByIdRequest{
			PostId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post by id", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCommentsByPostId get comments by post id
// @Summary GetCommentsByPostId
// @Security ApiKeyAuth
// @Description Api for getting comment by post id
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "PostID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/commentpost [get]
func (h *handlerV1) GetCommentsByPostId(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetCommentsByPostId(
		ctx, &pbc.IdRequst{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get comment by post id", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCommentByOwner gets post by id
// @Summary GetCommentByOwner
// @Security ApiKeyAuth
// @Description Api for getting comment by owner
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "UserID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment [get]
func (h *handlerV1) GetCommentByOwner(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetCommentsByOwnerId(
		ctx, &pbc.IdRequst{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetUserById gets post by id
// @Summary GetUserById
// @Security ApiKeyAuth
// @Description Api for getting user data
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "UserID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user [get]
func (h *handlerV1) GetUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().GetUserById(
		ctx, &pbc.GetUserByIdRequest{
			OwnerId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
