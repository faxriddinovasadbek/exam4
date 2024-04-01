package handlers

import (
	"api-gateway/api_test/storage"
	"api-gateway/api_test/storage/kv"
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"github.com/spf13/cast"
)

// User crud
func RegisterUser(c *gin.Context) {
	var newUser storage.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()
	newUser.Email = strings.ToLower(newUser.Email)
	err := newUser.Validate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	auth := smtp.PlainAuth("", "asadfaxriddinov611@gmail.com", "drkeagdlwrfanrdp", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "asadfaxriddinov611@gmail.com", []string{newUser.Email}, []byte(newUser.Email))
	if err != nil {
		log.Fatalf("Error sending otp to email: %v", err)
	}
	log.Println("Email sent successfully")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "One time password sent to your email",
	})
}

func Verify(c *gin.Context) {
	userCode := c.Param("code")

	if userCode != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect code",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func CreateUser(c *gin.Context) {
	var newUser storage.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 770); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func GetUser(c *gin.Context) {
	userID := c.Query("id")
	userString, err := kv.Get(userID)
	pp.Println(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.User
	if err := json.Unmarshal([]byte(userString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	userId := c.Query("id")
	if err := kv.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted successfully",
	})
}

func ListUsers(c *gin.Context) {
	usersStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	pp.Println(usersStrings)

	var users []*storage.User
	for _, userString := range usersStrings {
		var user storage.User
		if err := json.Unmarshal([]byte(userString), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, users)
}

// Post crud
func CreatePost(c *gin.Context) {
	var newPost storage.Post

	err := c.ShouldBindJSON(&newPost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userJson, err := json.Marshal(newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newPost.Id), string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newPost)
}

func GetPost(c *gin.Context) {
	PostID := c.Query("id")
	poststr, err := kv.Get(PostID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.Post
	if err := json.Unmarshal([]byte(poststr), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeletePost(c *gin.Context) {
	postId := c.Query("id")
	if err := kv.Delete(postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post was deleted successfully",
	})
}

func ListPost(c *gin.Context) {
	commentstrs, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var posts []*storage.Post
	for _, poststr := range commentstrs {
		var post storage.Post

		err := json.Unmarshal([]byte(poststr), &posts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		posts = append(posts, &post)
	}

	c.JSON(http.StatusOK, posts)
}


// Comment crud
func CreateComment(c *gin.Context) {
	var newComment storage.Comment

	err := c.ShouldBindJSON(&newComment)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	CommentJson, err := json.Marshal(newComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newComment.Id), string(CommentJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newComment)
}

func GetComment(c *gin.Context) {
	commentID := c.Query("id")
	poststr, err := kv.Get(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.Comment
	if err := json.Unmarshal([]byte(poststr), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteComment(c *gin.Context) {
	commentId := c.Query("id")
	if err := kv.Delete(commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment was deleted successfully",
	})
}

func ListComment(c *gin.Context) {
	commentstrs, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var comments []*storage.Comment
	for _, commentstr := range commentstrs {
		var post storage.Comment

		err := json.Unmarshal([]byte(commentstr), &comments)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		comments = append(comments, &post)
	}

	c.JSON(http.StatusOK, comments)
}
