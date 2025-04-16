// Пример изменения схемы токена (auth-service/ent/schema/token.go)
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Genre struct {
	ent.Schema
}

func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description").Default(""),
		field.Time("created_at").Default(time.Now()),
	}
}
