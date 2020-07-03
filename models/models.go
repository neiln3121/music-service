package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Artist struct {
	gorm.Model
	Name   string
	Albums []*Album
}

type Album struct {
	ID       uint
	ArtistID uint
	Title    string
	Year     uint
	Tracks   []*Track
}

type Track struct {
	ID      uint
	AlbumID uint
	Title   string
}

type Config struct {
	Port int `config:"PORT"`

	DBServer   string `config:"DB_SERVER"`
	DBPort     int    `config:"DB_PORT"`
	DBName     string `config:"DB_NAME"`
	DBUser     string `config:"DB_USER"`
	DBPassword string `config:"DB_PASSWORD"`
}

func ConnectDatabase(config *Config) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	retries := 5
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		config.DBServer, config.DBUser, config.DBPassword, config.DBPort, config.DBName)

	for i := 0; i < retries; i++ {
		db, err = gorm.Open("mssql", connectionString)
		// Give container more time to start
		log.Println("Could not connect, retrying...")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, err
	}
	gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")

	db.AutoMigrate(&Artist{})
	db.AutoMigrate(&Album{})
	db.AutoMigrate(&Track{})

	return db, nil
}
