package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("Laeeq_Ahmad")

// var users=map[string]string{"dsnc":"jsdn"}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Users
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, yes := service.MatchPassword(credentials.Email, credentials.Password)
	if !yes || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 12)
	claims := &Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logout successful")
}
func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenstr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello,%s", claims.Email)))
}
func Refresh(w http.ResponseWriter, r http.Request) {

}
func GetMailFromCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}
	tokenstr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusInternalServerError)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		return "", fmt.Errorf("user does not have access")
	}
	return claims.Email, nil
}

//func AddUser(){}
