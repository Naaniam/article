package drivers

import (

	// built-in package
	"article/logs"
	"fmt"
	"os"

	// third party package

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SQLDriver() *gorm.DB {
	var err error
	log := logs.Log()

	err = godotenv.Load(".env")
	if err != nil {
		log.Error.Println("Error : 'Error at loading '.env' file'")
		return nil
	}

	//Database connection establishment
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error.Println("Error : 'Invalid Database connection' ", err)
		panic(err)
	}

	log.Info.Printf("Message : 'Established a successful connection to the database!!!'\n")
	return db
}

func TestSQLDriver() *gorm.DB {
	log := logs.Log()

	//Database connection establishment
	dsn := fmt.Sprint("root:password@tcp(localhost:3306)/article?parseTime=true")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Info.Println("Message : 'Established a successful connection to database!!!'")
	return db
}
