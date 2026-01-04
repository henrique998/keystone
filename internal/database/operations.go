package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/henrique998/keystone/orm"
)

func (db *db) QueryRaw(query string, args ...any) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *db) FindMany(table string, filters ...any) orm.QueryBuilder {
	qb := newQueryBuilder(*db, table)
	whereClause, args := buildWhereClause(filters, nil)
	qb.whereClause = whereClause
	qb.args = args

	if qb.tableDef != nil && qb.tableDef.softDeleteCol != "" {
		qb.includeDeleted = false
		qb.onlyDeleted = false
	}

	return qb
}

func (db *db) FindOne(table string, filters ...any) orm.QueryBuilder {
	qb := newQueryBuilder(*db, table)
	whereClause, args := buildWhereClause(filters, nil)
	qb.whereClause = whereClause
	qb.args = args
	qb.limit = 1

	if qb.tableDef != nil && qb.tableDef.softDeleteCol != "" {
		qb.includeDeleted = false
		qb.onlyDeleted = false
	}

	return qb
}

func (db *db) Create(table string, data interface{}) error {
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Slice:
		return db.insertBatch(table, v)

	case reflect.Struct:
		return db.insertRow(table, data)

	default:
		return fmt.Errorf("invalid data type: expected struct or slice, got %s", v.Kind())
	}
}

func (db *db) Update(table string, filters ...any) orm.UpdateBuilder {
	return &updateBuilder{
		db:       *db,
		table:    table,
		tableDef: getTable(db, table),
		filters:  filters,
		sets:     make(map[string]any),
	}
}

func (db *db) Delete(table string, filters ...any) orm.DeleteBuilder {
	return &deleteBuilder{
		db:      *db,
		table:   table,
		filters: filters,
	}
}

func (db *db) DeleteBatch(table string, filters ...any) orm.DeleteBatchBuilder {
	return &deleteBatchBuilder{
		db:      *db,
		table:   table,
		filters: filters,
	}
}

func buildCondition(f filter, argIndex *int) (string, []interface{}) {
	op := strings.ToUpper(f.op)
	var args []interface{}

	switch {
	case strings.Contains(op, "IN"):
		values := toInterfaceSlice(f.value)
		placeholders := make([]string, len(values))
		for i, v := range values {
			placeholders[i] = fmt.Sprintf("$%d", *argIndex)
			*argIndex++
			args = append(args, v)
		}
		return fmt.Sprintf(`"%s" IN (%s)`, f.field, strings.Join(placeholders, ",")), args

	case op == "IS NULL" || op == "IS NOT NULL":
		return fmt.Sprintf(`"%s" %s`, f.field, op), args

	case op == "BETWEEN":
		rangeVals := toInterfaceSlice(f.value)
		if len(rangeVals) == 2 {
			sql := fmt.Sprintf(`"%s" BETWEEN $%d AND $%d`, f.field, *argIndex, *argIndex+1)
			args = append(args, rangeVals[0], rangeVals[1])
			*argIndex += 2
			return sql, args
		}

	default:
		sql := fmt.Sprintf(`"%s" %s $%d`, f.field, f.op, *argIndex)
		args = append(args, f.value)
		*argIndex++
		return sql, args
	}

	return "", nil
}

func buildWhereClause(filters []any, argIndex *int) (string, []any) {
	var conditions []string
	var args []interface{}

	if argIndex == nil {
		argIndex = new(int)
		*argIndex = 1
	}

	for _, f := range filters {
		switch c := f.(type) {
		case filter:
			sqlPart, sqlArgs := buildCondition(c, argIndex)
			conditions = append(conditions, sqlPart)
			args = append(args, sqlArgs...)

		case compoundFilter:
			groupParts := []string{}
			for _, sub := range c.filters {
				sqlPart, sqlArgs := buildCondition(sub, argIndex)
				groupParts = append(groupParts, sqlPart)
				args = append(args, sqlArgs...)
			}
			group := fmt.Sprintf("(%s)", strings.Join(groupParts, fmt.Sprintf(" %s ", c.op)))
			conditions = append(conditions, group)
		}
	}

	if len(conditions) == 0 {
		return "", args
	}

	return "WHERE " + strings.Join(conditions, " AND "), args
}

func (db *db) insertRow(table string, data interface{}) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")

		if tag == "" {
			tag = field.Tag.Get("json")
		}

		if tag == "" {
			tag = field.Tag.Get("ks")
		}

		if tag == "" {
			tag = field.Name
		}

		columns = append(columns, tag)
		placeholders = append(placeholders, "$"+strconv.Itoa(i+1))
		values = append(values, v.Field(i).Interface())
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.conn.Exec(query, values...)

	return err
}

func (db *db) insertBatch(table string, slice reflect.Value) error {
	if slice.Len() == 0 {
		return nil
	}

	firstElem := slice.Index(0)
	elemType := firstElem.Type()

	var columns []string
	var placeholders []string
	var allValues []interface{}
	argIndex := 1

	acceptedTags := []string{"db", "json", "ks"}

	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)

		var tag string
		for _, acceptedTag := range acceptedTags {
			if tag == "" {
				tag = field.Tag.Get(acceptedTag)
			}
		}

		if tag == "" {
			tag = field.Name
		}

		columns = append(columns, tag)
	}

	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)

		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		rowValues := make([]interface{}, len(columns))
		rowPlaceholders := make([]string, len(columns))

		for j := 0; j < len(columns); j++ {
			rowValues[j] = elem.Field(j).Interface()
			rowPlaceholders[j] = fmt.Sprintf("$%d", argIndex)
			argIndex++
		}

		allValues = append(allValues, rowValues...)
		placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.conn.Exec(query, allValues...)

	return err
}
