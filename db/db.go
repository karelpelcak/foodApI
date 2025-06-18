package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/karelpelcak/foodApI/ent"
	_ "github.com/lib/pq"
)

func SetupDB() (*ent.Client, error) {
		dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

		client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	// Zkus připojení
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL and migrated schema")
	return client, nil
}
