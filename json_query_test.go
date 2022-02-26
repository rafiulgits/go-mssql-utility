package mssql_test

import (
	"testing"

	"github.com/rafiulgits/go-mssql-utility"
)

//
// Query example from SQL Server 2019 official documentation
// https://docs.microsoft.com/en-us/sql/relational-databases/json/json-data-sql-server?view=sql-server-ver15
//
/*
SELECT Name, Surname,
  JSON_VALUE(jsonCol, '$.info.address.PostCode') AS PostCode,
  JSON_VALUE(jsonCol, '$.info.address."Address Line 1"') + ' '
  + JSON_VALUE(jsonCol, '$.info.address."Address Line 2"') AS Address,
  JSON_QUERY(jsonCol, '$.info.skills') AS Skills
FROM People
WHERE ISJSON(jsonCol) > 0
  AND JSON_VALUE(jsonCol, '$.info.address.Town') = 'Belgrade'
  AND Status = 'Active'
ORDER BY JSON_VALUE(jsonCol, '$.info.address.PostCode')
*/

func TestJsonWhereClauseBuild(t *testing.T) {
	expectedQuery := `ISJSON(jsonCol) > 0  AND (JSON_VALUE(jsonCol, '$.info.address.Town') = 'Belgrade' )`
	query := mssql.NewJsonQueryBuilder("jsonCol").Where("info", "address", "Town").Is("=", "Belgrade").Build().WhereQuery

	if expectedQuery != query {
		t.Log("expected: ", expectedQuery)
		t.Log("found: ", query)
		t.Error("failed to build expected query")
	}
}

func TestJsonSelectClauseBuild(t *testing.T) {
	expectedQuery1 := `JSON_VALUE(jsonCol, '$.info.address.PostCode') AS PostCode`
	query1 := mssql.NewJsonQueryBuilder("jsonCol").Select("info", "address", "PostCode").AsValue("PostCode").Build().SelectQuery

	if expectedQuery1 != query1 {
		t.Log("expected: ", expectedQuery1)
		t.Log("found: ", query1)
		t.Error("failed to build expected query")
	}

	expectedQuery2 := `JSON_QUERY(jsonCol, '$.info.skills') AS Skills`
	query2 := mssql.NewJsonQueryBuilder("jsonCol").Select("info", "skills").AsJson("Skills").Build().SelectQuery

	if expectedQuery2 != query2 {
		t.Log("expected: ", expectedQuery2)
		t.Log("found: ", query2)
		t.Error("failed to build expected query")
	}
}

func TestJsonKeyExistingQuery(t *testing.T) {
	expectedQuery := `ISJSON(jsonCol) > 0  AND (JSON_VALUE(jsonCol, '$.info.address.Town') != NULL )`
	query := mssql.NewJsonQueryBuilder("jsonCol").Where("info", "address", "Town").Is("!=", nil).Build().WhereQuery

	if expectedQuery != query {
		t.Log("expected: ", expectedQuery)
		t.Log("found: ", query)
		t.Error("failed to build expected query")
	}
}
