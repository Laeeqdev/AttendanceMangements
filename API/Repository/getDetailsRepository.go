package repository

import (
	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
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
func GetStudentDetails(user *models.Details) {

}
