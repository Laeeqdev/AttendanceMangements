package resthandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
)

func GetTeacherDetails(w http.ResponseWriter, r *http.Request) {
	email, err := auth.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok, yes := service.IsPermissibleForTeacherAndPrincipal(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Details{}
	if yes {
		user = &models.Details{Email: email}
	}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err, data := service.GetDetailsOfATeacher(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func GetStudentDetails(w http.ResponseWriter, r *http.Request) {
	email, err := auth.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok, yes := service.IsPermissibleForTeacherAndStudent(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Details{}
	if yes {
		user = &models.Details{Email: email}
	}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err, data := service.GetDetailsOfAStudent(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func GetStudentDetailsByClass(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data like date class")
		return
	}
	email, err := auth.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok := service.IsPermissibleForTeacher(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Details{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err, data := service.GetDetailsOfAStudentByClass(user)
	if err != nil {
		log.Println("Error getting details of user:", err)
		fmt.Fprint(w, "getting data failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}
