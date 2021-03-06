// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-event-go/ent/charm"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CharmCreate is the builder for creating a Charm entity.
type CharmCreate struct {
	config
	mutation *CharmMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CharmCreate) SetCreatedAt(t time.Time) *CharmCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CharmCreate) SetNillableCreatedAt(t *time.Time) *CharmCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CharmCreate) SetUpdatedAt(t time.Time) *CharmCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CharmCreate) SetNillableUpdatedAt(t *time.Time) *CharmCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CharmCreate) SetDeletedAt(t time.Time) *CharmCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CharmCreate) SetNillableDeletedAt(t *time.Time) *CharmCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CharmCreate) SetName(s string) *CharmCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetModelID sets the "model_id" field.
func (cc *CharmCreate) SetModelID(tmi types.CharmModelID) *CharmCreate {
	cc.mutation.SetModelID(tmi)
	return cc
}

// SetOwnerID sets the "owner_id" field.
func (cc *CharmCreate) SetOwnerID(ti types.LilyID) *CharmCreate {
	cc.mutation.SetOwnerID(ti)
	return cc
}

// SetID sets the "id" field.
func (cc *CharmCreate) SetID(ti types.CharmID) *CharmCreate {
	cc.mutation.SetID(ti)
	return cc
}

// Mutation returns the CharmMutation object of the builder.
func (cc *CharmCreate) Mutation() *CharmMutation {
	return cc.mutation
}

// Save creates the Charm in the database.
func (cc *CharmCreate) Save(ctx context.Context) (*Charm, error) {
	var (
		err  error
		node *Charm
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CharmMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CharmCreate) SaveX(ctx context.Context) *Charm {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CharmCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CharmCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CharmCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := charm.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := charm.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CharmCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := charm.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ModelID(); !ok {
		return &ValidationError{Name: "model_id", err: errors.New(`ent: missing required field "model_id"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "owner_id"`)}
	}
	return nil
}

func (cc *CharmCreate) sqlSave(ctx context.Context) (*Charm, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = types.CharmID(id)
	}
	return _node, nil
}

func (cc *CharmCreate) createSpec() (*Charm, *sqlgraph.CreateSpec) {
	var (
		_node = &Charm{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: charm.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: charm.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charm.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charm.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charm.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charm.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.ModelID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: charm.FieldModelID,
		})
		_node.ModelID = value
	}
	if value, ok := cc.mutation.OwnerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: charm.FieldOwnerID,
		})
		_node.OwnerID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Charm.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CharmUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CharmCreate) OnConflict(opts ...sql.ConflictOption) *CharmUpsertOne {
	cc.conflict = opts
	return &CharmUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Charm.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CharmCreate) OnConflictColumns(columns ...string) *CharmUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CharmUpsertOne{
		create: cc,
	}
}

type (
	// CharmUpsertOne is the builder for "upsert"-ing
	//  one Charm node.
	CharmUpsertOne struct {
		create *CharmCreate
	}

	// CharmUpsert is the "OnConflict" setter.
	CharmUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CharmUpsert) SetCreatedAt(v time.Time) *CharmUpsert {
	u.Set(charm.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CharmUpsert) UpdateCreatedAt() *CharmUpsert {
	u.SetExcluded(charm.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CharmUpsert) SetUpdatedAt(v time.Time) *CharmUpsert {
	u.Set(charm.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CharmUpsert) UpdateUpdatedAt() *CharmUpsert {
	u.SetExcluded(charm.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CharmUpsert) SetDeletedAt(v time.Time) *CharmUpsert {
	u.Set(charm.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CharmUpsert) UpdateDeletedAt() *CharmUpsert {
	u.SetExcluded(charm.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *CharmUpsert) ClearDeletedAt() *CharmUpsert {
	u.SetNull(charm.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *CharmUpsert) SetName(v string) *CharmUpsert {
	u.Set(charm.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CharmUpsert) UpdateName() *CharmUpsert {
	u.SetExcluded(charm.FieldName)
	return u
}

// SetModelID sets the "model_id" field.
func (u *CharmUpsert) SetModelID(v types.CharmModelID) *CharmUpsert {
	u.Set(charm.FieldModelID, v)
	return u
}

// UpdateModelID sets the "model_id" field to the value that was provided on create.
func (u *CharmUpsert) UpdateModelID() *CharmUpsert {
	u.SetExcluded(charm.FieldModelID)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *CharmUpsert) SetOwnerID(v types.LilyID) *CharmUpsert {
	u.Set(charm.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *CharmUpsert) UpdateOwnerID() *CharmUpsert {
	u.SetExcluded(charm.FieldOwnerID)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Charm.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(charm.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CharmUpsertOne) UpdateNewValues() *CharmUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(charm.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Charm.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CharmUpsertOne) Ignore() *CharmUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CharmUpsertOne) DoNothing() *CharmUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CharmCreate.OnConflict
// documentation for more info.
func (u *CharmUpsertOne) Update(set func(*CharmUpsert)) *CharmUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CharmUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CharmUpsertOne) SetCreatedAt(v time.Time) *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CharmUpsertOne) UpdateCreatedAt() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CharmUpsertOne) SetUpdatedAt(v time.Time) *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CharmUpsertOne) UpdateUpdatedAt() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CharmUpsertOne) SetDeletedAt(v time.Time) *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CharmUpsertOne) UpdateDeletedAt() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *CharmUpsertOne) ClearDeletedAt() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *CharmUpsertOne) SetName(v string) *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CharmUpsertOne) UpdateName() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateName()
	})
}

