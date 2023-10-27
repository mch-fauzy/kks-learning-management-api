package model

import (
	"time"

	"github.com/guregu/null"
)

type Student struct {
	Id             string      `db:"id"`
	Name           string      `db:"name"`
	Origin         string      `db:"origin"`
	EnrollmentDate time.Time   `db:"enrollment_date"`
	GPA            null.Float  `db:"gpa"`
	CreatedAt      time.Time   `db:"created_at"`
	CreatedBy      string      `db:"created_by"`
	UpdatedAt      time.Time   `db:"updated_at"`
	UpdatedBy      string      `db:"updated_by"`
	DeletedAt      null.Time   `db:"deleted_at"`
	DeletedBy      null.String `db:"deleted_by"`
}

type StudentList []*Student

type StudentPrimaryID struct {
	Id string `db:"id"`
}

func (s Student) ToStudentPrimaryID() StudentPrimaryID {
	return StudentPrimaryID{
		Id: s.Id,
	}
}

func (s Student) ToEnrollmentStudentID() StudentEnrollmentStudentID {
	return StudentEnrollmentStudentID{
		StudentId: s.Id,
	}
}
