package resthandler

import (
	"encoding/json"
	"fmt"
	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
	"log"
	"net/http"
)

type DetailsHandler interface {
	GetTeacherDetails(w http.ResponseWriter, r *http.Request)
	GetDetails(w http.ResponseWriter, r *http.Request)
	GetStudentDetails(w http.ResponseWriter, r *http.Request)
	GetStudentDetailsByClass(w http.ResponseWriter, r *http.Request)
}
type DetailsHandlerImpl struct {
	getDeatilsService service.GetDeatilsService
	userAuthHandler   auth.UserAuthHandler
}

func NewDetailsHandlerImpl(getDeatilsService service.GetDeatilsService, userAuthHandler auth.UserAuthHandler) *DetailsHandlerImpl {
	return &DetailsHandlerImpl{getDeatilsService: getDeatilsService, userAuthHandler: userAuthHandler}
}
func (impl *DetailsHandlerImpl) GetTeacherDetails(w http.ResponseWriter, r *http.Request) {
	email, err := impl.userAuthHandler.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok, yes := impl.getDeatilsService.IsPermissibleForTeacherAndPrincipal(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Details{}

	_ = json.NewDecoder(r.Body).Decode(&user)
	if yes {
		fmt.Println("hi im yes", email)
		user.Email = email
	}
	err, data := impl.getDeatilsService.GetDetailsOfATeacher(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func (impl *DetailsHandlerImpl) GetDetails(w http.ResponseWriter, r *http.Request) {
	email, err := impl.userAuthHandler.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.PunchInPunchOutDetails{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err, data := impl.getDeatilsService.GetDeatilsOfPunch(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func (impl *DetailsHandlerImpl) GetStudentDetails(w http.ResponseWriter, r *http.Request) {
	email, err := impl.userAuthHandler.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok, yes := impl.getDeatilsService.IsPermissibleForTeacherAndStudent(email)
	if err != nil || !ok {
		fmt.Println("hey I am inside ok")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Details{}

	_ = json.NewDecoder(r.Body).Decode(&user)
	if yes {
		fmt.Println("hey I am inside yes")
		user.Email = email
	}
	err, data := impl.getDeatilsService.GetDetailsOfAStudent(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func (impl *DetailsHandlerImpl) GetStudentDetailsByClass(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data like date class")
		return
	}
	email, err := impl.userAuthHandler.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok := impl.getDeatilsService.IsPermissibleForTeacher(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Details{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err, data := impl.getDeatilsService.GetDetailsOfAStudentByClass(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
