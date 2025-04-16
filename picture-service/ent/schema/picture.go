package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Picture holds the schema definition for the Picture entity.
type Picture struct {
	ent.Schema
}

// Fields of the Picture.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		// Название картины – обязательное.
		field.String("name").
			NotEmpty().Unique(),
		// Дата написания картины в формате строки (например, "2025-04-12").
		field.String("date_of_painting").
			NotEmpty(),
		// Массив идентификаторов жанров, хранится как JSON-массив строк.
		field.Strings("genres_ids").
			Optional(),
		// Массив идентификаторов авторов.
		field.Strings("authors_ids").
			Optional(),
		// Идентификатор выставки; если картина не выставлена, может быть пустым.
		field.String("exhibition_id").
			Optional(),
		// Примерная стоимость картины, по умолчанию 0.
		field.Float("cost").
			Default(0),
		// Enum для локации картины.
		field.Enum("location").
			Values("IN_STORAGE", "IN_EXHIBITION", "IN_RESTORATION").
			Default("IN_STORAGE"),
		// Время создания записи.
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges for Picture.
// Если связи с авторами и жанрами организованы через отдельные микросервисы,
// здесь можно не определять ребра, а просто хранить внешние идентификаторы в виде массива.
func (Picture) Edges() []ent.Edge {
	return nil
}
