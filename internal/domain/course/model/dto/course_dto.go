package dto

import (
	"time"

	"github.com/guregu/null"
)

type CourseResponse struct {
	Id         string      `json:"id"`
	LecturerId string      `json:"lecturer_id"`
	Name       string      `json:"name"`
	Credit     int         `json:"credit"`
	CreatedAt  time.Time   `json:"createdAt"`
	CreatedBy  string      `json:"createdBy"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	UpdatedBy  string      `json:"updatedBy"`
	DeletedAt  null.Time   `json:"deletedAt"`
	DeletedBy  null.String `json:"deletedBy"`
}

type CourseListResponse []*CourseResponse
