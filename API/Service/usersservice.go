package service

import (
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
func CheckRole(email string) (error, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false
	}
	if role == "Principal" { //todo move all names to constant file
		return nil, true
	}
	return nil, false
}
