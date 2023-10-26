package service

import (
	"github.com/kks-learning-management-api/internal/domain/student/model/dto"
	"github.com/rs/zerolog/log"
)

type StudentManagementService interface {
	GetStudentList(req dto.ViewStudentRequest) (dto.StudentListResponse, dto.Pagination, error)
}

func (s *StudentServiceImpl) GetStudentList(req dto.ViewStudentRequest) (dto.StudentListResponse, dto.Pagination, error) {

	paginationFilter := req.ToPaginationModel()
	student, err := s.StudentRepository.GetStudents(paginationFilter)
	if err != nil {
		log.Error().Err(err).Msg("[GetStudentList] Failed to retrieve student list")
		return dto.StudentListResponse{}, dto.Pagination{}, err
	}

	result := dto.BuildStudentListResponse(student)
	paginationMetadata := dto.BuildMetadata(req.Page, req.PageSize)
	return result, paginationMetadata, nil
}
