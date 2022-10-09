// Code generated by ent, DO NOT EDIT.

package pokemon

import (
	"time"
)

const (
	// Label holds the string label denoting the pokemon type in the database.
	Label = "pokemon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeFights holds the string denoting the fights edge name in mutations.
	EdgeFights = "fights"
	// EdgeOpponents holds the string denoting the opponents edge name in mutations.
	EdgeOpponents = "opponents"
	// Table holds the table name of the pokemon in the database.
	Table = "pokemons"
	// FightsTable is the table that holds the fights relation/edge.
	FightsTable = "battles"
	// FightsInverseTable is the table name for the Battle entity.
	// It exists in this package in order to avoid circular dependency with the "battle" package.
	FightsInverseTable = "battles"
	// FightsColumn is the table column denoting the fights relation/edge.
	FightsColumn = "pokemon_fights"
	// OpponentsTable is the table that holds the opponents relation/edge.
	OpponentsTable = "battles"
	// OpponentsInverseTable is the table name for the Battle entity.
	// It exists in this package in order to avoid circular dependency with the "battle" package.
	OpponentsInverseTable = "battles"
	// OpponentsColumn is the table column denoting the opponents relation/edge.
	OpponentsColumn = "pokemon_opponents"
)

// Columns holds all SQL columns for pokemon fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldWeight,
	FieldHeight,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
