package column

import "strings"

type EnumColumn struct {
	column  Common
	options []interface{}
}

func NewEnum(name string, options []interface{}) Common {
	return &EnumColumn{
		column:  newColumn(name, Enum),
		options: options,
	}
}

func (e *EnumColumn) Size(size uint8) Common {
	e.column.Size(size)
	return e
}

func (e *EnumColumn) PrimaryKey() Common {
	e.column.PrimaryKey()
	return e
}

func (e *EnumColumn) Unsigned() Common {
	e.column.Unsigned()
	return e
}

func (e *EnumColumn) Unique() Common {
	e.column.Unique()
	return e
}

// should the user specify index name?
func (e *EnumColumn) Index() Common {
	e.column.Index()
	return e
}

func (e *EnumColumn) AutoIncrement() Common {
	e.column.AutoIncrement()
	return e
}

// TODO change this to use []byte instead of interface{}
func (e *EnumColumn) Default(defaultValue interface{}) Common {
	e.column.Default(defaultValue)
	return e
}

func (e *EnumColumn) GenerateSchema(b *strings.Builder) string {
	return b.String()
}
