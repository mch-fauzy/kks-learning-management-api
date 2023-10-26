package service

import (
	enrollmentRepository "github.com/kks-learning-management-api/internal/domain/enrollment/repository"
	studentRepository "github.com/kks-learning-management-api/internal/domain/student/repository"
)

type StudentService interface {
	StudentManagementService
}

type StudentServiceImpl struct {
	StudentRepository    studentRepository.StudentRepository
	EnrollmentRepository enrollmentRepository.EnrollmentRepository
}

func ProvideStudentServiceImpl(studentRepository studentRepository.StudentRepository, enrollmentRepository enrollmentRepository.EnrollmentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		StudentRepository:    studentRepository,
		EnrollmentRepository: enrollmentRepository,
	}
}
