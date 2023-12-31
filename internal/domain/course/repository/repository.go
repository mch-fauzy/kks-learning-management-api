package repository

import (
	"github.com/kks-learning-management-api/infras"
)

type CourseRepository interface {
}

type CourseRepositoryMySQL struct {
	DB *infras.MySQLConn
}

func ProvideCourseRepositoryMySQL(db *infras.MySQLConn) *CourseRepositoryMySQL {
	return &CourseRepositoryMySQL{
		DB: db,
	}
}
