package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Connect returns a pointer to a Postgres database connection and an error.
func Connect() (*pgx.Conn, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.TODO(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf(`error connecting to database %v`, err)
	}
	// defer conn.Close(context.TODO())

	return conn, nil
}
