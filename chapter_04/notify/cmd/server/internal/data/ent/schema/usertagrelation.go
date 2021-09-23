package schema

import "entgo.io/ent"

// UserTagRelation holds the schema definition for the UserTagRelation entity.
type UserTagRelation struct {
	ent.Schema
}

// Fields of the UserTagRelation.
func (UserTagRelation) Fields() []ent.Field {
	return nil
}

// Edges of the UserTagRelation.
func (UserTagRelation) Edges() []ent.Edge {
	return nil
}
