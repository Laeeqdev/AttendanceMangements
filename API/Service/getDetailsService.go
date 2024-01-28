package service

import (
	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
)

type GetDeatilsService interface {
	GetDetailsOfATeacher(email string, user *models.Details) (error, []*models.AttendanceSchema)
}

func GetDetailsOfATeacher(user *models.Details) (error, []models.AttendanceSchema) {
	err, details := repository.GetUserDatails(user, constants.TEACHER)
	if err != nil {
		return err, nil
	}
	return nil, details
}

func IsPermissibleForTeacherAndPrincipal(email string) (error, bool, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false, false
	}
	if role == constants.PRINCIPAL { //todo move all names to constant file
		return nil, true, false
	}
	if role == constants.TEACHER { //todo move all names to constant file
		return nil, true, true
	}
	return nil, false, false
}

func GetDetailsOfAStudent(user *models.Details) (error, []models.AttendanceSchema) {
	err, details := repository.GetUserDatails(user, constants.STUDENT)
	if err != nil {
		return err, nil
	}
	return nil, details
}
func IsPermissibleForTeacherAndStudent(email string) (error, bool, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false, false
	}
	if role == constants.TEACHER {
		return nil, true, false
	}
	if role == constants.STUDENT {
	}
	return nil, false, false
}
func IsPermissibleForTeacher(email string) (error, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false
	}
	if role == constants.TEACHER {
		return nil, true
	}
	return nil, false
}
func GetDetailsOfAStudentByClass(user *models.Details) (error, []models.AttendanceSchema) {
	err, details := repository.GetUserDatails(user, "")
	if err != nil {
		return err, nil
	}
	return nil, details
}
