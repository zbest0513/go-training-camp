// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"notify-server/internal/data/ent/predicate"
	"notify-server/internal/data/ent/templatetagrelation"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TemplateTagRelationUpdate is the builder for updating TemplateTagRelation entities.
type TemplateTagRelationUpdate struct {
	config
	hooks    []Hook
	mutation *TemplateTagRelationMutation
}

// Where appends a list predicates to the TemplateTagRelationUpdate builder.
func (ttru *TemplateTagRelationUpdate) Where(ps ...predicate.TemplateTagRelation) *TemplateTagRelationUpdate {
	ttru.mutation.Where(ps...)
	return ttru
}

// SetUpdatedAt sets the "updated_at" field.
func (ttru *TemplateTagRelationUpdate) SetUpdatedAt(t time.Time) *TemplateTagRelationUpdate {
	ttru.mutation.SetUpdatedAt(t)
	return ttru
}

// SetTemplateUUID sets the "template_uuid" field.
func (ttru *TemplateTagRelationUpdate) SetTemplateUUID(s string) *TemplateTagRelationUpdate {
	ttru.mutation.SetTemplateUUID(s)
	return ttru
}

// SetNillableTemplateUUID sets the "template_uuid" field if the given value is not nil.
func (ttru *TemplateTagRelationUpdate) SetNillableTemplateUUID(s *string) *TemplateTagRelationUpdate {
	if s != nil {
		ttru.SetTemplateUUID(*s)
	}
	return ttru
}

// ClearTemplateUUID clears the value of the "template_uuid" field.
func (ttru *TemplateTagRelationUpdate) ClearTemplateUUID() *TemplateTagRelationUpdate {
	ttru.mutation.ClearTemplateUUID()
	return ttru
}

// SetTagUUID sets the "tag_uuid" field.
func (ttru *TemplateTagRelationUpdate) SetTagUUID(s string) *TemplateTagRelationUpdate {
	ttru.mutation.SetTagUUID(s)
	return ttru
}

// SetNillableTagUUID sets the "tag_uuid" field if the given value is not nil.
func (ttru *TemplateTagRelationUpdate) SetNillableTagUUID(s *string) *TemplateTagRelationUpdate {
	if s != nil {
		ttru.SetTagUUID(*s)
	}
	return ttru
}

// ClearTagUUID clears the value of the "tag_uuid" field.
func (ttru *TemplateTagRelationUpdate) ClearTagUUID() *TemplateTagRelationUpdate {
	ttru.mutation.ClearTagUUID()
	return ttru
}

// SetStatus sets the "status" field.
func (ttru *TemplateTagRelationUpdate) SetStatus(i int) *TemplateTagRelationUpdate {
	ttru.mutation.ResetStatus()
	ttru.mutation.SetStatus(i)
	return ttru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ttru *TemplateTagRelationUpdate) SetNillableStatus(i *int) *TemplateTagRelationUpdate {
	if i != nil {
		ttru.SetStatus(*i)
	}
	return ttru
}

// AddStatus adds i to the "status" field.
func (ttru *TemplateTagRelationUpdate) AddStatus(i int) *TemplateTagRelationUpdate {
	ttru.mutation.AddStatus(i)
	return ttru
}

