package dto

import (
	"time"

	enrollmentModel "github.com/kks-learning-management-api/internal/domain/enrollment/model"
)

type StudentEnrollmentResponse struct {
	CourseId             string    `json:"courseId"`
	CourseEnrollmentDate time.Time `json:"courseEnrollmentDate"`
}

type StudentEnrollmentListResponse []StudentEnrollmentResponse

func NewStudentEnrollmentResponse(enrollment enrollmentModel.Enrollment) StudentEnrollmentResponse {
	return StudentEnrollmentResponse{
		CourseId:             enrollment.CourseId,
		CourseEnrollmentDate: enrollment.CourseEnrollmentDate,
	}
}

func BuildStudentEnrollmentListResponse(enrollmentList enrollmentModel.EnrollmentList) StudentEnrollmentListResponse {
	results := StudentEnrollmentListResponse{}
	for _, enrollment := range enrollmentList {
		results = append(results, NewStudentEnrollmentResponse(*enrollment))
	}
	return results
}
