package model

import (
	"time"
)

type StudentEnrollment struct {
	Id                   int       `db:"id"`
	StudentId            string    `db:"student_id"`
	CourseId             string    `db:"course_id"`
	CourseEnrollmentDate time.Time `db:"course_enrollment_date"`
}

type StudentEnrollmentStudentID struct {
	StudentId string `db:"student_id"`
}

func (se StudentEnrollment) ToStudentCoursePrimaryId() StudentCoursePrimaryID {
	return StudentCoursePrimaryID{
		Id: se.CourseId,
	}
}

type StudentEnrollmentList []*StudentEnrollment

// func (seList StudentEnrollmentList) ToStudentCoursePrimaryIdList() StudentCoursePrimaryIDList {
// 	courseIdList := StudentCoursePrimaryIDList{}

// 	for _, studentEnrollment := range seList {
// 		courseIdList = append(courseIdList, &StudentCoursePrimaryID{Id: studentEnrollment.CourseId})
// 	}
// 	return courseIdList
// }

func (seList StudentEnrollmentList) ToCourseIdSlice() []string {
	results := []string{}
	for _, studentEnrollment := range seList {
		results = append(results, studentEnrollment.CourseId)
	}
	return results
}
