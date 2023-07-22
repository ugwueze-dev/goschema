package goschema

import (
	"strings"

	"github.com/ugwueze-dev/goschema/column"
)

type TableFunc func(table *Table)

type Table struct {
	name    string
	columns []column.Common
	//references []*Reference
}

func newTable(tableName string) *Table {
	return &Table{
		name: tableName,
	}
}

func (t *Table) Increments(columnName string) column.Common {
	c := column.New(columnName, column.Int)
	t.columns = append(t.columns, c)
	c.AutoIncrement()

	return c
}

func (t *Table) String(columnName string) column.Common {
	c := column.New(columnName, column.String)
	t.columns = append(t.columns, c)

	return c
}

// TODO set correct default char size
func (t *Table) Char(columnName string) column.Common {
	c := column.New(columnName, column.Char).Size(4)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Enum(columnName string, options []interface{}) column.Common {
	c := column.NewEnum(columnName, options)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Binary(columnName string) column.Common {
	c := column.New(columnName, column.Binary)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Blob(columnName string) column.Common {
	c := column.New(columnName, column.Binary)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Text(columnName string) column.Common {
	c := column.New(columnName, column.Text)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Boolean(columnName string) column.Common {
	c := column.New(columnName, column.Boolean)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Int(columnName string) column.Common {
	c := column.New(columnName, column.Int)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) TinyInt(columnName string) column.Common {
	c := column.New(columnName, column.TinyInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) SmallInt(columnName string) column.Common {
	c := column.New(columnName, column.SmallInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) MediumInt(columnName string) column.Common {
	c := column.New(columnName, column.MediumInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) BigInt(columnName string) column.Common {
	c := column.New(columnName, column.BigInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Float(columnName string) column.Common {
	c := column.New(columnName, column.Float)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Double(columnName string) column.Common {
	c := column.New(columnName, column.Double)
	t.columns = append(t.columns, c)

	return c
}

// TODO Handle decimal
func (t *Table) Decimal(columnName string, precision, scale int) column.Common {
	c := column.New(columnName, column.Decimal)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Bit(columnName string) column.Common {
	c := column.New(columnName, column.Bit)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) DateTime(columnName string) column.Common {
	c := column.New(columnName, column.DateTime)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Timestamp(columnName string) column.Common {
	c := column.New(columnName, column.Timestamp)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Date(columnName string) column.Common {
	c := column.New(columnName, column.Date)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Time(columnName string) column.Common {
	c := column.New(columnName, column.Time)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Year(columnName string) column.Common {
	c := column.New(columnName, column.Year)
	t.columns = append(t.columns, c)

	return c
}

/**func (t *Table) Foriegn(columnName string) *Reference {
	ref := newReference(t.name, columnName)
	t.references = append(t.references, ref)

	return t.references[len(t.references)-1]
}**/

func (t *Table) generateSchema(b *strings.Builder) {
	b.WriteString("CREATE TABLE `")
	b.WriteString(t.name)
	b.WriteString("` (\n")

	for _, column := range t.columns {
		column.GenerateSchema(b)
	}

	b.WriteString(");\n\n")
}
