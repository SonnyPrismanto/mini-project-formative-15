package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {

	// fmt.Println(dbParam)
	migration := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migration, migrate.Up)
	fmt.Println(n)
	if errs != nil {
		panic(errs)
	}
	DbConnection = dbParam
	fmt.Println("Applied", n, "migations!")
}
