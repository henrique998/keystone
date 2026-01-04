package database

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/henrique998/keystone/orm"
)

func (qb *queryBuilder) Select(columns ...string) orm.QueryBuilder {
	if len(columns) > 0 {
		qb.columns = columns
	}

	return qb
}

func (qb *queryBuilder) OrderBy(column string, direction string) orm.QueryBuilder {
	qb.orderBy = fmt.Sprintf("%s %s", column, strings.ToUpper(direction))
	return qb
}

func (qb *queryBuilder) Limit(limit int) orm.QueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *queryBuilder) Offset(offset int) orm.QueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *queryBuilder) Exec(dest any) error {
	query, args := qb.buildSQL()

	rows, err := qb.db.conn.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr {
		return fmt.Errorf("dest must be a pointer")
	}

	destVal = destVal.Elem()

	isSlice := destVal.Kind() == reflect.Slice
	var elemType reflect.Type

	if isSlice {
		elemType = destVal.Type().Elem()
	} else {
		elemType = destVal.Type()
	}

	for rows.Next() {
		// Criar nova instância do elemento
		elem := reflect.New(elemType).Elem()

		// Mapear campos da struct por tag db:"alias"
		fieldMap := make(map[string]reflect.Value)

		for i := 0; i < elem.NumField(); i++ {
			field := elemType.Field(i)
			tag := field.Tag.Get("db")

			if tag == "" || tag == "-" {
				continue
			}

			fv := elem.Field(i)
			if fv.CanAddr() {
				fieldMap[tag] = fv
			}
		}

		// Preparar ponteiros para Scan
		fieldPtrs := make([]any, len(columns))

		for i, col := range columns {
			if f, ok := fieldMap[col]; ok {
				fieldPtrs[i] = f.Addr().Interface()
			} else {
				var dummy any
				fieldPtrs[i] = &dummy
			}
		}

		// Scan da linha
		if err := rows.Scan(fieldPtrs...); err != nil {
			return err
		}

		// Append ou set
		if isSlice {
			destVal.Set(reflect.Append(destVal, elem))
		} else {
			destVal.Set(elem)
			break
		}
	}

	return rows.Err()
}

func (qb *queryBuilder) generateSelectColumns() string {
	cols := []string{}

	if qb.tableDef != nil {
		for _, colInfo := range qb.tableDef.columns {
			alias := fmt.Sprintf("%s__%s", qb.table, colInfo.name)

			cols = append(cols, fmt.Sprintf(
				`"%s"."%s" AS "%s"`,
				qb.table, colInfo.name, alias,
			))
		}
	}

	if qb.hasJoins {
		for _, j := range qb.joins {
			refTable := getTable(&qb.db, j.refTable)
			if refTable == nil {
				continue
			}

			for _, col := range refTable.columns {
				alias := fmt.Sprintf("%s__%s", j.refTable, col.name)

				cols = append(cols, fmt.Sprintf(
					`"%s"."%s" AS "%s"`,
					j.refTable,
					col.name,
					alias,
				))
			}
		}
	}

	if len(cols) == 0 {
		return "*"
	}

	return strings.Join(cols, ", ")
}

func (qb *queryBuilder) buildSQL() (string, []any) {
	cols := qb.generateSelectColumns()

	query := fmt.Sprintf(`SELECT %s FROM "%s"`, cols, qb.table)

	if qb.hasJoins {
		for _, j := range qb.joins {
			query += fmt.Sprintf(` %s JOIN "%s" ON "%s"."%s" = "%s"."%s"`,
				j.joinType,
				j.refTable,
				qb.table,
				j.localColumn,
				j.refTable,
				j.refColumn,
			)
		}
	}

	if qb.tableDef != nil && qb.tableDef.softDeleteCol != "" {

		soft := qb.tableDef.softDeleteCol

		if !qb.includeDeleted && !qb.onlyDeleted {
			if qb.whereClause == "" {
				qb.whereClause = fmt.Sprintf(`WHERE "%s"."%s" IS NULL`, qb.table, soft)
			} else {
				qb.whereClause += fmt.Sprintf(` AND "%s"."%s" IS NULL`, qb.table, soft)
			}
		}

		if qb.onlyDeleted {
			if qb.whereClause == "" {
				qb.whereClause = fmt.Sprintf(`WHERE "%s"."%s" IS NOT NULL`, qb.table, soft)
			} else {
				qb.whereClause += fmt.Sprintf(` AND "%s"."%s" IS NOT NULL`, qb.table, soft)
			}
		}
	}

	if qb.whereClause != "" {
		query += " " + qb.whereClause
	}

	if qb.orderBy != "" {
		query += " ORDER BY " + qb.orderBy
	}

	if qb.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", qb.limit)
	}

	if qb.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", qb.offset)
	}

	return query, qb.args
}

func (b *queryBuilder) WithDeleted() orm.QueryBuilder {
	b.includeDeleted = true
	return b
}

func (b *queryBuilder) OnlyDeleted() orm.QueryBuilder {
	b.onlyDeleted = true
	return b
}

func (qb *queryBuilder) Include(opt orm.WithOption) orm.QueryBuilder {
	r, ok := opt.(withRelation)
	if !ok {
		return qb
	}

	qb.joins = append(qb.joins, joinDef{
		joinType:    "LEFT",
		localTable:  qb.table,
		localColumn: r.meta.LocalColumn,
		refTable:    r.meta.RefTable,
		refColumn:   r.meta.RefColumn,
	})

	qb.hasJoins = true

	return qb
}

func (qb *queryBuilder) Require(opt orm.WithOption) orm.QueryBuilder {
	r, ok := opt.(withRelation)
	if !ok {
		return qb
	}

	qb.joins = append(qb.joins, joinDef{
		joinType:    "INNER",
		localTable:  qb.table,
		localColumn: r.meta.LocalColumn,
		refTable:    r.meta.RefTable,
		refColumn:   r.meta.RefColumn,
	})

	qb.hasJoins = true

	return qb
}

func (qb *queryBuilder) IncludeFrom(opt orm.WithOption) orm.QueryBuilder {
	r, ok := opt.(withRelation)
	if !ok {
		return qb
	}

	qb.joins = append(qb.joins, joinDef{
		joinType:    "RIGHT",
		localTable:  qb.table,
		localColumn: r.meta.LocalColumn,
		refTable:    r.meta.RefTable,
		refColumn:   r.meta.RefColumn,
	})

	qb.hasJoins = true

	return qb
}

func NewQueryBuilder(db db, table string) orm.QueryBuilder {
	tableDef := getTable(&db, table)

	return &queryBuilder{
		db:       db,
		table:    table,
		tableDef: tableDef,
	}
}
