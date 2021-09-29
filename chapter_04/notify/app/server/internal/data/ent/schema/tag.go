package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Annotations of the Tag.
func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notify_tag"},
	}
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).Immutable(),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).
			UpdateDefault(time.Now),
		field.String("uuid").Unique(),
		field.String("desc").
			NotEmpty().MaxLen(300),
		field.String("name").
			NotEmpty().MaxLen(100),
		field.Int("status").
			Default(1).NonNegative(),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	//return []ent.Edge{
	//	edge.From("users", User.Type).
	//		Ref("tags").Annotations(entsql.Annotation{
	//		OnDelete: entsql.Cascade,
	//		Table:"notify_user_tag",
	//	}),
	//	edge.From("templates", Template.Type).
	//		Ref("tags").Annotations(entsql.Annotation{
	//		OnDelete: entsql.Cascade,
	//		Table:"notify_tag_template",
	//	}),
	//}
	return nil
}
