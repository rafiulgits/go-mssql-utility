# MSSQL Utility

Target of this package is to providing high level go lang api to interact with SQL Server api in sophisticated way.



## Features

 - JSON query builder (required SQL Server 2016 or letter)



## How To Use

Install in your project

`go get github.com/rafiulgits/go-mssql-utility`



Sample user table

| ID   | Name    | Info                                                         |
| ---- | ------- | ------------------------------------------------------------ |
| 1    | Tony    | `{ "age":  24,  "gender": "male", "likes":  ["technology", "programming", "travelling"] }` |
| 2    | Stark   | `{ "age":  26,  "gender": "male", "likes":  ["baseball", "cricket", "travelling"] }` |
| 3    | Jenifer | `{ "age":  25,  "gender": "female", "likes":  ["dance", "reading"]}` |




Build JSON query for `SELECT` and `WHERE` clause

```go
builder := mssql.NewJsonQueryBuilder("info").Where("age").Between(
  					24, 30).Where("gender").In("female").Select("likes").AsValue("likes")

buildResult := builder.Build()
```



Use build clauses in SQL query. Here is an example of using this build clauses inside `GORM`  function

```go
var data []map[string]interface{}

queryErr :=db.Table(UserTableName).Where(buildResult.WhereQuery).Select(
  					buildResult.SelectQuery).Scan(&data).Error
```



