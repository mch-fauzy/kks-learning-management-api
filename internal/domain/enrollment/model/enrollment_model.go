package model

import (
	"time"
)

type Enrollment struct {
	Id                   int       `db:"id"`
	StudentId            string    `db:"student_id"`
	CourseId             string    `db:"course_id"`
	CourseEnrollmentDate time.Time `db:"course_enrollment_date"`
}

type EnrollmentList []*Enrollment
