package dto

import (
	"time"

	"github.com/kks-learning-management-api/internal/domain/student/model"
)

type StudentEnrollmentResponse struct {
	CourseId             string    `json:"courseId"`
	CourseEnrollmentDate time.Time `json:"courseEnrollmentDate"`
}

type StudentEnrollmentListResponse []StudentEnrollmentResponse

func NewStudentEnrollmentResponse(enrollment model.StudentEnrollment) StudentEnrollmentResponse {
	return StudentEnrollmentResponse{
		CourseId:             enrollment.CourseId,
		CourseEnrollmentDate: enrollment.CourseEnrollmentDate,
	}
}

func BuildStudentEnrollmentListResponse(enrollmentList model.StudentEnrollmentList) StudentEnrollmentListResponse {
	results := StudentEnrollmentListResponse{}
	for _, enrollment := range enrollmentList {
		results = append(results, NewStudentEnrollmentResponse(*enrollment))
	}
	return results
}
