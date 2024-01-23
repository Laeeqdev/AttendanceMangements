package resthandler

import (
	"encoding/json"
	"log"
	"net/http"

	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
	//"github.com/go-pg/pg"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	email, err := auth.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok := service.CheckRole(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Users{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err = service.Adduser(user)
	if err != nil {
		log.Println("Error inserting user:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
}

// package resthandler
// func hashPassword(password string) (string, error) {
// 	// Hash the password using bcrypt
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedPassword), nil
// }
