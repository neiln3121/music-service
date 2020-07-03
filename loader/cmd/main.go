package main

import (
	"fmt"
	"log"

	"github.com/neiln3121/music-service/loader"
	"github.com/neiln3121/music-service/models"
)

func main() {
	defaultConfig := models.Config{
		DBServer:   "localhost",
		DBPort:     1433,
		DBUser:     "SA",
		DBPassword: "Pass@word123",
		DBName:     "master",
	}

	db, err := models.ConnectDatabase(&defaultConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	// Create test data
	fmt.Println("Creating streams...")
	loader.LoadData(db)
}