// Mutation returns the TemplateTagRelationMutation object of the builder.
func (ttru *TemplateTagRelationUpdate) Mutation() *TemplateTagRelationMutation {
	return ttru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttru *TemplateTagRelationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ttru.defaults()
	if len(ttru.hooks) == 0 {
		if err = ttru.check(); err != nil {
			return 0, err
		}
		affected, err = ttru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateTagRelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttru.check(); err != nil {
				return 0, err
			}
			ttru.mutation = mutation
			affected, err = ttru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttru.hooks) - 1; i >= 0; i-- {
			if ttru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttru *TemplateTagRelationUpdate) SaveX(ctx context.Context) int {
	affected, err := ttru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttru *TemplateTagRelationUpdate) Exec(ctx context.Context) error {
	_, err := ttru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttru *TemplateTagRelationUpdate) ExecX(ctx context.Context) {
	if err := ttru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttru *TemplateTagRelationUpdate) defaults() {
	if _, ok := ttru.mutation.UpdatedAt(); !ok {
		v := templatetagrelation.UpdateDefaultUpdatedAt()
		ttru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttru *TemplateTagRelationUpdate) check() error {
	if v, ok := ttru.mutation.TemplateUUID(); ok {
		if err := templatetagrelation.TemplateUUIDValidator(v); err != nil {
			return &ValidationError{Name: "template_uuid", err: fmt.Errorf("ent: validator failed for field \"template_uuid\": %w", err)}
		}
	}
	if v, ok := ttru.mutation.TagUUID(); ok {
		if err := templatetagrelation.TagUUIDValidator(v); err != nil {
			return &ValidationError{Name: "tag_uuid", err: fmt.Errorf("ent: validator failed for field \"tag_uuid\": %w", err)}
		}
	}
	if v, ok := ttru.mutation.Status(); ok {
		if err := templatetagrelation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ttru *TemplateTagRelationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   templatetagrelation.Table,
			Columns: templatetagrelation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: templatetagrelation.FieldID,
			},
		},
	}
	if ps := ttru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: templatetagrelation.FieldUpdatedAt,
		})
	}
	if value, ok := ttru.mutation.TemplateUUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: templatetagrelation.FieldTemplateUUID,
		})
	}
	if ttru.mutation.TemplateUUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: templatetagrelation.FieldTemplateUUID,
		})
	}
	if value, ok := ttru.mutation.TagUUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: templatetagrelation.FieldTagUUID,
		})
	}
	if ttru.mutation.TagUUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: templatetagrelation.FieldTagUUID,
		})
	}
	if value, ok := ttru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: templatetagrelation.FieldStatus,
		})
	}
	if value, ok := ttru.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: templatetagrelation.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templatetagrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TemplateTagRelationUpdateOne is the builder for updating a single TemplateTagRelation entity.
type TemplateTagRelationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TemplateTagRelationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ttruo *TemplateTagRelationUpdateOne) SetUpdatedAt(t time.Time) *TemplateTagRelationUpdateOne {
	ttruo.mutation.SetUpdatedAt(t)
	return ttruo
}

// SetTemplateUUID sets the "template_uuid" field.
func (ttruo *TemplateTagRelationUpdateOne) SetTemplateUUID(s string) *TemplateTagRelationUpdateOne {
	ttruo.mutation.SetTemplateUUID(s)
	return ttruo
}

// SetNillableTemplateUUID sets the "template_uuid" field if the given value is not nil.
func (ttruo *TemplateTagRelationUpdateOne) SetNillableTemplateUUID(s *string) *TemplateTagRelationUpdateOne {
	if s != nil {
		ttruo.SetTemplateUUID(*s)
	}
	return ttruo
}

// ClearTemplateUUID clears the value of the "template_uuid" field.
func (ttruo *TemplateTagRelationUpdateOne) ClearTemplateUUID() *TemplateTagRelationUpdateOne {
	ttruo.mutation.ClearTemplateUUID()
	return ttruo
}

// SetTagUUID sets the "tag_uuid" field.
func (ttruo *TemplateTagRelationUpdateOne) SetTagUUID(s string) *TemplateTagRelationUpdateOne {
	ttruo.mutation.SetTagUUID(s)
	return ttruo
}

// SetNillableTagUUID sets the "tag_uuid" field if the given value is not nil.
func (ttruo *TemplateTagRelationUpdateOne) SetNillableTagUUID(s *string) *TemplateTagRelationUpdateOne {
	if s != nil {
		ttruo.SetTagUUID(*s)
	}
	return ttruo
}

// ClearTagUUID clears the value of the "tag_uuid" field.
func (ttruo *TemplateTagRelationUpdateOne) ClearTagUUID() *TemplateTagRelationUpdateOne {
	ttruo.mutation.ClearTagUUID()
	return ttruo
}

