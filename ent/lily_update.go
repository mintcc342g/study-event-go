// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"study-event-go/ent/lily"
	"study-event-go/ent/predicate"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LilyUpdate is the builder for updating Lily entities.
type LilyUpdate struct {
	config
	hooks    []Hook
	mutation *LilyMutation
}

// Where appends a list predicates to the LilyUpdate builder.
func (lu *LilyUpdate) Where(ps ...predicate.Lily) *LilyUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LilyUpdate) SetCreatedAt(t time.Time) *LilyUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableCreatedAt(t *time.Time) *LilyUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LilyUpdate) SetUpdatedAt(t time.Time) *LilyUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetDeletedAt sets the "deleted_at" field.
func (lu *LilyUpdate) SetDeletedAt(t time.Time) *LilyUpdate {
	lu.mutation.SetDeletedAt(t)
	return lu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableDeletedAt(t *time.Time) *LilyUpdate {
	if t != nil {
		lu.SetDeletedAt(*t)
	}
	return lu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (lu *LilyUpdate) ClearDeletedAt() *LilyUpdate {
	lu.mutation.ClearDeletedAt()
	return lu
}

// SetDeletionReason sets the "deletion_reason" field.
func (lu *LilyUpdate) SetDeletionReason(tr types.DeletionReason) *LilyUpdate {
	lu.mutation.ResetDeletionReason()
	lu.mutation.SetDeletionReason(tr)
	return lu
}

// SetNillableDeletionReason sets the "deletion_reason" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableDeletionReason(tr *types.DeletionReason) *LilyUpdate {
	if tr != nil {
		lu.SetDeletionReason(*tr)
	}
	return lu
}

// AddDeletionReason adds tr to the "deletion_reason" field.
func (lu *LilyUpdate) AddDeletionReason(tr types.DeletionReason) *LilyUpdate {
	lu.mutation.AddDeletionReason(tr)
	return lu
}

// SetFirstName sets the "first_name" field.
func (lu *LilyUpdate) SetFirstName(s string) *LilyUpdate {
	lu.mutation.SetFirstName(s)
	return lu
}

// SetMiddleName sets the "middle_name" field.
func (lu *LilyUpdate) SetMiddleName(s string) *LilyUpdate {
	lu.mutation.SetMiddleName(s)
	return lu
}

// SetLastName sets the "last_name" field.
func (lu *LilyUpdate) SetLastName(s string) *LilyUpdate {
	lu.mutation.SetLastName(s)
	return lu
}

// SetRank sets the "rank" field.
func (lu *LilyUpdate) SetRank(u uint32) *LilyUpdate {
	lu.mutation.ResetRank()
	lu.mutation.SetRank(u)
	return lu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableRank(u *uint32) *LilyUpdate {
	if u != nil {
		lu.SetRank(*u)
	}
	return lu
}

// AddRank adds u to the "rank" field.
func (lu *LilyUpdate) AddRank(u uint32) *LilyUpdate {
	lu.mutation.AddRank(u)
	return lu
}

// SetMainCharmID sets the "main_charm_id" field.
func (lu *LilyUpdate) SetMainCharmID(ti types.CharmID) *LilyUpdate {
	lu.mutation.ResetMainCharmID()
	lu.mutation.SetMainCharmID(ti)
	return lu
}

// SetNillableMainCharmID sets the "main_charm_id" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableMainCharmID(ti *types.CharmID) *LilyUpdate {
	if ti != nil {
		lu.SetMainCharmID(*ti)
	}
	return lu
}

// AddMainCharmID adds ti to the "main_charm_id" field.
func (lu *LilyUpdate) AddMainCharmID(ti types.CharmID) *LilyUpdate {
	lu.mutation.AddMainCharmID(ti)
	return lu
}

// SetSubCharmID sets the "sub_charm_id" field.
func (lu *LilyUpdate) SetSubCharmID(ti types.CharmID) *LilyUpdate {
	lu.mutation.ResetSubCharmID()
	lu.mutation.SetSubCharmID(ti)
	return lu
}

// SetNillableSubCharmID sets the "sub_charm_id" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableSubCharmID(ti *types.CharmID) *LilyUpdate {
	if ti != nil {
		lu.SetSubCharmID(*ti)
	}
	return lu
}

