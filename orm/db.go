package orm

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5"
)

type DB interface {
	QueryRaw(query string, args ...any) (*sql.Rows, error)
	FindMany(table string, filters ...any) QueryBuilder
	FindOne(table string, filters ...any) QueryBuilder
	Create(table string, data interface{}) error
	Update(table string, filters ...any) UpdateBuilder
	Delete(table string, filters ...any) DeleteBuilder
	DeleteBatch(table string, filters ...any) DeleteBatchBuilder
	NewTable(name string, fn func(TableBuilder)) Table
	SyncSchemas() error
	Close()
}

type Credentials struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type Args struct {
	Credentials
	Dialect string
}
