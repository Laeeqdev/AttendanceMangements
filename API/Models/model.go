package models

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
	ID        int    `json:"id",pg:",pk"`
	UserId    int    `json:"user_id",pg:"user_id"`
	ClassName string `json:"class",pg:"class_name"`
}

// Student represents the Students table in PostgreSQL.
type Student struct {
	ID     int    `json:"id",pg:",pk"`
	UserId int    `json:"user_id",pg:"user_id"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
	Class  string `json:"class",pg:"class"`
}

// Teacher represents the Teacher table in PostgreSQL.
type Teacher struct {
	ID     int    `json:"id",pg:",pk"`
	UserId int    `json:"user_id",pg:"user_id"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
}

// Attendance represents the Attendance table in PostgreSQL.
type Attendance struct {
	ID       int    `json:"id",pg:",pk"`
	UserId   int    `json:"user_id",pg:"user_id"`
	Date     string `json:"date"`
	PunchIn  string `json:"punch_in",pg:"punch_in"`
	PunchOut string `json:"punch_out",pg:"punch_out"`
	Class    string `json:"class",pg:"class"`
}

// just for data
type Details struct {
	Email string `json:email`
	Date  string `json:date`
	Class string `json:class`
}

// just for data fetching
type AttendanceSchema struct {
	Name   string `json:"name,omitempty"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
}
