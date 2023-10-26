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
	filterByIdQuery = `
	WHERE
		id = ?
	`
	paginationQuery = `
	LIMIT ? OFFSET ?
	`
)

type StudentManagementRepository interface {
	ResolveStudents(pagination model.Pagination) (model.StudentList, error)
	ResolveStudentById(filter model.StudentPrimaryID) (model.Student, error)
}

func (r *StudentRepositoryMySQL) ResolveStudents(pagination model.Pagination) (model.StudentList, error) {
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
			Msg("[ResolveStudents] Failed to get student")
		err = failure.InternalError(err)
		return model.StudentList{}, err
	}

	return student, nil
}

func (r *StudentRepositoryMySQL) ResolveStudentById(filter model.StudentPrimaryID) (model.Student, error) {
	query := fmt.Sprintf(selectStudentQuery)

	var args []interface{}
	query += filterByIdQuery
	args = append(args, filter.Id)

	var student model.Student
	err := r.DB.Read.Get(&student, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("Student with id '%s' not found", fmt.Sprint(filter.Id)))
			return model.Student{}, err
		}
		log.Error().
			Err(err).
			Msg("[ResolveStudentById] Failed to get student by id")
		err = failure.InternalError(err)
		return model.Student{}, err
	}

	return student, nil
}
