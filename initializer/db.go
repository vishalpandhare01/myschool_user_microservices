package initializer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	var dbUrl = os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(mysql.Open(dbUrl))
	if err != nil {
		log.Fatal("Error in Databse Connection: ", err.Error())
	}
	fmt.Println("Database connected successfully")
}
