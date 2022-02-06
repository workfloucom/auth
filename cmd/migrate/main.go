package main

import (
	"fmt"
	"os"

	"workflou.com/auth/pkg/database"
)

func main() {
	cfg := database.Config{
		Dsn:    os.Getenv("DB_DSN"),
		Driver: os.Getenv("DB_DRIVER"),
	}

	db := database.New(cfg)

	db.Migrate()

	fmt.Print("Migrated.")
}
