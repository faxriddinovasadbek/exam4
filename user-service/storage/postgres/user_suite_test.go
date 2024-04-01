package postgres

import (
	"log"
	"testing"
	"user-service/config"
	"user-service/pkg/db"
	pbu "user-service/protos/user-service"
	"user-service/storage/repo"

	"github.com/google/uuid"

	"github.com/stretchr/testify/suite"
)

type UserRepositorySuiteTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.UserStorageI
}

func (s *UserRepositorySuiteTest) SetupSuite() {
	pgPoll, err, cleanUp := db.ConnectToDB(config.Load())
	if err != nil {
		log.Fatal("Error while connecting database with suite test")
		return
	}
	s.CleanUpFunc = cleanUp
	s.Repository = NewUserRepo(pgPoll)
}

// test func
func (s *UserRepositorySuiteTest) TestUserCRUD() {
	// struct for create user
	user := &pbu.User{
		Name:     "test name",
		LastName: "test last name",
	}

	// uuid generating
	user.Id = uuid.New().String()

	// check create user method
	createdUser, err := s.Repository.Create(user)
	s.Suite.NotNil(createdUser)
	s.Suite.NoError(err)

	// struct for get user method
	userRequst := &pbu.UserRequest{
		UserId: createdUser.Id,
	}
	// check get user method
	getCreatedUser, err := s.Repository.Get(userRequst)
	s.Suite.NoError(err)
	s.Suite.NotNil(getCreatedUser)
	s.Suite.Equal(getCreatedUser, createdUser)

	// check update user method
	createdUser.Name = "new updated name"
	createdUser.LastName = "new updated last name"
	updatedUser, err := s.Repository.Update(createdUser)
	s.Suite.NoError(err)
	s.Suite.NotNil(updatedUser)
	s.Suite.NotEqual(updatedUser, getCreatedUser)

	// check get all users method
	getAllRequest := &pbu.GetAllUsersRequest{
		Page:  1,
		Limit: 30,
	}
	getAllUsers, err := s.Repository.GetAll(getAllRequest)
	s.Suite.NoError(err)
	s.Suite.NotNil(getAllUsers)

	// check delete user method
	deletedUser, err := s.Repository.Delete(userRequst)
	s.Suite.NoError(err)
	s.Suite.NotNil(deletedUser)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuiteTest))
}
