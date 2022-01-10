// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"study-event-go/ent/lilyskill"
	"study-event-go/ent/predicate"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LilySkillUpdate is the builder for updating LilySkill entities.
type LilySkillUpdate struct {
	config
	hooks    []Hook
	mutation *LilySkillMutation
}

// Where appends a list predicates to the LilySkillUpdate builder.
func (lsu *LilySkillUpdate) Where(ps ...predicate.LilySkill) *LilySkillUpdate {
	lsu.mutation.Where(ps...)
	return lsu
}

// SetLilyID sets the "lily_id" field.
func (lsu *LilySkillUpdate) SetLilyID(ti types.LilyID) *LilySkillUpdate {
	lsu.mutation.ResetLilyID()
	lsu.mutation.SetLilyID(ti)
	return lsu
}

// AddLilyID adds ti to the "lily_id" field.
func (lsu *LilySkillUpdate) AddLilyID(ti types.LilyID) *LilySkillUpdate {
	lsu.mutation.AddLilyID(ti)
	return lsu
}

// SetSkillID sets the "skill_id" field.
func (lsu *LilySkillUpdate) SetSkillID(ti types.SkillID) *LilySkillUpdate {
	lsu.mutation.ResetSkillID()
	lsu.mutation.SetSkillID(ti)
	return lsu
}

// AddSkillID adds ti to the "skill_id" field.
func (lsu *LilySkillUpdate) AddSkillID(ti types.SkillID) *LilySkillUpdate {
	lsu.mutation.AddSkillID(ti)
	return lsu
}

// SetCreatedAt sets the "created_at" field.
func (lsu *LilySkillUpdate) SetCreatedAt(t time.Time) *LilySkillUpdate {
	lsu.mutation.SetCreatedAt(t)
	return lsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lsu *LilySkillUpdate) SetNillableCreatedAt(t *time.Time) *LilySkillUpdate {
	if t != nil {
		lsu.SetCreatedAt(*t)
	}
	return lsu
}

// SetUpdatedAt sets the "updated_at" field.
func (lsu *LilySkillUpdate) SetUpdatedAt(t time.Time) *LilySkillUpdate {
	lsu.mutation.SetUpdatedAt(t)
	return lsu
}

// Mutation returns the LilySkillMutation object of the builder.
func (lsu *LilySkillUpdate) Mutation() *LilySkillMutation {
	return lsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lsu *LilySkillUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lsu.defaults()
	if len(lsu.hooks) == 0 {
		affected, err = lsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LilySkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lsu.mutation = mutation
			affected, err = lsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lsu.hooks) - 1; i >= 0; i-- {
			if lsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lsu *LilySkillUpdate) SaveX(ctx context.Context) int {
	affected, err := lsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lsu *LilySkillUpdate) Exec(ctx context.Context) error {
	_, err := lsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsu *LilySkillUpdate) ExecX(ctx context.Context) {
	if err := lsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lsu *LilySkillUpdate) defaults() {
	if _, ok := lsu.mutation.UpdatedAt(); !ok {
		v := lilyskill.UpdateDefaultUpdatedAt()
		lsu.mutation.SetUpdatedAt(v)
	}
}

func (lsu *LilySkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lilyskill.Table,
			Columns: lilyskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lilyskill.FieldID,
			},
		},
	}
	if ps := lsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsu.mutation.LilyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldLilyID,
		})
	}
	if value, ok := lsu.mutation.AddedLilyID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldLilyID,
		})
	}
	if value, ok := lsu.mutation.SkillID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldSkillID,
		})
	}
	if value, ok := lsu.mutation.AddedSkillID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldSkillID,
		})
	}
	if value, ok := lsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lilyskill.FieldCreatedAt,
		})
	}
	if value, ok := lsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lilyskill.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lilyskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LilySkillUpdateOne is the builder for updating a single LilySkill entity.
type LilySkillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LilySkillMutation
}

