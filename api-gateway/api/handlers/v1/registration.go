package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"

	"strings"
	"time"

	"api-gateway/api/handlers/models"
	token "api-gateway/api/handlers/tokens"
	"api-gateway/pkg/etc"
	l "api-gateway/pkg/logger"
	pbu "api-gateway/protos/user-service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
)

// jwt

// Register ...
// @Summary Register
// @Description Api for registration
// @Tags register
// @Accept json
// @Produce json
// @Param User body models.User true "createUserModel"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/register/ [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		body        models.RegisterUser
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	if err = body.Validate(); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This password is already in use or email error, please choose another",
		})
		h.log.Error("failed to check email uniques", l.Error(err))
		return
	}

	exists, err := h.serviceManager.UserService().CheckUniques(ctx, &pbu.CheckUniquesRequest{
		Field: "email",
		Value: body.Email,
	})

	if err != nil && exists.Check {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check email uniques", l.Error(err))
		return
	}

	if exists.Check {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email is already in use, please choose another",
		})
		h.log.Error("failed to check email uniques", l.Error(err))
		return
	}

	exists, err = h.serviceManager.UserService().CheckUniques(ctx, &pbu.CheckUniquesRequest{
		Field: "username",
		Value: body.UserName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check username uniques", l.Error(err))
	}

	if exists.Check {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This username is already in use, please choose another",
		})
		h.log.Error("failed to check username uniques", l.Error(err))
		return

	}

	byteData, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not marshaled code",
		})
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	code := etc.GenerateCode(6)

	err = client.Set(context.Background(), code, byteData, time.Hour*2).Err()
	if err != nil {
		log.Fatal(err)
	}

	auth := smtp.PlainAuth("", "asadfaxriddinov611@gmail.com", "drkeagdlwrfanrdp", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "asadfaxriddinov611@gmail.com", []string{body.Email}, []byte(code))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	c.JSON(http.StatusOK, true)
}

// LogIn
// @Summary LogIn User
// @Description LogIn - Api for login users
// @Tags register
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/login [get]
func (h *handlerV1) LogIn(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetUserByEmail(ctx, &pbu.ByEmail{Email: email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user by email", l.Error(err))
		return
	}

	if response.Email != email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not found",
		})
		h.log.Error("email is not found")
		return
	}
	if !etc.CheckPasswordHash(password, response.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is not correct",
		})
		h.log.Error("password is not correct")
		return
	}

	h.jwthandler = token.JWTHandler{
		Sub:       response.Id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "user",
		SigninKey: h.cfg.SigningKey,
		Timeot:    h.cfg.AccessTokenTimout,
	}

	access, refresh_token, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		log.Fatal("error while generating auth token")
	}

	_, err = h.serviceManager.UserService().Update(ctx, &pbu.User{
		Id:           response.Id,
		Name:         response.Name,
		LastName:     response.LastName,
		Username:     response.Username,
		Email:        response.Email,
		RefreshToken: refresh_token,
		Password:     response.Password,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	var respModel models.UserBYtokens
	respModel.Id = response.Id
	respModel.Name = response.Name
	respModel.LastName = response.LastName
	respModel.UserName = response.Username
	respModel.Email = response.Email
	respModel.RefreshToken = refresh_token
	respModel.Password = response.Password
	respModel.AccessToken = access

	c.JSON(http.StatusOK, respModel)
}

// Verification
// @Summary Verification User
// @Description LogIn - Api for verification users
// @Tags register
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param code query string true "Code"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/verification [get]
func (h *handlerV1) Verification(c *gin.Context) {

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	code := c.Query("code")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	val, err := rdb.Get(ctx, code).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Verification code is expired",
		})
		h.log.Error("Verification code is expired", l.Error(err))
		return
	}

	var userdetail models.User
	if err := json.Unmarshal([]byte(val), &userdetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unmarshiling error",
		})
		h.log.Error("error unmarshalling userdetail", l.Error(err))
		return
	}

	if email != userdetail.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect email. Try again",
		})
		return
	}

	id := uuid.New().String()

	hashPassword, err := etc.HashPassword(userdetail.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error message",
		})
		h.log.Error("Error hashing possword", l.Error(err))
		return
	}

	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "user",
		SigninKey: h.cfg.SigningKey,
		Timeot:    h.cfg.AccessTokenTimout,
	}
	// aksestoken bn refreshtokeni generatsa qiliah
	access, refresh, err := h.jwthandler.GenerateAuthJWT()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "error generating token")
		return
	}

	createdUser, err := h.serviceManager.UserService().Create(ctx, &pbu.User{
		Id:           id,
		Name:         userdetail.Name,
		Bio:          userdetail.Bio,
		Website:      userdetail.Website,
		LastName:     userdetail.LastName,
		Email:        userdetail.Email,
		Password:     hashPassword,
		Username:     userdetail.UserName,
		RefreshToken: refresh,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating user",
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	response := &models.UserBYtokens{
		Id:           id,
		Name:         createdUser.Name,
		Bio:          userdetail.Bio,
		Website:      userdetail.Website,
		LastName:     createdUser.LastName,
		UserName:     createdUser.Username,
		Email:        createdUser.Email,
		Password:     hashPassword,
		AccessToken:  access,
		RefreshToken: refresh,
	}

	c.JSON(http.StatusOK, response)
}
