package main

import (
	"github.com/henrique998/keystone/internal/database"
	"github.com/henrique998/keystone/orm"
)

func main() {
	db := database.New(orm.Args{
		Credentials: orm.Credentials{
			Host:     "localhost",
			User:     "postgres",
			Password: "postgres",
			DBName:   "postgres",
			Port:     5432,
		},
		Dialect: "postgres",
	})

	db.NewTable("cars", func(tb orm.TableBuilder) {
		tb.Varchar("name", 50).NotNull()
	}).UseSoftDelete("dissmised_at")

	err := db.SyncSchemas()
	if err != nil {
		panic(err)
	}
}
