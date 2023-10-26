package service

import (
	courseRepository "github.com/kks-learning-management-api/internal/domain/course/repository"
	enrollmentRepository "github.com/kks-learning-management-api/internal/domain/enrollment/repository"
	studentRepository "github.com/kks-learning-management-api/internal/domain/student/repository"
)

type StudentService interface {
	StudentManagementService
}

type StudentServiceImpl struct {
	StudentRepository    studentRepository.StudentRepository
	EnrollmentRepository enrollmentRepository.EnrollmentRepository
	CourseRepository     courseRepository.CourseRepository
}

func ProvideStudentServiceImpl(studentRepository studentRepository.StudentRepository, enrollmentRepository enrollmentRepository.EnrollmentRepository, courseRepository courseRepository.CourseRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		StudentRepository:    studentRepository,
		EnrollmentRepository: enrollmentRepository,
		CourseRepository:     courseRepository,
	}
}
