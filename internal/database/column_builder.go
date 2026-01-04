package database

import (
	"fmt"

	"github.com/henrique998/keystone/orm"
)

func (c *columnBuilder) PrimaryKey() orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, "PRIMARY KEY")
	return c
}

func (c *columnBuilder) Unique() orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, "UNIQUE")
	return c
}

func (c *columnBuilder) NotNull() orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, "NOT NULL")
	return c
}

func (c *columnBuilder) Check(expr string) orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, fmt.Sprintf("CHECK (%s)", expr))
	return c
}

func (c *columnBuilder) DefaultNow() orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, "DEFAULT NOW()")
	return c
}

func (c *columnBuilder) DefaultUUID() orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, "DEFAULT uuid_generate_v4()")
	return c
}

func (c *columnBuilder) AutoIncrement() orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, "GENERATED ALWAYS AS IDENTITY")
	return c
}

func (c *columnBuilder) References(table, column string) orm.ColumnBuilder {
	c.col.ForeignKey = &foreignKeyDef{RefTable: table, RefColumn: column}

	return c
}

func (c *columnBuilder) OnDelete(action string) orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, fmt.Sprintf("ON DELETE %s", action))
	return c
}

func (c *columnBuilder) OnUpdate(action string) orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, fmt.Sprintf("ON UPDATE %s", action))
	return c
}

func (c *columnBuilder) Default(value string) orm.ColumnBuilder {
	c.col.constraints = append(c.col.constraints, fmt.Sprintf("DEFAULT %s", value))
	return c
}

func (c *columnBuilder) BelongsTo(name string, refTable string) orm.ColumnBuilder {
	c.tb.belongsTo = append(c.tb.belongsTo, &relationBelongsTo{
		name:       name,
		columnName: c.col.name,
		refTable:   refTable,
	})

	return c
}
