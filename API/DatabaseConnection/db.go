package databaseconnection

import (
	"github.com/go-pg/pg"
	"log"
	"os"
	"sync"
)

var (
	db   *pg.DB
	once sync.Once
)

// Connect initializes and returns a database connection.
func Connect() *pg.DB {
	once.Do(func() {
		db = initializeConnection()
	})
	return db
}

// initializeConnection sets up the database connection.
func initializeConnection() *pg.DB {
	// Your database connection setup code here
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
	//var dbb *pg.DB= *pg.DB;
	return db
	//return nil
}
func Close() {
	if db != nil {
		db.Close()
	}
}
