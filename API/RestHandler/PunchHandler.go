package resthandler

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"

	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	//models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
)

func PunchInUser(w http.ResponseWriter, r *http.Request) {
	email, err := auth.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err, ok := service.Punch_in(email)
	if err != nil {
		log.Println("Error inserting attendance:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err == nil && ok == false {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Fprint(w, "already punched in")
		return
	}
	fmt.Fprint(w, "punch_in successfully")
}
func PunchOutUser(w http.ResponseWriter, r *http.Request) {
	email, err := auth.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err, ok := service.Punch_out(email)
	if err != nil {
		log.Println("Error inserting attendance:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err == nil && ok == false {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Fprint(w, "punch_in before making punchout")
		return
	}
	fmt.Fprint(w, "punch_out successfully")
}
