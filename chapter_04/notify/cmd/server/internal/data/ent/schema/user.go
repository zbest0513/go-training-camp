package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notify_user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).Immutable(),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Default(time.Now).
			UpdateDefault(time.Now),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("mobile").
			NotEmpty().MaxLen(11).MinLen(11),
		field.String("name").
			NotEmpty().MaxLen(100),
		field.String("email").
			NotEmpty().MaxLen(100),
		field.Int("status").
			Default(0).NonNegative(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	//return []ent.Edge{
	//	edge.To("tags", Tag.Type).Annotations(entsql.Annotation{
	//		OnDelete: entsql.Cascade,
	//		Table:"notify_user_tag",
	//	}),
	//}
	return nil
}
