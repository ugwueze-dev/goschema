package goschema

type Constraint string

func (c Constraint) String() string {
	return string(c)
}

type Reference struct {
	// tableName is the name of the referenced table
	tableName string
	// columnName is the name of the column on the reference table we are referencing
	columnName string
	// onUpdate defines the action to perform when the parent column is updated
	onUpdate Constraint
	// onUpdate defines the action to perform when the parent column is deleted
	onDelete Constraint
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
