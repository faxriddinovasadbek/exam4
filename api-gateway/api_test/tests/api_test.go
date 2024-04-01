package tests

import (
	"api-gateway/api_test/handlers"
	"api-gateway/api_test/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {
	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	jsonUser, err := OpenFile("user.json")

	require.NoError(t, err)
	// User Create
	req := NewRequest(http.MethodPost, "/user/create", jsonUser)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/user/create", handlers.CreateUser)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var user storage.User

	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))

	require.NotNil(t, user.Id)
	require.Equal(t, "Asadbek", user.Name)
	require.Equal(t, "Faxriddinov", user.LastName)
	require.Equal(t, "asadfaxriddinov611@gmail.com", user.Email)
	require.Equal(t, "12345678", user.Password)
	require.Equal(t, "Asad15576", user.UserName)
	require.Equal(t, "Salom", user.Website)
	require.Equal(t, "Salom", user.Bio)

	// User Create
	// req := NewRequest(http.MethodPost, "/users", jsonUser)
	// r := gin.New()
	// res, err := Serve(v1.New(&v1.HandlerV1Config{}).CreateUser, "/users", req , r)
	// r.POST("/users", handlers.CreateUser)
	// r.ServeHTTP(res, req)
	// assert.NoError(t, err)
	// assert.Equal(t, http.StatusOK, res.Code)

	// var user storage.User

	// require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))

	// require.NotNil(t, user.Id)
	// require.Equal(t, "Asadbek", user.Name)
	// require.Equal(t, "Faxriddinov", user.LastName)
	// require.Equal(t, "asadfaxriddinov611@gmail.com", user.Email)
	// require.Equal(t, "12345678", user.Password)
	// require.Equal(t, "Asad15576", user.UserName)

	// GetUser
	getReq := NewRequest(http.MethodGet, "/users/get", jsonUser)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/get", handlers.GetUser)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp storage.User
	bodyBytes, err := io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.Name, getUserResp.Name)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Email, getUserResp.Email)
	assert.Equal(t, user.Password, getUserResp.Password)
	assert.Equal(t, user.UserName, getUserResp.UserName)

	// User List
	listReq := NewRequest(http.MethodGet, "/users", jsonUser)
	listRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users", handlers.ListUsers)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// User Delete
	delReq := NewRequest(http.MethodDelete, "/user/delete?id="+user.Id, jsonUser)
	delRes := httptest.NewRecorder()
	r.DELETE("/user/delete", handlers.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respm storage.Message
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &respm))
	require.Equal(t, "user was deleted successfully", respm.Message)

	// User Register
	regReq := NewRequest(http.MethodPost, "/user/register", jsonUser)
	regRes := httptest.NewRecorder()
	r.POST("/user/register", handlers.RegisterUser)
	r.ServeHTTP(regRes, regReq)
	assert.Equal(t, http.StatusOK, regRes.Code)
	var resp storage.Message
	bodyBytes, err = io.ReadAll(regRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resp))
	require.NotNil(t, resp.Message)
	require.Equal(t, "One time password sent to your email", resp.Message)

	// User Verify with correct code
	verURLCorrect := "/user/verify/12345"
	verReqCorrect := NewRequest(http.MethodGet, verURLCorrect, jsonUser)
	verResCorrect := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/user/verify/:code", handlers.Verify)
	r.ServeHTTP(verResCorrect, verReqCorrect)

	// Debugging
	// fmt.Println("Requested URL with correct code:", verURLCorrect)
	// fmt.Println("Response Code with correct code:", verResCorrect.Code)

	assert.Equal(t, http.StatusOK, verResCorrect.Code)
	var responseCorrect storage.Message
	bodyBytesCorrect, err := io.ReadAll(verResCorrect.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytesCorrect, &responseCorrect))
	require.Equal(t, "Success", responseCorrect.Message)

	// User Verify with incorrect code
	verURLIncorrect := "/user/verify/54321" // 54321 ni xato kod deb oldim va xatoliklar uchun ham tekshirib ketdim
	verReqIncorrect := NewRequest(http.MethodGet, verURLIncorrect, jsonUser)
	verResIncorrect := httptest.NewRecorder()
	r.ServeHTTP(verResIncorrect, verReqIncorrect)

	// Debugging
	// fmt.Println("Requested URL with incorrect code:", verURLIncorrect)
	// fmt.Println("Response Code with incorrect code:", verResIncorrect.Code)

	assert.Equal(t, http.StatusBadRequest, verResIncorrect.Code)
	var responseIncorrect storage.Message
	bodyBytesIncorrect, err := io.ReadAll(verResIncorrect.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytesIncorrect, &responseIncorrect))
	require.Equal(t, "Incorrect code", responseIncorrect.Message)

	//POST TEST

	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	jsonPost, err := OpenFile("post.json")

	require.NoError(t, err)

	//post create
	req = NewRequest(http.MethodPost, "/post/create", jsonPost)
	res = httptest.NewRecorder()
	r = gin.Default()
	r.POST("/post/create", handlers.CreatePost)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	var post storage.Post
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &post))
	require.NotNil(t, post.Id)
	require.Equal(t, post.Title, "Hello Golang")
	require.Equal(t, post.Content, "Hello Golang")
	require.Equal(t, post.Dislikes, "10")
	require.Equal(t, post.Views, "10")
	require.Equal(t, post.Category, "Hello Golang")
	require.NotNil(t, post.OwnerId)

	// Get Post
	getReq = NewRequest(http.MethodGet, "/post/get", jsonPost)
	q = getReq.URL.Query()
	q.Add("id", string(post.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/post/get", handlers.GetPost)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var getPost storage.Post
	bodyBytes, err = io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getPost))
	require.Equal(t, post.Id, getPost.Id)
	require.Equal(t, post.Title, getPost.Title)
	require.Equal(t, post.Content, getPost.Content)
	require.Equal(t, post.Like, getPost.Like)
	require.Equal(t, post.Dislikes, getPost.Dislikes)
	require.Equal(t, post.Views, getPost.Views)
	require.Equal(t, post.Category, getPost.Category)
	require.Equal(t, post.OwnerId, getPost.OwnerId)

	// List post
	listReq = NewRequest(http.MethodGet, "/posts", jsonPost)
	listRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/posts", handlers.ListPost)
	r.ServeHTTP(listRes, listReq)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// Delete post
	delReq = NewRequest(http.MethodDelete, "/post/delete?id="+post.Id, jsonPost)
	delRes = httptest.NewRecorder()
	r.DELETE("/post/delete", handlers.DeletePost)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respmessage storage.Message
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &respmessage))
	require.Equal(t, "post was deleted successfully", respmessage.Message)

	//COMMENT TEST

	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	jsonComment, err := OpenFile("comment.json")

	require.NoError(t, err)

	//comment create
	req = NewRequest(http.MethodPost, "/comment/create", jsonComment)
	res = httptest.NewRecorder()
	r = gin.Default()
	r.POST("/comment/create", handlers.CreateComment)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	var comment storage.Comment
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &comment))
	require.NotNil(t, comment.Id)
	require.Equal(t, comment.Content, "Hello Golang")
	require.NotNil(t, comment.OwnerId)
	require.NotNil(t, comment.PostId)


	// Get Comment
	getReq = NewRequest(http.MethodGet, "/comment/get", jsonComment)
	q = getReq.URL.Query()
	q.Add("id", string(comment.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/comment/get", handlers.GetComment)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var GetComment storage.Comment
	bodyBytes, err = io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &GetComment))
	require.Equal(t, comment.Id, GetComment.Id)
	require.Equal(t, comment.Content, GetComment.Content)
	require.Equal(t, comment.PostId, GetComment.PostId)
	require.Equal(t, comment.OwnerId, GetComment.OwnerId)

	// List comment
	listReq = NewRequest(http.MethodGet, "/comments", jsonComment)
	listRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/comments", handlers.ListComment)
	r.ServeHTTP(listRes, listReq)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// Delete comment
	delReq = NewRequest(http.MethodDelete, "/comment/delete?id="+comment.Id, jsonComment)
	delRes = httptest.NewRecorder()
	r.DELETE("/comment/delete", handlers.DeleteComment)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respmessage1 storage.Message
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &respmessage1))
	require.Equal(t, "comment was deleted successfully", respmessage1.Message)
}
