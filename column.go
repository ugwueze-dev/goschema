package goschema

import (
	"fmt"
	"strings"
)

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
	IsPrimaryKey() bool
	IsUnique() bool
	HasReferences() bool
	GetReferences() []*Reference
	GenerateSchema(*strings.Builder, bool)
	GenerateIndex(*strings.Builder)
}

type Column struct {
	tableName     string
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
	references    []*Reference
}

func newColumn(tableName, name string, dataType DataType) Common {
	return &Column{
		tableName:  tableName,
		name:       name,
		dataType:   dataType,
		size:       0,
		nullStatus: "NOT NULL",
	}
}

func (c *Column) IsPrimaryKey() bool {
	return c.primaryKey
}

func (c *Column) IsUnique() bool {
	return c.unique
}

func (c *Column) GetReferences() []*Reference {
	return c.references
}

func (c *Column) HasReferences() bool {
	return len(c.references) > 0
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
func (c *Column) Default(defaultValue string) Common {
	c.defaultValue = defaultValue
	return c
}

func (c *Column) Nullable() Common {
	c.nullStatus = "NULL"
	return c
}

func (c *Column) References(columnName, tableName string) Common {
	ref := newReference(columnName, tableName)
	c.references = append(c.references, ref)

	return c
}

func (c *Column) GenerateSchema(b *strings.Builder, isLastColumn bool) {
	sizeOrOptionsWriterFunc := func(b *strings.Builder) {
		// add size if set
		if c.size > 0 {
			b.WriteString(fmt.Sprintf("(%d)", c.size))
		}
	}
	c.generateColumnSchema(b, isLastColumn, sizeOrOptionsWriterFunc)
}

func (c *Column) generateColumnSchema(b *strings.Builder, isLastColumn bool, sw func(b *strings.Builder)) {
	// column name
	b.WriteString("  `")
	b.WriteString(c.name)
	b.WriteString("` ")

	// datatype
	b.WriteString(c.dataType.String())

	// write size or enum options or decimal precision and scale
	sw(b)

	// add null status
	b.WriteString(" ")
	b.WriteString(c.nullStatus)

	// set default value if it exists
	/**if c.defaultValue != "" {
		b.WriteString(" DEFAULT `")
		b.WriteString(c.defaultValue)
		b.WriteString("`")
	}**/

	// only add comma (,) if it's not last column
	if !isLastColumn {
		b.WriteString(",")
	}
	b.WriteString("\n")
}

func (c *Column) GenerateIndex(b *strings.Builder) {
	if c.IsPrimaryKey() {
		b.WriteString("\n  ADD PRIMARY KEY (`")
		b.WriteString(c.name)
		b.WriteString("`),")
	}

	if c.IsUnique() {
		b.WriteString("\n  ADD UNIQUE KEY `")
		b.WriteString(c.tableName)
		b.WriteString("_")
		b.WriteString(c.name)
		b.WriteString("_unique`")
		b.WriteString(" (`")
		b.WriteString(c.name)
		b.WriteString("`),")
	}

	references := c.GetReferences()
	for _, reference := range references {
		b.WriteString("\n  ADD CONSTRAINT `fk_")
		b.WriteString(c.name)
		b.WriteString("` FOREIGN KEY (`")
		b.WriteString(c.name)
		b.WriteString("`) REFERENCES `")
		b.WriteString(reference.tableName)
		b.WriteString("`(`")
		b.WriteString(reference.columnName)
		b.WriteString("`),")
	}
}
