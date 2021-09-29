package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// TemplateTagRelation holds the schema definition for the TemplateTagRelation entity.
type TemplateTagRelation struct {
	ent.Schema
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
		field.String("template_uuid").Nillable().NotEmpty(),
		field.String("tag_uuid").Nillable().NotEmpty(),
		field.Int("status").
			Default(1).NonNegative(),
	}
}

// Edges of the TemplateTagRelation.
func (TemplateTagRelation) Edges() []ent.Edge {
	return nil
}
