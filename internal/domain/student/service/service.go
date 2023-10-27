package service

import (
	studentRepository "github.com/kks-learning-management-api/internal/domain/student/repository"
)

type StudentService interface {
	StudentManagementService
}

type StudentServiceImpl struct {
	StudentRepository studentRepository.StudentRepository
}

func ProvideStudentServiceImpl(studentRepository studentRepository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
	}
}
