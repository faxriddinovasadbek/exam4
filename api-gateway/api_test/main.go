package main

import (
	"database/sql"
	"log"
	"net/http"
	"api-gateway/api_test/handlers"
	"api-gateway/api_test/storage/kv"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := "user=postgres password=asadbek dbname=userdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	kv.Init(kv.NewPostgres(db))

	router := gin.New()

	router.POST("/user/register", handlers.RegisterUser)
	router.POST("/user/verify/:code", handlers.Verify)
	router.GET("/user/get", handlers.GetUser)
	router.POST("/user/create", handlers.CreateUser)
	router.DELETE("/user/delete", handlers.DeleteUser)
	router.GET("/users", handlers.ListUsers)

	router.POST("/post/create", handlers.CreatePost)
	router.GET("/post/get", handlers.GetPost)
	router.GET("/posts", handlers.ListPost)
	router.DELETE("/post/delete", handlers.DeletePost)

	router.POST("/comment/create", handlers.CreateComment)
	router.GET("/comment/get", handlers.GetComment)
	router.GET("/commnets", handlers.ListComment)
	router.DELETE("/comment/delete", handlers.DeleteComment)

	log.Fatal(http.ListenAndServe(":9999", router))
}
