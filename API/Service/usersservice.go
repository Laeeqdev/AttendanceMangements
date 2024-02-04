package service

import (
	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
)

type ImageTaggingRepository interface {
	Adduser(user *models.Users) error
	MatchPassword(password string) (error, bool)
}

func Adduser(user *models.Users) error {
	err := repository.Adduser(user)
	return err
}
func MatchPassword(email string, expectedpassword string) (error, bool) {
	err, password := repository.Findpassword(email)
	if err != nil {
		return err, false
	}
	if password == expectedpassword {
		return nil, true
	}
	return nil, false
}
func IsPrincipal(email string) (error, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false
	}
	if role == constants.PRINCIPAL {
		return nil, true
	}
	return nil, false
}
func GetDataForHome(email string) (error, []string) {
	err, data := repository.GetNameAndRole(email)
	if err != nil {
		return err, nil
	}
	return nil, data
}
