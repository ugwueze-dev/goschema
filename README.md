# goschema
Simple library to create MySQL database and tables written in Go.

## Example Usage 
```go
package main

import (
	"github.com/codemaestro64/goschema"
)

func main() {
	cfg := goschema.DBConfig{
		Name: "goschema",
		Host: "localhost",
		Username: "root",
		Password: "password",
	}
	schema := goschema.NewSchema(cfg, true)

	userTable := schema.NewTable("users")
	userIDCol := userTable.Int("ID", 11).IsPrimaryKey()
	userTable.Varchar("username", 50)
	userTable.Timestamp("created_at").SetDefaultValue("CURRENT_TIMESTAMP")

	profileTable := schema.NewTable("profiles")
	profileTable.Int("ID", 11).IsPrimaryKey().Reference(userIDCol, goschema.Cascade, goschema.Cascade)
	profileTable.Varchar("firstname", 70)
	profileTable.Varchar("lastname", 70)
	profileTable.Timestamp("created_at").SetDefaultValue("CURRENT_TIMESTAMP")

	err := schema.Create()
	if err != nil {
		log.Fatal(err)
	}
}
```

## API 
[API DOC](https://godoc.org/github.com/codemaestro64/goschema).
