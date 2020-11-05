package persistence

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

const DATABASE_URL = "postgres://qaa_user:123@localhost:5432/training?currentSchema=qaa"

func GetDataBaseConnection() *pgx.Conn {

	connection, err := pgx.Connect(context.Background(), os.Getenv(DATABASE_URL))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return connection
}
