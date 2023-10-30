package model

import (
	"time"

	"github.com/guregu/null"
)

type StudentCourse struct {
	Id         string      `db:"id"`
	LecturerId string      `db:"lecturer_id"`
	Name       string      `db:"name"`
	Credit     int         `db:"credit"`
	CreatedAt  time.Time   `db:"created_at"`
	CreatedBy  string      `db:"created_by"`
	UpdatedAt  time.Time   `db:"updated_at"`
	UpdatedBy  string      `db:"updated_by"`
	DeletedAt  null.Time   `db:"deleted_at"`
	DeletedBy  null.String `db:"deleted_by"`
}

type StudentCourseList []*StudentCourse

type StudentCoursePrimaryID struct {
	Id string `db:"id"`
}

type StudentCoursePrimaryIDList []*StudentCoursePrimaryID
