package mssql

import (
	"fmt"
	"reflect"
	"strings"
)

type WhereClause struct {
	builder *JsonQueryBuilder
	query   string
}

func (clause *WhereClause) Is(expression string, value interface{}) *JsonQueryBuilder {
	clause.query += fmt.Sprintf("%s %v ", expression, formatSqlTypeValue(value))
	return clause.builder
}

func (clause *WhereClause) Between(from interface{}, to interface{}) *JsonQueryBuilder {
	clause.query += fmt.Sprintf("BETWEEN %v AND %v ", formatSqlTypeValue(from), formatSqlTypeValue(to))
	return clause.builder
}

func (clause *WhereClause) Like(value interface{}) *JsonQueryBuilder {
	clause.query += fmt.Sprintf("LIKE %v ", formatSqlTypeValue(value))
	return clause.builder
}

func (clause *WhereClause) In(values ...interface{}) *JsonQueryBuilder {
	statement := ""
	for index, item := range values {
		if index > 0 {
			statement += fmt.Sprintf(", %v", formatSqlTypeValue(item))
		} else {
			statement += fmt.Sprintf("%v", formatSqlTypeValue(item))
		}
	}
	clause.query += "IN (" + statement + ")"
	return clause.builder
}

func formatSqlTypeValue(val interface{}) interface{} {
	if val == nil {
		return "NULL"
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.String:
		return fmt.Sprintf("'%s'", val)
	default:
		return val
	}
}

type SelectClause struct {
	builder    *JsonQueryBuilder
	fieldChain []string
	query      string
}

func (clause *SelectClause) AsJson(col string) *JsonQueryBuilder {
	clause.query = "JSON_QUERY(" + clause.builder.col + ", '$." + strings.Join(clause.fieldChain, ".") + "') AS " + col
	return clause.builder
}

func (clause *SelectClause) AsValue(col string) *JsonQueryBuilder {
	clause.query = "JSON_VALUE(" + clause.builder.col + ", '$." + strings.Join(clause.fieldChain, ".") + "') AS " + col
	return clause.builder
}
