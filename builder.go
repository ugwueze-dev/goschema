package goschema

import (
	"strings"
)

func (g *GoSchema) buildSchema() string {
	var b strings.Builder
	var r strings.Builder

	for _, table := range g.tables {
		references := g.buildTable(table, &b)

		g.buildTableReferences(table.name, references, &r)
	}

	b.WriteString(r.String())
	return b.String()
}

func (g *GoSchema) buildTable(t *Table, b *strings.Builder) []*commonData {
	var references []*commonData

	b.WriteString("CREATE TABLE `")
	b.WriteString(t.name)
	b.WriteString("` (\n")

	for index, column := range t.columns {
		c := column.getCommonData()
		if c.reference != nil {
			references = append(references, &c)
		}

		// column name
		b.WriteString("  `")
		b.WriteString(c.name)
		b.WriteString("` ")

		// datatype
		b.WriteString(c.dataType.String())

		// write size or option
		column.writeSizeOrOption(b)

		// add null status
		b.WriteString(" ")
		b.WriteString(c.nullStatus)

		// write primary key
		if c.primaryKey {
			b.WriteString(" PRIMARY KEY")
		}

		// write unique
		if c.unique {
			b.WriteString(" UNIQUE")
		}

		// write auto increment
		if c.autoIncrement {
			b.WriteString(" AUTO_INCREMENT")
		}

		// only add comma (,) if it's not last column
		if index != len(t.columns)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}

	b.WriteString(")")
	b.WriteString(" ENGINE=")
	b.WriteString(InnoDB)
	//b.WriteString(" DEFAULT CHARSET=")
	//b.WriteString(DefaultCharset)
	b.WriteString(";\n\n")

	return references
}

func (g *GoSchema) buildTableReferences(tableName string, references []*commonData, b *strings.Builder) {
	if len(references) == 0 {
		return
	}
	b.WriteString("ALTER TABLE `")
	b.WriteString(tableName)
	b.WriteString("`")

	for _, col := range references {
		b.WriteString("\n  ADD CONSTRAINT `fk_")
		b.WriteString(col.name)
		b.WriteString("` FOREIGN KEY (`")
		b.WriteString(col.name)
		b.WriteString("`) REFERENCES `")
		b.WriteString(col.reference.tableName)
		b.WriteString("`(`")
		b.WriteString(col.reference.columnName)
		b.WriteString("`)")
		b.WriteString(" ON UPDATE ")
		b.WriteString(col.reference.onUpdate.String())
		b.WriteString(" ON DELETE ")
		b.WriteString(col.reference.onDelete.String())
		b.WriteString(",")
	}

	str := strings.TrimSuffix(b.String(), ",")
	b.Reset()
	b.WriteString(str)
	b.WriteString(";\n\n")
}
