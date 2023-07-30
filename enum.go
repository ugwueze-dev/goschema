package goschema

import (
	"strings"
)

type EnumColumn struct {
	*Column
	options []string
}

func NewEnum(tableName, name string, options []string) Common {
	return &EnumColumn{
		Column:  newColumn(tableName, name, Enum).(*Column),
		options: options,
	}
}

func (e *EnumColumn) writeSizeOrOption(b *strings.Builder) {
	b.WriteString("(")
	for index, option := range e.options {
		b.WriteString("'")
		b.WriteString(option)
		b.WriteString("'")

		if index != len(e.options)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")
}
