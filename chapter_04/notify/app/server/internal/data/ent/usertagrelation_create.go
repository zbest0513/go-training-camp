// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"notify-server/internal/data/ent/usertagrelation"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserTagRelationCreate is the builder for creating a UserTagRelation entity.
type UserTagRelationCreate struct {
	config
	mutation *UserTagRelationMutation
	hooks    []Hook
}

// Mutation returns the UserTagRelationMutation object of the builder.
func (utrc *UserTagRelationCreate) Mutation() *UserTagRelationMutation {
	return utrc.mutation
}

// Save creates the UserTagRelation in the database.
func (utrc *UserTagRelationCreate) Save(ctx context.Context) (*UserTagRelation, error) {
	var (
		err  error
		node *UserTagRelation
	)
	if len(utrc.hooks) == 0 {
		if err = utrc.check(); err != nil {
			return nil, err
		}
		node, err = utrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTagRelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = utrc.check(); err != nil {
				return nil, err
			}
			utrc.mutation = mutation
			if node, err = utrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(utrc.hooks) - 1; i >= 0; i-- {
			if utrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = utrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (utrc *UserTagRelationCreate) SaveX(ctx context.Context) *UserTagRelation {
	v, err := utrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utrc *UserTagRelationCreate) Exec(ctx context.Context) error {
	_, err := utrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utrc *UserTagRelationCreate) ExecX(ctx context.Context) {
	if err := utrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utrc *UserTagRelationCreate) check() error {
	return nil
}

func (utrc *UserTagRelationCreate) sqlSave(ctx context.Context) (*UserTagRelation, error) {
	_node, _spec := utrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, utrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (utrc *UserTagRelationCreate) createSpec() (*UserTagRelation, *sqlgraph.CreateSpec) {
	var (
		_node = &UserTagRelation{config: utrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usertagrelation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertagrelation.FieldID,
			},
		}
	)
	return _node, _spec
}

// UserTagRelationCreateBulk is the builder for creating many UserTagRelation entities in bulk.
type UserTagRelationCreateBulk struct {
	config
	builders []*UserTagRelationCreate
}

// Save creates the UserTagRelation entities in the database.
func (utrcb *UserTagRelationCreateBulk) Save(ctx context.Context) ([]*UserTagRelation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(utrcb.builders))
	nodes := make([]*UserTagRelation, len(utrcb.builders))
	mutators := make([]Mutator, len(utrcb.builders))
	for i := range utrcb.builders {
		func(i int, root context.Context) {
			builder := utrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserTagRelationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, utrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, utrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, utrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (utrcb *UserTagRelationCreateBulk) SaveX(ctx context.Context) []*UserTagRelation {
	v, err := utrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utrcb *UserTagRelationCreateBulk) Exec(ctx context.Context) error {
	_, err := utrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utrcb *UserTagRelationCreateBulk) ExecX(ctx context.Context) {
	if err := utrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
