package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/qrm"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	pg *sql.DB
}

func NewDatabaseClient(conn string) *Database {
	db, err := sql.Open("pgx", conn)
	if err != nil {
		panic(err)
	}

	// config
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("⚡️ [postgresql]: connected \n")

	return &Database{
		pg: db,
	}
}

func (d Database) GetDB() *sql.DB {
	return d.pg
}

func (Database) IsNoRowsError(err error) bool {
	return errors.Is(err, qrm.ErrNoRows)
}
