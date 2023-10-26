package repository

import (
	"github.com/kks-learning-management-api/infras"
)

type EnrollmentRepository interface {
	EnrollmentManagementRepository
}

type EnrollmentRepositoryMySQL struct {
	DB *infras.MySQLConn
}

func ProvideEnrollmentRepositoryMySQL(db *infras.MySQLConn) *EnrollmentRepositoryMySQL {
	return &EnrollmentRepositoryMySQL{
		DB: db,
	}
}
