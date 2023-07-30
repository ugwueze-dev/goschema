package goschema

import (
	"strconv"
	"strings"
)

type DecimalColumn struct {
	*Column
	precision int
	scale     int
}

func NewDecimal(tableName, name string, precision, scale int) Common {
	return &DecimalColumn{
		Column:    newColumn(tableName, name, Decimal).(*Column),
		precision: precision,
		scale:     scale,
	}
}

func (d *DecimalColumn) writeSizeOrOption(b *strings.Builder) {
	b.WriteString("(")
	b.WriteString(strconv.Itoa(d.precision))
	b.WriteString(", ")
	b.WriteString(strconv.Itoa(d.scale))
	b.WriteString(")")
}
