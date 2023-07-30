package goschema

import (
	"strings"
)

func (g *GoSchema) buildSchema() string {
	var b strings.Builder
	var r strings.Builder

	for _, table := range g.tables {
		indexed, references := g.buildTable(table, &b)
		g.buildTableConstraints(table.name, indexed, references, &r)
	}

	b.WriteString(r.String())
	return b.String()
}

func (g *GoSchema) writePrimaryKey(name string, b *strings.Builder) {
	b.WriteString("\n  ADD PRIMARY KEY (`")
	b.WriteString(name)
	b.WriteString("`),")
}

func (g *GoSchema) writeAutoIncrement(tableName string, col *commonData, b *strings.Builder) {
	b.WriteString("ALTER TABLE `")
	b.WriteString(tableName)
	b.WriteString("`\n  MODIFY `")
	b.WriteString(col.name)
	b.WriteString("` ")
	b.WriteString(col.dataType.String())
	b.WriteString(" ")
	b.WriteString(col.nullStatus)
	b.WriteString(" AUTO_INCREMENT;\n")

}

func (g *GoSchema) writeUniqueKey(tableName, name string, b *strings.Builder) {
	b.WriteString("\n  ADD UNIQUE KEY `")
	b.WriteString(tableName)
	b.WriteString("_")
	b.WriteString(name)
	b.WriteString("_unique`")
	b.WriteString(" (`")
	b.WriteString(name)
	b.WriteString("`),")
}

func (g *GoSchema) writeReference(colName string, ref *Reference, b *strings.Builder) {
	b.WriteString("\n  ADD CONSTRAINT `fk_")
	b.WriteString(colName)
	b.WriteString("` FOREIGN KEY (`")
	b.WriteString(colName)
	b.WriteString("`) REFERENCES `")
	b.WriteString(ref.tableName)
	b.WriteString("`(`")
	b.WriteString(ref.columnName)
	b.WriteString("`)")
	b.WriteString(" ON UPDATE ")
	b.WriteString(ref.onUpdate.String())
	b.WriteString(" ON DELETE ")
	b.WriteString(ref.onDelete.String())
	b.WriteString(",")
}

func (g *GoSchema) buildTableConstraints(tableName string, indexed, references []*commonData, b *strings.Builder) {
	if len(indexed) == 0 && len(references) == 0 {
		return
	}
	b.WriteString("ALTER TABLE `")
	b.WriteString(tableName)
	b.WriteString("`")

	for _, col := range indexed {
		if col.primaryKey {
			g.writePrimaryKey(col.name, b)
		}

		if col.unique {
			g.writeUniqueKey(tableName, col.name, b)
		}
	}

	for _, col := range references {
		g.writeReference(col.name, col.reference, b)
	}

	str := strings.TrimSuffix(b.String(), ",")
	b.Reset()
	b.WriteString(str)
	b.WriteString(";\n\n")

	for _, col := range indexed {
		if col.autoIncrement {
			g.writeAutoIncrement(tableName, col, b)
		}
	}
}

func (g *GoSchema) buildTable(t *Table, b *strings.Builder) ([]*commonData, []*commonData) {
	var indexed, references []*commonData

	b.WriteString("CREATE TABLE `")
	b.WriteString(t.name)
	b.WriteString("` (\n")

	for index, column := range t.columns {
		c := column.getCommonData()
		if c.autoIncrement || c.unique || c.primaryKey {
			indexed = append(indexed, &c)
		}

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

	return indexed, references
}
