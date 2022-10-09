package main

import (
	"GoEntFiberPokeman/ent"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	ctx := context.Background()

	app := fiber.New()

	//url := fmt.Sprintf(
	//	"%s:%s@(%s:%s)/%s?parseTime=True", viper.Get("DB_USER"), viper.Get("DB_PASSWORD"),
	//	viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"),
	//)
	//client, err := ent.Open(dialect.Postgres, url) // connect to MySQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_NAME"), viper.Get("DB_PASSWORD"),
	)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi, World!")
	})
	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.Get("APP_PORT"))))
}
