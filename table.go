package goschema

type table struct {
	name    string
	columns []*column
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

func newTable(tableName string) *table {
	return &table{
		name: tableName,
	}
}

// add a new column to the current table
func (t *table) newColumn(columnName string, dataType dataType, setSize bool, size int) *column {
	column := newColumn(columnName, t.name, dataType, setSize, size)
	t.columns = append(t.columns, column)

	return column
}

// Int represents a database column of datattype INT
func (t *table) Int(columnName string, size int) *column {
	return t.newColumn(columnName, Int, true, size)
}

// TinyInt represents a database column of datattype TINYINT
func (t *table) TinyInt(columnName string, size int) *column {
	return t.newColumn(columnName, tinyInt, true, size)
}

// MediumInt represents a database column of datattype MEDIUMINT
func (t *table) MediumInt(columnName string, size int) *column {
	return t.newColumn(columnName, mediumInt, true, size)
}

// BigInt represents a database column of datattype BIGINT
func (t *table) BigInt(columnName string, size int) *column {
	return t.newColumn(columnName, bigInt, true, size)
}

// Float represents a database column of datattype FLOAT
func (t *table) Float(columnName string, length, numDecimals int) *column {
	column := t.newColumn(columnName, float, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

// Double represents a database column of datattype DOUBLE
func (t *table) Double(columnName string, length, numDecimals int) *column {
	column := t.newColumn(columnName, double, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

// Decimal represents a database column of datattype DECIMAL
func (t *table) Decimal(columnName string, length, numDecimals int) *column {
	column := t.newColumn(columnName, decimal, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

// Char represents a database column of datattype CHAR
func (t *table) Char(columnName string, size uint) *column {
	return t.newColumn(columnName, char, true, int(size))
}

// Varchar represents a database column of datattype VARCHAR
func (t *table) Varchar(columnName string, size uint) *column {
	return t.newColumn(columnName, varChar, true, int(size))
}

// Blog represents a database column of datattype BLOB
func (t *table) Blob(columnName string, size uint) *column {
	return t.newColumn(columnName, blob, true, int(size))
}

// TinyBlob represents a database column of datattype TINYBLOB
func (t *table) TinyBlob(columnName string) *column {
	return t.newColumn(columnName, tinyBlob, false, 0)
}

// MediumBlob represents a database column of datattype MEDIUMBLOB
func (t *table) MediumBlob(columnName string) *column {
	return t.newColumn(columnName, mediumBlob, false, 0)
}

// LongBlob represents a database column of datattype LONGBLOB
func (t *table) LongBlob(columnName string) *column {
	return t.newColumn(columnName, longBlob, false, 0)
}

// Text represents a database column of datattype TEXT
func (t *table) Text(columnName string, size uint) *column {
	return t.newColumn(columnName, text, true, int(size))
}

// TinyText represents a database column of datattype TINYTEXT
func (t *table) TinyText(columnName string) *column {
	return t.newColumn(columnName, tinyText, false, 0)
}

// MediumText represents a database column of datattype MEDIUMTEXT
func (t *table) MediumText(columnName string) *column {
	return t.newColumn(columnName, mediumText, false, 0)
}

// LongText represents a database column of datattype LONGTEXT
func (t *table) LongText(columnName string) *column {
	return t.newColumn(columnName, longText, false, 0)
}

// Enum represents a database column of datattype ENUM
func (t *table) Enum(columnName string, variants []string) *column {
	return t.newColumn(columnName, enum, false, 0)
}

// Date represents a database column of datattype DATE
func (t *table) Date(columnName string) *column {
	return t.newColumn(columnName, date, false, 0)
}

// DateTime represents a database column of datattype DATETIME
func (t *table) DateTime(columnName string) *column {
	return t.newColumn(columnName, dateTime, false, 0)
}

// Timestamp represents a database column of datattype TIMESTAMP
func (t *table) Timestamp(columnName string) *column {
	return t.newColumn(columnName, timestamp, false, 0)
}

// Time represents a database column of datattype TIME
func (t *table) Time(columnName string) *column {
	return t.newColumn(columnName, time, false, 0)
}

// Year represents a database column of datattype YEAR
func (t *table) Year(columnName string) *column {
	return t.newColumn(columnName, year, false, 0)
}
