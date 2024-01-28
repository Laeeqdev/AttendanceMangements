package service

import (
	"fmt"
	//models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
)

type PunchinService interface {
	Punch_in(email string) (error, bool)
	Punch_out(email string) (error, bool)
}

func Punch_in(email string) (error, bool) {
	err, ok := repository.AlreadyPunch(email)
	if err != nil {
		fmt.Println("yes while cheking")
		return err, false
	}
	if !ok {
		err = repository.PunchInUser(email)
		if err != nil {
			return err, false
		}
		return nil, true
	}
	fmt.Println("already punched_in")
	return nil, false
}

func Punch_out(email string) (error, bool) {
	err, ok := repository.AlreadyPunch(email)
	if err != nil {
		fmt.Println("while punchout")
		return err, false
	}
	if ok {
		err = repository.PunchOutUser(email)
		if err != nil {
			return err, false
		}
		return nil, true
	}
	fmt.Println("punch in before making punch out")
	return nil, false
}
