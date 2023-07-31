package goschema

type TableFunc func(table *Table)

// Table represents a database table
// Each table will hold a slice of structs representing every column in it
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

// String creates a new Varchar column on the table
func (t *Table) String(columnName string) Common {
	c := newColumn(t.name, columnName, String)
	t.columns = append(t.columns, c)

	return c
}

// Char creates a new char column on the table
func (t *Table) Char(columnName string) Common {
	c := newColumn(t.name, columnName, Char).Size(4)
	t.columns = append(t.columns, c)

	return c
}

// Enum creates a new Enum column on the table.
// the allowed options are passed in as a string slice as the second argument,
func (t *Table) Enum(columnName string, options []string) Common {
	c := NewEnum(t.name, columnName, options)
	t.columns = append(t.columns, c)

	return c
}

// Binary creates a new Binary column on the table
func (t *Table) Binary(columnName string) Common {
	c := newColumn(t.name, columnName, Binary)
	t.columns = append(t.columns, c)

	return c
}

// Blob creates a new Varchar column on the table
func (t *Table) Blob(columnName string) Common {
	c := newColumn(t.name, columnName, Binary)
	t.columns = append(t.columns, c)

	return c
}

// Text creates a new text column on the table
func (t *Table) Text(columnName string) Common {
	c := newColumn(t.name, columnName, Text)
	t.columns = append(t.columns, c)

	return c
}

// Boolean creates a new boolean column on the table
func (t *Table) Boolean(columnName string) Common {
	c := newColumn(t.name, columnName, Boolean)
	t.columns = append(t.columns, c)

	return c
}

// Int creates a new interger column on the table
func (t *Table) Int(columnName string) Common {
	c := newColumn(t.name, columnName, Int)
	t.columns = append(t.columns, c)

	return c
}

// TinyInt creates a new tiny integer column on the table
func (t *Table) TinyInt(columnName string) Common {
	c := newColumn(t.name, columnName, TinyInt)
	t.columns = append(t.columns, c)

	return c
}

// SmallInt creates a small integer column on the table
func (t *Table) SmallInt(columnName string) Common {
	c := newColumn(t.name, columnName, SmallInt)
	t.columns = append(t.columns, c)

	return c
}

// MediumInt creates a new medium integer column on the table
func (t *Table) MediumInt(columnName string) Common {
	c := newColumn(t.name, columnName, MediumInt)
	t.columns = append(t.columns, c)

	return c
}

// BigInt creates a new big integer column on the table
func (t *Table) BigInt(columnName string) Common {
	c := newColumn(t.name, columnName, BigInt)
	t.columns = append(t.columns, c)

	return c
}

// Float creates a new float column on the table
func (t *Table) Float(columnName string) Common {
	c := newColumn(t.name, columnName, Float)
	t.columns = append(t.columns, c)

	return c
}

// Double creates a new double column on the table
func (t *Table) Double(columnName string) Common {
	c := newColumn(t.name, columnName, Double)
	t.columns = append(t.columns, c)

	return c
}

// Decimal creates a new decimal column on the table
// precision and scale is required as arguments
func (t *Table) Decimal(columnName string, precision, scale int) Common {
	c := NewDecimal(t.name, columnName, precision, scale)
	t.columns = append(t.columns, c)

	return c
}

// Bit creates a new bit column on the table
func (t *Table) Bit(columnName string) Common {
	c := newColumn(t.name, columnName, Bit)
	t.columns = append(t.columns, c)

	return c
}

// DateTime creates a new datetime column on the table
func (t *Table) DateTime(columnName string) Common {
	c := newColumn(t.name, columnName, DateTime).Nullable()
	t.columns = append(t.columns, c)

	return c
}

// Timestamp creates a new timestamp column on the table
func (t *Table) Timestamp(columnName string) Common {
	c := newColumn(t.name, columnName, Timestamp).Nullable()
	t.columns = append(t.columns, c)

	return c
}

// Date creates a new date column on the table
func (t *Table) Date(columnName string) Common {
	c := newColumn(t.name, columnName, Date).Nullable()
	t.columns = append(t.columns, c)

	return c
}

// Time creates a new time column on the table
func (t *Table) Time(columnName string) Common {
	c := newColumn(t.name, columnName, Time)
	t.columns = append(t.columns, c)

	return c
}

// Year creates a new year column on the table
func (t *Table) Year(columnName string) Common {
	c := newColumn(t.name, columnName, Year)
	t.columns = append(t.columns, c)

	return c
}
