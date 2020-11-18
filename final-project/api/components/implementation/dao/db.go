package dao

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

const DATABASE_URL = "postgres://postgres:admin@localhost:5432/training?search_path=qaa"

func GetDataBaseConnection() *pgx.Conn {
	connection, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return connection
}
