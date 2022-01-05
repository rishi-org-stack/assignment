package main

import (
	"context"
	"database/sql"
	"fmt"
	"portfolio/internal/auth"
	"portfolio/internal/portfolio"
	"portfolio/internal/user"
	"portfolio/util/cache"

	// "gitub.com/lib/pq"
	"log"

	_ "github.com/lib/pq"

	gpsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	psn := "postgresql://pgadmin:password@localhost/portfolio?sslmode=disable"
	sqlDB, err := sql.Open("postgres", psn)
	checkErr(err)
	gdb, err := gorm.Open(gpsql.New(gpsql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	tx := gdb.Exec("DROP SCHEMA IF EXISTS auth CASCADE")
	checkErr(tx.Error)
	tx = gdb.Exec("CREATE SCHEMA IF NOT EXISTS auth")
	checkErr(tx.Error)
	ok := gdb.AutoMigrate(&auth.AuthRequest{})
	checkErr(ok)
	tx = gdb.Exec("DROP SCHEMA IF EXISTS usr CASCADE")
	checkErr(tx.Error)

	tx = gdb.Exec("CREATE SCHEMA IF NOT EXISTS usr")
	checkErr(tx.Error)
	ok = gdb.AutoMigrate(
		&user.Usr{},
	)
	checkErr(ok)

	tx = gdb.Exec("DROP SCHEMA IF EXISTS portfolio CASCADE")
	checkErr(tx.Error)

	tx = gdb.Exec("CREATE SCHEMA IF NOT  EXISTS portfolio ")
	checkErr(tx.Error)
	ok = gdb.AutoMigrate(
		&portfolio.Portfolio{},
		&portfolio.Entry{},
	)
	checkErr(ok)
	// portfolios
	clearCache()
}
func clearCache() {
	rdb, err := cache.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	rdb.DB.FlushAll(context.TODO())
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
