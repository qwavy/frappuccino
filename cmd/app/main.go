package main

import (
	"context"
	handlers2 "frappucchino/internal/adapters/primary/http-adapter"
	repositories2 "frappucchino/internal/adapters/secondary/postgres"
	services2 "frappucchino/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to DataBase", err)
	}
	defer pool.Close()

	repositories := repositories2.NewRepositories(pool)
	services := services2.NewServices(repositories)
	handlers := handlers2.NewHandlers(services)
	mux := http.NewServeMux()

	handlers2.Router(handlers, mux)
}
