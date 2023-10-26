package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kks-learning-management-api/internal/domain/course/model"
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
	filterByIdQuery = `
	WHERE
		id = ?
	`
)

type CourseManagementRepository interface {
	ResolveCourseById(filter model.CoursePrimaryID) (model.Course, error)
}

func (r *CourseRepositoryMySQL) ResolveCourseById(filter model.CoursePrimaryID) (model.Course, error) {
	query := fmt.Sprintf(selectCourseQuery)

	var args []interface{}
	query += filterByIdQuery
	args = append(args, filter.Id)

	var course model.Course
	err := r.DB.Read.Get(&course, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("Course with id '%s' not found", fmt.Sprint(filter.Id)))
			return model.Course{}, err
		}
		log.Error().
			Err(err).
			Msg("[ResolveCourseById] Failed to get course by id")
		err = failure.InternalError(err)
		return model.Course{}, err
	}

	return course, nil
}
