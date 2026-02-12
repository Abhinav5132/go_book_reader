package main
import (

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"GoReader/models"

)

func setUpDb() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("GoReader.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} 
	
	database.AutoMigrate(&models.Book{}, &models.Franchise{}, &models.Library{}) // add the models here

	return database, nil
}