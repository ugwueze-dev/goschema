package goschema

type Constraint string

func (c Constraint) String() string {
	return string(c)
}

type Reference struct {
	tableName  string
	columnName string
	onUpdate   Constraint
	onDelete   Constraint
}

const (
	Cascade  Constraint = "CASCADE"
	Restrict Constraint = "RESTRICT"
	NoAction Constraint = "NO ACTION"
	SetNull  Constraint = "SET NULL"
)

func newReference(columnName, tableName string) *Reference {
	return &Reference{
		tableName:  tableName,
		columnName: columnName,
		onUpdate:   NoAction,
		onDelete:   NoAction,
	}
}

func (r *Reference) OnUpdate(c Constraint) *Reference {
	r.onUpdate = c
	return r
}

func (r *Reference) OnDelete(c Constraint) *Reference {
	r.onDelete = c
	return r
}
