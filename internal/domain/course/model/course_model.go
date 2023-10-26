package model

type Course struct {
	Id         int    `db:"id"`
	LecturerId string `db:"lecturer_id"`
	Name       string `db:"name"`
	Credit     int    `db:"credit"`
}

type CourseList []*Course
