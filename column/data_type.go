package column

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
