package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kks-learning-management-api/internal/domain/enrollment/model"
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
	filterByStudentIdQuery = `
	WHERE
		student_id = ?
	`
)

type EnrollmentManagementRepository interface {
	ResolveEnrollmentByStudentId(filter model.EnrollmentStudentID) (model.EnrollmentList, error)
}

func (r *EnrollmentRepositoryMySQL) ResolveEnrollmentByStudentId(filter model.EnrollmentStudentID) (model.EnrollmentList, error) {
	query := fmt.Sprintf(selectEnrollmentQuery)

	var args []interface{}
	query += filterByStudentIdQuery
	args = append(args, filter.StudentId)

	var enrollment model.EnrollmentList
	err := r.DB.Read.Select(&enrollment, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("Enrollment with student id '%s' not found", fmt.Sprint(filter.StudentId)))
			return model.EnrollmentList{}, err
		}
		log.Error().
			Err(err).
			Msg("[ResolveEnrollmentByStudentId] Failed to get enrollment by student id")
		err = failure.InternalError(err)
		return model.EnrollmentList{}, err
	}

	return enrollment, nil
}
