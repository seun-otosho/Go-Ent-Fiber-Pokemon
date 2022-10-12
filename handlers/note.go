package handlers

import (
	"GoEntFiberPokeman/database"
	"context"
	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
	//Parse the request body
	note := new(struct {
		Title   string
		Content string
		Private bool
	})

	if err := c.BodyParser(&note); err != nil {
		c.Status(400).JSON("Error  Parsing Input")
		return err
	}

	//Save to the database
	createdNote, err := database.DbClient.Note.
		Create().
		SetTitle(note.Title).
		SetContent(note.Content).
		SetPrivate(note.Private).
		Save(context.Background())

	if err != nil {
		c.Status(500).JSON("Unable to save note")
		return err
	}
	//Send the created note back with the appropriate code.
	c.Status(200).JSON(createdNote)

	return nil
}
