package orm

type Column interface {
	Name() string
	Table() string
	SQLType() string
}

type ForeignKeyDef interface {
	RefTable() string
	RefColumn() string
}

type ColInfo interface {
	Name() string
	SQLType() string
	Constraints() []string
	ForeignKey() ForeignKeyDef
}
