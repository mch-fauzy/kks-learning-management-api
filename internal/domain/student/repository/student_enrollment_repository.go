package repository

import (
	"fmt"
	"strings"

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
	query, args, err := composeStudentEnrollmentFilterQuery(filter)
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

func composeStudentEnrollmentFilterQuery(filter model.Filter) (string, []interface{}, error) {
	var err error
	var args []interface{}
	query := fmt.Sprintf(selectEnrollmentQuery)

	if len(filter.FilterFields) > 0 {
		var (
			whereQueries []string
			whereArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				// valueArray, ok := filterField.Value.([]interface{})
				// if !ok {
				// 	err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
				// 	return query, args, err
				// }
				whereQueries = append(whereQueries, fmt.Sprintf("`%s` = ?", filterField.Field))
				whereArgs = append(whereArgs, filterField.Value)
				// whereArgs = append(whereArgs, valueArray...)
			case model.OperatorRange:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}
				whereQueries = append(whereQueries, fmt.Sprintf("`%s` BETWEEN ? AND ?", filterField.Field))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIn:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}

				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, "?")
				}

				whereQueries = append(whereQueries, fmt.Sprintf("`%s` IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIsNull:
				value, ok := filterField.Value.(bool)
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}
				if value {
					whereQueries = append(whereQueries, fmt.Sprintf("`%s` IS NULL", filterField.Field))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("`%s` IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				whereQueries = append(whereQueries, fmt.Sprintf("`%s` != ?", filterField.Field))
				whereArgs = append(whereArgs, filterField.Value)
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))
		args = append(args, whereArgs...)
	}

	if len(filter.Sorts) > 0 {
		sortQuery := []string{}
		for _, sort := range filter.Sorts {
			sortQuery = append(sortQuery, fmt.Sprintf("`%s` %s", sort.Field, sort.Order))
		}
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
	}

	if filter.Pagination.PageSize > 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}
	return query, args, err
}