// SetStatus sets the "status" field.
func (ttruo *TemplateTagRelationUpdateOne) SetStatus(i int) *TemplateTagRelationUpdateOne {
	ttruo.mutation.ResetStatus()
	ttruo.mutation.SetStatus(i)
	return ttruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ttruo *TemplateTagRelationUpdateOne) SetNillableStatus(i *int) *TemplateTagRelationUpdateOne {
	if i != nil {
		ttruo.SetStatus(*i)
	}
	return ttruo
}

// AddStatus adds i to the "status" field.
func (ttruo *TemplateTagRelationUpdateOne) AddStatus(i int) *TemplateTagRelationUpdateOne {
	ttruo.mutation.AddStatus(i)
	return ttruo
}

// Mutation returns the TemplateTagRelationMutation object of the builder.
func (ttruo *TemplateTagRelationUpdateOne) Mutation() *TemplateTagRelationMutation {
	return ttruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttruo *TemplateTagRelationUpdateOne) Select(field string, fields ...string) *TemplateTagRelationUpdateOne {
	ttruo.fields = append([]string{field}, fields...)
	return ttruo
}

// Save executes the query and returns the updated TemplateTagRelation entity.
func (ttruo *TemplateTagRelationUpdateOne) Save(ctx context.Context) (*TemplateTagRelation, error) {
	var (
		err  error
		node *TemplateTagRelation
	)
	ttruo.defaults()
	if len(ttruo.hooks) == 0 {
		if err = ttruo.check(); err != nil {
			return nil, err
		}
		node, err = ttruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateTagRelationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttruo.check(); err != nil {
				return nil, err
			}
			ttruo.mutation = mutation
			node, err = ttruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttruo.hooks) - 1; i >= 0; i-- {
			if ttruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ttruo *TemplateTagRelationUpdateOne) SaveX(ctx context.Context) *TemplateTagRelation {
	node, err := ttruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttruo *TemplateTagRelationUpdateOne) Exec(ctx context.Context) error {
	_, err := ttruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttruo *TemplateTagRelationUpdateOne) ExecX(ctx context.Context) {
	if err := ttruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttruo *TemplateTagRelationUpdateOne) defaults() {
	if _, ok := ttruo.mutation.UpdatedAt(); !ok {
		v := templatetagrelation.UpdateDefaultUpdatedAt()
		ttruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttruo *TemplateTagRelationUpdateOne) check() error {
	if v, ok := ttruo.mutation.TemplateUUID(); ok {
		if err := templatetagrelation.TemplateUUIDValidator(v); err != nil {
			return &ValidationError{Name: "template_uuid", err: fmt.Errorf("ent: validator failed for field \"template_uuid\": %w", err)}
		}
	}
	if v, ok := ttruo.mutation.TagUUID(); ok {
		if err := templatetagrelation.TagUUIDValidator(v); err != nil {
			return &ValidationError{Name: "tag_uuid", err: fmt.Errorf("ent: validator failed for field \"tag_uuid\": %w", err)}
		}
	}
	if v, ok := ttruo.mutation.Status(); ok {
		if err := templatetagrelation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ttruo *TemplateTagRelationUpdateOne) sqlSave(ctx context.Context) (_node *TemplateTagRelation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   templatetagrelation.Table,
			Columns: templatetagrelation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: templatetagrelation.FieldID,
			},
		},
	}
	id, ok := ttruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TemplateTagRelation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ttruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, templatetagrelation.FieldID)
		for _, f := range fields {
			if !templatetagrelation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != templatetagrelation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: templatetagrelation.FieldUpdatedAt,
		})
	}
	if value, ok := ttruo.mutation.TemplateUUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: templatetagrelation.FieldTemplateUUID,
		})
	}
	if ttruo.mutation.TemplateUUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: templatetagrelation.FieldTemplateUUID,
		})
	}
	if value, ok := ttruo.mutation.TagUUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: templatetagrelation.FieldTagUUID,
		})
	}
	if ttruo.mutation.TagUUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: templatetagrelation.FieldTagUUID,
		})
	}
	if value, ok := ttruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: templatetagrelation.FieldStatus,
		})
	}
	if value, ok := ttruo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: templatetagrelation.FieldStatus,
		})
	}
	_node = &TemplateTagRelation{config: ttruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templatetagrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
