package model

import (
	"time"

	"github.com/guregu/null"
)

type studentDBFieldName struct {
	Id             string
	Name           string
	Origin         string
	EnrollmentDate string
	GPA            string
	CreatedAt      string
	CreatedBy      string
	UpdatedAt      string
	UpdatedBy      string
	DeletedAt      string
	DeletedBy      string
}

var StudentDBFieldName = studentDBFieldName{
	Id:             "id",
	Name:           "name",
	Origin:         "origin",
	EnrollmentDate: "enrollment_date",
	GPA:            "gpa",
	CreatedAt:      "created_at",
	CreatedBy:      "created_by",
	UpdatedAt:      "updated_at",
	UpdatedBy:      "updated_by",
	DeletedAt:      "deleted_at",
	DeletedBy:      "deleted_by",
}

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

func (s Student) ToStudentEnrollmentStudentID() StudentEnrollmentStudentID {
	return StudentEnrollmentStudentID{
		StudentId: s.Id,
	}
}
