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
				Field:    model.StudentEnrollmentDBFieldName.StudentId,
				Operator: model.OperatorEqual,
				Value:    studentById.Id,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve enrollment list by student id")
		return dto.StudentResponse{}, err
	}

	courseIdSlice := enrollmentListByStudentId.ToCourseIdSlice()
	courseListById, err := s.StudentRepository.ResolveCourseListByFilter(model.Filter{
		FilterFields: []model.FilterField{
			{
				Field:    model.StudentCourseDBFieldName.Id,
				Operator: model.OperatorIn,
				Value:    shared.SliceStringToInterfaces(courseIdSlice),
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

	studentList, err := s.StudentRepository.ResolveStudentListByFilter(model.Filter{
		Pagination: model.Pagination{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Sorts: []model.Sort{
			{
				Field: model.StudentDBFieldName.Id,
				Order: model.SortDesc,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentList] Failed to retrieve student list")
		return dto.StudentListResponse{}, dto.Pagination{}, err
	}

	result := dto.BuildStudentListResponse(studentList)
	paginationMetadata := dto.BuildMetadata(req.Page, req.PageSize)
	return result, paginationMetadata, nil
}
