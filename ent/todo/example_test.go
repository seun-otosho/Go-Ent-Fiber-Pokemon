package todo

import (
	"GoEntFiberPokeman/ent"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"

	_ "github.com/lib/pq"
)

func Example_Todo() {
	DSN := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USER"),
		viper.Get("DB_NAME"), viper.Get("DB_PASSWORD"),
	)
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open("postgres", DSN)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:
}