// SetModelID sets the "model_id" field.
func (u *CharmUpsertOne) SetModelID(v types.CharmModelID) *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.SetModelID(v)
	})
}

// UpdateModelID sets the "model_id" field to the value that was provided on create.
func (u *CharmUpsertOne) UpdateModelID() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateModelID()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *CharmUpsertOne) SetOwnerID(v types.LilyID) *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *CharmUpsertOne) UpdateOwnerID() *CharmUpsertOne {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateOwnerID()
	})
}

// Exec executes the query.
func (u *CharmUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CharmCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CharmUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CharmUpsertOne) ID(ctx context.Context) (id types.CharmID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CharmUpsertOne) IDX(ctx context.Context) types.CharmID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CharmCreateBulk is the builder for creating many Charm entities in bulk.
type CharmCreateBulk struct {
	config
	builders []*CharmCreate
	conflict []sql.ConflictOption
}

// Save creates the Charm entities in the database.
func (ccb *CharmCreateBulk) Save(ctx context.Context) ([]*Charm, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Charm, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharmMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
					nodes[i].ID = types.CharmID(id)
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CharmCreateBulk) SaveX(ctx context.Context) []*Charm {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CharmCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CharmCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Charm.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CharmUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CharmCreateBulk) OnConflict(opts ...sql.ConflictOption) *CharmUpsertBulk {
	ccb.conflict = opts
	return &CharmUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Charm.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CharmCreateBulk) OnConflictColumns(columns ...string) *CharmUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CharmUpsertBulk{
		create: ccb,
	}
}

// CharmUpsertBulk is the builder for "upsert"-ing
// a bulk of Charm nodes.
type CharmUpsertBulk struct {
	create *CharmCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Charm.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(charm.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CharmUpsertBulk) UpdateNewValues() *CharmUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(charm.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Charm.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CharmUpsertBulk) Ignore() *CharmUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CharmUpsertBulk) DoNothing() *CharmUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CharmCreateBulk.OnConflict
// documentation for more info.
func (u *CharmUpsertBulk) Update(set func(*CharmUpsert)) *CharmUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CharmUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CharmUpsertBulk) SetCreatedAt(v time.Time) *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CharmUpsertBulk) UpdateCreatedAt() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CharmUpsertBulk) SetUpdatedAt(v time.Time) *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CharmUpsertBulk) UpdateUpdatedAt() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CharmUpsertBulk) SetDeletedAt(v time.Time) *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CharmUpsertBulk) UpdateDeletedAt() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *CharmUpsertBulk) ClearDeletedAt() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *CharmUpsertBulk) SetName(v string) *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CharmUpsertBulk) UpdateName() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateName()
	})
}

// SetModelID sets the "model_id" field.
func (u *CharmUpsertBulk) SetModelID(v types.CharmModelID) *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.SetModelID(v)
	})
}

// UpdateModelID sets the "model_id" field to the value that was provided on create.
func (u *CharmUpsertBulk) UpdateModelID() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateModelID()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *CharmUpsertBulk) SetOwnerID(v types.LilyID) *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *CharmUpsertBulk) UpdateOwnerID() *CharmUpsertBulk {
	return u.Update(func(s *CharmUpsert) {
		s.UpdateOwnerID()
	})
}

// Exec executes the query.
func (u *CharmUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CharmCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CharmCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CharmUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
