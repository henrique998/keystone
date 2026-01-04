package orm

type Table interface {
	UseSoftDelete(column string)
}
