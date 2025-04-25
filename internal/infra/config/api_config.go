package config

import (
	"database/sql"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
	"github.com/murphy6867/silly_server_go/internal/database"
)

type APIConfig struct {
	FileServerHits atomic.Int32
	DB             *database.Queries
	JWTSecret      string
	TokenExpireIn  time.Duration
	CloseDB        *sql.DB
}

func Load() *APIConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using real env vars")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET environment variable is not set")
	}

	expireTk := os.Getenv("TOKEN_EXPIRE")
	if expireTk == "" {
		log.Fatal("TOKEN_EXPIRE environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	parsedExpire, err := time.ParseDuration(expireTk)
	if err != nil {
		log.Fatalf("failed to parse TOKEN_EXPIRE: %v", err)
	}
	return &APIConfig{
		DB:            database.New(db),
		JWTSecret:     secret,
		TokenExpireIn: parsedExpire,
		CloseDB:       db,
	}
}
