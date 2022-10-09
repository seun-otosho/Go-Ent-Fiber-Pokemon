package main

import (
	"GoEntFiberPokeman/ent"
	elk "GoEntFiberPokeman/ent/http"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_NAME"), viper.Get("DB_PASSWORD"),
	)
	c, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer c.Close()
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Run the auto migration tool.
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Router and Logger.
	r, l := chi.NewRouter(), zap.NewExample()
	// Create the pet handler.
	r.Route("/pets", func(r chi.Router) {
		elk.NewPetHandler(c, l).Mount(r, elk.PetRoutes)
	})
	// Create the user handler.
	r.Route("/users", func(r chi.Router) {
		elk.NewUserHandler(c, l).Mount(r, elk.UserRoutes)
	})
	// Start listen to incoming requests.
	fmt.Println("Server running")
	defer fmt.Println("Server stopped")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
