package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"notify-server/internal/pkg/enum"
	"time"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

// Annotations of the Template.
func (Template) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notify_template"},
	}
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
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
		field.String("content"),
		field.Int("status").
			Default(enum.TEMPLATE_STATUS_AVAILABLE).NonNegative(),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	//return []ent.Edge{
	//	edge.To("tags", Tag.Type).Annotations(entsql.Annotation{
	//		OnDelete: entsql.Cascade,
	//		Table:"notify_tag_template",
	//	}),
	//}
	return nil
}
