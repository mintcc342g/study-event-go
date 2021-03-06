// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-event-go/ent/lilyskill"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LilySkillCreate is the builder for creating a LilySkill entity.
type LilySkillCreate struct {
	config
	mutation *LilySkillMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetLilyID sets the "lily_id" field.
func (lsc *LilySkillCreate) SetLilyID(ti types.LilyID) *LilySkillCreate {
	lsc.mutation.SetLilyID(ti)
	return lsc
}

// SetSkillID sets the "skill_id" field.
func (lsc *LilySkillCreate) SetSkillID(ti types.SkillID) *LilySkillCreate {
	lsc.mutation.SetSkillID(ti)
	return lsc
}

// SetCreatedAt sets the "created_at" field.
func (lsc *LilySkillCreate) SetCreatedAt(t time.Time) *LilySkillCreate {
	lsc.mutation.SetCreatedAt(t)
	return lsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lsc *LilySkillCreate) SetNillableCreatedAt(t *time.Time) *LilySkillCreate {
	if t != nil {
		lsc.SetCreatedAt(*t)
	}
	return lsc
}

// SetUpdatedAt sets the "updated_at" field.
func (lsc *LilySkillCreate) SetUpdatedAt(t time.Time) *LilySkillCreate {
	lsc.mutation.SetUpdatedAt(t)
	return lsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lsc *LilySkillCreate) SetNillableUpdatedAt(t *time.Time) *LilySkillCreate {
	if t != nil {
		lsc.SetUpdatedAt(*t)
	}
	return lsc
}

// Mutation returns the LilySkillMutation object of the builder.
func (lsc *LilySkillCreate) Mutation() *LilySkillMutation {
	return lsc.mutation
}

// Save creates the LilySkill in the database.
func (lsc *LilySkillCreate) Save(ctx context.Context) (*LilySkill, error) {
	var (
		err  error
		node *LilySkill
	)
	lsc.defaults()
	if len(lsc.hooks) == 0 {
		if err = lsc.check(); err != nil {
			return nil, err
		}
		node, err = lsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LilySkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lsc.check(); err != nil {
				return nil, err
			}
			lsc.mutation = mutation
			if node, err = lsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lsc.hooks) - 1; i >= 0; i-- {
			if lsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lsc *LilySkillCreate) SaveX(ctx context.Context) *LilySkill {
	v, err := lsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lsc *LilySkillCreate) Exec(ctx context.Context) error {
	_, err := lsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsc *LilySkillCreate) ExecX(ctx context.Context) {
	if err := lsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lsc *LilySkillCreate) defaults() {
	if _, ok := lsc.mutation.CreatedAt(); !ok {
		v := lilyskill.DefaultCreatedAt()
		lsc.mutation.SetCreatedAt(v)
	}
	if _, ok := lsc.mutation.UpdatedAt(); !ok {
		v := lilyskill.DefaultUpdatedAt()
		lsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lsc *LilySkillCreate) check() error {
	if _, ok := lsc.mutation.LilyID(); !ok {
		return &ValidationError{Name: "lily_id", err: errors.New(`ent: missing required field "lily_id"`)}
	}
	if _, ok := lsc.mutation.SkillID(); !ok {
		return &ValidationError{Name: "skill_id", err: errors.New(`ent: missing required field "skill_id"`)}
	}
	if _, ok := lsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := lsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (lsc *LilySkillCreate) sqlSave(ctx context.Context) (*LilySkill, error) {
	_node, _spec := lsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lsc *LilySkillCreate) createSpec() (*LilySkill, *sqlgraph.CreateSpec) {
	var (
		_node = &LilySkill{config: lsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lilyskill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lilyskill.FieldID,
			},
		}
	)
	_spec.OnConflict = lsc.conflict
	if value, ok := lsc.mutation.LilyID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldLilyID,
		})
		_node.LilyID = value
	}
	if value, ok := lsc.mutation.SkillID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lilyskill.FieldSkillID,
		})
		_node.SkillID = value
	}
	if value, ok := lsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lilyskill.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := lsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lilyskill.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LilySkill.Create().
//		SetLilyID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LilySkillUpsert) {
//			SetLilyID(v+v).
//		}).
//		Exec(ctx)
//
func (lsc *LilySkillCreate) OnConflict(opts ...sql.ConflictOption) *LilySkillUpsertOne {
	lsc.conflict = opts
	return &LilySkillUpsertOne{
		create: lsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LilySkill.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lsc *LilySkillCreate) OnConflictColumns(columns ...string) *LilySkillUpsertOne {
	lsc.conflict = append(lsc.conflict, sql.ConflictColumns(columns...))
	return &LilySkillUpsertOne{
		create: lsc,
	}
}

type (
	// LilySkillUpsertOne is the builder for "upsert"-ing
	//  one LilySkill node.
	LilySkillUpsertOne struct {
		create *LilySkillCreate
	}

	// LilySkillUpsert is the "OnConflict" setter.
	LilySkillUpsert struct {
		*sql.UpdateSet
	}
)

// SetLilyID sets the "lily_id" field.
func (u *LilySkillUpsert) SetLilyID(v types.LilyID) *LilySkillUpsert {
	u.Set(lilyskill.FieldLilyID, v)
	return u
}

// UpdateLilyID sets the "lily_id" field to the value that was provided on create.
func (u *LilySkillUpsert) UpdateLilyID() *LilySkillUpsert {
	u.SetExcluded(lilyskill.FieldLilyID)
	return u
}

// SetSkillID sets the "skill_id" field.
func (u *LilySkillUpsert) SetSkillID(v types.SkillID) *LilySkillUpsert {
	u.Set(lilyskill.FieldSkillID, v)
	return u
}

// UpdateSkillID sets the "skill_id" field to the value that was provided on create.
func (u *LilySkillUpsert) UpdateSkillID() *LilySkillUpsert {
	u.SetExcluded(lilyskill.FieldSkillID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LilySkillUpsert) SetCreatedAt(v time.Time) *LilySkillUpsert {
	u.Set(lilyskill.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LilySkillUpsert) UpdateCreatedAt() *LilySkillUpsert {
	u.SetExcluded(lilyskill.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LilySkillUpsert) SetUpdatedAt(v time.Time) *LilySkillUpsert {
	u.Set(lilyskill.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LilySkillUpsert) UpdateUpdatedAt() *LilySkillUpsert {
	u.SetExcluded(lilyskill.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.LilySkill.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *LilySkillUpsertOne) UpdateNewValues() *LilySkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.LilySkill.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LilySkillUpsertOne) Ignore() *LilySkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LilySkillUpsertOne) DoNothing() *LilySkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LilySkillCreate.OnConflict
// documentation for more info.
func (u *LilySkillUpsertOne) Update(set func(*LilySkillUpsert)) *LilySkillUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LilySkillUpsert{UpdateSet: update})
	}))
	return u
}

