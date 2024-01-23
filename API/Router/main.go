package main

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	resthandler "github.com/Laeeqdev/AttendanceMangements/API/RestHandler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("hi im error")
		log.Fatal(err)
	}
	// Db := dbb.Connect()
	r := mux.NewRouter()
	//fmt.Println(Db)
	//login
	r.HandleFunc("/Login", auth.Login).Methods("POST")
	r.HandleFunc("/home", auth.Home).Methods("GET")
	r.HandleFunc("/adduser", resthandler.AddUser).Methods("POST")
	r.HandleFunc("/logout", resthandler.Logout).Methods("POST")
	// r.HandleFunc("/",Login).Methods("POST")
	// r.HandleFunc("/Login",Login).Methods("POST")
	defer log.Fatal(http.ListenAndServe(":8081", r))

}

// {
// 	"name":"rajeev ranjan",
// 	"email":"Rajeev123@gmail.com",
// 	"password":"Rajeev#012",
// 	"role":"Student",
// 	"class":"4"
//   }
