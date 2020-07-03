package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JeremyLoy/config"
	"github.com/gorilla/mux"
	"github.com/neiln3121/music-service/delivery"
	"github.com/neiln3121/music-service/models"
	"github.com/neiln3121/music-service/repository"
	"github.com/urfave/negroni"
)

func main() {
	// Read in generic config from environment variables
	var conf models.Config
	if err := config.FromEnv().To(&conf); err != nil {
		log.Fatalf("Bad config: %v", err)
	}

	db, err := models.ConnectDatabase(&conf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	repo := repository.CreateRepository(db)

	router := mux.NewRouter()
	neg := negroni.New(
		negroni.NewRecovery(),
	)

	neg.UseHandler(router)
	router.HandleFunc("/api/artists", delivery.GetArtists(repo)).Methods(http.MethodGet)
	router.HandleFunc("/api/albums/{id}", delivery.GetAlbum(repo)).Methods(http.MethodGet)

	log.Println("Starting up")
	server := &http.Server{Addr: fmt.Sprintf(":%d", conf.Port), Handler: neg}
	err = server.ListenAndServe()
}
