package database

import (
	"fmt"

	"github.com/henrique998/keystone/orm"
)

func (b *deleteBatchBuilder) AllowUnsafe() orm.DeleteBatchBuilder {
	b.allowUnsafe = true
	return b
}

func (b *deleteBatchBuilder) HardDelete() orm.DeleteBatchBuilder {
	b.hardDelete = true
	return b
}

func (b *deleteBatchBuilder) Exec() error {
	tableDef := getTable(&b.db, b.table)
	useSoftDelete := tableDef != nil && tableDef.softDeleteCol != "" && !b.hardDelete

	if len(b.filters) == 0 && !b.allowUnsafe {
		return fmt.Errorf("unsafe DeleteBatch(): no filters provided and AllowUnsafe() not enabled")
	}

	for _, f := range b.filters {
		err := b.execSingle(f, useSoftDelete)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *deleteBatchBuilder) execSingle(filter any, useSoftDelete bool) error {
	var args []any
	argIndex := 1

	where, whereArgs := buildWhereClause([]any{filter}, &argIndex)
	args = append(args, whereArgs...)

	var query string

	if useSoftDelete {
		table := b.table
		softCol := getTable(&b.db, b.table).softDeleteCol

		query = fmt.Sprintf(
			`UPDATE %s SET %s = NOW() %s`,
			table,
			softCol,
			where,
		)

	} else {
		// Hard delete
		table := b.table
		query = fmt.Sprintf(`DELETE FROM %s %s`, table, where)
	}

	_, err := b.db.conn.Exec(query, args...)
	return err
}
