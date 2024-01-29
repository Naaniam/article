package drivers

import (

	// built-in package
	"article/logs"
	"fmt"
	"os"

	// third party package

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SQLDriver() *gorm.DB {
	var err error

	logrusEntry := logrus.WithFields(logrus.Fields{
		"function": "DB connection",
	})

	err = godotenv.Load(".env")
	if err != nil {
		logrusEntry.Errorf("Error : 'Error at loading '.env' file'")
		return nil
	}

	//Database connection establishment
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrusEntry.Error("Error : 'Invalid Database connection' ", err)
		panic(err)
	}

	logrusEntry.Info("Message : 'Established a successful connection to the database!!!'\n")
	return db
}

func TestSQLDriver() *gorm.DB {
	log := logs.Log()

	//Database connection establishment
	dsn := fmt.Sprint("mitrah135:password@tcp(127.0.0.1:3306)/article?parseTime=true")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	log.Info.Println("Message : 'Established a successful connection to database!!!'")
	return db
}
