package repository

import (
	"github.com/kks-learning-management-api/internal/domain/student/model"
	"github.com/kks-learning-management-api/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	selectEnrollmentQuery = `
	SELECT
		id,
		student_id,
		course_id,
		course_enrollment_date
	FROM
		enrollment
	`
)

type StudentEnrollmentRepository interface {
	ResolveEnrollmentListByFilter(filter model.Filter) (model.StudentEnrollmentList, error)
}

func (r *StudentRepositoryMySQL) ResolveEnrollmentListByFilter(filter model.Filter) (model.StudentEnrollmentList, error) {
	query, args, err := composeFilterQuery(selectEnrollmentQuery, filter)
	if err != nil {
		return model.StudentEnrollmentList{}, err
	}

	var studentEnrollment model.StudentEnrollmentList
	err = r.DB.Read.Select(&studentEnrollment, query, args...)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[ResolveEnrollmentListByFilter] Failed to get enrollment list by student id")
		err = failure.InternalError(err)
		return model.StudentEnrollmentList{}, err
	}

	return studentEnrollment, nil
}
