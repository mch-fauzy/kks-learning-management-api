package service

import (
	"github.com/kks-learning-management-api/internal/domain/student/model"
	"github.com/kks-learning-management-api/internal/domain/student/model/dto"
	"github.com/kks-learning-management-api/shared"
	"github.com/rs/zerolog/log"
)

type StudentManagementService interface {
	GetStudentById(req dto.ViewStudentByIdRequest) (dto.StudentResponse, error)
	GetStudentList(req dto.ViewStudentRequest) (dto.StudentListResponse, dto.Pagination, error)
}

func (s *StudentServiceImpl) GetStudentById(req dto.ViewStudentByIdRequest) (dto.StudentResponse, error) {

	primaryId := req.ToModel()
	studentById, err := s.StudentRepository.ResolveStudentById(primaryId.ToStudentPrimaryID())
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve student by id")
		return dto.StudentResponse{}, err
	}

	enrollmentListByStudentId, err := s.StudentRepository.ResolveEnrollmentListByFilter(model.Filter{
		FilterFields: []model.FilterField{
			{
				Field:    "student_id",
				Operator: model.OperatorEqual,
				Value:    req.StudentId,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve enrollment list by student id")
		return dto.StudentResponse{}, err
	}

	courseListById, err := s.StudentRepository.ResolveCourseListByFilter(model.Filter{
		FilterFields: []model.FilterField{
			{
				Field:    "id",
				Operator: model.OperatorIn,
				Value:    shared.SliceStringToInterfaces(enrollmentListByStudentId.ToCourseIdSlice()),
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve course list by id")
		return dto.StudentResponse{}, err
	}

	result := dto.BuildStudentByIdResponse(studentById, enrollmentListByStudentId, courseListById)

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
