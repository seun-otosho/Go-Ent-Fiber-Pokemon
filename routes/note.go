package routes

import (
	"GoEntFiberPokeman/handlers"
	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(app *fiber.App) {
	api := app.Group("/api/v1/notes")

	api.Post("/", handlers.CreateNote)
	//api.Get("/", handlers.ReadNote)
	//api.Get("/by_title/:title", handlers.SearchNotes)
	//api.Put("/:id", handlers.UpdateNote)
	//api.Delete("/:id", handlers.DeleteNotes)
}
