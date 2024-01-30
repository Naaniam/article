package drivers

import (

	// built-in package

	"article/helpers"
	"fmt"
	"os"
	"time"

	// third party package

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SQLDriver() *gorm.DB {
	var err error
	helpers.Log.WithFields(logrus.Fields{
		"service":    "article",
		"function":   "DB connection",
		"started_at": time.Now(),
	}).Info("DB connection establishment -started")

	err = godotenv.Load(".env")
	if err != nil {
		helpers.Log.WithFields(logrus.Fields{
			"service": "article",
			"error":   err.Error(),
		}).Errorf("Error loading .env file")
		return nil
	}

	//Database connection establishment
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.Log.WithFields(logrus.Fields{
			"service": "article",
			"message": err.Error(),
		}).Error("Failed to connect to the database")
		panic(err)
	}

	helpers.Log.WithFields(logrus.Fields{
		"service": "article",
		"message": "Established a successful connection to the database!",
	}).Info("Message : Established a successful connection to the database!")
	return db
}

func TestSQLDriver() *gorm.DB {
	helpers.Log.WithFields(logrus.Fields{
		"service":    "article",
		"function":   "DB connection",
		"started_at": time.Now(),
	}).Info("DB connection establishment -started")

	//Database connection establishment
	dsn := fmt.Sprint("mitrah135:password@tcp(127.0.0.1:3306)/article?parseTime=true")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.Log.Fatal("error:", err)
	}

	helpers.Log.WithFields(logrus.Fields{
		"service": "article",
		"message": "Established a successful connection to database!!!",
	}).Info("Message : 'Established a successful connection to database!!!'")
	return db
}
