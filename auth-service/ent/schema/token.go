// Пример изменения схемы токена (auth-service/ent/schema/token.go)
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Token описывает токен авторизации.
type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique(),
		// Вместо связи с сущностью User, здесь поле user_id в виде строки:
		field.String("user_id"),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now),
	}
}
