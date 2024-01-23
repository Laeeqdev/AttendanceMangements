package databaseconnection

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}
	//x:=9
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Database connection failed.\n")
		os.Exit(100)
	} else {
		log.Printf("connected")
	}
	return db
}
