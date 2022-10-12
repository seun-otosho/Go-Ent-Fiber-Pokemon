package main

import (
	"GoEntFiberPokeman/database"
	"GoEntFiberPokeman/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	app := fiber.New()

	client := database.DbClient
	defer client.Close()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	routes.NoteRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.Get("APP_PORT"))))

}
