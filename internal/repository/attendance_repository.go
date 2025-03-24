package repository

import (
	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

// create attendance by teacher staff
func CreateAttendanceRepository(body *model.Attendance) (*model.Attendance, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

// get attendance by date + school id , subject , teacher id ,student id , classid
func GetAttendanceRepository(
	date string,
	classId string,
	schoolId string,
	subject string,
	teacherId string,
	studentId string) (*[]model.Attendance, error) {
	var attendance *[]model.Attendance
	query := initializer.DB.Where("date = ? AND school_id", date, schoolId).Find(&attendance)
	if subject != "" {
		query = query.Where("subject = ?", subject)
	}
	if teacherId != "" {
		query = query.Where("teacher_id = ?", teacherId)
	}
	if studentId != "" {
		query = query.Where("student_id = ?", studentId)
	}

	if err := query.Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

// update attendance
func UpdateAttendanceRepository(body *model.Attendance) (*model.Attendance, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}
