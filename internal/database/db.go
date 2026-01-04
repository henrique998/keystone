package database

import (
	"database/sql"
	"fmt"

	"github.com/henrique998/keystone/orm"
)

type db struct {
	conn   *sql.DB
	tables map[string]orm.TableDefinition // TODO: replace with map[string]*tableDefinition
}

func New(args orm.Args) orm.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		args.Host,
		args.User,
		args.Password,
		args.DBName,
		args.Port,
	)

	conn, err := sql.Open("pgx", connStr)
	if err != nil {
		panic(err)
	}

	return &db{
		conn:   conn,
		tables: make(map[string]orm.TableDefinition),
	}
}
