package database

import "fmt"

func NewColumn[T any](tableName string, columnName string, sqlType string) column[T] {
	return column[T]{
		tableName:  tableName,
		columnName: columnName,
		sqlType:    sqlType,
	}
}

func (c column[T]) Eq(value T) condition {
	return condition{
		sql:  fmt.Sprintf(`"%s"."%s" = $1`, c.tableName, c.columnName),
		args: []any{value},
	}
}

func (c column[T]) Gt(value T) condition {
	return condition{
		sql:  fmt.Sprintf(`"%s"."%s" > $1`, c.tableName, c.columnName),
		args: []any{value},
	}
}

func (c column[T]) Lt(value T) condition {
	return condition{
		sql:  fmt.Sprintf(`"%s"."%s" < $1`, c.tableName, c.columnName),
		args: []any{value},
	}
}
