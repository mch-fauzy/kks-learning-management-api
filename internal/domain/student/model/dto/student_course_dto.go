package dto

import (
	courseModel "github.com/kks-learning-management-api/internal/domain/course/model"
)

type StudentCourseResponse struct {
	Id         string `json:"id"`
	LecturerId string `json:"lecturer_id"`
	Name       string `json:"name"`
	Credit     int    `json:"credit"`
}

type StudentCourseListResponse []StudentCourseResponse

func NewStudentCourseResponse(course courseModel.Course) StudentCourseResponse {
	return StudentCourseResponse{
		Id:         course.Id,
		LecturerId: course.LecturerId,
		Name:       course.Name,
		Credit:     course.Credit,
	}
}

func BuildStudentCourseListResponse(courseList courseModel.CourseList) StudentCourseListResponse {
	results := StudentCourseListResponse{}
	for _, course := range courseList {
		results = append(results, NewStudentCourseResponse(*course))
	}
	return results
}
