// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"notify-server/internal/data/ent/predicate"
	"notify-server/internal/data/ent/template"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TemplateUpdate is the builder for updating Template entities.
type TemplateUpdate struct {
	config
	hooks    []Hook
	mutation *TemplateMutation
}

// Where appends a list predicates to the TemplateUpdate builder.
func (tu *TemplateUpdate) Where(ps ...predicate.Template) *TemplateUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TemplateUpdate) SetUpdatedAt(t time.Time) *TemplateUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetUUID sets the "uuid" field.
func (tu *TemplateUpdate) SetUUID(s string) *TemplateUpdate {
	tu.mutation.SetUUID(s)
	return tu
}

// SetDesc sets the "desc" field.
func (tu *TemplateUpdate) SetDesc(s string) *TemplateUpdate {
	tu.mutation.SetDesc(s)
	return tu
}

// SetName sets the "name" field.
func (tu *TemplateUpdate) SetName(s string) *TemplateUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetContent sets the "content" field.
func (tu *TemplateUpdate) SetContent(s string) *TemplateUpdate {
	tu.mutation.SetContent(s)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TemplateUpdate) SetStatus(i int) *TemplateUpdate {
	tu.mutation.ResetStatus()
	tu.mutation.SetStatus(i)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableStatus(i *int) *TemplateUpdate {
	if i != nil {
		tu.SetStatus(*i)
	}
	return tu
}

// AddStatus adds i to the "status" field.
func (tu *TemplateUpdate) AddStatus(i int) *TemplateUpdate {
	tu.mutation.AddStatus(i)
	return tu
}

// Mutation returns the TemplateMutation object of the builder.
func (tu *TemplateUpdate) Mutation() *TemplateMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TemplateUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TemplateUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TemplateUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := template.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TemplateUpdate) check() error {
	if v, ok := tu.mutation.Desc(); ok {
		if err := template.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf("ent: validator failed for field \"desc\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := template.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (tu *TemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   template.Table,
			Columns: template.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: template.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: template.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldUUID,
		})
	}
	if value, ok := tu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldDesc,
		})
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldName,
		})
	}
	if value, ok := tu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldContent,
		})
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldStatus,
		})
	}
	if value, ok := tu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TemplateUpdateOne is the builder for updating a single Template entity.
type TemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TemplateMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TemplateUpdateOne) SetUpdatedAt(t time.Time) *TemplateUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetUUID sets the "uuid" field.
func (tuo *TemplateUpdateOne) SetUUID(s string) *TemplateUpdateOne {
	tuo.mutation.SetUUID(s)
	return tuo
}

// SetDesc sets the "desc" field.
func (tuo *TemplateUpdateOne) SetDesc(s string) *TemplateUpdateOne {
	tuo.mutation.SetDesc(s)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TemplateUpdateOne) SetName(s string) *TemplateUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetContent sets the "content" field.
func (tuo *TemplateUpdateOne) SetContent(s string) *TemplateUpdateOne {
	tuo.mutation.SetContent(s)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TemplateUpdateOne) SetStatus(i int) *TemplateUpdateOne {
	tuo.mutation.ResetStatus()
	tuo.mutation.SetStatus(i)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableStatus(i *int) *TemplateUpdateOne {
	if i != nil {
		tuo.SetStatus(*i)
	}
	return tuo
}

// AddStatus adds i to the "status" field.
func (tuo *TemplateUpdateOne) AddStatus(i int) *TemplateUpdateOne {
	tuo.mutation.AddStatus(i)
	return tuo
}

// Mutation returns the TemplateMutation object of the builder.
func (tuo *TemplateUpdateOne) Mutation() *TemplateMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TemplateUpdateOne) Select(field string, fields ...string) *TemplateUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Template entity.
func (tuo *TemplateUpdateOne) Save(ctx context.Context) (*Template, error) {
	var (
		err  error
		node *Template
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TemplateUpdateOne) SaveX(ctx context.Context) *Template {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplateUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TemplateUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := template.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TemplateUpdateOne) check() error {
	if v, ok := tuo.mutation.Desc(); ok {
		if err := template.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf("ent: validator failed for field \"desc\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := template.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (tuo *TemplateUpdateOne) sqlSave(ctx context.Context) (_node *Template, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   template.Table,
			Columns: template.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: template.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Template.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, template.FieldID)
		for _, f := range fields {
			if !template.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != template.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: template.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldUUID,
		})
	}
	if value, ok := tuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldDesc,
		})
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldName,
		})
	}
	if value, ok := tuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldContent,
		})
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldStatus,
		})
	}
	if value, ok := tuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldStatus,
		})
	}
	_node = &Template{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
