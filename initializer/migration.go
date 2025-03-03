package initializer

import (
	"fmt"
	"log"

	"github.com/vishalpandhare01/internal/model"
)

func Migration() {
	err := DB.AutoMigrate(
		model.User{},
		model.User_Otp{},
		model.ClassAndStandrd{},
		model.Student{},
	)
	if err != nil {
		log.Fatal("Migration Failed: ", err)
	}

	fmt.Println("All table migrate successfully: ")
}
