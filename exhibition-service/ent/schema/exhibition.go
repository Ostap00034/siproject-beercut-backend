package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Exhibition struct {
	ent.Schema
}

func (Exhibition) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().Unique(),
		field.String("description"),
		field.Strings("pictures_ids").
			Optional(),
		field.Enum("status").
			Values("OPENED", "CLOSED").
			Default("CLOSED"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Exhibition) Edges() []ent.Edge {
	return nil
}
