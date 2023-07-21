package goschema

import (
	"fmt"
	"strings"
)

func (g *GoSchema) build() string {
	var b strings.Builder

	var references []*Reference

	for _, table := range g.tables {
		if table.references != nil {
			references = append(references, table.references...)
		}

		b.WriteString("CREATE TABLE `")
		b.WriteString(table.name)
		b.WriteString("` (\n")

		for _, column := range table.columns {
			g.buildColumn(&b, column)
		}

		b.WriteString(");\n\n")
	}

	g.buildReferences(&b, references)
	return b.String()
}

func (g *GoSchema) buildReferences(b *strings.Builder, references []*Reference) {
	for _, reference := range references {
		b.WriteString("\n\nALTER TABLE `")
		b.WriteString(reference.tableName)
		b.WriteString("`\n  ")
		b.WriteString("ADD CONSTRAINT")
	}
}

func (g *GoSchema) buildColumn(b *strings.Builder, col *Column) {
	// column name
	b.WriteString("  `")
	b.WriteString(col.name)
	b.WriteString("` ")

	// datatype
	b.WriteString(col.dataType)

	// add size if set
	if col.size != -1 {
		b.WriteString(fmt.Sprintf("(%d)", col.size))
	}

	// add null status
	if col.nullable {
		b.WriteString(" NULL")
	} else {
		b.WriteString(" NOT NULL")
	}

	// set default value if it exists
	if col.defaultValue != nil {
		b.WriteString(fmt.Sprintf(" DEFAULT %v", col.defaultValue))
	}

	b.WriteString(",\n")
}
