package model

import (
	"time"

	"github.com/guregu/null"
)

type Enrollment struct {
	Id                   int         `db:"id"`
	StudentId            string      `db:"student_id"`
	CourseId             string      `db:"course_id"`
	CourseEnrollmentDate time.Time   `db:"course_enrollment_date"`
	CreatedAt            time.Time   `db:"created_at"`
	CreatedBy            string      `db:"created_by"`
	UpdatedAt            time.Time   `db:"updated_at"`
	UpdatedBy            string      `db:"updated_by"`
	DeletedAt            null.Time   `db:"deleted_at"`
	DeletedBy            null.String `db:"deleted_by"`
}

type EnrollmentList []*Enrollment

type EnrollmentStudentID struct {
	StudentId string `db:"student_id"`
}
