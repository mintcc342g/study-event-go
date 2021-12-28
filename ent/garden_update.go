// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"study-event-go/ent/garden"
	"study-event-go/ent/mentorshipsystem"
	"study-event-go/ent/predicate"
	"study-event-go/types"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GardenUpdate is the builder for updating Garden entities.
type GardenUpdate struct {
	config
	hooks    []Hook
	mutation *GardenMutation
}

// Where appends a list predicates to the GardenUpdate builder.
func (gu *GardenUpdate) Where(ps ...predicate.Garden) *GardenUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GardenUpdate) SetName(s string) *GardenUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetLocation sets the "location" field.
func (gu *GardenUpdate) SetLocation(s string) *GardenUpdate {
	gu.mutation.SetLocation(s)
	return gu
}

// SetMentorshipSystemID sets the "mentorship_system_id" field.
func (gu *GardenUpdate) SetMentorshipSystemID(tsi types.MentorshipSystemID) *GardenUpdate {
	gu.mutation.SetMentorshipSystemID(tsi)
	return gu
}

// SetNillableMentorshipSystemID sets the "mentorship_system_id" field if the given value is not nil.
func (gu *GardenUpdate) SetNillableMentorshipSystemID(tsi *types.MentorshipSystemID) *GardenUpdate {
	if tsi != nil {
		gu.SetMentorshipSystemID(*tsi)
	}
	return gu
}

// ClearMentorshipSystemID clears the value of the "mentorship_system_id" field.
func (gu *GardenUpdate) ClearMentorshipSystemID() *GardenUpdate {
	gu.mutation.ClearMentorshipSystemID()
	return gu
}

// SetMentorshipSystem sets the "mentorship_system" edge to the MentorshipSystem entity.
func (gu *GardenUpdate) SetMentorshipSystem(m *MentorshipSystem) *GardenUpdate {
	return gu.SetMentorshipSystemID(m.ID)
}

// Mutation returns the GardenMutation object of the builder.
func (gu *GardenUpdate) Mutation() *GardenMutation {
	return gu.mutation
}

// ClearMentorshipSystem clears the "mentorship_system" edge to the MentorshipSystem entity.
func (gu *GardenUpdate) ClearMentorshipSystem() *GardenUpdate {
	gu.mutation.ClearMentorshipSystem()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GardenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GardenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GardenUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GardenUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GardenUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GardenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   garden.Table,
			Columns: garden.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldName,
		})
	}
	if value, ok := gu.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldLocation,
		})
	}
	if gu.mutation.MentorshipSystemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   garden.MentorshipSystemTable,
			Columns: []string{garden.MentorshipSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mentorshipsystem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MentorshipSystemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   garden.MentorshipSystemTable,
			Columns: []string{garden.MentorshipSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mentorshipsystem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{garden.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GardenUpdateOne is the builder for updating a single Garden entity.
type GardenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GardenMutation
}

// SetName sets the "name" field.
func (guo *GardenUpdateOne) SetName(s string) *GardenUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetLocation sets the "location" field.
func (guo *GardenUpdateOne) SetLocation(s string) *GardenUpdateOne {
	guo.mutation.SetLocation(s)
	return guo
}

// SetMentorshipSystemID sets the "mentorship_system_id" field.
func (guo *GardenUpdateOne) SetMentorshipSystemID(tsi types.MentorshipSystemID) *GardenUpdateOne {
	guo.mutation.SetMentorshipSystemID(tsi)
	return guo
}

// SetNillableMentorshipSystemID sets the "mentorship_system_id" field if the given value is not nil.
func (guo *GardenUpdateOne) SetNillableMentorshipSystemID(tsi *types.MentorshipSystemID) *GardenUpdateOne {
	if tsi != nil {
		guo.SetMentorshipSystemID(*tsi)
	}
	return guo
}

// ClearMentorshipSystemID clears the value of the "mentorship_system_id" field.
func (guo *GardenUpdateOne) ClearMentorshipSystemID() *GardenUpdateOne {
	guo.mutation.ClearMentorshipSystemID()
	return guo
}

// SetMentorshipSystem sets the "mentorship_system" edge to the MentorshipSystem entity.
func (guo *GardenUpdateOne) SetMentorshipSystem(m *MentorshipSystem) *GardenUpdateOne {
	return guo.SetMentorshipSystemID(m.ID)
}

// Mutation returns the GardenMutation object of the builder.
func (guo *GardenUpdateOne) Mutation() *GardenMutation {
	return guo.mutation
}

// ClearMentorshipSystem clears the "mentorship_system" edge to the MentorshipSystem entity.
func (guo *GardenUpdateOne) ClearMentorshipSystem() *GardenUpdateOne {
	guo.mutation.ClearMentorshipSystem()
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GardenUpdateOne) Select(field string, fields ...string) *GardenUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Garden entity.
func (guo *GardenUpdateOne) Save(ctx context.Context) (*Garden, error) {
	var (
		err  error
		node *Garden
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GardenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GardenUpdateOne) SaveX(ctx context.Context) *Garden {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GardenUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GardenUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GardenUpdateOne) sqlSave(ctx context.Context) (_node *Garden, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   garden.Table,
			Columns: garden.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Garden.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, garden.FieldID)
		for _, f := range fields {
			if !garden.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != garden.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldName,
		})
	}
	if value, ok := guo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: garden.FieldLocation,
		})
	}
	if guo.mutation.MentorshipSystemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   garden.MentorshipSystemTable,
			Columns: []string{garden.MentorshipSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mentorshipsystem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MentorshipSystemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   garden.MentorshipSystemTable,
			Columns: []string{garden.MentorshipSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: mentorshipsystem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Garden{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{garden.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
