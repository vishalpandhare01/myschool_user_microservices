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
		model.PastStudent{},
		model.Staff{},
		model.FeeType{},
		model.StudentFees{},
		model.FeesStructure{},
		model.TimeTable{},
		model.Attendance{},
	)
	if err != nil {
		log.Fatal("Migration Failed: ", err)
	}

	fmt.Println("All table migrate successfully: ")
}
