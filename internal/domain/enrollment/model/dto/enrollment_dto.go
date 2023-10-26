package dto

import (
	"time"

	"github.com/guregu/null"
)

type EnrollmentResponse struct {
	Id                   int         `json:"id"`
	StudentId            string      `json:"studentId"`
	CourseId             string      `json:"courseId"`
	CourseEnrollmentDate time.Time   `json:"courseEnrollmentDate"`
	CreatedAt            time.Time   `json:"createdAt"`
	CreatedBy            string      `json:"createdBy"`
	UpdatedAt            time.Time   `json:"updatedAt"`
	UpdatedBy            string      `json:"updatedBy"`
	DeletedAt            null.Time   `json:"deletedAt"`
	DeletedBy            null.String `json:"DeletedBy"`
}

type EnrollmentListResponse []*EnrollmentResponse
