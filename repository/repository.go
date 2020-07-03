package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/neiln3121/music-service/models"
)

// Repository - interface for retrieving artists and albums
type Repository interface {
	GetArtists(offset, limit int) ([]*models.Artist, error)
	GetAlbum(id string) (*models.Album, error)
	Close()
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) GetArtists(offset, limit int) (streams []*models.Artist, err error) {
	if err = r.DB.Order("created_at asc").Offset(offset).Limit(limit).Find(&streams).Error; err != nil {
		err = fmt.Errorf("Connection error: %v", err)
		return
	}
	for _, stream := range streams {
		r.DB.Model(&stream).Related(&stream.Albums)
		for _, buff := range stream.Albums {
			r.DB.Model(&buff).Related(&buff.Tracks)
		}
	}
	return
}

func (r *repo) GetAlbum(id string) (*models.Album, error) {
	var buff models.Album
	var err error
	if err = r.DB.Where("id = ?", id).First(&buff).Error; err != nil {
		return nil, err
	}
	r.DB.Model(&buff).Related(&buff.Tracks)

	return &buff, nil
}

func (r *repo) Close() {
	r.DB.Close()
}

// CreateRepository - create instance
func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
