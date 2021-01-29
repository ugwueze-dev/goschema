package goschema

type Table struct {
	name    string
	columns []*Column
}

type dataType string

func (d dataType) String() string {
	return string(d)
}

const (
	Int        dataType = "INT"
	tinyInt             = "TINYINT"
	smallInt            = "SMALLINT"
	mediumInt           = "MEDIUMINT"
	bigInt              = "BIGINT"
	float               = "FLOAT"
	double              = "DOUBLE"
	decimal             = "DECIMAL"
	char                = "CHAR"
	varChar             = "VARCHAR"
	blob                = "BLOB"
	tinyBlob            = "TINYBLOB"
	mediumBlob          = "MEDIUMBLOB"
	longBlob            = "LONGBLOB"
	text                = "TEXT"
	tinyText            = "TINYTEXT"
	mediumText          = "MEDIUMTEXT"
	longText            = "LONGTEXT"
	enum                = "ENUM"
	date                = "DATE"
	dateTime            = "DATETIME"
	timestamp           = "TIMESTAMP"
	time                = "time"
	year                = "year"
)

func newTable(tableName string) *Table {
	return &Table{
		name: tableName,
	}
}

// add a new column to the current table
func (t *Table) newColumn(columnName string, dataType dataType, setSize bool, size int) *Column {
	column := newColumn(columnName, t.name, dataType, setSize, size)
	t.columns = append(t.columns, column)

	return column
}

// Int represents a database column of datattype INT
func (t *Table) Int(columnName string, size int) *Column {
	return t.newColumn(columnName, Int, true, size)
}

// TinyInt represents a database column of datattype TINYINT
func (t *Table) TinyInt(columnName string, size int) *Column {
	return t.newColumn(columnName, tinyInt, true, size)
}

// MediumInt represents a database column of datattype MEDIUMINT
func (t *Table) MediumInt(columnName string, size int) *Column {
	return t.newColumn(columnName, mediumInt, true, size)
}

// BigInt represents a database column of datattype BIGINT
func (t *Table) BigInt(columnName string, size int) *Column {
	return t.newColumn(columnName, bigInt, true, size)
}

// Float represents a database column of datattype FLOAT
func (t *Table) Float(columnName string, length, numDecimals int) *Column {
	column := t.newColumn(columnName, float, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

// Double represents a database column of datattype DOUBLE
func (t *Table) Double(columnName string, length, numDecimals int) *Column {
	column := t.newColumn(columnName, double, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

// Decimal represents a database column of datattype DECIMAL
func (t *Table) Decimal(columnName string, length, numDecimals int) *Column {
	column := t.newColumn(columnName, decimal, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

// Char represents a database column of datattype CHAR
func (t *Table) Char(columnName string, size uint) *Column {
	return t.newColumn(columnName, char, true, int(size))
}

// Varchar represents a database column of datattype VARCHAR
func (t *Table) Varchar(columnName string, size uint) *Column {
	return t.newColumn(columnName, varChar, true, int(size))
}

// Blog represents a database column of datattype BLOB
func (t *Table) Blob(columnName string, size uint) *Column {
	return t.newColumn(columnName, blob, true, int(size))
}

// TinyBlob represents a database column of datattype TINYBLOB
func (t *Table) TinyBlob(columnName string) *Column {
	return t.newColumn(columnName, tinyBlob, false, 0)
}

// MediumBlob represents a database column of datattype MEDIUMBLOB
func (t *Table) MediumBlob(columnName string) *Column {
	return t.newColumn(columnName, mediumBlob, false, 0)
}

// LongBlob represents a database column of datattype LONGBLOB
func (t *Table) LongBlob(columnName string) *Column {
	return t.newColumn(columnName, longBlob, false, 0)
}

// Text represents a database column of datattype TEXT
func (t *Table) Text(columnName string, size uint) *Column {
	return t.newColumn(columnName, text, true, int(size))
}

// TinyText represents a database column of datattype TINYTEXT
func (t *Table) TinyText(columnName string) *Column {
	return t.newColumn(columnName, tinyText, false, 0)
}

// MediumText represents a database column of datattype MEDIUMTEXT
func (t *Table) MediumText(columnName string) *Column {
	return t.newColumn(columnName, mediumText, false, 0)
}

// LongText represents a database column of datattype LONGTEXT
func (t *Table) LongText(columnName string) *Column {
	return t.newColumn(columnName, longText, false, 0)
}

// Enum represents a database column of datattype ENUM
func (t *Table) Enum(columnName string, variants []string) *Column {
	return t.newColumn(columnName, enum, false, 0)
}

// Date represents a database column of datattype DATE
func (t *Table) Date(columnName string) *Column {
	return t.newColumn(columnName, date, false, 0)
}

// DateTime represents a database column of datattype DATETIME
func (t *Table) DateTime(columnName string) *Column {
	return t.newColumn(columnName, dateTime, false, 0)
}

// Timestamp represents a database column of datattype TIMESTAMP
func (t *Table) Timestamp(columnName string) *Column {
	return t.newColumn(columnName, timestamp, false, 0)
}

// Time represents a database column of datattype TIME
func (t *Table) Time(columnName string) *Column {
	return t.newColumn(columnName, time, false, 0)
}

// Year represents a database column of datattype YEAR
func (t *Table) Year(columnName string) *Column {
	return t.newColumn(columnName, year, false, 0)
}
