package keystone

import (
	"github.com/henrique998/keystone/internal/database"
	"github.com/henrique998/keystone/orm"
)

func NewConnection(args orm.Args) orm.DB {
	return database.New(args)
}
