package db

import (
	"context"
	"database/sql"
	"fmt"
	"portfolio/util/config"

	gpsql "gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func Connect(ctx context.Context, env *config.Env) *gorm.DB {
	sqlDB, err := sql.Open("postgres", env.DB)
	if err != nil {
		fmt.Println(err)
	}
	gdb, err := gorm.Open(gpsql.New(gpsql.Config{
		// DSN:                  "user=pgadmin password=password dbname=ganymede port=5432 sslmode=disable",
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return gdb
}
