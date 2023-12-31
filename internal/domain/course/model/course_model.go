package model

import (
	"time"

	"github.com/guregu/null"
)

type Course struct {
	Id         string      `db:"id"`
	LecturerId string      `db:"lecturer_id"`
	Name       string      `db:"name"`
	Credit     int         `db:"credit"`
	CreatedAt  time.Time   `json:"createdAt"`
	CreatedBy  string      `json:"createdBy"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	UpdatedBy  string      `json:"updatedBy"`
	DeletedAt  null.Time   `json:"deletedAt"`
	DeletedBy  null.String `json:"deletedBy"`
}

type CourseList []*Course

type CoursePrimaryID struct {
	Id string `db:"id"`
}
