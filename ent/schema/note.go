package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("Title").
			Unique(),
		field.String("Content"),
		field.Bool("Private").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	return nil
}
