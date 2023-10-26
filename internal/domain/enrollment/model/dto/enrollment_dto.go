package dto

import (
	"time"

	"github.com/guregu/null"
)

type EnrollmentResponse struct {
	Id                   int         `json:"id,omitempty"`
	StudentId            string      `json:"studentId,omitempty"`
	CourseId             string      `json:"courseId,omitempty"`
	CourseEnrollmentDate time.Time   `json:"courseEnrollmentDate,omitempty"`
	CreatedAt            time.Time   `json:"createdAt,omitempty"`
	CreatedBy            string      `json:"createdBy,omitempty"`
	UpdatedAt            time.Time   `json:"updatedAt,omitempty"`
	UpdatedBy            string      `json:"updatedBy,omitempty"`
	DeletedAt            null.Time   `json:"deletedAt,omitempty"`
	DeletedBy            null.String `json:"deletedBy,omitempty"`
}

type EnrollmentListResponse []EnrollmentResponse
