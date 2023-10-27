package repository

import (
	"github.com/kks-learning-management-api/infras"
)

type StudentRepository interface {
	StudentManagementRepository
	StudentEnrollmentRepository
	StudentCourseRepository
}

type StudentRepositoryMySQL struct {
	DB *infras.MySQLConn
}

func ProvideStudentRepositoryMySQL(db *infras.MySQLConn) *StudentRepositoryMySQL {
	return &StudentRepositoryMySQL{
		DB: db,
	}
}
