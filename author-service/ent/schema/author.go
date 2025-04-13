package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Author struct {
	ent.Schema
}

func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Unique(),
		field.String("date_of_birth"),
		field.String("date_of_death").Optional(),
		field.Time("created_at").Default(time.Now()),
	}
}