// SetLilyID sets the "lily_id" field.
func (lsuo *LilySkillUpdateOne) SetLilyID(ti types.LilyID) *LilySkillUpdateOne {
	lsuo.mutation.ResetLilyID()
	lsuo.mutation.SetLilyID(ti)
	return lsuo
}

// AddLilyID adds ti to the "lily_id" field.
func (lsuo *LilySkillUpdateOne) AddLilyID(ti types.LilyID) *LilySkillUpdateOne {
	lsuo.mutation.AddLilyID(ti)
	return lsuo
}

// SetSkillID sets the "skill_id" field.
func (lsuo *LilySkillUpdateOne) SetSkillID(ti types.SkillID) *LilySkillUpdateOne {
	lsuo.mutation.ResetSkillID()
	lsuo.mutation.SetSkillID(ti)
	return lsuo
}

// AddSkillID adds ti to the "skill_id" field.
func (lsuo *LilySkillUpdateOne) AddSkillID(ti types.SkillID) *LilySkillUpdateOne {
	lsuo.mutation.AddSkillID(ti)
	return lsuo
}

// SetCreatedAt sets the "created_at" field.
func (lsuo *LilySkillUpdateOne) SetCreatedAt(t time.Time) *LilySkillUpdateOne {
	lsuo.mutation.SetCreatedAt(t)
	return lsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lsuo *LilySkillUpdateOne) SetNillableCreatedAt(t *time.Time) *LilySkillUpdateOne {
	if t != nil {
		lsuo.SetCreatedAt(*t)
	}
	return lsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (lsuo *LilySkillUpdateOne) SetUpdatedAt(t time.Time) *LilySkillUpdateOne {
	lsuo.mutation.SetUpdatedAt(t)
	return lsuo
}

// Mutation returns the LilySkillMutation object of the builder.
func (lsuo *LilySkillUpdateOne) Mutation() *LilySkillMutation {
	return lsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lsuo *LilySkillUpdateOne) Select(field string, fields ...string) *LilySkillUpdateOne {
	lsuo.fields = append([]string{field}, fields...)
	return lsuo
}

// Save executes the query and returns the updated LilySkill entity.
func (lsuo *LilySkillUpdateOne) Save(ctx context.Context) (*LilySkill, error) {
	var (
		err  error
		node *LilySkill
	)
	lsuo.defaults()
	if len(lsuo.hooks) == 0 {
		node, err = lsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LilySkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lsuo.mutation = mutation
			node, err = lsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lsuo.hooks) - 1; i >= 0; i-- {
			if lsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lsuo *LilySkillUpdateOne) SaveX(ctx context.Context) *LilySkill {
	node, err := lsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lsuo *LilySkillUpdateOne) Exec(ctx context.Context) error {
	_, err := lsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsuo *LilySkillUpdateOne) ExecX(ctx context.Context) {
	if err := lsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lsuo *LilySkillUpdateOne) defaults() {
	if _, ok := lsuo.mutation.UpdatedAt(); !ok {
		v := lilyskill.UpdateDefaultUpdatedAt()
		lsuo.mutation.SetUpdatedAt(v)
	}
}

func (lsuo *LilySkillUpdateOne) sqlSave(ctx context.Context) (_node *LilySkill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lilyskill.Table,
			Columns: lilyskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lilyskill.FieldID,
			},
		},
	}
	id, ok := lsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LilySkill.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := lsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lilyskill.FieldID)
		for _, f := range fields {
			if !lilyskill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lilyskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsuo.mutation.LilyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldLilyID,
		})
	}
	if value, ok := lsuo.mutation.AddedLilyID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldLilyID,
		})
	}
	if value, ok := lsuo.mutation.SkillID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldSkillID,
		})
	}
	if value, ok := lsuo.mutation.AddedSkillID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldSkillID,
		})
	}
	if value, ok := lsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lilyskill.FieldCreatedAt,
		})
	}
	if value, ok := lsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lilyskill.FieldUpdatedAt,
		})
	}
	_node = &LilySkill{config: lsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lilyskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
