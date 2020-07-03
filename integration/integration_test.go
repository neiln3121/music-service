package integration_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/neiln3121/music-service/delivery"
	"github.com/neiln3121/music-service/loader"
	"github.com/neiln3121/music-service/models"
	"github.com/neiln3121/music-service/repository"
	"github.com/stretchr/testify/suite"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationTestSuite struct {
	suite.Suite
	container testcontainers.Container
	ctx       context.Context
	repo      repository.Repository
}

func Test_IntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupSuite() {
	var err error
	suite.ctx = context.Background()

	// Request an nginx container that exposes port 80
	req := testcontainers.ContainerRequest{
		Image:        "microsoft/mssql-server-linux:2017-latest",
		ExposedPorts: []string{"1433/tcp"},
		Env:          map[string]string{"SA_PASSWORD": "Pass@word", "ACCEPT_EULA": "Y"},
		WaitingFor:   wait.ForListeningPort("1433"),
	}
	suite.container, err = testcontainers.GenericContainer(suite.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	suite.NoError(err)

	// Retrieve the container IP
	ip, err := suite.container.Host(suite.ctx)
	suite.NoError(err)
	// Retrieve the port mapped to port
	port, err := suite.container.MappedPort(suite.ctx, "1433")
	suite.NoError(err)

	defaultConfig := models.Config{
		DBServer:   ip,
		DBPort:     port.Int(),
		DBUser:     "sa",
		DBPassword: "Pass@word",
		DBName:     "master",
	}

	db, err := models.ConnectDatabase(&defaultConfig)
	suite.NoError(err)

	suite.repo = repository.CreateRepository(db)

	// Load in test data
	loader.LoadData(db)
}

func (suite *IntegrationTestSuite) TearDownSuite() {
	suite.repo.Close()
	// At the end of the test remove the container
	suite.container.Terminate(suite.ctx)
}

func (suite *IntegrationTestSuite) Test_GetArtists_Valid() {
	server := httptest.NewServer(http.HandlerFunc(delivery.GetArtists(suite.repo)))
	defer server.Close()

	resp, err := http.Get(server.URL)
	suite.NoError(err)

	suite.Equal(http.StatusOK, resp.StatusCode, "Wrong status")
	actual, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	response := new([]*models.Artist)
	json.Unmarshal(actual, response)

	suite.Equal(3, len(*response))

	streams := *response
	suite.Equal(uint(1), streams[0].ID)
	// Should be 3
	suite.Equal(3, len(streams[0].Albums))
	suite.Equal(uint(1), streams[0].Albums[0].ArtistID)
	suite.Equal(uint(1), streams[0].Albums[1].ArtistID)
	suite.Equal(uint(1), streams[0].Albums[2].ArtistID)

	suite.Equal(9, len(streams[0].Albums[0].Tracks))
}

func (suite *IntegrationTestSuite) Test_GetArtists_Valid_Limit() {

	param := make(url.Values)
	param["limit"] = []string{"2"}

	server := httptest.NewServer(http.HandlerFunc(delivery.GetArtists(suite.repo)))
	defer server.Close()

	resp, err := http.Get(server.URL + "?" + param.Encode())
	suite.NoError(err)

	suite.Equal(http.StatusOK, resp.StatusCode, "Wrong status")
	actual, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	response := new([]*models.Artist)
	json.Unmarshal(actual, response)

	suite.Equal(2, len(*response))
}

func (suite *IntegrationTestSuite) Test_GetArtists_Valid_Offset() {

	param := make(url.Values)
	param["offset"] = []string{"1"}

	server := httptest.NewServer(http.HandlerFunc(delivery.GetArtists(suite.repo)))
	defer server.Close()

	resp, err := http.Get(server.URL + "?" + param.Encode())
	suite.NoError(err)

	suite.Equal(http.StatusOK, resp.StatusCode, "Wrong status")
	actual, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	response := new([]*models.Artist)
	json.Unmarshal(actual, response)

	suite.Equal(2, len(*response))

	streams := *response
	suite.Equal(uint(2), streams[0].ID)
}

func (suite *IntegrationTestSuite) Test_GetArtists_Invalid_Offset() {

	param := make(url.Values)
	param["offset"] = []string{"NOT-A-NUMBER"}

	server := httptest.NewServer(http.HandlerFunc(delivery.GetArtists(suite.repo)))
	defer server.Close()

	resp, err := http.Get(server.URL + "?" + param.Encode())
	suite.NoError(err)

	suite.Equal(http.StatusBadRequest, resp.StatusCode, "Wrong status")
	actual, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	response := new([]*models.Artist)
	json.Unmarshal(actual, response)

	suite.Equal(0, len(*response))
}

func (suite *IntegrationTestSuite) Test_GetArtists_Invalid_Limit() {

	param := make(url.Values)
	param["limit"] = []string{"NOT-A-NUMBER"}

	server := httptest.NewServer(http.HandlerFunc(delivery.GetArtists(suite.repo)))
	defer server.Close()

	resp, err := http.Get(server.URL + "?" + param.Encode())
	suite.NoError(err)

	suite.Equal(http.StatusBadRequest, resp.StatusCode, "Wrong status")
	actual, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	response := new([]*models.Artist)
	json.Unmarshal(actual, response)

	suite.Equal(0, len(*response))
}

func (suite *IntegrationTestSuite) Test_GetAlbum_Valid() {

	router := mux.NewRouter()
	router.HandleFunc("/api/albums/{id}", delivery.GetAlbum(suite.repo)).Methods(http.MethodGet)
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/api/albums/1")
	suite.NoError(err)

	suite.Equal(http.StatusOK, resp.StatusCode, "Wrong status")
	actual, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	response := new(models.Album)
	json.Unmarshal(actual, response)

	suite.Equal(uint(1), response.ID)
	suite.Equal("Purple Rain", response.Title)

	suite.Equal(9, len(response.Tracks))
	suite.Equal(uint(1), response.Tracks[0].AlbumID)
	suite.Equal(uint(1), response.Tracks[1].AlbumID)
	suite.Equal(uint(1), response.Tracks[2].AlbumID)
	suite.Equal(uint(1), response.Tracks[3].AlbumID)

	suite.Equal("Let's Go Crazy", response.Tracks[0].Title)
}

func (suite *IntegrationTestSuite) Test_GetAlbum_NotFound() {

	router := mux.NewRouter()
	router.HandleFunc("/api/albums/{id}", delivery.GetAlbum(suite.repo)).Methods(http.MethodGet)
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/api/albums/100")
	suite.NoError(err)

	suite.Equal(http.StatusNotFound, resp.StatusCode, "Wrong status")
	suite.NoError(err)
}
