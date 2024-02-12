package service

import (
	constants "github.com/Laeeqdev/AttendanceMangements/API/Constant"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
)

type UserService interface {
	Adduser(user *models.Users) error
	MatchPassword(email string, expectedpassword string) (error, bool)
	IsPrincipal(email string) (error, bool)
	GetDataForHome(email string) (error, []string)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
func (impl *UserServiceImpl) Adduser(user *models.Users) error {
	err := impl.userRepository.Adduser(user)
	return err
}
func (impl *UserServiceImpl) MatchPassword(email string, expectedpassword string) (error, bool) {
	err, password := impl.userRepository.Findpassword(email)
	if err != nil {
		return err, false
	}
	if password == expectedpassword {
		return nil, true
	}
	return nil, false
}
func (impl *UserServiceImpl) IsPrincipal(email string) (error, bool) {
	err, role := repository.GetRole(email)
	if err != nil {
		return err, false
	}
	if role == constants.PRINCIPAL {
		return nil, true
	}
	return nil, false
}
func (impl *UserServiceImpl) GetDataForHome(email string) (error, []string) {
	err, data := repository.GetNameAndRole(email)
	if err != nil {
		return err, nil
	}
	return nil, data
}
