package goschema

type Table struct {
	name       string
	columns    []*Column
	references []*Reference
}

func newTable(tableName string) *Table {
	return &Table{
		name:    tableName,
		columns: []*Column{},
	}
}

func (t *Table) Increments(columnName string) *Column {
	c := newColumn(columnName, Int).PrimaryKey()
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) String(columnName string) *Column {
	c := newColumn(columnName, String)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Char(columnName string) *Column {
	c := newColumn(columnName, Char)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Binary(columnName string) *Column {
	c := newColumn(columnName, Binary)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Blob(columnName string) *Column {
	c := newColumn(columnName, Binary)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Text(columnName string) *Column {
	c := newColumn(columnName, Text)
	t.columns = append(t.columns, c)

	return c
}

// TODO handle enum options
func (t *Table) Enum(columnName string, options []interface{}) *Column {
	c := newColumn(columnName, Enum)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Boolean(columnName string) *Column {
	c := newColumn(columnName, Boolean)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Int(columnName string) *Column {
	c := newColumn(columnName, Int)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) TinyInt(columnName string) *Column {
	c := newColumn(columnName, TinyInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) SmallInt(columnName string) *Column {
	c := newColumn(columnName, SmallInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) MediumInt(columnName string) *Column {
	c := newColumn(columnName, MediumInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) BigInt(columnName string) *Column {
	c := newColumn(columnName, BigInt)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Float(columnName string) *Column {
	c := newColumn(columnName, Float)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Double(columnName string) *Column {
	c := newColumn(columnName, Double)
	t.columns = append(t.columns, c)

	return c
}

// TODO Handle decimal
func (t *Table) Decimal(columnName string, precision, scale int) *Column {
	c := newColumn(columnName, Decimal)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Bit(columnName string) *Column {
	c := newColumn(columnName, Bit)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) DateTime(columnName string) *Column {
	c := newColumn(columnName, DateTime)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Timestamp(columnName string) *Column {
	c := newColumn(columnName, Timestamp)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Date(columnName string) *Column {
	c := newColumn(columnName, Date)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Time(columnName string) *Column {
	c := newColumn(columnName, Time)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Year(columnName string) *Column {
	c := newColumn(columnName, Year)
	t.columns = append(t.columns, c)

	return c
}

func (t *Table) Foriegn(columnName string) *Reference {
	ref := newReference(t.name, columnName)
	t.references = append(t.references, ref)

	return t.references[len(t.references)-1]
}
