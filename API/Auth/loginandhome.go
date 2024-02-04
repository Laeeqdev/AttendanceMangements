package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/Laeeqdev/AttendanceMangements/API/Models"
	//repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("Laeeq_Ahmad")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// login
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Users
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("error in data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err, yes := service.MatchPassword(credentials.Email, credentials.Password)
	if !yes {
		fmt.Println("password or email not found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 5)
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
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: false,
			Secure:   false,
			Domain:   "",
			Path:     "/v1",
		})
	err, role := service.GetDataForHome(credentials.Email)
	if err != nil {
		fmt.Println("error while fetching role and name")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(role)
}

// logout
func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logout successful")
}

// home
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
	err, data := service.GetDataForHome(claims.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: false,
			Secure:   false,
			Domain:   "",
			Path:     "/v1",
		})

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
