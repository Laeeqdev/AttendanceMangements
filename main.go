package main

import (
	"fmt"
	databaseconnection "github.com/Laeeqdev/AttendanceMangements/API/DatabaseConnection"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var version string = "/v1"

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("hi im error")
		log.Fatal(err)
	}
	db := databaseconnection.Connect()
	myrouter := InitializeApp(db)
	r := myrouter.MyRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://attendancemanagementwithlaeeq.netlify.app/"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	defer log.Fatal(http.ListenAndServe(":8089", handler))
	defer databaseconnection.Close()
}
