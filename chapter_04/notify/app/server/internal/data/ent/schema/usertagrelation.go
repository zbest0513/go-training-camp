package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"notify-server/internal/pkg/enum"
	"time"
)

// UserTagRelation holds the schema definition for the UserTagRelation entity.
type UserTagRelation struct {
	ent.Schema
}

// Fields of the UserTagRelation.
func (UserTagRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).Immutable(),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).
			UpdateDefault(time.Now),
		field.String("user_uuid").Optional().Nillable().NotEmpty(),
		field.String("tag_uuid").Optional().Nillable().NotEmpty(),
		field.Int("status").
			Default(enum.RELATION_USER_TAG_AVAILABLE).NonNegative(),
	}
}

// Edges of the UserTagRelation.
func (UserTagRelation) Edges() []ent.Edge {
	return nil
}
