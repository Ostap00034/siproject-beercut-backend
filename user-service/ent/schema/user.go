package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User описывает сущность пользователя в User Service.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password_hash"),
		field.String("full_name"),
		field.Enum("role").Values("DIRECTOR", "ADMIN", "RESTORER", "EMPLOYEE").Default("EMPLOYEE"),
		field.Time("created_at").Default(time.Now),
	}
}
