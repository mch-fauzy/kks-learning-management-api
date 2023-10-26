package service

import (
	"github.com/kks-learning-management-api/internal/domain/student/model/dto"
	"github.com/rs/zerolog/log"
)

type StudentManagementService interface {
	GetStudentById(req dto.ViewStudentByIdRequest) (dto.StudentResponse, error)
	GetStudentList(req dto.ViewStudentRequest) (dto.StudentListResponse, dto.Pagination, error)
}

func (s *StudentServiceImpl) GetStudentById(req dto.ViewStudentByIdRequest) (dto.StudentResponse, error) {

	filter := req.ToModel()
	studentById, err := s.StudentRepository.ResolveStudentById(filter.ToStudentPrimaryID())
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve student by id")
		return dto.StudentResponse{}, err
	}

	enrollmentByStudentId, err := s.EnrollmentRepository.ResolveEnrollmentByStudentId(studentById.ToEnrollmentStudentID())
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve enrollment by student id")
		return dto.StudentResponse{}, err
	}

	result := dto.BuildStudentByIdResponse(studentById, enrollmentByStudentId)

	return result, nil
}

func (s *StudentServiceImpl) GetStudentList(req dto.ViewStudentRequest) (dto.StudentListResponse, dto.Pagination, error) {

	paginationFilter := req.ToPaginationModel()
	student, err := s.StudentRepository.ResolveStudents(paginationFilter)
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentList] Failed to retrieve student list")
		return dto.StudentListResponse{}, dto.Pagination{}, err
	}

	result := dto.BuildStudentListResponse(student)
	paginationMetadata := dto.BuildMetadata(req.Page, req.PageSize)
	return result, paginationMetadata, nil
}
