package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	pgURL := os.Getenv("POSTGRESQL_URL")
	conn, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal("Erreur lors de la connexion à PostgreSQL:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal("Erreur lors du ping de PostgreSQL:", err)
	}

	db = conn
	fmt.Println("Connecté à la base de données PostgreSQL!")
}

func GetDB() *sql.DB {
	return db
}
