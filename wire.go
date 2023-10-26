//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kks-learning-management-api/configs"
	"github.com/kks-learning-management-api/infras"
	enrollmentRepository "github.com/kks-learning-management-api/internal/domain/enrollment/repository"
	studentRepository "github.com/kks-learning-management-api/internal/domain/student/repository"
	studentService "github.com/kks-learning-management-api/internal/domain/student/service"
	studentHandler "github.com/kks-learning-management-api/internal/handlers/student"
	"github.com/kks-learning-management-api/transport/http"
	"github.com/kks-learning-management-api/transport/http/router"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvideMySQLConn,
)

// Wiring for domain.
var domainStudent = wire.NewSet(
	// Service interface and implementation
	studentService.ProvideStudentServiceImpl,
	wire.Bind(new(studentService.StudentService), new(*studentService.StudentServiceImpl)),
	// Repository interface and implementation
	studentRepository.ProvideStudentRepositoryMySQL,
	wire.Bind(new(studentRepository.StudentRepository), new(*studentRepository.StudentRepositoryMySQL)),
)

var domainEnrollment = wire.NewSet(
	// Repository interface and implementation
	enrollmentRepository.ProvideEnrollmentRepositoryMySQL,
	wire.Bind(new(enrollmentRepository.EnrollmentRepository), new(*enrollmentRepository.EnrollmentRepositoryMySQL)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainStudent,
	domainEnrollment,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "StudentHandler"),
	studentHandler.ProvideStudentHandler,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
