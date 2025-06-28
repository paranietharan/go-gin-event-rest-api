package main

import (
	"database/sql"
	"fmt"
	"go-gin-event-rest-api/internal/database"
	"go-gin-event-rest-api/internal/env"
	"log"

	_ "github.com/lib/pq"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	host := env.GetEnvString("HOST", "localhost")
	port := env.GetEnvInt("PORT", 5432)
	user := env.GetEnvString("USER", "postgres")
	password := env.GetEnvString("PASSWORD", "root")
	dbname := env.GetEnvString("DB_NAME", "event_db")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-123456"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
