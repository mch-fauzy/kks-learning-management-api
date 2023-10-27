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
	selectEnrollmentQuery = `
	SELECT
		id,
		student_id,
		course_id,
		course_enrollment_date
	FROM
		enrollment
	`
	filterEnrollmentByStudentIdQuery = `
	WHERE
		student_id = ?
	`
)

type StudentEnrollmentRepository interface {
	ResolveEnrollmentByStudentId(filter model.StudentEnrollmentStudentID) (model.StudentEnrollmentList, error)
}

func (r *StudentRepositoryMySQL) ResolveEnrollmentByStudentId(filter model.StudentEnrollmentStudentID) (model.StudentEnrollmentList, error) {
	query := fmt.Sprintf(selectEnrollmentQuery)

	var args []interface{}
	query += filterEnrollmentByStudentIdQuery
	args = append(args, filter.StudentId)

	var studentEnrollment model.StudentEnrollmentList
	err := r.DB.Read.Select(&studentEnrollment, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("Enrollment with student id '%s' not found", fmt.Sprint(filter.StudentId)))
			return model.StudentEnrollmentList{}, err
		}
		log.Error().
			Err(err).
			Msg("[ResolveEnrollmentByStudentId] Failed to get enrollment by student id")
		err = failure.InternalError(err)
		return model.StudentEnrollmentList{}, err
	}

	return studentEnrollment, nil
}
