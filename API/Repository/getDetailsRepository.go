package repository

import (
	"fmt"

	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	"github.com/go-pg/pg"
	//"github.com/go-pg/pg"
)

type DetailsRepository interface {
	GetUserDatails(user *models.Details, role string) (error, []*models.AttendanceSchema)
	AlreadyPunchIn(user *models.Users) (error, string, string)
}

// punch in code starts from here
func GetUserDatails(user *models.Details, role string) (error, []models.AttendanceSchema) {
	var DbConnection = databaseconnection.Connect()
	dbMutex.Lock()
	defer dbMutex.Unlock()
	err, user_Id := GetUserId(user.Email, DbConnection)
	if err != nil {
		return err, nil
	}
	var details []models.AttendanceSchema
	date := user.Date + "%"
	if role == constants.STUDENT {
		var studentmodel []models.Student
		err := DbConnection.Model(&studentmodel).
			Where("user_id = ?", user_Id).
			Where("date ILIKE ? ", date).
			Select()
		if err != nil {
			return err, nil
		}
		for _, val := range studentmodel {
			d := models.AttendanceSchema{
				Date:   val.Date,
				Status: val.Status,
			}
			details = append(details, d)
		}
	} else if role == constants.TEACHER {
		var teachermodel []models.Teacher
		err := DbConnection.Model(&teachermodel).
			Where("user_id = ?", user_Id).
			Where("date ILIKE ? ", date).
			Select()
		if err != nil {
			return err, nil
		}
		for _, val := range teachermodel {
			d := models.AttendanceSchema{
				Date:   val.Date,
				Status: val.Status,
			}
			details = append(details, d)
		}
	}
	return nil, details
}
func GetStudentDetails(user *models.Details) (error, []models.AttendanceSchema) {
	var DbConnection = databaseconnection.Connect()
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var details []models.AttendanceSchema
	var studentmodel []models.Student
	err := DbConnection.Model(&studentmodel).
		Where("date = ? ", user.Date).
		Where("class = ? ", user.Class).
		Select()
	if err != nil {
		fmt.Println("line 2")
		return err, nil
	}

	if err != nil {
		fmt.Println("line 3")
		return err, nil
	}
	for _, val := range studentmodel {
		err, name := GetUserName(val.UserId, DbConnection)
		if err != nil {
			fmt.Println("line 3")
			return err, nil
		}
		d := models.AttendanceSchema{
			Name:   name,
			Date:   val.Date,
			Status: val.Status,
		}
		details = append(details, d)
	}
	return nil, details
}
func GetUserName(userId int, db *pg.DB) (error, string) {
	var name string
	err := db.Model(&models.Users{}).Column("name").Where(" id = ? ", userId).
		Select(&name)
	if err != nil {
		return err, name
	}
	return err, name
}
