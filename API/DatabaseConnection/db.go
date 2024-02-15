package databaseconnection

import (
	"log"
	"os"
	"sync"

	"github.com/go-pg/pg"
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
	Dbaddress, _ := pg.ParseURL(os.Getenv("DB_ADDR1"))
	opts := &pg.Options{}
	if os.Getenv("Type") != "local" {
		opts = &pg.Options{
			User:     os.Getenv("DB_USER1"),
			Password: os.Getenv("DB_PASSWORD1"),
			Addr:     Dbaddress.Addr,
			Database: os.Getenv("DB_DATABASE1"),
		}
	} else {
		opts = &pg.Options{
			User:     os.Getenv("DB_USER2"),
			Password: os.Getenv("DB_PASSWORD2"),
			Addr:     os.Getenv("DB_ADDR2"),
			Database: os.Getenv("DB_DATABASE2"),
		}
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
