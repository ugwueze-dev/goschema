package goschema

type Column struct {
	name         string
	dataType     DataType
	nullable     bool
	defaultValue interface{}
	//enumOptions  []interface{}
	size int

	isPrimaryKey bool
	isUnique     bool
	isUnsigned   bool
}

type DataType string

func (d DataType) String() string {
	return string(d)
}

const (
	String DataType = "varchar"
	Char   DataType = "char"
	Binary DataType = "binary"
	Blob   DataType = "blob"
	Text   DataType = "text"
	Enum   DataType = "enum"

	Boolean DataType = "boolean"

	Int       DataType = "int"
	TinyInt   DataType = "tinyint"
	SmallInt  DataType = "smallint"
	MediumInt DataType = "mediumint"
	BigInt    DataType = "bigint"

	Float   DataType = "float"
	Double  DataType = "double"
	Decimal DataType = "decimal"

	Bit DataType = "bit"

	DateTime  DataType = "datetime"
	Timestamp DataType = "timestamp"
	Date      DataType = "date"
	Time      DataType = "time"
	Year      DataType = "year"
)

const (
	CurrentTimestamp = "CURRENT_TIMESTAMP"
)

func newColumn(columnName string, dataType DataType) *Column {
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

/**func (c *Column) setEnumOptions(options []interface{}) {
	c.enumOptions = options
}**/