// AddSubCharmID adds ti to the "sub_charm_id" field.
func (lu *LilyUpdate) AddSubCharmID(ti types.CharmID) *LilyUpdate {
	lu.mutation.AddSubCharmID(ti)
	return lu
}

// SetGardenID sets the "garden_id" field.
func (lu *LilyUpdate) SetGardenID(ti types.GardenID) *LilyUpdate {
	lu.mutation.ResetGardenID()
	lu.mutation.SetGardenID(ti)
	return lu
}

// SetNillableGardenID sets the "garden_id" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableGardenID(ti *types.GardenID) *LilyUpdate {
	if ti != nil {
		lu.SetGardenID(*ti)
	}
	return lu
}

// AddGardenID adds ti to the "garden_id" field.
func (lu *LilyUpdate) AddGardenID(ti types.GardenID) *LilyUpdate {
	lu.mutation.AddGardenID(ti)
	return lu
}

// SetLegionID sets the "legion_id" field.
func (lu *LilyUpdate) SetLegionID(ti types.LegionID) *LilyUpdate {
	lu.mutation.ResetLegionID()
	lu.mutation.SetLegionID(ti)
	return lu
}

// SetNillableLegionID sets the "legion_id" field if the given value is not nil.
func (lu *LilyUpdate) SetNillableLegionID(ti *types.LegionID) *LilyUpdate {
	if ti != nil {
		lu.SetLegionID(*ti)
	}
	return lu
}

// AddLegionID adds ti to the "legion_id" field.
func (lu *LilyUpdate) AddLegionID(ti types.LegionID) *LilyUpdate {
	lu.mutation.AddLegionID(ti)
	return lu
}

// Mutation returns the LilyMutation object of the builder.
func (lu *LilyUpdate) Mutation() *LilyMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LilyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lu.defaults()
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LilyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LilyUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LilyUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LilyUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LilyUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := lily.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LilyUpdate) check() error {
	if v, ok := lu.mutation.FirstName(); ok {
		if err := lily.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf("ent: validator failed for field \"first_name\": %w", err)}
		}
	}
	if v, ok := lu.mutation.LastName(); ok {
		if err := lily.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf("ent: validator failed for field \"last_name\": %w", err)}
		}
	}
	return nil
}

func (lu *LilyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lily.Table,
			Columns: lily.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: lily.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldCreatedAt,
		})
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldUpdatedAt,
		})
	}
	if value, ok := lu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldDeletedAt,
		})
	}
	if lu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: lily.FieldDeletedAt,
		})
	}
	if value, ok := lu.mutation.DeletionReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldDeletionReason,
		})
	}
	if value, ok := lu.mutation.AddedDeletionReason(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldDeletionReason,
		})
	}
	if value, ok := lu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldFirstName,
		})
	}
	if value, ok := lu.mutation.MiddleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldMiddleName,
		})
	}
	if value, ok := lu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldLastName,
		})
	}
	if value, ok := lu.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldRank,
		})
	}
	if value, ok := lu.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldRank,
		})
	}
	if value, ok := lu.mutation.MainCharmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldMainCharmID,
		})
	}
	if value, ok := lu.mutation.AddedMainCharmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldMainCharmID,
		})
	}
	if value, ok := lu.mutation.SubCharmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldSubCharmID,
		})
	}
	if value, ok := lu.mutation.AddedSubCharmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldSubCharmID,
		})
	}
	if value, ok := lu.mutation.GardenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldGardenID,
		})
	}
	if value, ok := lu.mutation.AddedGardenID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldGardenID,
		})
	}
	if value, ok := lu.mutation.LegionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldLegionID,
		})
	}
	if value, ok := lu.mutation.AddedLegionID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldLegionID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lily.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LilyUpdateOne is the builder for updating a single Lily entity.
type LilyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LilyMutation
}

