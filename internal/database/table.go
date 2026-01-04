package database

import (
	"fmt"

	"github.com/henrique998/keystone/orm"
)

func (t *tableDefinition) UseSoftDelete(column string) {
	t.softDeleteCol = column
}

func (t *tableBuilder) Char(name string, size int) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: fmt.Sprintf("CHAR(%d)", size),
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Varchar(name string, size int) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: fmt.Sprintf("VARCHAR(%d)", size),
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Text(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "TEXT",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) SmallInt(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "SMALLINT",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Int(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "INTEGER",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) BigInt(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "BIGINT",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) SmallSerial(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "SMALLSERIAL",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Serial(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "SERIAL",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) BigSerial(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "BIGSERIAL",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Numeric(name string, precision ...int) orm.ColumnBuilder {
	sqlType := "NUMERIC"

	if len(precision) == 1 {
		sqlType = fmt.Sprintf("NUMERIC(%d)", precision[0])
	} else if len(precision) == 2 {
		sqlType = fmt.Sprintf("NUMERIC(%d, %d)", precision[0], precision[1])
	}

	columnInfo := columnInfo{name: name, sqlType: sqlType}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Decimal(name string, precision, scale int) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: fmt.Sprintf("DECIMAL(%d, %d)", precision, scale),
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Real(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "REAL",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Double(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "DOUBLE PRECISION",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Float(name string, precision ...int) orm.ColumnBuilder {
	sqlType := "FLOAT"

	if len(precision) > 0 {
		sqlType = fmt.Sprintf("FLOAT(%d)", precision[0])
	}

	columnInfo := columnInfo{name: name, sqlType: sqlType}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Bool(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "BOOLEAN",
	}

	t.columns[name] = &columnInfo

	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Date(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "DATE",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Timestamp(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "TIMESTAMP",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Time(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "TIME",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Timestamptz(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "TIMESTAMPTZ",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) TimeWithTimeZone(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "TIME WITH TIME ZONE",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) JSON(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "JSON",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) JSONB(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "JSONB",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) UUID(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "UUID",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}

func (t *tableBuilder) Bytea(name string) orm.ColumnBuilder {
	columnInfo := columnInfo{
		name:    name,
		sqlType: "BYTEA",
	}
	t.columns[name] = &columnInfo
	return &columnBuilder{col: &columnInfo, tb: t}
}
