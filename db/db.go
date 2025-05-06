package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DBConnection() {
	var err error
	dbURL := "postgres://root:secret@localhost:5432/SimpleBankDB"
	// Conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	Conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err.Error())
	}
}
