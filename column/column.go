package column

import "strings"

type Common interface {
	Size(size uint8) Common
	PrimaryKey() Common
	Unsigned() Common
	Unique() Common
	Index() Common // should user specify an index key/name?
	AutoIncrement() Common
	Default(defaultValue interface{}) Common
	buildQuery(*strings.Builder) string
}

type Column struct {
	name          string
	dataType      DataType
	size          uint8
	defaultValue  interface{} // change this to use []byte
	unsigned      bool
	primaryKey    bool
	unique        bool
	index         bool
	autoIncrement bool
}

func New(name string, dataType DataType) Common {
	return newColumn(name, dataType)
}

func newColumn(name string, dataType DataType) Common {
	return &Column{
		name:     name,
		dataType: dataType,
		size:     0,
	}
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
	return c
}

// TODO change this to use []byte instead of interface{}
func (c *Column) Default(defaultValue interface{}) Common {
	c.defaultValue = defaultValue
	return c
}

func (c *Column) buildQuery(b *strings.Builder) string {
	return b.String()
}
