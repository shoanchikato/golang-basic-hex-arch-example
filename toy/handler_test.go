package toy_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/shoanchikato/demo/hex/mocks"
	"github.com/shoanchikato/demo/hex/toy"
	"github.com/stretchr/testify/suite"
)

func TestToyHandlerSuite(t *testing.T) {
	suite.Run(t, new(ToyHandlerTestSuite))
}

type ToyHandlerTestSuite struct {
	suite.Suite
	toyService *mocks.MockService
	underTest  toy.Handler
}

func (suite *ToyHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.toyService = mocks.NewMockService(mockCtrl)
	suite.underTest = toy.NewToyHandler(suite.toyService)
}

func (suite *ToyHandlerTestSuite) TestCreate() {
	t := &toy.Toy{
		Name: "Boat",
		Owner: "Joe",
	}
	suite.toyService.EXPECT().CreateToy(gomock.Eq(t)).Return(nil)

	body, _ := json.Marshal(t)
	r, _ := http.NewRequest("POST", "/toys", bytes.NewBuffer(body))

	w := httptest.NewRecorder()
	suite.underTest.Create(w, r)

	response := w.Result()
	suite.Equal("201 Created", response.Status)

	defer response.Body.Close()
	result := new(toy.Toy)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal("Boat", result.Name)
	suite.Equal("Joe", result.Owner)
}

func (suite *ToyHandlerTestSuite) TestFindToyByID() {
	t := &toy.Toy{
		Name: "Boat",
		Owner: "Joe",
	}
	suite.toyService.EXPECT().FindToyByID("test").Return(t, nil)

	vars := map[string]string{
		"id": "test",
	}

	r, _ := http.NewRequest("GET", "/toys/test", nil)
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()
	suite.underTest.GetByID(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)
	
	defer response.Body.Close()
	result := new(toy.Toy)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal("Boat", result.Name)
	suite.Equal("Joe", result.Owner)
}

func (suite *ToyHandlerTestSuite) TestFindAll() {
	toys := []*toy.Toy{
		&toy.Toy{
			ID: "test1",
			Name: "Boat",
			Owner: "Joe",
		},
		&toy.Toy{
			ID: "test2",
			Name: "Other",
			Owner: "Joe",
		},
	}
	suite.toyService.EXPECT().FindAllToys().Return(toys, nil)

	r, _ := http.NewRequest("GET", "/toys", nil)

	w := httptest.NewRecorder()
	suite.underTest.Get(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	result := new([]toy.Toy)
	json.NewDecoder(response.Body).Decode(result)
	suite.Len(*result, 2, "Should get two results")
}
