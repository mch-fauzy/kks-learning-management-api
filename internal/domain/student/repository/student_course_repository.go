package repository

import (
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
)

type StudentCourseRepository interface {
	ResolveCourseListByFilter(filter model.Filter) (model.StudentCourseList, error)
}

func (r *StudentRepositoryMySQL) ResolveCourseListByFilter(filter model.Filter) (model.StudentCourseList, error) {
	query, args, err := composeFilterQuery(selectCourseQuery, filter)
	if err != nil {
		return model.StudentCourseList{}, err
	}

	var course model.StudentCourseList
	err = r.DB.Read.Select(&course, query, args...)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[ResolveCourseListById] Failed to get course list by id")
		err = failure.InternalError(err)
		return model.StudentCourseList{}, err
	}

	return course, nil
}