// SetCreatedAt sets the "created_at" field.
func (luo *LilyUpdateOne) SetCreatedAt(t time.Time) *LilyUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableCreatedAt(t *time.Time) *LilyUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LilyUpdateOne) SetUpdatedAt(t time.Time) *LilyUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetDeletedAt sets the "deleted_at" field.
func (luo *LilyUpdateOne) SetDeletedAt(t time.Time) *LilyUpdateOne {
	luo.mutation.SetDeletedAt(t)
	return luo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableDeletedAt(t *time.Time) *LilyUpdateOne {
	if t != nil {
		luo.SetDeletedAt(*t)
	}
	return luo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (luo *LilyUpdateOne) ClearDeletedAt() *LilyUpdateOne {
	luo.mutation.ClearDeletedAt()
	return luo
}

// SetDeletionReason sets the "deletion_reason" field.
func (luo *LilyUpdateOne) SetDeletionReason(tr types.DeletionReason) *LilyUpdateOne {
	luo.mutation.ResetDeletionReason()
	luo.mutation.SetDeletionReason(tr)
	return luo
}

// SetNillableDeletionReason sets the "deletion_reason" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableDeletionReason(tr *types.DeletionReason) *LilyUpdateOne {
	if tr != nil {
		luo.SetDeletionReason(*tr)
	}
	return luo
}

// AddDeletionReason adds tr to the "deletion_reason" field.
func (luo *LilyUpdateOne) AddDeletionReason(tr types.DeletionReason) *LilyUpdateOne {
	luo.mutation.AddDeletionReason(tr)
	return luo
}

// SetFirstName sets the "first_name" field.
func (luo *LilyUpdateOne) SetFirstName(s string) *LilyUpdateOne {
	luo.mutation.SetFirstName(s)
	return luo
}

// SetMiddleName sets the "middle_name" field.
func (luo *LilyUpdateOne) SetMiddleName(s string) *LilyUpdateOne {
	luo.mutation.SetMiddleName(s)
	return luo
}

// SetLastName sets the "last_name" field.
func (luo *LilyUpdateOne) SetLastName(s string) *LilyUpdateOne {
	luo.mutation.SetLastName(s)
	return luo
}

// SetRank sets the "rank" field.
func (luo *LilyUpdateOne) SetRank(u uint32) *LilyUpdateOne {
	luo.mutation.ResetRank()
	luo.mutation.SetRank(u)
	return luo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableRank(u *uint32) *LilyUpdateOne {
	if u != nil {
		luo.SetRank(*u)
	}
	return luo
}

// AddRank adds u to the "rank" field.
func (luo *LilyUpdateOne) AddRank(u uint32) *LilyUpdateOne {
	luo.mutation.AddRank(u)
	return luo
}

// SetMainCharmID sets the "main_charm_id" field.
func (luo *LilyUpdateOne) SetMainCharmID(ti types.CharmID) *LilyUpdateOne {
	luo.mutation.ResetMainCharmID()
	luo.mutation.SetMainCharmID(ti)
	return luo
}

// SetNillableMainCharmID sets the "main_charm_id" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableMainCharmID(ti *types.CharmID) *LilyUpdateOne {
	if ti != nil {
		luo.SetMainCharmID(*ti)
	}
	return luo
}

// AddMainCharmID adds ti to the "main_charm_id" field.
func (luo *LilyUpdateOne) AddMainCharmID(ti types.CharmID) *LilyUpdateOne {
	luo.mutation.AddMainCharmID(ti)
	return luo
}

// SetSubCharmID sets the "sub_charm_id" field.
func (luo *LilyUpdateOne) SetSubCharmID(ti types.CharmID) *LilyUpdateOne {
	luo.mutation.ResetSubCharmID()
	luo.mutation.SetSubCharmID(ti)
	return luo
}

// SetNillableSubCharmID sets the "sub_charm_id" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableSubCharmID(ti *types.CharmID) *LilyUpdateOne {
	if ti != nil {
		luo.SetSubCharmID(*ti)
	}
	return luo
}

// AddSubCharmID adds ti to the "sub_charm_id" field.
func (luo *LilyUpdateOne) AddSubCharmID(ti types.CharmID) *LilyUpdateOne {
	luo.mutation.AddSubCharmID(ti)
	return luo
}

