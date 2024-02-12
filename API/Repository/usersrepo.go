package repository

import (
	"fmt"
	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	"github.com/go-pg/pg"
	"sync"
)

type UserRepository interface {
	Adduser(user *models.Users) error
	Findpassword(email string) (error, string)
}
type UserRepositoryImpl struct {
	DbConnection *pg.DB
}

func NewUserRepositoryImpl(db *pg.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DbConnection: db}
}

var dbMutex sync.Mutex

func (impl *UserRepositoryImpl) Adduser(user *models.Users) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	_, err := impl.DbConnection.Model(user).Insert()

	if err != nil {
		return err
	}
	if user.Role == constants.STUDENT {
		err, userId := GetUserId(user.Email, impl.DbConnection)
		if err != nil {
			return err
		}
		studclass := &models.Studclass{
			UserId:    userId,
			ClassName: user.Class,
		}
		_, er := impl.DbConnection.Model(studclass).Insert()
		if er != nil {
			return er
		}
	}
	return nil
}
func (impl *UserRepositoryImpl) Findpassword(email string) (error, string) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var pswd string
	err := impl.DbConnection.Model(&models.Users{}).Column("password").Where("email = ?", email).Select(&pswd)
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
func GetNameAndRole(email string) (error, []string) {
	var DbConnection = databaseconnection.Connect()
	var name, role string
	err := DbConnection.Model(&models.Users{}).
		Column("name", "role").
		Where("email = ?", email).Select(&name, &role)
	if err != nil {
		return err, nil
	}
	fmt.Println(name, role)
	return nil, []string{name, role}
}
