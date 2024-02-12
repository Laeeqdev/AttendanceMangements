package repository

import (
	"fmt"
	"time"

	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	"github.com/go-pg/pg"
	//"github.com/go-pg/pg"
)

type DetailsRepository interface {
	GetUserDatails(user *models.Details, role string) (error, []models.AttendanceSchema)
	GetStudentDetailsByclass(user *models.Details) (error, []models.AttendanceSchema)
	GetDetails(user *models.PunchInPunchOutDetails) (error, []models.PunchInPunchOutDetails)
}
type DetailsRepositoryImpl struct {
	DbConnection *pg.DB
}

func NewDetailsRepositoryImpl(db *pg.DB) *DetailsRepositoryImpl {
	return &DetailsRepositoryImpl{
		DbConnection: db,
	}
}

// punch in code starts from here
func (impl *DetailsRepositoryImpl) GetUserDatails(user *models.Details, role string) (error, []models.AttendanceSchema) {

	dbMutex.Lock()
	defer dbMutex.Unlock()
	err, user_Id := GetUserId(user.Email, impl.DbConnection)
	fmt.Println("my user id ", user_Id)
	if err != nil {
		fmt.Println("hey while getting userId", user.Date)
		return err, nil
	}
	var details []models.AttendanceSchema
	date := user.Date + "%"
	if role == constants.STUDENT {
		fmt.Println("hello I am inside the student")
		var studentmodel []models.Student
		err := impl.DbConnection.Model(&studentmodel).
			Where("user_id = ?", user_Id).
			Where("date ILIKE ? ", date).
			Select()
		if err != nil {
			return err, nil
		}
		for _, val := range studentmodel {
			d := models.AttendanceSchema{
				UserId: user_Id,
				Date:   val.Date,
				Status: val.Status,
			}
			details = append(details, d)
		}

	} else if role == constants.TEACHER {
		fmt.Print("hey I ma working", user.Date)
		var teachermodel []models.Teacher
		err := impl.DbConnection.Model(&teachermodel).
			Where("user_id = ?", user_Id).
			Where("date ILIKE ? ", date).
			Select()
		if err != nil {
			fmt.Println("hello")
			return err, nil
		}
		for _, val := range teachermodel {
			if err != nil {
				return err, nil
			}
			d := models.AttendanceSchema{
				UserId: user_Id,
				Date:   val.Date,
				Status: val.Status,
			}
			details = append(details, d)
		}
	}
	return nil, details
}

func (impl *DetailsRepositoryImpl) GetStudentDetailsByclass(user *models.Details) (error, []models.AttendanceSchema) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var details []models.AttendanceSchema
	var studentmodel []models.Student
	err := impl.DbConnection.Model(&studentmodel).
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
		err, name := GetUserName(val.UserId, impl.DbConnection)
		if err != nil {
			fmt.Println("line 3")
			return err, nil
		}
		d := models.AttendanceSchema{
			UserId: val.UserId,
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
func (impl *DetailsRepositoryImpl) GetDetails(user *models.PunchInPunchOutDetails) (error, []models.PunchInPunchOutDetails) {

	dbMutex.Lock()
	defer dbMutex.Unlock()
	var details []models.PunchInPunchOutDetails
	var attendancemodel []models.Attendance
	err := impl.DbConnection.Model(&attendancemodel).
		Where("user_id = ? ", user.UserId).
		Where("date = ? ", user.Date).
		Select()
	if err != nil {
		fmt.Println("line 2")
		return err, nil
	}

	if err != nil {
		fmt.Println("line 3")
		return err, nil
	}
	for _, val := range attendancemodel {
		if err != nil {
			fmt.Println("line 3")
			return err, nil
		}
		if val.PunchOut == "" {
			continue
		}
		err, duration := calculateDiffrence(val.PunchIn, val.PunchOut)
		if err != nil {
			return err, nil
		}
		d := models.PunchInPunchOutDetails{
			PunchIn:  val.PunchIn,
			PunchOut: val.PunchOut,
			Duartion: duration,
		}
		details = append(details, d)
	}

	return nil, details
}
func calculateDiffrence(in string, out string) (error, string) {

	punchInTime, err := time.Parse("15:04:05", in)
	if err != nil {
		fmt.Println("Error parsing punch in time:", err)
		return err, ""
	}
	punchOutTime, err := time.Parse("15:04:05", out)
	if err != nil {
		fmt.Println("Error parsing punch out time:", err)
		return err, ""
	}
	// Calculate duration
	duration := punchOutTime.Sub(punchInTime)
	return nil, duration.String()
}
