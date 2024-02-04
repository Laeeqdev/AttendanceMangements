package main

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	resthandler "github.com/Laeeqdev/AttendanceMangements/API/RestHandler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var version string = "/v1"

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("hi im error")
		log.Fatal(err)
	}
	//Db = databaseconnection.Connect()
	r1 := mux.NewRouter()

	//fmt.Println(Db)
	//login
	r := r1.PathPrefix(version).Subrouter()
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/home", auth.Home).Methods("GET")
	r.HandleFunc("/adduser", resthandler.AddUser).Methods("POST")
	r.HandleFunc("/logout", resthandler.Logout).Methods("POST")
	r.HandleFunc("/punchin", resthandler.PunchInUser).Methods("POST")
	r.HandleFunc("/punchout", resthandler.PunchOutUser).Methods("POST")
	r.HandleFunc("/getteacherdetails", resthandler.GetTeacherDetails).Methods("POST")
	r.HandleFunc("/refresh", auth.Refresh).Methods("GET")
	r.HandleFunc("/getstudentdetails", resthandler.GetStudentDetails).Methods("POST")
	r.HandleFunc("/getstudentdetailsbyclass", resthandler.GetStudentDetailsByClass).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	defer log.Fatal(http.ListenAndServe(":8081", handler))
	defer databaseconnection.Close()
}

// access control and  CORS middleware
// func GetConnection() *pg.DB {
// 	return Db
// }

// {
// 	"name":"rajeev ranjan",
// 	"email":"Rajeev123@gmail.com",
// 	"password":"Rajeev#012",
// 	"role":"Student",
// 	"class":"4"
//   }
