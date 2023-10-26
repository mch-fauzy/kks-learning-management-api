package repository

import (
	"fmt"

	"github.com/kks-learning-management-api/internal/domain/student/model"
	"github.com/kks-learning-management-api/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	selectStudentQuery = `
	SELECT
		id,
		name,
		origin,
		enrollment_date,
		gpa,
		created_at,
		created_by,
		updated_at,
		updated_by,
		deleted_at,
		deleted_by
	FROM
		student
	`
	paginationQuery = `
	LIMIT ? OFFSET ?
	`
)

type StudentManagementRepository interface {
	GetStudents(pagination model.Pagination) (model.StudentList, error)
}

func (r *StudentRepositoryMySQL) GetStudents(pagination model.Pagination) (model.StudentList, error) {
	query := fmt.Sprintf(selectStudentQuery)

	var args []interface{}
	query += paginationQuery
	offset := (pagination.Page - 1) * pagination.PageSize
	args = append(args, pagination.PageSize, offset)

	var student model.StudentList
	err := r.DB.Read.Select(&student, query, args...)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[GetStudents] Failed to get student")
		err = failure.InternalError(err)
		return model.StudentList{}, err
	}

	return student, nil
}
