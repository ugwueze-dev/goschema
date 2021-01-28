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

func (t *table) newColumn(columnName string, dataType dataType, setSize bool, size int) *column {
	column := newColumn(columnName, t.name, dataType, setSize, size)
	t.columns = append(t.columns, column)

	return column
}

func (t *table) Int(columnName string, size int) *column {
	return t.newColumn(columnName, Int, true, size)
}

func (t *table) TinyInt(columnName string, size int) *column {
	return t.newColumn(columnName, tinyInt, true, size)
}

func (t *table) MediumInt(columnName string, size int) *column {
	return t.newColumn(columnName, mediumInt, true, size)
}

func (t *table) BigInt(columnName string, size int) *column {
	return t.newColumn(columnName, bigInt, true, size)
}

func (t *table) Float(columnName string, length, numDecimals int) *column {
	column := t.newColumn(columnName, float, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

func (t *table) Double(columnName string, length, numDecimals int) *column {
	column := t.newColumn(columnName, double, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

func (t *table) Decimal(columnName string, length, numDecimals int) *column {
	column := t.newColumn(columnName, decimal, false, 0)
	column.length = length
	column.numDecimals = numDecimals

	return column
}

func (t *table) Char(columnName string, size uint) *column {
	return t.newColumn(columnName, char, true, int(size))
}

func (t *table) Varchar(columnName string, size uint) *column {
	return t.newColumn(columnName, varChar, true, int(size))
}

func (t *table) Blob(columnName string, size uint) *column {
	return t.newColumn(columnName, blob, true, int(size))
}

func (t *table) TinyBlob(columnName string) *column {
	return t.newColumn(columnName, tinyBlob, false, 0)
}

func (t *table) MediumBlob(columnName string) *column {
	return t.newColumn(columnName, mediumBlob, false, 0)
}

func (t *table) LongBlob(columnName string) *column {
	return t.newColumn(columnName, longBlob, false, 0)
}

func (t *table) Text(columnName string, size uint) *column {
	return t.newColumn(columnName, text, true, int(size))
}

func (t *table) TinyText(columnName string) *column {
	return t.newColumn(columnName, tinyText, false, 0)
}

func (t *table) MediumText(columnName string) *column {
	return t.newColumn(columnName, mediumText, false, 0)
}

func (t *table) LongText(columnName string) *column {
	return t.newColumn(columnName, longText, false, 0)
}

func (t *table) Enum(columnName string, variants []string) *column {
	return t.newColumn(columnName, enum, false, 0)
}

func (t *table) Date(columnName string) *column {
	return t.newColumn(columnName, date, false, 0)
}

func (t *table) DateTime(columnName string) *column {
	return t.newColumn(columnName, dateTime, false, 0)
}

func (t *table) Timestamp(columnName string) *column {
	return t.newColumn(columnName, timestamp, false, 0)
}

func (t *table) Time(columnName string) *column {
	return t.newColumn(columnName, time, false, 0)
}

func (t *table) Year(columnName string) *column {
	return t.newColumn(columnName, year, false, 0)
}
