package goschema

type TableFunc func(table *Table)

type Table struct {
	name    string
	columns []Common
}

const (
	InnoDB         = "InnoDB"
	DefaultCharset = "utf8mb4"
)

func newTable(tableName string) *Table {
	return &Table{
		name: tableName,
	}
}

func (t *Table) Increments(columnName string) Common {
	c := newColumn(t.name, columnName, Int)
	t.columns = append(t.columns, c)
	c.AutoIncrement()

	return c
}

func (t *Table) String(columnName string) Common {
	c := newColumn(t.name, columnName, String)
	t.columns = append(t.columns, c)

	return c
}

// TODO set correct default char size
func (t *Table) Char(columnName string) Common {
	c := newColumn(t.name, columnName, Char).Size(4)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Enum(columnName string, options []string) Common {
	c := NewEnum(t.name, columnName, options)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Binary(columnName string) Common {
	c := newColumn(t.name, columnName, Binary)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Blob(columnName string) Common {
	c := newColumn(t.name, columnName, Binary)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Text(columnName string) Common {
	c := newColumn(t.name, columnName, Text)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Boolean(columnName string) Common {
	c := newColumn(t.name, columnName, Boolean)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Int(columnName string) Common {
	c := newColumn(t.name, columnName, Int)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) TinyInt(columnName string) Common {
	c := newColumn(t.name, columnName, TinyInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) SmallInt(columnName string) Common {
	c := newColumn(t.name, columnName, SmallInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) MediumInt(columnName string) Common {
	c := newColumn(t.name, columnName, MediumInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) BigInt(columnName string) Common {
	c := newColumn(t.name, columnName, BigInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Float(columnName string) Common {
	c := newColumn(t.name, columnName, Float)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Double(columnName string) Common {
	c := newColumn(t.name, columnName, Double)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Decimal(columnName string, precision, scale int) Common {
	c := NewDecimal(t.name, columnName, precision, scale)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Bit(columnName string) Common {
	c := newColumn(t.name, columnName, Bit)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) DateTime(columnName string) Common {
	c := newColumn(t.name, columnName, DateTime)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Timestamp(columnName string) Common {
	c := newColumn(t.name, columnName, Timestamp)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Date(columnName string) Common {
	c := newColumn(t.name, columnName, Date)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Time(columnName string) Common {
	c := newColumn(t.name, columnName, Time)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Year(columnName string) Common {
	c := newColumn(t.name, columnName, Year)
	t.columns = append(t.columns, c)

	return c
}

/**
func (t *Table) generateSchema(b *strings.Builder) bool {

	/**b.WriteString("CREATE TABLE `")
	b.WriteString(t.name)
	b.WriteString("` (\n")

	var hasIndexes bool

	for index, column := range t.columns {
		isLastColumn := index == len(t.columns)-1
		column.GenerateSchema(b, isLastColumn)

		if !hasIndexes {
			if column.IsPrimaryKey() || column.IsUnique() || column.HasReference() {
				hasIndexes = true
			}
		}
	}

	b.WriteString(");\n\n")

	return hasIndexes
}

/**func (t *Table) generateIndexes(b *strings.Builder) {

	/**b.WriteString("ALTER TABLE `")
	b.WriteString(t.name)
	b.WriteString("`")

	for _, col := range t.columns {
		col.GenerateIndex(b)
		fmt.Println(col.name)
	}

	str := strings.TrimSuffix(b.String(), ",")
	b.Reset()
	b.WriteString(str)
	b.WriteString(";\n\n")
}
**/
