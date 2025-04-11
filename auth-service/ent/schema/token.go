// ent/schema/token.go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique(),
		field.String("role"),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Unique().Required(),
	}
}
