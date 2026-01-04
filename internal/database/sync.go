package database

import "github.com/henrique998/keystone/orm"

func (db *db) NewTable(name string, fn func(orm.TableBuilder)) orm.TableDefinition {
	tb := tableBuilder{
		tableName: name,
		columns:   make(map[string]*columnInfo),
	}

	fn(&tb)

	def := &tableDefinition{
		name:          tb.tableName,
		columns:       tb.columns,
		softDeleteCol: "",
		belongsTo:     tb.belongsTo,
	}

	db.tables[name] = def

	return def
}

func (db *db) SyncSchemas() error {
	return nil
}
