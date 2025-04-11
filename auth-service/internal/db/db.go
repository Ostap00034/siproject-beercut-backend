package db

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent"
	_ "github.com/lib/pq" // драйвер для PostgreSQL
)

func NewClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Создание таблиц (автоматическая миграция)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
