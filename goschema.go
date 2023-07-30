package goschema

import (
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type GoSchema struct {
	tables []*Table
}

func New() *GoSchema {
	return &GoSchema{}
}

func (g *GoSchema) Table(tableName string, cb TableFunc) {
	table := newTable(tableName)
	g.tables = append(g.tables, table)
	cb(table)
}

func (g *GoSchema) GenerateSchema() string {
	var b strings.Builder

	for _, table := range g.tables {
		table.hasIndexes = table.generateSchema(&b)
	}

	for _, table := range g.tables {
		table.generateIndexes(&b)
	}

	ioutil.WriteFile("schema.sql", []byte(b.String()), os.ModePerm)

	return b.String()
}
