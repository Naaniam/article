package migrations

import (

	//user defined package
	"article/models"

	//third party package
	"gorm.io/gorm"
)

// Function for table migration
func Migrations(db *gorm.DB) {
	db.AutoMigrate(&models.Article{}, &models.Comment{}, models.Reply{})
}
