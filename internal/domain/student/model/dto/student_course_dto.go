package dto

import (
	"github.com/kks-learning-management-api/internal/domain/student/model"
)

type StudentCourseResponse struct {
	Id         string `json:"id"`
	LecturerId string `json:"lecturerId"`
	Name       string `json:"name"`
	Credit     int    `json:"credit"`
}

type StudentCourseListResponse []StudentCourseResponse

func NewStudentCourseResponse(course model.StudentCourse) StudentCourseResponse {
	return StudentCourseResponse{
		Id:         course.Id,
		LecturerId: course.LecturerId,
		Name:       course.Name,
		Credit:     course.Credit,
	}
}

func BuildStudentCourseListResponse(courseList model.StudentCourseList) StudentCourseListResponse {
	results := StudentCourseListResponse{}
	for _, course := range courseList {
		results = append(results, NewStudentCourseResponse(*course))
	}
	return results
}
