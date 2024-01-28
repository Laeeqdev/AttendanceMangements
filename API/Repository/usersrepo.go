package repository

import (
	"sync"

	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	"github.com/go-pg/pg"
)

type ImageTaggingRepository interface {
	AddTeacher(user *models.Users) error
}

var dbMutex sync.Mutex

func Adduser(user *models.Users) error {
	var DbConnection = databaseconnection.Connect()
	dbMutex.Lock()
	defer dbMutex.Unlock()
	_, err := DbConnection.Model(user).Insert()

	if err != nil {
		return err
	}
	if user.Role == constants.PRINCIPAL {
		err, userId := GetUserId(user.Email, DbConnection)
		if err != nil {
			return err
		}
		studclass := &models.Studclass{
			UserId:    userId,
			ClassName: user.Class,
		}
		_, er := DbConnection.Model(studclass).Insert()
		if er != nil {
			return er
		}
	}
	return nil
}
func Findpassword(email string) (error, string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var DbConnection = databaseconnection.Connect()
	var pswd string
	err := DbConnection.Model(&models.Users{}).Column("password").Where("email = ?", email).Select(&pswd)
	if err != nil {
		return err, ""
	}
	return nil, pswd
}
func GetRole(email string) (error, string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var DbConnection = databaseconnection.Connect()
	var role string
	err := DbConnection.Model(&models.Users{}).Column("role").Where("email = ?", email).Select(&role)
	if err != nil {
		return err, ""
	}
	return nil, role
}
func GetUserId(email string, DbConnection *pg.DB) (error, int) {
	var userId int
	err := DbConnection.Model(&models.Users{}).Column("id").Where("email = ?", email).Select(&userId)
	if err != nil {
		return err, 0
	}
	return nil, userId
}
