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
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Database connection failed.\n")
		os.Exit(100)
	} else {
		log.Printf("connected")
	}

	return db
}
func Close() {
	if db != nil {
		db.Close()
	}
}
