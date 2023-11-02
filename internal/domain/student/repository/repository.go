package repository

import (
	"fmt"
	"strings"

	"github.com/kks-learning-management-api/infras"
	"github.com/kks-learning-management-api/internal/domain/student/model"
	"github.com/kks-learning-management-api/shared/failure"
)

type StudentRepository interface {
	StudentManagementRepository
	StudentEnrollmentRepository
	StudentCourseRepository
}

type StudentRepositoryMySQL struct {
	DB *infras.MySQLConn
}

func ProvideStudentRepositoryMySQL(db *infras.MySQLConn) *StudentRepositoryMySQL {
	return &StudentRepositoryMySQL{
		DB: db,
	}
}

func composePrimaryIdQuery(queryStr, primaryID string) (string, []interface{}) {
	var args []interface{}
	query := fmt.Sprintf(queryStr)
	query += fmt.Sprintf(" WHERE `id` = ?")
	args = append(args, primaryID)
	return query, args
}

func composeFilterQuery(queryStr string, filter model.Filter) (string, []interface{}, error) {
	var (
		err  error
		args []interface{}
	)

	query := fmt.Sprintf(queryStr)

	if len(filter.FilterFields) > 0 {
		var (
			filterQueries []string
			filterArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				// Input can be string, int, float (singular)
				filterQueries = append(filterQueries, fmt.Sprintf("`%s` = ?", filterField.Field))
				filterArgs = append(filterArgs, filterField.Value)
			case model.OperatorRange:
				// Input must be []interface{}
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}
				filterQueries = append(filterQueries, fmt.Sprintf("`%s` BETWEEN ? AND ?", filterField.Field))
				filterArgs = append(filterArgs, valueArray...)
			case model.OperatorIn:
				// Input must be []interface{}
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}

				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, "?")
				}

				filterQueries = append(filterQueries, fmt.Sprintf("`%s` IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				filterArgs = append(filterArgs, valueArray...)
			case model.OperatorIsNull:
				// Input must be boolean
				value, ok := filterField.Value.(bool)
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return query, args, err
				}
				if value {
					filterQueries = append(filterQueries, fmt.Sprintf("`%s` IS NULL", filterField.Field))
				} else {
					filterQueries = append(filterQueries, fmt.Sprintf("`%s` IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				// Input can be string, int, float (singular)
				filterQueries = append(filterQueries, fmt.Sprintf("`%s` != ?", filterField.Field))
				filterArgs = append(filterArgs, filterField.Value)
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(filterQueries, " AND "))
		args = append(args, filterArgs...)
	}

	if len(filter.Sorts) > 0 {
		var sortQueries []string
		for _, sort := range filter.Sorts {
			sortQueries = append(sortQueries, fmt.Sprintf("`%s` %s", sort.Field, sort.Order))
		}
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQueries, ","))
	}

	if filter.Pagination.PageSize > 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}
	return query, args, err
}
