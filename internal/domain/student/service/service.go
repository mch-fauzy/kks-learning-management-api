package service

import (
	"github.com/kks-learning-management-api/internal/domain/student/repository"
)

type StudentService interface {
	StudentManagementService
}

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
}

func ProvideStudentServiceImpl(studentRepository repository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
	}
}
