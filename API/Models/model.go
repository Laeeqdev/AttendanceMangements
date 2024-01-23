package models

import (
	"time"
)

type Users struct {
	ID       int    `json:"id",pg:",pk"` // Annotation for primary key
	Name     string `json:"name"`
	Email    string `json:"email",sql:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Class    string `json:"class" sql:"-"`
}

// Class represents the Class table in PostgreSQL.
type Studclass struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id",pg:"user_id"`
	ClassName string `json:"class",pg:"class_name"`
}

// Student represents the Students table in PostgreSQL.
type Student struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Date      time.Time `json:"date"`
	Status    bool      `json:"status"`
	ClassName string    `json:"class"`
}

// Teacher represents the Teacher table in PostgreSQL.
type Teacher struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Status bool      `json:"status"`
}

// Attendance represents the Attendance table in PostgreSQL.
type Attendance struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Date      time.Time `json:"date"`
	PunchIn   time.Time `json:"punch_in"`
	PunchOut  time.Time `json:"punch_out"`
	ClassName string    `json:"class"`
}
