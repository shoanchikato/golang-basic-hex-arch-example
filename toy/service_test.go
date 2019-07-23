package toy_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shoanchikato/demo/hex/mocks"
	"github.com/shoanchikato/demo/hex/toy"
	"github.com/stretchr/testify/suite"
)

func TestToyServiceSuite(t *testing.T) {
	suite.Run(t, new(ToyServiceTestSuite))
}

type ToyServiceTestSuite struct {
	suite.Suite
	toyRepo   *mocks.MockRepository
	underTest toy.Service
}

func (suite *ToyServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.toyRepo = mocks.NewMockRepository(mockCtrl)
	suite.underTest = toy.NewToyService(suite.toyRepo)
}

func (suite *ToyServiceTestSuite) TestCreate() {
	// Arrange
	t := &toy.Toy{
		Name: "Ball",
	}
	suite.toyRepo.EXPECT().Create(gomock.Any()).Return(nil)

	// Act
	err := suite.underTest.CreateToy(t)
	fmt.Printf("this is the toy, %+v", t)

	//Assert
	suite.NoError(err, "Shouldn't error")
	suite.NotNil(t.ID, "should not be null")
	suite.NotNil(t.CreatedAt, "should not be null")
	suite.NotNil(t.UpdatedAt, "should not be null")
	suite.NotNil(t.Name, "should not be null")
	suite.NotNil(t.Owner, "should not be null")
}

func (suite *ToyServiceTestSuite) TestFindToyById() {
	t := &toy.Toy{
		ID:   "test",
		Name: "Ball",
	}
	suite.toyRepo.EXPECT().FindByID("test").Return(t, nil)

	result, err := suite.underTest.FindToyByID("test")

	suite.NoError(err, "Shouldn't error")
	suite.Equal(t, result, "should be pushing value returned from repo")
}

func (suite *ToyServiceTestSuite) TestFindAll() {
	ts := []*toy.Toy{
		&toy.Toy{
			ID:   "test1",
			Name: "Ball",
		},
		&toy.Toy{
			ID:   "test2",
			Name: "Cars",
		},
	}
	suite.toyRepo.EXPECT().FindAll().Return(ts, nil)

	result, err := suite.underTest.FindAllToys()

	suite.NoError(err, "Shouldn't error")
	suite.Len(result, 2, "Should get two results")
}
