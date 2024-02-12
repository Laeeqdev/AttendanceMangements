package service

import (
	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
)

type GetDeatilsService interface {
	GetDetailsOfATeacher(user *models.Details) (error, []models.AttendanceSchema)
	IsPermissibleForTeacherAndPrincipal(email string) (error, bool, bool)
	GetDetailsOfAStudent(user *models.Details) (error, []models.AttendanceSchema)
	IsPermissibleForTeacherAndStudent(email string) (error, bool, bool)
	IsPermissibleForTeacher(email string) (error, bool)
	GetDetailsOfAStudentByClass(user *models.Details) (error, []models.AttendanceSchema)
	GetDeatilsOfPunch(user *models.PunchInPunchOutDetails) (error, []models.PunchInPunchOutDetails)
}

type GetDeatilsServiceImpl struct {
	detailsRepository repository.DetailsRepository
	userRepository    repository.UserRepository
}

func NewGetDeatilsServiceImpl(detailsRepository repository.DetailsRepository, userRepository repository.UserRepository) *GetDeatilsServiceImpl {
	return &GetDeatilsServiceImpl{detailsRepository: detailsRepository, userRepository: userRepository}
}
func (impl *GetDeatilsServiceImpl) GetDetailsOfATeacher(user *models.Details) (error, []models.AttendanceSchema) {
	err, details := impl.detailsRepository.GetUserDatails(user, constants.TEACHER)
	if err != nil {
		return err, nil
	}
	return nil, details
}

func (impl *GetDeatilsServiceImpl) IsPermissibleForTeacherAndPrincipal(email string) (error, bool, bool) {
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
	return err, false, false
}

func (impl *GetDeatilsServiceImpl) GetDetailsOfAStudent(user *models.Details) (error, []models.AttendanceSchema) {
	err, details := impl.detailsRepository.GetUserDatails(user, constants.STUDENT)
	if err != nil {
		return err, nil
	}
	return nil, details
}
func (impl *GetDeatilsServiceImpl) IsPermissibleForTeacherAndStudent(email string) (error, bool, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false, false
	}
	if role == constants.TEACHER {
		return nil, true, false
	}
	if role == constants.STUDENT {
		return nil, true, true
	}
	return nil, false, false
}
func (impl *GetDeatilsServiceImpl) IsPermissibleForTeacher(email string) (error, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false
	}
	if role == constants.TEACHER {
		return nil, true
	}
	return nil, false
}
func (impl *GetDeatilsServiceImpl) GetDetailsOfAStudentByClass(user *models.Details) (error, []models.AttendanceSchema) {
	err, details := impl.detailsRepository.GetStudentDetailsByclass(user)
	if err != nil {
		return err, nil
	}
	return nil, details
}
func (impl *GetDeatilsServiceImpl) GetDeatilsOfPunch(user *models.PunchInPunchOutDetails) (error, []models.PunchInPunchOutDetails) {
	err, details := impl.detailsRepository.GetDetails(user)
	if err != nil {
		return err, nil
	}
	return nil, details
}
