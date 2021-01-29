package goschema

import (
	"fmt"
)

type constraint string

func (c constraint) String() string {
	return string(c)
}

type reference struct {
	column   *column
	onUpdate constraint
	onDelete constraint
}

type column struct {
	name         string
	tableName    string
	dataType     dataType
	nullable     bool
	setSize      bool
	size         int
	length       int
	numDecimals  int
	defaultValue string
	isPrimaryKey bool
	key          string

	references []*reference
}

const (
	Cascade  constraint = "CASCADE"
	Restrict            = "RESTRICT"
	SetNull             = "SET NULL"
	NoAction            = "NO ACTION"
)

func newColumn(name, tableName string, dataType dataType, setSize bool, size int) *column {
	return &column{
		name:      name,
		dataType:  dataType,
		tableName: tableName,
		setSize:   setSize,
		size:      size,
	}
}

// IsPrimaryKey sets the column as the table's primary key. 
func (c *column) IsPrimaryKey() *column {
	c.isPrimaryKey = true
	c.IsIndex()
	return c
}

// IsIndex sets this column to be an index
func (c *column) IsIndex() *column {
	c.key = fmt.Sprintf("%s_%s_key", c.tableName, c.name)
	return c
}

// IsNull sets this column to accept NULL values
func (c *column) IsNULL() *column {
	c.nullable = true
	return c
}

// SetDefaultValue sets the column's default value if there's no value passed for this column when inserting
func (c *column) SetDefaultValue(defaultValue string) *column {
	c.defaultValue = defaultValue
	return c
}

// Reference sets this colum to reference another column in the database
func (c *column) Reference(column *column, onUpdate, onDelete constraint) *column {
	reference := &reference{
		column:   column,
		onUpdate: onUpdate,
		onDelete: onDelete,
	}
	c.references = append(c.references, reference)
	return c
}
