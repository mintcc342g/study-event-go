// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-event-go/ent/mentorship"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MentorshipCreate is the builder for creating a Mentorship entity.
type MentorshipCreate struct {
	config
	mutation *MentorshipMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MentorshipCreate) SetCreatedAt(t time.Time) *MentorshipCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MentorshipCreate) SetNillableCreatedAt(t *time.Time) *MentorshipCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MentorshipCreate) SetUpdatedAt(t time.Time) *MentorshipCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MentorshipCreate) SetNillableUpdatedAt(t *time.Time) *MentorshipCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MentorshipCreate) SetDeletedAt(t time.Time) *MentorshipCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MentorshipCreate) SetNillableDeletedAt(t *time.Time) *MentorshipCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MentorshipCreate) SetName(s string) *MentorshipCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MentorshipCreate) SetID(ti types.MentorshipID) *MentorshipCreate {
	mc.mutation.SetID(ti)
	return mc
}

// Mutation returns the MentorshipMutation object of the builder.
func (mc *MentorshipCreate) Mutation() *MentorshipMutation {
	return mc.mutation
}

// Save creates the Mentorship in the database.
func (mc *MentorshipCreate) Save(ctx context.Context) (*Mentorship, error) {
	var (
		err  error
		node *Mentorship
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MentorshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MentorshipCreate) SaveX(ctx context.Context) *Mentorship {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MentorshipCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MentorshipCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MentorshipCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := mentorship.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := mentorship.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MentorshipCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := mentorship.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	return nil
}

func (mc *MentorshipCreate) sqlSave(ctx context.Context) (*Mentorship, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = types.MentorshipID(id)
	}
	return _node, nil
}

func (mc *MentorshipCreate) createSpec() (*Mentorship, *sqlgraph.CreateSpec) {
	var (
		_node = &Mentorship{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mentorship.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: mentorship.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mentorship.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mentorship.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mentorship.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mentorship.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mentorship.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MentorshipUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (mc *MentorshipCreate) OnConflict(opts ...sql.ConflictOption) *MentorshipUpsertOne {
	mc.conflict = opts
	return &MentorshipUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mentorship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mc *MentorshipCreate) OnConflictColumns(columns ...string) *MentorshipUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MentorshipUpsertOne{
		create: mc,
	}
}

type (
	// MentorshipUpsertOne is the builder for "upsert"-ing
	//  one Mentorship node.
	MentorshipUpsertOne struct {
		create *MentorshipCreate
	}

	// MentorshipUpsert is the "OnConflict" setter.
	MentorshipUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MentorshipUpsert) SetCreatedAt(v time.Time) *MentorshipUpsert {
	u.Set(mentorship.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MentorshipUpsert) UpdateCreatedAt() *MentorshipUpsert {
	u.SetExcluded(mentorship.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MentorshipUpsert) SetUpdatedAt(v time.Time) *MentorshipUpsert {
	u.Set(mentorship.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MentorshipUpsert) UpdateUpdatedAt() *MentorshipUpsert {
	u.SetExcluded(mentorship.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MentorshipUpsert) SetDeletedAt(v time.Time) *MentorshipUpsert {
	u.Set(mentorship.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MentorshipUpsert) UpdateDeletedAt() *MentorshipUpsert {
	u.SetExcluded(mentorship.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MentorshipUpsert) ClearDeletedAt() *MentorshipUpsert {
	u.SetNull(mentorship.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *MentorshipUpsert) SetName(v string) *MentorshipUpsert {
	u.Set(mentorship.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MentorshipUpsert) UpdateName() *MentorshipUpsert {
	u.SetExcluded(mentorship.FieldName)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Mentorship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mentorship.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MentorshipUpsertOne) UpdateNewValues() *MentorshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mentorship.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Mentorship.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MentorshipUpsertOne) Ignore() *MentorshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MentorshipUpsertOne) DoNothing() *MentorshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MentorshipCreate.OnConflict
// documentation for more info.
func (u *MentorshipUpsertOne) Update(set func(*MentorshipUpsert)) *MentorshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MentorshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MentorshipUpsertOne) SetCreatedAt(v time.Time) *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MentorshipUpsertOne) UpdateCreatedAt() *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MentorshipUpsertOne) SetUpdatedAt(v time.Time) *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MentorshipUpsertOne) UpdateUpdatedAt() *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MentorshipUpsertOne) SetDeletedAt(v time.Time) *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MentorshipUpsertOne) UpdateDeletedAt() *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MentorshipUpsertOne) ClearDeletedAt() *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *MentorshipUpsertOne) SetName(v string) *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MentorshipUpsertOne) UpdateName() *MentorshipUpsertOne {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *MentorshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MentorshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MentorshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MentorshipUpsertOne) ID(ctx context.Context) (id types.MentorshipID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MentorshipUpsertOne) IDX(ctx context.Context) types.MentorshipID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MentorshipCreateBulk is the builder for creating many Mentorship entities in bulk.
type MentorshipCreateBulk struct {
	config
	builders []*MentorshipCreate
	conflict []sql.ConflictOption
}

// Save creates the Mentorship entities in the database.
func (mcb *MentorshipCreateBulk) Save(ctx context.Context) ([]*Mentorship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mentorship, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MentorshipMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = types.MentorshipID(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MentorshipCreateBulk) SaveX(ctx context.Context) []*Mentorship {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MentorshipCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MentorshipCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mentorship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MentorshipUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (mcb *MentorshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *MentorshipUpsertBulk {
	mcb.conflict = opts
	return &MentorshipUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mentorship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mcb *MentorshipCreateBulk) OnConflictColumns(columns ...string) *MentorshipUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MentorshipUpsertBulk{
		create: mcb,
	}
}

// MentorshipUpsertBulk is the builder for "upsert"-ing
// a bulk of Mentorship nodes.
type MentorshipUpsertBulk struct {
	create *MentorshipCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Mentorship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mentorship.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MentorshipUpsertBulk) UpdateNewValues() *MentorshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mentorship.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Mentorship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MentorshipUpsertBulk) Ignore() *MentorshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MentorshipUpsertBulk) DoNothing() *MentorshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MentorshipCreateBulk.OnConflict
// documentation for more info.
func (u *MentorshipUpsertBulk) Update(set func(*MentorshipUpsert)) *MentorshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MentorshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MentorshipUpsertBulk) SetCreatedAt(v time.Time) *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MentorshipUpsertBulk) UpdateCreatedAt() *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MentorshipUpsertBulk) SetUpdatedAt(v time.Time) *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MentorshipUpsertBulk) UpdateUpdatedAt() *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MentorshipUpsertBulk) SetDeletedAt(v time.Time) *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MentorshipUpsertBulk) UpdateDeletedAt() *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MentorshipUpsertBulk) ClearDeletedAt() *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *MentorshipUpsertBulk) SetName(v string) *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MentorshipUpsertBulk) UpdateName() *MentorshipUpsertBulk {
	return u.Update(func(s *MentorshipUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *MentorshipUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MentorshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MentorshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MentorshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
