package goschema

import (
	"strconv"
	"strings"
)

type commonData struct {
	name          string
	dataType      DataType
	size          uint8
	defaultValue  string // change this to use []byte
	nullStatus    string
	unsigned      bool
	primaryKey    bool
	unique        bool
	index         bool
	autoIncrement bool
	reference     *Reference
}

type Common interface {
	Size(size uint8) Common
	PrimaryKey() Common
	Unsigned() Common
	Unique() Common
	Index() Common // should user specify an index key/name?
	AutoIncrement() Common
	Default(defaultValue string) Common
	Nullable() Common
	References(columnName, tableName string) Common
	getCommonData() commonData
	writeSizeOrOption(b *strings.Builder)
}

type Column struct {
	commonData
}

func newColumn(tableName, name string, dataType DataType) Common {
	c := &Column{}
	c.dataType = dataType
	c.nullStatus = "NOT NULL"
	c.name = name

	return c
}

func (c *Column) Size(size uint8) Common {
	c.size = size
	return c
}

func (c *Column) PrimaryKey() Common {
	c.primaryKey = true
	return c
}

func (c *Column) Unsigned() Common {
	c.unsigned = true
	return c
}

func (c *Column) Unique() Common {
	c.unique = true
	return c
}

// should the user specify index name?
func (c *Column) Index() Common {
	c.index = true
	return c
}

func (c *Column) AutoIncrement() Common {
	c.autoIncrement = true
	c.PrimaryKey()

	return c
}

// TODO change this to use []byte instead of interface{}
func (c *Column) Default(defaultValue string) Common {
	c.defaultValue = defaultValue
	return c
}

func (c *Column) Nullable() Common {
	c.nullStatus = "NULL"
	return c
}

func (c *Column) getCommonData() commonData {
	return c.commonData
}

func (c *Column) writeSizeOrOption(b *strings.Builder) {
	if c.size > 0 {
		//b.WriteString(fmt.Sprintf("(%d)", c.size))
		b.WriteString("(")
		b.WriteString(strconv.Itoa(int(c.size)))
		b.WriteString(")")
	}
}

func (c *Column) References(columnName, tableName string) Common {
	c.reference = newReference(columnName, tableName)

	return c
}
