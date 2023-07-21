package goschema

type Column struct {
	name         string
	dataType     string
	nullable     bool
	defaultValue interface{}
	size         int

	isPrimaryKey bool
	isUnique     bool
	isUnsigned   bool
}

const (
	String = "varchar"
	Char   = "char"
	Binary = "binary"
	Blob   = "blob"
	Text   = "text"
	Enum   = "enum"

	Boolean = "boolean"

	Int       = "int"
	TinyInt   = "tinyint"
	SmallInt  = "smallint"
	MediumInt = "mediumint"
	BigInt    = "bigint"

	Float   = "float"
	Double  = "double"
	Decimal = "decimal"

	Bit = "bit"

	DateTime  = "datetime"
	Timestamp = "timestamp"
	Date      = "date"
	Time      = "time"
	Year      = "year"
)

const (
	CurrentTimestamp = "CURRENT_TIMESTAMP"
)

func newColumn(columnName string, dataType string) *Column {
	return &Column{
		name:     columnName,
		dataType: dataType,
		size:     -1,
	}
}

func (c *Column) Size(size int) *Column {
	c.size = size
	return c
}

func (c *Column) PrimaryKey() *Column {
	c.isPrimaryKey = true
	return c
}

func (c *Column) Unique() *Column {
	c.isUnique = true
	return c
}

func (c *Column) Nullable() *Column {
	c.nullable = true
	return c
}

func (c *Column) Default(value interface{}) *Column {
	c.defaultValue = value
	return c
}

func (c *Column) Unsigned() *Column {
	c.isUnsigned = true
	return c
}
