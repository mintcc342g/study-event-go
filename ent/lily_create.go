// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"study-event-go/ent/lily"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LilyCreate is the builder for creating a Lily entity.
type LilyCreate struct {
	config
	mutation *LilyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (lc *LilyCreate) SetCreatedAt(t time.Time) *LilyCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LilyCreate) SetNillableCreatedAt(t *time.Time) *LilyCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LilyCreate) SetUpdatedAt(t time.Time) *LilyCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LilyCreate) SetNillableUpdatedAt(t *time.Time) *LilyCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetDeletedAt sets the "deleted_at" field.
func (lc *LilyCreate) SetDeletedAt(t time.Time) *LilyCreate {
	lc.mutation.SetDeletedAt(t)
	return lc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lc *LilyCreate) SetNillableDeletedAt(t *time.Time) *LilyCreate {
	if t != nil {
		lc.SetDeletedAt(*t)
	}
	return lc
}

// SetDeletionReason sets the "deletion_reason" field.
func (lc *LilyCreate) SetDeletionReason(tr types.DeletionReason) *LilyCreate {
	lc.mutation.SetDeletionReason(tr)
	return lc
}

// SetNillableDeletionReason sets the "deletion_reason" field if the given value is not nil.
func (lc *LilyCreate) SetNillableDeletionReason(tr *types.DeletionReason) *LilyCreate {
	if tr != nil {
		lc.SetDeletionReason(*tr)
	}
	return lc
}

// SetFirstName sets the "first_name" field.
func (lc *LilyCreate) SetFirstName(s string) *LilyCreate {
	lc.mutation.SetFirstName(s)
	return lc
}

// SetMiddleName sets the "middle_name" field.
func (lc *LilyCreate) SetMiddleName(s string) *LilyCreate {
	lc.mutation.SetMiddleName(s)
	return lc
}

// SetLastName sets the "last_name" field.
func (lc *LilyCreate) SetLastName(s string) *LilyCreate {
	lc.mutation.SetLastName(s)
	return lc
}

