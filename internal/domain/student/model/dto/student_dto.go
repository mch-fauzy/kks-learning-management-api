package dto

import (
	"time"

	"github.com/guregu/null"
	courseDTO "github.com/kks-learning-management-api/internal/domain/course/model/dto"
	enrollmentDTO "github.com/kks-learning-management-api/internal/domain/enrollment/model/dto"
	"github.com/kks-learning-management-api/internal/domain/student/model"
	"github.com/kks-learning-management-api/shared"
	"github.com/kks-learning-management-api/shared/failure"
)

type ViewStudentRequest struct {
	Page     int `json:"-"`
	PageSize int `json:"-"`
}

func BuildViewProductRequest(page, pageSize int) ViewStudentRequest {
	if page == 0 {
		page = shared.DefaultPage
	}

	if pageSize == 0 {
		pageSize = shared.DefaultPageSize
	}

	return ViewStudentRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

func (v ViewStudentRequest) Validate() error {
	if v.Page < 0 {
		return failure.BadRequestFromString("Page must be a positive integer")
	}

	if v.PageSize < 0 {
		return failure.BadRequestFromString("PageSize must be a positive integer")
	}

	return nil
}

func (v ViewStudentRequest) ToPaginationModel() model.Pagination {
	return model.Pagination{
		Page:     v.Page,
		PageSize: v.PageSize,
	}
}

type StudentResponse struct {
	Id             string                               `json:"id"`
	Name           string                               `json:"name"`
	Origin         string                               `json:"origin"`
	EnrollmentDate time.Time                            `json:"enrollmentDate"`
	GPA            null.Float                           `json:"gpa"`
	CreatedAt      time.Time                            `json:"createdAt"`
	CreatedBy      string                               `json:"createdBy"`
	UpdatedAt      time.Time                            `json:"updatedAt"`
	UpdatedBy      string                               `json:"updatedBy"`
	DeletedAt      null.Time                            `json:"deletedAt"`
	DeletedBy      null.String                          `json:"deletedBy"`
	Enrollment     enrollmentDTO.EnrollmentListResponse `json:"enrollment"`
	Courses        courseDTO.CourseListResponse         `json:"courses"`
}

type StudentListResponse []StudentResponse

func NewStudentListResponse(student model.Student) StudentResponse {
	return StudentResponse{
		Id:             student.Id,
		Name:           student.Name,
		Origin:         student.Origin,
		EnrollmentDate: student.EnrollmentDate,
		GPA:            student.GPA,
		CreatedAt:      student.CreatedAt,
		CreatedBy:      student.CreatedBy,
		UpdatedAt:      student.UpdatedAt,
		UpdatedBy:      student.UpdatedBy,
		DeletedAt:      student.DeletedAt,
		DeletedBy:      student.DeletedBy,
		Enrollment:     enrollmentDTO.EnrollmentListResponse{},
		Courses:        courseDTO.CourseListResponse{},
	}
}

func BuildStudentListResponse(studentList model.StudentList) StudentListResponse {
	results := StudentListResponse{}
	for _, student := range studentList {
		results = append(results, NewStudentListResponse(*student))
	}
	return results
}
