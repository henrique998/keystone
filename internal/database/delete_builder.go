package database

import (
	"fmt"

	"github.com/henrique998/keystone/orm"
)

func (b *deleteBuilder) buildSQL() (string, []any, error) {
	tableDef := getTable(&b.db, b.table)

	if len(b.filters) == 0 && !b.allowUnsafe {
		return "", nil, fmt.Errorf("DELETE without WHERE requires .AllowUnsafe()")
	}

	if tableDef != nil && tableDef.softDeleteCol != "" && !b.hardDelete {
		return b.buildSoftDeleteSQL(tableDef)
	}

	return b.buildHardDeleteSQL()
}

func (b *deleteBuilder) buildSoftDeleteSQL(t *tableDefinition) (string, []any, error) {
	var args []any
	argIndex := 1

	where, whereArgs := buildWhereClause(b.filters, &argIndex)
	args = append(args, whereArgs...)

	query := fmt.Sprintf(
		`UPDATE "%s" SET %s = NOW() %s`,
		b.table,
		t.softDeleteCol,
		where,
	)

	return query, args, nil
}

func (b *deleteBuilder) buildHardDeleteSQL() (string, []any, error) {
	var args []any
	argIndex := 1

	query := fmt.Sprintf(`DELETE FROM "%s"`, b.table)

	if len(b.filters) > 0 {
		where, whereArgs := buildWhereClause(b.filters, &argIndex)
		query += " " + where
		args = append(args, whereArgs...)
	}

	return query, args, nil
}

func (b *deleteBuilder) AllowUnsafe() orm.DeleteBuilder {
	b.allowUnsafe = true
	return b
}

func (b *deleteBuilder) Force() orm.DeleteBuilder {
	b.hardDelete = true
	return b
}

func (b *deleteBuilder) Exec() error {
	if len(b.filters) == 0 && !b.allowUnsafe {
		return fmt.Errorf("unsafe DELETE blocked: no WHERE clause. Use .AllowUnsafe() to override")
	}

	query, args, err := b.buildSQL()
	if err != nil {
		return err
	}

	_, err = b.db.conn.Exec(query, args...)
	return err
}

func (b *deleteBuilder) ExecReturningCount() (int64, error) {
	if len(b.filters) == 0 && !b.allowUnsafe {
		return 0, fmt.Errorf("unsafe DELETE blocked: no WHERE clause. Use .AllowUnsafe() to override")
	}

	query, args, err := b.buildSQL()
	if err != nil {
		return 0, err
	}

	res, err := b.db.conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rows, _ := res.RowsAffected()

	return rows, nil
}
