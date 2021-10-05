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

// TemplateTagRelation holds the schema definition for the TemplateTagRelation entity.
type TemplateTagRelation struct {
	ent.Schema
}

// Annotations of the TemplateTagRelation.
func (TemplateTagRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notify_tag_template"},
	}
}

// Fields of the TemplateTagRelation.
func (TemplateTagRelation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).Immutable(),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).
			UpdateDefault(time.Now),
		field.String("template_uuid").Optional().Nillable().NotEmpty(),
		field.String("tag_uuid").Optional().Nillable().NotEmpty(),
		field.Int("status").
			Default(enum.RELATION_TEMPLATE_TAG_AVAILABLE).NonNegative(),
	}
}

// Edges of the TemplateTagRelation.
func (TemplateTagRelation) Edges() []ent.Edge {
	return nil
}
