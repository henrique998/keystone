package database

import "github.com/henrique998/keystone/orm"

func getTable(db *db, tableName string) orm.TableDefinition {
	return db.tables[tableName]
}
