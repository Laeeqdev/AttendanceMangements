package repository

import (
	"time"

	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	"github.com/go-pg/pg"
)

type PunchinRepository interface {
	PunchInUser(email string) error
	AlreadyPunch(email string) (error, bool)
	PunchOutUser(email string) error
}
type PunchinRepositoryImpl struct {
	DbConnection *pg.DB
}

func NewPunchinRepositoryImpl(db *pg.DB) *PunchinRepositoryImpl {
	return &PunchinRepositoryImpl{DbConnection: db}
}

// punch in code starts from here
func (impl *PunchinRepositoryImpl) PunchInUser(email string) error {
	var currentTime = time.Now()
	var Date string = currentTime.Format("2006-01-02") // Format for date: Year-Month-Day
	var Time string = currentTime.Format("15:04:05")
	dbMutex.Lock()
	defer dbMutex.Unlock()
	err, user_Id, role := GetRoleandUserId(email, impl.DbConnection)
	if err != nil {
		return err
	}
	var class string
	class = ""
	//hour:minutes:second
	if role == constants.STUDENT {
		err, class := GetClassByUserId(user_Id, impl.DbConnection)
		if err != nil {
			return err
		}
		var yes bool
		yes = false
		err = impl.DbConnection.Model(&models.Student{}).
			Column("status").
			Where("user_id = ?", user_Id).
			Where("date = ?", Date).
			Select(&yes)
		if err == nil && !yes || err == pg.ErrNoRows {
			studentmodel := &models.Student{
				UserId: user_Id,
				Date:   Date,
				Status: true,
				Class:  class,
			}
			_, er := impl.DbConnection.Model(studentmodel).Insert()
			if er != nil {
				return er
			}
		}
	} else if role == constants.TEACHER {
		var yes bool
		yes = false
		err = impl.DbConnection.Model(&models.Teacher{}).
			Column("status").
			Where("user_id = ?", user_Id).
			Where("date = ?", Date).
			Select(&yes)
		if err == nil && !yes || err == pg.ErrNoRows {
			teachermodel := &models.Teacher{
				UserId: user_Id,
				Date:   Date,
				Status: true,
			}
			_, er := impl.DbConnection.Model(teachermodel).Insert()
			if er != nil {
				return er
			}
		}
	}
	attendancemodel := &models.Attendance{
		UserId:  user_Id,
		Date:    Date,
		PunchIn: Time,
		Class:   class,
	}
	_, er := impl.DbConnection.Model(attendancemodel).Insert()
	if er != nil {

		return er
	}
	return nil
}
func (impl *PunchinRepositoryImpl) AlreadyPunch(email string) (error, bool) {
	var currentTime = time.Now()
	var Date string = currentTime.Format("2006-01-02") // Format for date: Year-Month-Day
	DbConnection := databaseconnection.Connect()
	err, userId := GetUserId(email, DbConnection)
	if err != nil {
		return err, false
	}
	var ii int
	ii = -1
	err = DbConnection.Model(&models.Attendance{}).
		Column("id").
		Where("user_id = ?", userId).
		Where("date = ?", Date).
		Where("punch_in IS NOT NULL AND punch_out IS NULL").
		Select(&ii)
	if err == pg.ErrNoRows {
		return nil, false
	}
	if err != nil {
		return err, true
	}
	if ii != 0 {
		return nil, true
	}
	return nil, false
}

// punch out code starts from here
func (impl *PunchinRepositoryImpl) PunchOutUser(email string) error {
	var currentTime = time.Now()
	var Date string = currentTime.Format("2006-01-02") // Format for date: Year-Month-Day
	var Time string = currentTime.Format("15:04:05")
	var DbConnection = databaseconnection.Connect()
	dbMutex.Lock()
	defer dbMutex.Unlock()
	err, user_Id := GetUserId(email, DbConnection)
	if err != nil {
		return err
	}
	_, er := DbConnection.Model(&models.Attendance{}).
		Where("user_id = ? ", user_Id).
		Where(" date = ? ", Date).
		Where("punch_in IS NOT NULL AND punch_out IS NULL").
		Set("punch_out = ? ", Time).Update()
	if er != nil {
		return er
	}
	return nil
}
func GetRoleandUserId(email string, DbConnection *pg.DB) (error, int, string) {
	var role string
	var userId int
	role = ""

	err := DbConnection.Model(&models.Users{}).
		Column("id", "role").
		Where("email = ?", email).
		Select(&userId, &role)
	if err != nil {
		return err, 0, role
	}
	return nil, userId, role
}
func GetClassByUserId(userId int, DbConnection *pg.DB) (error, string) {
	var class string
	err := DbConnection.Model(&models.Studclass{}).
		Column("class_name").
		Where("user_id = ?", userId).
		Select(&class)
	if err != nil {
		return err, class
	}
	return nil, class
}
