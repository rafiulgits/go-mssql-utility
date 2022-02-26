package mssql

import "strings"

type JsonQueryBuildResult struct {
	WhereQuery  string
	SelectQuery string
}

type JsonQueryBuilder struct {
	col     string
	wheres  []*WhereClause
	selects []*SelectClause
}

func NewJsonQueryBuilder(col string) *JsonQueryBuilder {
	return &JsonQueryBuilder{
		col:     col,
		wheres:  make([]*WhereClause, 0),
		selects: make([]*SelectClause, 0),
	}
}

func (builder *JsonQueryBuilder) Where(fieldChain ...string) *WhereClause {
	clause := &WhereClause{
		builder: builder,
		query:   "(JSON_VALUE(" + builder.col + ", '$." + strings.Join(fieldChain, ".") + "') ",
	}
	builder.wheres = append(builder.wheres, clause)
	return clause
}

func (builder *JsonQueryBuilder) Select(fieldChain ...string) *SelectClause {
	clause := &SelectClause{
		builder:    builder,
		fieldChain: fieldChain,
	}
	builder.selects = append(builder.selects, clause)
	return clause
}

func (builder *JsonQueryBuilder) Build() *JsonQueryBuildResult {
	return &JsonQueryBuildResult{
		WhereQuery:  builder.buildWhereQueryResult(),
		SelectQuery: builder.buildSelectQueryResult(),
	}
}

func (builder *JsonQueryBuilder) buildWhereQueryResult() string {
	query := "ISJSON(" + builder.col + ") > 0 "
	for _, where := range builder.wheres {
		query += " AND " + where.query + ")"
	}
	return query
}

func (builder *JsonQueryBuilder) buildSelectQueryResult() string {
	query := ""
	for index, clause := range builder.selects {
		if index > 0 {
			query += ", " + clause.query
		} else {
			query += clause.query
		}
	}
	return query
}
