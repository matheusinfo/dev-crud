package database

import (
	"dev-crud/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Instance *gorm.DB
var databaseError error

func Connect(connectionString string){
	Instance, databaseError = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if databaseError != nil {
		log.Fatal("Error connecting to database")
		return
	}

	log.Println("Database connected successfully")
}

func Migrate(){
	Instance.AutoMigrate(&models.Dev{})
	log.Println("Database migrated successfully")
}