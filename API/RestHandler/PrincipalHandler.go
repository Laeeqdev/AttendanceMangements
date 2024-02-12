package resthandler

import (
	"encoding/json"
	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
	"log"
	"net/http"
)

type PrincipalHandler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
}
type PrincipalHandlerImpl struct {
	userService     service.UserService
	userAuthHandler auth.UserAuthHandler
}

func NewPrincipalHandlerImpl(userService service.UserService, userAuthHandler auth.UserAuthHandler) *PrincipalHandlerImpl {
	return &PrincipalHandlerImpl{userService: userService, userAuthHandler: userAuthHandler}
}
func (impl PrincipalHandlerImpl) AddUser(w http.ResponseWriter, r *http.Request) {
	email, err := impl.userAuthHandler.GetMailFromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, ok := impl.userService.IsPrincipal(email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := &models.Users{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err = impl.userService.Adduser(user)
	if err != nil {
		log.Println("Error inserting user:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
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
