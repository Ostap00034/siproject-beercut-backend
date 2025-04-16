package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MovementHistory holds the schema definition for the MovementHistory entity.
type MovementHistory struct {
	ent.Schema
}

// Fields of the MovementHistory.
func (MovementHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("picture_id").
			NotEmpty(),
		field.String("user_id").NotEmpty(),
		field.Enum("from").
			Values("IN_STORAGE", "IN_EXHIBITION", "IN_RESTORATION"),
		field.Enum("to").
			Values("IN_STORAGE", "IN_EXHIBITION", "IN_RESTORATION"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges for MovementHistory.
// Если связи с авторами и жанрами организованы через отдельные микросервисы,
// здесь можно не определять ребра, а просто хранить внешние идентификаторы в виде массива.
func (MovementHistory) Edges() []ent.Edge {
	return nil
}
