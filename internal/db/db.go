package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectToDatabase() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(context.TODO(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf(`error connecting to database %v`, err)
	}

	defer conn.Close(context.Background())
	fmt.Println("Connected to database")
	return nil
}
