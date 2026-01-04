package orm

type TableDefinition interface {
	UseSoftDelete(column string)
}
