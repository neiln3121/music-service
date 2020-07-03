package delivery

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/neiln3121/music-service/repository"
)

// GetArtists is the handler for retrieving all artists
func GetArtists(repo repository.Repository) http.HandlerFunc {
	if repo == nil {
		log.Fatal("Bad configuration: no repository provided")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		offset := 0
		limit := -1
		formValueOffSet := r.FormValue("offset")
		if len(formValueOffSet) > 0 {
			offset, err = strconv.Atoi(formValueOffSet)
			if err != nil {
				respondJSON(w, http.StatusBadRequest, newAppError("Invalid Offset", err))
				return
			}
		}

		formValueLimit := r.FormValue("limit")
		if len(formValueLimit) > 0 {
			limit, err = strconv.Atoi(formValueLimit)
			if err != nil {
				respondJSON(w, http.StatusBadRequest, newAppError("Invalid Limit", err))
				return
			}
		}

		streams, err := repo.GetArtists(offset, limit)
		if err != nil {
			respondJSON(w, http.StatusBadGateway, newAppError("Database error", err))
			return
		}

		respondJSON(w, http.StatusOK, streams)
	}
}

// GetAlbum is the handler for a single album
func GetAlbum(repo repository.Repository) http.HandlerFunc {
	if repo == nil {
		log.Fatal("Bad configuration: no repository provided")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		buff, err := repo.GetAlbum(id)
		if err != nil {
			respondJSON(w, http.StatusNotFound, newAppError("Record not found!", err))
			return
		}
		respondJSON(w, http.StatusOK, buff)
	}
}
