package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kks-learning-management-api/internal/domain/student/model"
	"github.com/kks-learning-management-api/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	selectCourseQuery = `
	SELECT
		id,
		lecturer_id,
		name,
		credit
	FROM
		course
	`
	filterCourseByIdQuery = `
	WHERE
		id = ?
	`
)

type StudentCourseRepository interface {
	ResolveCourseById(filter model.StudentCoursePrimaryID) (model.StudentCourse, error)
}

func (r *StudentRepositoryMySQL) ResolveCourseById(filter model.StudentCoursePrimaryID) (model.StudentCourse, error) {
	query := fmt.Sprintf(selectCourseQuery)

	var args []interface{}
	query += filterCourseByIdQuery
	args = append(args, filter.Id)

	var course model.StudentCourse
	err := r.DB.Read.Get(&course, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("Course with id '%s' not found", fmt.Sprint(filter.Id)))
			return model.StudentCourse{}, err
		}
		log.Error().
			Err(err).
			Msg("[ResolveCourseById] Failed to get course by id")
		err = failure.InternalError(err)
		return model.StudentCourse{}, err
	}

	return course, nil
}