// SetLilyID sets the "lily_id" field.
func (u *LilySkillUpsertOne) SetLilyID(v types.LilyID) *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetLilyID(v)
	})
}

// UpdateLilyID sets the "lily_id" field to the value that was provided on create.
func (u *LilySkillUpsertOne) UpdateLilyID() *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateLilyID()
	})
}

// SetSkillID sets the "skill_id" field.
func (u *LilySkillUpsertOne) SetSkillID(v types.SkillID) *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetSkillID(v)
	})
}

// UpdateSkillID sets the "skill_id" field to the value that was provided on create.
func (u *LilySkillUpsertOne) UpdateSkillID() *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateSkillID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *LilySkillUpsertOne) SetCreatedAt(v time.Time) *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LilySkillUpsertOne) UpdateCreatedAt() *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LilySkillUpsertOne) SetUpdatedAt(v time.Time) *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LilySkillUpsertOne) UpdateUpdatedAt() *LilySkillUpsertOne {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *LilySkillUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LilySkillCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LilySkillUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LilySkillUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LilySkillUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LilySkillCreateBulk is the builder for creating many LilySkill entities in bulk.
type LilySkillCreateBulk struct {
	config
	builders []*LilySkillCreate
	conflict []sql.ConflictOption
}

// Save creates the LilySkill entities in the database.
func (lscb *LilySkillCreateBulk) Save(ctx context.Context) ([]*LilySkill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lscb.builders))
	nodes := make([]*LilySkill, len(lscb.builders))
	mutators := make([]Mutator, len(lscb.builders))
	for i := range lscb.builders {
		func(i int, root context.Context) {
			builder := lscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LilySkillMutation)
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
					_, err = mutators[i+1].Mutate(root, lscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lscb *LilySkillCreateBulk) SaveX(ctx context.Context) []*LilySkill {
	v, err := lscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lscb *LilySkillCreateBulk) Exec(ctx context.Context) error {
	_, err := lscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lscb *LilySkillCreateBulk) ExecX(ctx context.Context) {
	if err := lscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LilySkill.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LilySkillUpsert) {
//			SetLilyID(v+v).
//		}).
//		Exec(ctx)
//
func (lscb *LilySkillCreateBulk) OnConflict(opts ...sql.ConflictOption) *LilySkillUpsertBulk {
	lscb.conflict = opts
	return &LilySkillUpsertBulk{
		create: lscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LilySkill.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lscb *LilySkillCreateBulk) OnConflictColumns(columns ...string) *LilySkillUpsertBulk {
	lscb.conflict = append(lscb.conflict, sql.ConflictColumns(columns...))
	return &LilySkillUpsertBulk{
		create: lscb,
	}
}

// LilySkillUpsertBulk is the builder for "upsert"-ing
// a bulk of LilySkill nodes.
type LilySkillUpsertBulk struct {
	create *LilySkillCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LilySkill.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *LilySkillUpsertBulk) UpdateNewValues() *LilySkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LilySkill.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LilySkillUpsertBulk) Ignore() *LilySkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LilySkillUpsertBulk) DoNothing() *LilySkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LilySkillCreateBulk.OnConflict
// documentation for more info.
func (u *LilySkillUpsertBulk) Update(set func(*LilySkillUpsert)) *LilySkillUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LilySkillUpsert{UpdateSet: update})
	}))
	return u
}

// SetLilyID sets the "lily_id" field.
func (u *LilySkillUpsertBulk) SetLilyID(v types.LilyID) *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetLilyID(v)
	})
}

// UpdateLilyID sets the "lily_id" field to the value that was provided on create.
func (u *LilySkillUpsertBulk) UpdateLilyID() *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateLilyID()
	})
}

// SetSkillID sets the "skill_id" field.
func (u *LilySkillUpsertBulk) SetSkillID(v types.SkillID) *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetSkillID(v)
	})
}

// UpdateSkillID sets the "skill_id" field to the value that was provided on create.
func (u *LilySkillUpsertBulk) UpdateSkillID() *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateSkillID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *LilySkillUpsertBulk) SetCreatedAt(v time.Time) *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LilySkillUpsertBulk) UpdateCreatedAt() *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LilySkillUpsertBulk) SetUpdatedAt(v time.Time) *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LilySkillUpsertBulk) UpdateUpdatedAt() *LilySkillUpsertBulk {
	return u.Update(func(s *LilySkillUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *LilySkillUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LilySkillCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LilySkillCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LilySkillUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