// SetGardenID sets the "garden_id" field.
func (luo *LilyUpdateOne) SetGardenID(ti types.GardenID) *LilyUpdateOne {
	luo.mutation.ResetGardenID()
	luo.mutation.SetGardenID(ti)
	return luo
}

// SetNillableGardenID sets the "garden_id" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableGardenID(ti *types.GardenID) *LilyUpdateOne {
	if ti != nil {
		luo.SetGardenID(*ti)
	}
	return luo
}

// AddGardenID adds ti to the "garden_id" field.
func (luo *LilyUpdateOne) AddGardenID(ti types.GardenID) *LilyUpdateOne {
	luo.mutation.AddGardenID(ti)
	return luo
}

// SetLegionID sets the "legion_id" field.
func (luo *LilyUpdateOne) SetLegionID(ti types.LegionID) *LilyUpdateOne {
	luo.mutation.ResetLegionID()
	luo.mutation.SetLegionID(ti)
	return luo
}

// SetNillableLegionID sets the "legion_id" field if the given value is not nil.
func (luo *LilyUpdateOne) SetNillableLegionID(ti *types.LegionID) *LilyUpdateOne {
	if ti != nil {
		luo.SetLegionID(*ti)
	}
	return luo
}

// AddLegionID adds ti to the "legion_id" field.
func (luo *LilyUpdateOne) AddLegionID(ti types.LegionID) *LilyUpdateOne {
	luo.mutation.AddLegionID(ti)
	return luo
}

// Mutation returns the LilyMutation object of the builder.
func (luo *LilyUpdateOne) Mutation() *LilyMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LilyUpdateOne) Select(field string, fields ...string) *LilyUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Lily entity.
func (luo *LilyUpdateOne) Save(ctx context.Context) (*Lily, error) {
	var (
		err  error
		node *Lily
	)
	luo.defaults()
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LilyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LilyUpdateOne) SaveX(ctx context.Context) *Lily {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LilyUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LilyUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LilyUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := lily.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LilyUpdateOne) check() error {
	if v, ok := luo.mutation.FirstName(); ok {
		if err := lily.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf("ent: validator failed for field \"first_name\": %w", err)}
		}
	}
	if v, ok := luo.mutation.LastName(); ok {
		if err := lily.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf("ent: validator failed for field \"last_name\": %w", err)}
		}
	}
	return nil
}

func (luo *LilyUpdateOne) sqlSave(ctx context.Context) (_node *Lily, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lily.Table,
			Columns: lily.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: lily.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Lily.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lily.FieldID)
		for _, f := range fields {
			if !lily.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lily.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldCreatedAt,
		})
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldUpdatedAt,
		})
	}
	if value, ok := luo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lily.FieldDeletedAt,
		})
	}
	if luo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: lily.FieldDeletedAt,
		})
	}
	if value, ok := luo.mutation.DeletionReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldDeletionReason,
		})
	}
	if value, ok := luo.mutation.AddedDeletionReason(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldDeletionReason,
		})
	}
	if value, ok := luo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldFirstName,
		})
	}
	if value, ok := luo.mutation.MiddleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldMiddleName,
		})
	}
	if value, ok := luo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lily.FieldLastName,
		})
	}
	if value, ok := luo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldRank,
		})
	}
	if value, ok := luo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: lily.FieldRank,
		})
	}
	if value, ok := luo.mutation.MainCharmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldMainCharmID,
		})
	}
	if value, ok := luo.mutation.AddedMainCharmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldMainCharmID,
		})
	}
	if value, ok := luo.mutation.SubCharmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldSubCharmID,
		})
	}
	if value, ok := luo.mutation.AddedSubCharmID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldSubCharmID,
		})
	}
	if value, ok := luo.mutation.GardenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldGardenID,
		})
	}
	if value, ok := luo.mutation.AddedGardenID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldGardenID,
		})
	}
	if value, ok := luo.mutation.LegionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldLegionID,
		})
	}
	if value, ok := luo.mutation.AddedLegionID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: lily.FieldLegionID,
		})
	}
	_node = &Lily{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lily.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}