package database

import (
	"database/sql"

	"github.com/henrique998/keystone/orm"
)

func (db *db) QueryRaw(query string, args ...any) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *db) FindMany(table string, filters ...any) orm.QueryBuilder {
	return nil
}

func (db *db) FindOne(table string, filters ...any) orm.QueryBuilder {
	return nil
}

func (db *db) Create(table string, data interface{}) error {
	return nil
}

func (db *db) Update(table string, filters ...any) orm.UpdateBuilder {
	return nil
}

func (db *db) Delete(table string, filters ...any) orm.DeleteBuilder {
	return nil
}

func (db *db) DeleteBatch(table string, filters ...any) orm.DeleteBuilder {
	return nil
}
