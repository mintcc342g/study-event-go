// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"study-event-go/ent/garden"
	"study-event-go/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GardenDelete is the builder for deleting a Garden entity.
type GardenDelete struct {
	config
	hooks    []Hook
	mutation *GardenMutation
}

// Where appends a list predicates to the GardenDelete builder.
func (gd *GardenDelete) Where(ps ...predicate.Garden) *GardenDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GardenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gd.hooks) == 0 {
		affected, err = gd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GardenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gd.mutation = mutation
			affected, err = gd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gd.hooks) - 1; i >= 0; i-- {
			if gd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GardenDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GardenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: garden.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		},
	}
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
}

// GardenDeleteOne is the builder for deleting a single Garden entity.
type GardenDeleteOne struct {
	gd *GardenDelete
}

// Exec executes the deletion query.
func (gdo *GardenDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{garden.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GardenDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}