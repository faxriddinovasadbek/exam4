package postgres

import (
	"comment-service/config"
	"comment-service/pkg/db"
	"comment-service/storage/repo"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"

	pbc "comment-service/protos/comment-service"
)

type CommentRepositrySuiteTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.CommentStorageI
}

func (s *CommentRepositrySuiteTest) SetupSuite() {
	pgPoll, err, cleanUp := db.ConnectToDB(config.Load())
	if err != nil {
		log.Fatal("Error while connecting database with suite test")
		return
	}
	s.CleanUpFunc = cleanUp
	s.Repository = NewCommentRepo(pgPoll)
}

// test func
func (s *CommentRepositrySuiteTest) TestCommnetCRUD() {
	// create comment
	commentReq := &pbc.Comment{
		Content: "suite content",
		PostId:  "7bbb6af9-d2ff-4c3f-a0a1-82b1f3f86964",
		OwnerId: "bff94bb9-0f5f-4893-acee-030ffd0df885",
	}
	createdComment, err := s.Repository.CreateCommment(commentReq)
	s.Suite.NoError(err)
	s.Suite.NotNil(createdComment)
	// update comment
	createdComment.Content = "new update content"
	updatedComment, err := s.Repository.UpdateComment(createdComment)
	s.Suite.NoError(err)
	s.Suite.NotNil(updatedComment)
	s.Suite.Equal(createdComment.Content, updatedComment.Content)
	// get comment
	getComment, err := s.Repository.GetComment(updatedComment.Id)
	s.Suite.NoError(err)
	s.Suite.NotNil(getComment)
	s.Suite.Equal(updatedComment, getComment)
	// get all comment
	allComment, err := s.Repository.GetAllComment(1, 20)
	s.Suite.NoError(err)
	s.Suite.NotNil(allComment)
	// get all comment by post_id
	commentsByPostId, err := s.Repository.GetAllCommentsByPostId(getComment.PostId)
	s.Suite.NoError(err)
	s.Suite.NotNil(commentsByPostId)
	// // get all comment by owner_id
	commentByOwnerId, err := s.Repository.GetAllCommentsByOwnerId(getComment.OwnerId)
	s.Suite.NoError(err)
	s.Suite.NotNil(commentByOwnerId)
	// get all comment by id
	commentById, err := s.Repository.GetCommentsById(getComment.Id)
	s.Suite.NoError(err)
	s.Suite.NotNil(commentById)
	// delete comment
	err = s.Repository.DeleteComment(updatedComment.Id)
	s.Suite.NoError(err)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CommentRepositrySuiteTest))
}
