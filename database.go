package main

import (
	"database/sql"
	"healthtracker/migrations"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func mustNewDatabase(dsn string) *sql.DB {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	goose.SetBaseFS(migrations.FS)

	if err = goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err = goose.Up(db, "."); err != nil {
		panic(err)
	}

	return db
}
