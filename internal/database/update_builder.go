package database

import (
	"fmt"
	"strings"

	"github.com/henrique998/keystone/orm"
)

func (ub *updateBuilder) Set(column string, value any) orm.UpdateBuilder {
	ub.sets[column] = value

	return ub
}

func (ub *updateBuilder) SetMap(values map[string]any) orm.UpdateBuilder {
	for k, v := range values {
		ub.sets[k] = v
	}

	return ub
}

func (ub *updateBuilder) buildSQL() (string, []any) {
	if len(ub.sets) == 0 {
		panic("no fields specified for update (use Set or SetMap before Exec)")
	}

	var (
		setClauses []string
		args       []any
	)
	argIndex := 1

	for col, val := range ub.sets {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, argIndex))
		args = append(args, val)
		argIndex++
	}

	query := fmt.Sprintf("UPDATE %s SET %s", ub.table, strings.Join(setClauses, ", "))

	softCol := ""
	if ub.tableDef != nil {
		softCol = ub.tableDef.softDeleteCol
	}

	whereClause := ""
	var whereArgs []any

	if len(ub.filters) > 0 {
		whereClause, whereArgs = buildWhereClause(ub.filters, &argIndex)
	}

	if softCol != "" {
		if whereClause == "" {
			whereClause = fmt.Sprintf(`WHERE "%s" IS NULL`, softCol)
		} else {
			whereClause += fmt.Sprintf(` AND "%s" IS NULL`, softCol)
		}
	}

	if whereClause != "" {
		query += " " + whereClause
		args = append(args, whereArgs...)
	}

	return query, args
}

func (ub *updateBuilder) Exec() error {
	query, args := ub.buildSQL()
	_, err := ub.db.conn.Exec(query, args...)

	return err
}
