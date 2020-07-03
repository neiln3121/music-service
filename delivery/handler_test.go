package delivery_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/neiln3121/music-service/delivery"
	"github.com/neiln3121/music-service/mocks"
	"github.com/neiln3121/music-service/models"
	"github.com/stretchr/testify/suite"
)

type Handler_TestSuite struct {
	suite.Suite
	mockRepository *mocks.Repository
}

func TestHandler_Tests(t *testing.T) {
	suite.Run(t, new(Handler_TestSuite))
}

func (suite *Handler_TestSuite) SetupTest() {
	suite.mockRepository = &mocks.Repository{}
}

func (suite *Handler_TestSuite) Test_GetArtists_Valid() {
	streams := make([]*models.Artist, 10)
	suite.mockRepository.On("GetArtists", 0, -1).Return(streams, nil)

	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(delivery.GetArtists(suite.mockRepository))
	suite.NotNil(testHandler)

	req, err := http.NewRequest("GET", "/streams", nil)
	suite.NoError(err)

	testHandler.ServeHTTP(rr, req)

	suite.Equal(http.StatusOK, rr.Code, "Wrong status")

	response := new([]*models.Artist)
	json.Unmarshal(rr.Body.Bytes(), response)

	suite.Equal(10, len(*response))
}

func (suite *Handler_TestSuite) Test_GetArtists_Error() {

	mockError := fmt.Errorf("Bad Database connection")
	suite.mockRepository.On("GetArtists", 0, -1).Return(nil, mockError)

	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(delivery.GetArtists(suite.mockRepository))
	suite.NotNil(testHandler)

	req, err := http.NewRequest("GET", "/streams", nil)
	suite.NoError(err)

	testHandler.ServeHTTP(rr, req)

	suite.Equal(http.StatusBadGateway, rr.Code, "Wrong status")
}

func (suite *Handler_TestSuite) Test_GetAlbum_Valid() {
	buff := &models.Album{
		ID:    1,
		Title: "Track 1",
	}

	suite.mockRepository.On("GetAlbum", "1").Return(buff, nil)

	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(delivery.GetAlbum(suite.mockRepository))
	suite.NotNil(testHandler)

	req, err := http.NewRequest("GET", "/buff/1", nil)
	suite.NoError(err)

	vars := map[string]string{
		"id": "1",
	}
	testHandler.ServeHTTP(rr, mux.SetURLVars(req, vars))

	suite.Equal(http.StatusOK, rr.Code, "Wrong status")

	response := new(models.Album)
	json.Unmarshal(rr.Body.Bytes(), response)

	suite.Equal(buff, response)
}

func (suite *Handler_TestSuite) Test_GetAlbum_NotFound() {
	req, err := http.NewRequest("GET", "/album/1", nil)
	suite.NoError(err)

	mockError := fmt.Errorf("Not found")

	suite.mockRepository.On("GetAlbum", "").Return(nil, mockError)

	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(delivery.GetAlbum(suite.mockRepository))
	suite.NotNil(testHandler)

	testHandler.ServeHTTP(rr, req)

	suite.Equal(http.StatusNotFound, rr.Code, "Wrong status")
}