// SetRank sets the "rank" field.
func (lc *LilyCreate) SetRank(u uint32) *LilyCreate {
	lc.mutation.SetRank(u)
	return lc
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (lc *LilyCreate) SetNillableRank(u *uint32) *LilyCreate {
	if u != nil {
		lc.SetRank(*u)
	}
	return lc
}

// SetMainCharmID sets the "main_charm_id" field.
func (lc *LilyCreate) SetMainCharmID(ti types.CharmID) *LilyCreate {
	lc.mutation.SetMainCharmID(ti)
	return lc
}

// SetNillableMainCharmID sets the "main_charm_id" field if the given value is not nil.
func (lc *LilyCreate) SetNillableMainCharmID(ti *types.CharmID) *LilyCreate {
	if ti != nil {
		lc.SetMainCharmID(*ti)
	}
	return lc
}

// SetSubCharmID sets the "sub_charm_id" field.
func (lc *LilyCreate) SetSubCharmID(ti types.CharmID) *LilyCreate {
	lc.mutation.SetSubCharmID(ti)
	return lc
}

// SetNillableSubCharmID sets the "sub_charm_id" field if the given value is not nil.
func (lc *LilyCreate) SetNillableSubCharmID(ti *types.CharmID) *LilyCreate {
	if ti != nil {
		lc.SetSubCharmID(*ti)
	}
	return lc
}

// SetGardenID sets the "garden_id" field.
func (lc *LilyCreate) SetGardenID(ti types.GardenID) *LilyCreate {
	lc.mutation.SetGardenID(ti)
	return lc
}

// SetNillableGardenID sets the "garden_id" field if the given value is not nil.
func (lc *LilyCreate) SetNillableGardenID(ti *types.GardenID) *LilyCreate {
	if ti != nil {
		lc.SetGardenID(*ti)
	}
	return lc
}

// SetLegionID sets the "legion_id" field.
func (lc *LilyCreate) SetLegionID(ti types.LegionID) *LilyCreate {
	lc.mutation.SetLegionID(ti)
	return lc
}

// SetNillableLegionID sets the "legion_id" field if the given value is not nil.
func (lc *LilyCreate) SetNillableLegionID(ti *types.LegionID) *LilyCreate {
	if ti != nil {
		lc.SetLegionID(*ti)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LilyCreate) SetID(ti types.LilyID) *LilyCreate {
	lc.mutation.SetID(ti)
	return lc
}

// Mutation returns the LilyMutation object of the builder.
func (lc *LilyCreate) Mutation() *LilyMutation {
	return lc.mutation
}

// Save creates the Lily in the database.
func (lc *LilyCreate) Save(ctx context.Context) (*Lily, error) {
	var (
		err  error
		node *Lily
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LilyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LilyCreate) SaveX(ctx context.Context) *Lily {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LilyCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LilyCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LilyCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := lily.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := lily.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.DeletionReason(); !ok {
		v := lily.DefaultDeletionReason
		lc.mutation.SetDeletionReason(v)
	}
	if _, ok := lc.mutation.Rank(); !ok {
		v := lily.DefaultRank
		lc.mutation.SetRank(v)
	}
	if _, ok := lc.mutation.MainCharmID(); !ok {
		v := lily.DefaultMainCharmID
		lc.mutation.SetMainCharmID(v)
	}
	if _, ok := lc.mutation.SubCharmID(); !ok {
		v := lily.DefaultSubCharmID
		lc.mutation.SetSubCharmID(v)
	}
	if _, ok := lc.mutation.GardenID(); !ok {
		v := lily.DefaultGardenID
		lc.mutation.SetGardenID(v)
	}
	if _, ok := lc.mutation.LegionID(); !ok {
		v := lily.DefaultLegionID
		lc.mutation.SetLegionID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LilyCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := lc.mutation.DeletionReason(); !ok {
		return &ValidationError{Name: "deletion_reason", err: errors.New(`ent: missing required field "deletion_reason"`)}
	}
	if _, ok := lc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "first_name"`)}
	}
	if v, ok := lc.mutation.FirstName(); ok {
		if err := lily.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "first_name": %w`, err)}
		}
	}
	if _, ok := lc.mutation.MiddleName(); !ok {
		return &ValidationError{Name: "middle_name", err: errors.New(`ent: missing required field "middle_name"`)}
	}
	if _, ok := lc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "last_name"`)}
	}
	if v, ok := lc.mutation.LastName(); ok {
		if err := lily.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "last_name": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Rank(); !ok {
		return &ValidationError{Name: "rank", err: errors.New(`ent: missing required field "rank"`)}
	}
	if _, ok := lc.mutation.MainCharmID(); !ok {
		return &ValidationError{Name: "main_charm_id", err: errors.New(`ent: missing required field "main_charm_id"`)}
	}
	if _, ok := lc.mutation.SubCharmID(); !ok {
		return &ValidationError{Name: "sub_charm_id", err: errors.New(`ent: missing required field "sub_charm_id"`)}
	}
	if _, ok := lc.mutation.GardenID(); !ok {
		return &ValidationError{Name: "garden_id", err: errors.New(`ent: missing required field "garden_id"`)}
	}
	if _, ok := lc.mutation.LegionID(); !ok {
		return &ValidationError{Name: "legion_id", err: errors.New(`ent: missing required field "legion_id"`)}
	}
	return nil
}

func (lc *LilyCreate) sqlSave(ctx context.Context) (*Lily, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = types.LilyID(id)
	}
	return _node, nil
}

func (lc *LilyCreate) createSpec() (*Lily, *sqlgraph.CreateSpec) {
	var (
		_node = &Lily{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lily.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: lily.FieldID,
			},
		}
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := lc.mutation.DeletionReason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldDeletionReason,
		})
		_node.DeletionReason = value
	}
	if value, ok := lc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := lc.mutation.MiddleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldMiddleName,
		})
		_node.MiddleName = value
	}
	if value, ok := lc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := lc.mutation.Rank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldRank,
		})
		_node.Rank = value
	}
	if value, ok := lc.mutation.MainCharmID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldMainCharmID,
		})
		_node.MainCharmID = value
	}
	if value, ok := lc.mutation.SubCharmID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldSubCharmID,
		})
		_node.SubCharmID = value
	}
	if value, ok := lc.mutation.GardenID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldGardenID,
		})
		_node.GardenID = value
	}
	if value, ok := lc.mutation.LegionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldLegionID,
		})
		_node.LegionID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lily.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LilyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (lc *LilyCreate) OnConflict(opts ...sql.ConflictOption) *LilyUpsertOne {
	lc.conflict = opts
	return &LilyUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lily.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lc *LilyCreate) OnConflictColumns(columns ...string) *LilyUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LilyUpsertOne{
		create: lc,
	}
}

type (
	// LilyUpsertOne is the builder for "upsert"-ing
	//  one Lily node.
	LilyUpsertOne struct {
		create *LilyCreate
	}

	// LilyUpsert is the "OnConflict" setter.
	LilyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *LilyUpsert) SetCreatedAt(v time.Time) *LilyUpsert {
	u.Set(lily.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LilyUpsert) UpdateCreatedAt() *LilyUpsert {
	u.SetExcluded(lily.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LilyUpsert) SetUpdatedAt(v time.Time) *LilyUpsert {
	u.Set(lily.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LilyUpsert) UpdateUpdatedAt() *LilyUpsert {
	u.SetExcluded(lily.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LilyUpsert) SetDeletedAt(v time.Time) *LilyUpsert {
	u.Set(lily.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LilyUpsert) UpdateDeletedAt() *LilyUpsert {
	u.SetExcluded(lily.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *LilyUpsert) ClearDeletedAt() *LilyUpsert {
	u.SetNull(lily.FieldDeletedAt)
	return u
}

// SetDeletionReason sets the "deletion_reason" field.
func (u *LilyUpsert) SetDeletionReason(v types.DeletionReason) *LilyUpsert {
	u.Set(lily.FieldDeletionReason, v)
	return u
}

// UpdateDeletionReason sets the "deletion_reason" field to the value that was provided on create.
func (u *LilyUpsert) UpdateDeletionReason() *LilyUpsert {
	u.SetExcluded(lily.FieldDeletionReason)
	return u
}

// SetFirstName sets the "first_name" field.
func (u *LilyUpsert) SetFirstName(v string) *LilyUpsert {
	u.Set(lily.FieldFirstName, v)
	return u
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *LilyUpsert) UpdateFirstName() *LilyUpsert {
	u.SetExcluded(lily.FieldFirstName)
	return u
}

// SetMiddleName sets the "middle_name" field.
func (u *LilyUpsert) SetMiddleName(v string) *LilyUpsert {
	u.Set(lily.FieldMiddleName, v)
	return u
}

// UpdateMiddleName sets the "middle_name" field to the value that was provided on create.
func (u *LilyUpsert) UpdateMiddleName() *LilyUpsert {
	u.SetExcluded(lily.FieldMiddleName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *LilyUpsert) SetLastName(v string) *LilyUpsert {
	u.Set(lily.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *LilyUpsert) UpdateLastName() *LilyUpsert {
	u.SetExcluded(lily.FieldLastName)
	return u
}

// SetRank sets the "rank" field.
func (u *LilyUpsert) SetRank(v uint32) *LilyUpsert {
	u.Set(lily.FieldRank, v)
	return u
}

// UpdateRank sets the "rank" field to the value that was provided on create.
func (u *LilyUpsert) UpdateRank() *LilyUpsert {
	u.SetExcluded(lily.FieldRank)
	return u
}

// SetMainCharmID sets the "main_charm_id" field.
func (u *LilyUpsert) SetMainCharmID(v types.CharmID) *LilyUpsert {
	u.Set(lily.FieldMainCharmID, v)
	return u
}

// UpdateMainCharmID sets the "main_charm_id" field to the value that was provided on create.
func (u *LilyUpsert) UpdateMainCharmID() *LilyUpsert {
	u.SetExcluded(lily.FieldMainCharmID)
	return u
}

// SetSubCharmID sets the "sub_charm_id" field.
func (u *LilyUpsert) SetSubCharmID(v types.CharmID) *LilyUpsert {
	u.Set(lily.FieldSubCharmID, v)
	return u
}

// UpdateSubCharmID sets the "sub_charm_id" field to the value that was provided on create.
func (u *LilyUpsert) UpdateSubCharmID() *LilyUpsert {
	u.SetExcluded(lily.FieldSubCharmID)
	return u
}

// SetGardenID sets the "garden_id" field.
func (u *LilyUpsert) SetGardenID(v types.GardenID) *LilyUpsert {
	u.Set(lily.FieldGardenID, v)
	return u
}

// UpdateGardenID sets the "garden_id" field to the value that was provided on create.
func (u *LilyUpsert) UpdateGardenID() *LilyUpsert {
	u.SetExcluded(lily.FieldGardenID)
	return u
}

// SetLegionID sets the "legion_id" field.
func (u *LilyUpsert) SetLegionID(v types.LegionID) *LilyUpsert {
	u.Set(lily.FieldLegionID, v)
	return u
}

// UpdateLegionID sets the "legion_id" field to the value that was provided on create.
func (u *LilyUpsert) UpdateLegionID() *LilyUpsert {
	u.SetExcluded(lily.FieldLegionID)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Lily.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lily.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LilyUpsertOne) UpdateNewValues() *LilyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(lily.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Lily.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LilyUpsertOne) Ignore() *LilyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LilyUpsertOne) DoNothing() *LilyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LilyCreate.OnConflict
// documentation for more info.
func (u *LilyUpsertOne) Update(set func(*LilyUpsert)) *LilyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LilyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LilyUpsertOne) SetCreatedAt(v time.Time) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateCreatedAt() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LilyUpsertOne) SetUpdatedAt(v time.Time) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateUpdatedAt() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LilyUpsertOne) SetDeletedAt(v time.Time) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateDeletedAt() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *LilyUpsertOne) ClearDeletedAt() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.ClearDeletedAt()
	})
}

// SetDeletionReason sets the "deletion_reason" field.
func (u *LilyUpsertOne) SetDeletionReason(v types.DeletionReason) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetDeletionReason(v)
	})
}

// UpdateDeletionReason sets the "deletion_reason" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateDeletionReason() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateDeletionReason()
	})
}

// SetFirstName sets the "first_name" field.
func (u *LilyUpsertOne) SetFirstName(v string) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateFirstName() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateFirstName()
	})
}

// SetMiddleName sets the "middle_name" field.
func (u *LilyUpsertOne) SetMiddleName(v string) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetMiddleName(v)
	})
}

// UpdateMiddleName sets the "middle_name" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateMiddleName() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateMiddleName()
	})
}

// SetLastName sets the "last_name" field.
func (u *LilyUpsertOne) SetLastName(v string) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateLastName() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateLastName()
	})
}

// SetRank sets the "rank" field.
func (u *LilyUpsertOne) SetRank(v uint32) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetRank(v)
	})
}

// UpdateRank sets the "rank" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateRank() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateRank()
	})
}

// SetMainCharmID sets the "main_charm_id" field.
func (u *LilyUpsertOne) SetMainCharmID(v types.CharmID) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetMainCharmID(v)
	})
}

// UpdateMainCharmID sets the "main_charm_id" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateMainCharmID() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateMainCharmID()
	})
}

// SetSubCharmID sets the "sub_charm_id" field.
func (u *LilyUpsertOne) SetSubCharmID(v types.CharmID) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetSubCharmID(v)
	})
}

// UpdateSubCharmID sets the "sub_charm_id" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateSubCharmID() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateSubCharmID()
	})
}

// SetGardenID sets the "garden_id" field.
func (u *LilyUpsertOne) SetGardenID(v types.GardenID) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetGardenID(v)
	})
}

// UpdateGardenID sets the "garden_id" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateGardenID() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateGardenID()
	})
}

// SetLegionID sets the "legion_id" field.
func (u *LilyUpsertOne) SetLegionID(v types.LegionID) *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.SetLegionID(v)
	})
}

// UpdateLegionID sets the "legion_id" field to the value that was provided on create.
func (u *LilyUpsertOne) UpdateLegionID() *LilyUpsertOne {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateLegionID()
	})
}

// Exec executes the query.
func (u *LilyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LilyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LilyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LilyUpsertOne) ID(ctx context.Context) (id types.LilyID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LilyUpsertOne) IDX(ctx context.Context) types.LilyID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LilyCreateBulk is the builder for creating many Lily entities in bulk.
type LilyCreateBulk struct {
	config
	builders []*LilyCreate
	conflict []sql.ConflictOption
}

// Save creates the Lily entities in the database.
func (lcb *LilyCreateBulk) Save(ctx context.Context) ([]*Lily, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lily, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LilyMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
					nodes[i].ID = types.LilyID(id)
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LilyCreateBulk) SaveX(ctx context.Context) []*Lily {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LilyCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LilyCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lily.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LilyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (lcb *LilyCreateBulk) OnConflict(opts ...sql.ConflictOption) *LilyUpsertBulk {
	lcb.conflict = opts
	return &LilyUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lily.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lcb *LilyCreateBulk) OnConflictColumns(columns ...string) *LilyUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LilyUpsertBulk{
		create: lcb,
	}
}

// LilyUpsertBulk is the builder for "upsert"-ing
// a bulk of Lily nodes.
type LilyUpsertBulk struct {
	create *LilyCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Lily.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lily.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LilyUpsertBulk) UpdateNewValues() *LilyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(lily.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Lily.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LilyUpsertBulk) Ignore() *LilyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LilyUpsertBulk) DoNothing() *LilyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LilyCreateBulk.OnConflict
// documentation for more info.
func (u *LilyUpsertBulk) Update(set func(*LilyUpsert)) *LilyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LilyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LilyUpsertBulk) SetCreatedAt(v time.Time) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateCreatedAt() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LilyUpsertBulk) SetUpdatedAt(v time.Time) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateUpdatedAt() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LilyUpsertBulk) SetDeletedAt(v time.Time) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateDeletedAt() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *LilyUpsertBulk) ClearDeletedAt() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.ClearDeletedAt()
	})
}

// SetDeletionReason sets the "deletion_reason" field.
func (u *LilyUpsertBulk) SetDeletionReason(v types.DeletionReason) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetDeletionReason(v)
	})
}

// UpdateDeletionReason sets the "deletion_reason" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateDeletionReason() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateDeletionReason()
	})
}

// SetFirstName sets the "first_name" field.
func (u *LilyUpsertBulk) SetFirstName(v string) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateFirstName() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateFirstName()
	})
}

// SetMiddleName sets the "middle_name" field.
func (u *LilyUpsertBulk) SetMiddleName(v string) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetMiddleName(v)
	})
}

// UpdateMiddleName sets the "middle_name" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateMiddleName() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateMiddleName()
	})
}

// SetLastName sets the "last_name" field.
func (u *LilyUpsertBulk) SetLastName(v string) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateLastName() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateLastName()
	})
}

// SetRank sets the "rank" field.
func (u *LilyUpsertBulk) SetRank(v uint32) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetRank(v)
	})
}

// UpdateRank sets the "rank" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateRank() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateRank()
	})
}

// SetMainCharmID sets the "main_charm_id" field.
func (u *LilyUpsertBulk) SetMainCharmID(v types.CharmID) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetMainCharmID(v)
	})
}

// UpdateMainCharmID sets the "main_charm_id" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateMainCharmID() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateMainCharmID()
	})
}

// SetSubCharmID sets the "sub_charm_id" field.
func (u *LilyUpsertBulk) SetSubCharmID(v types.CharmID) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetSubCharmID(v)
	})
}

// UpdateSubCharmID sets the "sub_charm_id" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateSubCharmID() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateSubCharmID()
	})
}

// SetGardenID sets the "garden_id" field.
func (u *LilyUpsertBulk) SetGardenID(v types.GardenID) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetGardenID(v)
	})
}

// UpdateGardenID sets the "garden_id" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateGardenID() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateGardenID()
	})
}

// SetLegionID sets the "legion_id" field.
func (u *LilyUpsertBulk) SetLegionID(v types.LegionID) *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.SetLegionID(v)
	})
}

// UpdateLegionID sets the "legion_id" field to the value that was provided on create.
func (u *LilyUpsertBulk) UpdateLegionID() *LilyUpsertBulk {
	return u.Update(func(s *LilyUpsert) {
		s.UpdateLegionID()
	})
}

// Exec executes the query.
func (u *LilyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LilyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LilyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LilyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
