package goschema

import (
	"fmt"
)

type constraint string

func (c constraint) String() string {
	return string(c)
}

type reference struct {
	column   *Column
	onUpdate constraint
	onDelete constraint
}

type Column struct {
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

func newColumn(name, tableName string, dataType dataType, setSize bool, size int) *Column {
	return &Column{
		name:      name,
		dataType:  dataType,
		tableName: tableName,
		setSize:   setSize,
		size:      size,
	}
}

// IsPrimaryKey sets the column as the table's primary key.
func (c *Column) IsPrimaryKey() *Column {
	c.isPrimaryKey = true
	c.IsIndex()
	return c
}

// IsIndex sets this column to be an index
func (c *Column) IsIndex() *Column {
	c.key = fmt.Sprintf("%s_%s_key", c.tableName, c.name)
	return c
}

// IsNull sets this column to accept NULL values
func (c *Column) IsNULL() *Column {
	c.nullable = true
	return c
}

// SetDefaultValue sets the column's default value if there's no value passed for this column when inserting
func (c *Column) SetDefaultValue(defaultValue string) *Column {
	c.defaultValue = defaultValue
	return c
}

// Reference sets this colum to reference another column in the database
func (c *Column) Reference(column *Column, onUpdate, onDelete constraint) *Column {
	reference := &reference{
		column:   column,
		onUpdate: onUpdate,
		onDelete: onDelete,
	}
	c.references = append(c.references, reference)
	return c
}
