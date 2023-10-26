package service

import (
	courseModel "github.com/kks-learning-management-api/internal/domain/course/model"
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

	courseDetails := make(courseModel.CourseList, 0)

	// Iterate through the enrollments and fetch course details for each course_id
	for _, enrollment := range enrollmentByStudentId {
		courseById, err := s.CourseRepository.ResolveCourseById(enrollment.ToCoursePrimaryId())
		if err != nil {
			log.Error().Err(err).Msg("[GetStudentById] Failed to retrieve course by id")
			return dto.StudentResponse{}, err
		}

		courseDetails = append(courseDetails, &courseById)
	}

	result := dto.BuildStudentByIdResponse(studentById, enrollmentByStudentId, courseDetails)

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
