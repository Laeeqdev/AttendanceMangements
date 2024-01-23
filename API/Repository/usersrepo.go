package repository

import (
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	//"github.com/go-pg/pg"
)

type ImageTaggingRepository interface {
	AddTeacher(user *models.Users) error
}

// type DbReference struct {
// 	DbCon *pg.DB
// }

//	func GetDBReference() *DbReference {
//		return &DbReference{
//			DbCon: DbConnection,
//		}
//	}
func Adduser(user *models.Users) error {
	var DbConnection = databaseconnection.Connect()
	_, err := DbConnection.Model(user).Insert()
	defer DbConnection.Close()
	if err != nil {
		return err
	}
	if user.Role == "Student" {
		var userId int
		err := DbConnection.Model(&models.Users{}).Column("id").Where("email = ?", user.Email).Select(&userId)
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
	var DbConnection = databaseconnection.Connect() //todo please define this globally
	var pswd string
	err := DbConnection.Model(&models.Users{}).Column("password").Where("email = ?", email).Select(&pswd)
	if err != nil {
		return err, ""
	}
	return nil, pswd
}
func GetRole(email string) (error, string) {
	var DbConnection = databaseconnection.Connect()
	var role string
	err := DbConnection.Model(&models.Users{}).Column("role").Where("email = ?", email).Select(&role)
	if err != nil {
		return err, ""
	}
	return nil, role
}
