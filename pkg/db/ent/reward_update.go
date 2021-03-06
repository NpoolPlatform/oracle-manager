// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/oracle-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/oracle-manager/pkg/db/ent/reward"
	"github.com/google/uuid"
)

// RewardUpdate is the builder for updating Reward entities.
type RewardUpdate struct {
	config
	hooks    []Hook
	mutation *RewardMutation
}

// Where appends a list predicates to the RewardUpdate builder.
func (ru *RewardUpdate) Where(ps ...predicate.Reward) *RewardUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RewardUpdate) SetCreatedAt(u uint32) *RewardUpdate {
	ru.mutation.ResetCreatedAt()
	ru.mutation.SetCreatedAt(u)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *RewardUpdate) SetNillableCreatedAt(u *uint32) *RewardUpdate {
	if u != nil {
		ru.SetCreatedAt(*u)
	}
	return ru
}

// AddCreatedAt adds u to the "created_at" field.
func (ru *RewardUpdate) AddCreatedAt(u int32) *RewardUpdate {
	ru.mutation.AddCreatedAt(u)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RewardUpdate) SetUpdatedAt(u uint32) *RewardUpdate {
	ru.mutation.ResetUpdatedAt()
	ru.mutation.SetUpdatedAt(u)
	return ru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ru *RewardUpdate) AddUpdatedAt(u int32) *RewardUpdate {
	ru.mutation.AddUpdatedAt(u)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RewardUpdate) SetDeletedAt(u uint32) *RewardUpdate {
	ru.mutation.ResetDeletedAt()
	ru.mutation.SetDeletedAt(u)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RewardUpdate) SetNillableDeletedAt(u *uint32) *RewardUpdate {
	if u != nil {
		ru.SetDeletedAt(*u)
	}
	return ru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ru *RewardUpdate) AddDeletedAt(u int32) *RewardUpdate {
	ru.mutation.AddDeletedAt(u)
	return ru
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ru *RewardUpdate) SetCoinTypeID(u uuid.UUID) *RewardUpdate {
	ru.mutation.SetCoinTypeID(u)
	return ru
}

// SetDailyReward sets the "daily_reward" field.
func (ru *RewardUpdate) SetDailyReward(u uint64) *RewardUpdate {
	ru.mutation.ResetDailyReward()
	ru.mutation.SetDailyReward(u)
	return ru
}

// AddDailyReward adds u to the "daily_reward" field.
func (ru *RewardUpdate) AddDailyReward(u int64) *RewardUpdate {
	ru.mutation.AddDailyReward(u)
	return ru
}

// Mutation returns the RewardMutation object of the builder.
func (ru *RewardUpdate) Mutation() *RewardMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RewardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ru.defaults(); err != nil {
		return 0, err
	}
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RewardUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RewardUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RewardUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RewardUpdate) defaults() error {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		if reward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized reward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := reward.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ru *RewardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reward.Table,
			Columns: reward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reward.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldCreatedAt,
		})
	}
	if value, ok := ru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldCreatedAt,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldDeletedAt,
		})
	}
	if value, ok := ru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldDeletedAt,
		})
	}
	if value, ok := ru.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: reward.FieldCoinTypeID,
		})
	}
	if value, ok := ru.mutation.DailyReward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: reward.FieldDailyReward,
		})
	}
	if value, ok := ru.mutation.AddedDailyReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: reward.FieldDailyReward,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RewardUpdateOne is the builder for updating a single Reward entity.
type RewardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RewardMutation
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RewardUpdateOne) SetCreatedAt(u uint32) *RewardUpdateOne {
	ruo.mutation.ResetCreatedAt()
	ruo.mutation.SetCreatedAt(u)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *RewardUpdateOne) SetNillableCreatedAt(u *uint32) *RewardUpdateOne {
	if u != nil {
		ruo.SetCreatedAt(*u)
	}
	return ruo
}

// AddCreatedAt adds u to the "created_at" field.
func (ruo *RewardUpdateOne) AddCreatedAt(u int32) *RewardUpdateOne {
	ruo.mutation.AddCreatedAt(u)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RewardUpdateOne) SetUpdatedAt(u uint32) *RewardUpdateOne {
	ruo.mutation.ResetUpdatedAt()
	ruo.mutation.SetUpdatedAt(u)
	return ruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ruo *RewardUpdateOne) AddUpdatedAt(u int32) *RewardUpdateOne {
	ruo.mutation.AddUpdatedAt(u)
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RewardUpdateOne) SetDeletedAt(u uint32) *RewardUpdateOne {
	ruo.mutation.ResetDeletedAt()
	ruo.mutation.SetDeletedAt(u)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RewardUpdateOne) SetNillableDeletedAt(u *uint32) *RewardUpdateOne {
	if u != nil {
		ruo.SetDeletedAt(*u)
	}
	return ruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ruo *RewardUpdateOne) AddDeletedAt(u int32) *RewardUpdateOne {
	ruo.mutation.AddDeletedAt(u)
	return ruo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ruo *RewardUpdateOne) SetCoinTypeID(u uuid.UUID) *RewardUpdateOne {
	ruo.mutation.SetCoinTypeID(u)
	return ruo
}

// SetDailyReward sets the "daily_reward" field.
func (ruo *RewardUpdateOne) SetDailyReward(u uint64) *RewardUpdateOne {
	ruo.mutation.ResetDailyReward()
	ruo.mutation.SetDailyReward(u)
	return ruo
}

// AddDailyReward adds u to the "daily_reward" field.
func (ruo *RewardUpdateOne) AddDailyReward(u int64) *RewardUpdateOne {
	ruo.mutation.AddDailyReward(u)
	return ruo
}

// Mutation returns the RewardMutation object of the builder.
func (ruo *RewardUpdateOne) Mutation() *RewardMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RewardUpdateOne) Select(field string, fields ...string) *RewardUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reward entity.
func (ruo *RewardUpdateOne) Save(ctx context.Context) (*Reward, error) {
	var (
		err  error
		node *Reward
	)
	if err := ruo.defaults(); err != nil {
		return nil, err
	}
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RewardUpdateOne) SaveX(ctx context.Context) *Reward {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RewardUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RewardUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RewardUpdateOne) defaults() error {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		if reward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized reward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := reward.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ruo *RewardUpdateOne) sqlSave(ctx context.Context) (_node *Reward, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reward.Table,
			Columns: reward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reward.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reward.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reward.FieldID)
		for _, f := range fields {
			if !reward.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldCreatedAt,
		})
	}
	if value, ok := ruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldCreatedAt,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldDeletedAt,
		})
	}
	if value, ok := ruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reward.FieldDeletedAt,
		})
	}
	if value, ok := ruo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: reward.FieldCoinTypeID,
		})
	}
	if value, ok := ruo.mutation.DailyReward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: reward.FieldDailyReward,
		})
	}
	if value, ok := ruo.mutation.AddedDailyReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: reward.FieldDailyReward,
		})
	}
	_node = &Reward{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
