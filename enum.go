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

func (e *EnumColumn) GenerateSchema(b *strings.Builder, isLastColumn bool) {
	optionsWriterFunc := func(b *strings.Builder) {
		lastOptionIndex := len(e.options) - 1
		b.WriteString("(")

		for idx, option := range e.options {
			b.WriteString("'")
			b.WriteString(option)
			b.WriteString("'")

			if idx != lastOptionIndex {
				b.WriteString(", ")
			}
		}

		b.WriteString(")")
	}
	e.Column.generateColumnSchema(b, isLastColumn, optionsWriterFunc)
}
