// ent/schema/user.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// Поле email, уникальное и обязательное.
		field.String("email").Unique(),
		field.String("password_hash"),
		field.String("full_name"),
		// Определяем поле role как enum: допустимые значения — DIRECTOR, ADMIN, RESTORER, EMPLOYEE.
		field.Enum("role").
			Values("DIRECTOR", "ADMIN", "RESTORER", "EMPLOYEE").
			Default("EMPLOYEE"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", Token.Type),
	}
}
