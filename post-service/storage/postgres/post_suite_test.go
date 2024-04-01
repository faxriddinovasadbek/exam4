package postgres

import (
	"log"
	"post-service/config"
	"post-service/pkg/db"
	"post-service/storage/repo"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	pbp "post-service/protos/post-service"
)

type PostRepositorySuiteTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.PostStorageI
}

func (s *PostRepositorySuiteTest) SetupSuite() {
	pgPoll, err, cleanUp := db.ConnectToDB(config.Load())
	if err != nil {
		log.Fatal("Error while connecting database with suite test")
		return
	}
	s.CleanUpFunc = cleanUp
	s.Repository = NewPostRepo(pgPoll)
}

// test func
func (s *PostRepositorySuiteTest) TestUserCRUD() {
	// create post
	postReq := &pbp.Post{
		Id:       uuid.New().String(),
		Title:    "Test Post Title",
		OwnerId:  "bff94bb9-0f5f-4893-acee-030ffd0df885",
	}
	createdPost, err := s.Repository.Create(postReq)
	s.Suite.NotNil(createdPost)
	s.Suite.NoError(err)
	s.Suite.Equal(createdPost, postReq)

	// get post
	getPost, err := s.Repository.GetPost(&pbp.GetRequest{PostId: createdPost.Id})
	s.Suite.NoError(err)
	s.Suite.NotNil(getPost)
	s.Suite.Equal(getPost.Id, createdPost.Id)

	
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(PostRepositorySuiteTest))
}
