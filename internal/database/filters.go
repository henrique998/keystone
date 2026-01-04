package database

import (
	"fmt"
	"strings"
)

// ============== COMPOUND FILTERS ===============

func Or(filters ...filter) compoundFilter {
	return compoundFilter{
		op:      "OR",
		filters: filters,
	}
}

func And(filters ...filter) compoundFilter {
	return compoundFilter{
		op:      "AND",
		filters: filters,
	}
}

// =============== COLUMN FILTERS ===============

func (c filterColumn[T]) Equals(value T) filter {
	return filter{field: c.name, op: "=", value: value}
}

func (c filterColumn[T]) NotEquals(value T) filter {
	return filter{field: c.name, op: "<>", value: value, table: c.table}
}

func (c filterColumn[T]) Gt(value T) filter {
	return filter{field: c.name, op: ">", value: value}
}

func (c filterColumn[T]) Lt(value T) filter {
	return filter{field: c.name, op: "<", value: value}
}

func (c filterColumn[T]) Like(value T) filter {
	return filter{field: c.name, op: "ILIKE", value: value}
}

func (c filterColumn[T]) In(values ...T) filter {
	placeholders := make([]string, len(values))

	for i, v := range values {
		if _, ok := any(v).(string); ok {
			placeholders[i] = fmt.Sprintf("'%v'", v)
		} else {
			placeholders[i] = fmt.Sprintf("%v", v)
		}
	}

	return filter{
		field: c.name,
		op:    fmt.Sprintf("IN (%s)", strings.Join(placeholders, ", ")),
		table: c.table,
	}
}

func (c filterColumn[T]) NotIn(values ...T) filter {
	placeholders := make([]string, len(values))

	for i, v := range values {
		if _, ok := any(v).(string); ok {
			placeholders[i] = fmt.Sprintf("'%v'", v)
		} else {
			placeholders[i] = fmt.Sprintf("%v", v)
		}
	}

	return filter{
		field: c.name,
		op:    fmt.Sprintf("NOT IN (%s)", strings.Join(placeholders, ", ")),
		table: c.table,
	}
}

func (c filterColumn[T]) Between(start, end T) filter {
	return filter{
		field: c.name,
		op:    "BETWEEN",
		value: fmt.Sprintf("%v AND %v", start, end),
		table: c.table,
	}
}

func (c filterColumn[T]) IsNull() filter {
	return filter{
		field: c.name,
		op:    "IS NULL",
		table: c.table,
	}
}

func (c filterColumn[T]) IsNotNull() filter {
	return filter{
		field: c.name,
		op:    "IS NOT NULL",
		table: c.table,
	}
}

func newFilterColumn[T any](table string, name string) filterColumn[T] {
	return filterColumn[T]{table: table, name: name}
}
