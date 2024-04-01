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
	pbp "api-gateway/protos/post-service"
)

// CreatePost ...
// @Summary CreatePost
// @Security ApiKeyAuth
// @Description Api for creating a new user
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.PostCreate true "createPostModel"
// @Success 200 {object} models.PostReq
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/ [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        models.PostCreate
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

	response, err := h.serviceManager.PostService().Create(ctx, &pbp.Post{
		Content:  body.Content,
		Title:    body.Title,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
		Views:    body.Views,
		Category: body.Category,
		OwnerId:  body.OwnerId,
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

// GetPost gets post by id
// @Summary GetPost
// @Security ApiKeyAuth
// @Description Api for getting post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.PostReq
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	// intID, err := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPost(
		ctx, &pbp.GetRequest{
			PostId: id,
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

// ListPosts returns list of posts
// @Summary GetListPosts
// @Security ApiKeyAuth
// @Description Api for getting post by page and limit
// @Tags post
// @Accept json
// @Produce json
// @Param page query string true "PAGE"
// @Param limit query string true "LIMIT"
// @Success 200 {object} models.PostReq
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/ [get]
func (h *handlerV1) ListPost(c *gin.Context) {
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

	response, err := h.serviceManager.PostService().GetAllPosts(
		ctx, &pbp.GetAllPostsRequest{
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

// UpdatePost updates user by id
// @Summary UpdatePost
// @Security ApiKeyAuth
// @Description Api for updating post by id
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.Post true "updatePostModel"
// @Success 200 {object} models.PostReq
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/ [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pbp.Post
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

	
	response, err := h.serviceManager.PostService().Update(ctx, &pbp.Post{
		Id:       body.Id,
		Content:  body.Content,
		Title:    body.Title,
		Likes:    body.Likes,
		Dislikes: body.Dislikes,
		Views:    body.Views,
		Category: body.Category,
		OwnerId:  body.OwnerId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}


	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes user by id
// @Summary UpdatePost
// @Security ApiKeyAuth
// @Description Api for deleting post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} bool
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Delete(ctx, &pbp.GetRequest{PostId: id})
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
// @Tags post
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/all/post/data [get]
func (h *handlerV1) GetAllPostData(c *gin.Context) {
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