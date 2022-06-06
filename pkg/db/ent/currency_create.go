// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/oracle-manager/pkg/db/ent/currency"
	"github.com/google/uuid"
)

// CurrencyCreate is the builder for creating a Currency entity.
type CurrencyCreate struct {
	config
	mutation *CurrencyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CurrencyCreate) SetCreatedAt(u uint32) *CurrencyCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableCreatedAt(u *uint32) *CurrencyCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CurrencyCreate) SetUpdatedAt(u uint32) *CurrencyCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableUpdatedAt(u *uint32) *CurrencyCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CurrencyCreate) SetDeletedAt(u uint32) *CurrencyCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableDeletedAt(u *uint32) *CurrencyCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cc *CurrencyCreate) SetCoinTypeID(u uuid.UUID) *CurrencyCreate {
	cc.mutation.SetCoinTypeID(u)
	return cc
}

// SetPriceVsUsdt sets the "price_vs_usdt" field.
func (cc *CurrencyCreate) SetPriceVsUsdt(u uint64) *CurrencyCreate {
	cc.mutation.SetPriceVsUsdt(u)
	return cc
}

// SetID sets the "id" field.
func (cc *CurrencyCreate) SetID(u uuid.UUID) *CurrencyCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableID(u *uuid.UUID) *CurrencyCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CurrencyMutation object of the builder.
func (cc *CurrencyCreate) Mutation() *CurrencyMutation {
	return cc.mutation
}

// Save creates the Currency in the database.
func (cc *CurrencyCreate) Save(ctx context.Context) (*Currency, error) {
	var (
		err  error
		node *Currency
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyMutation)
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
func (cc *CurrencyCreate) SaveX(ctx context.Context) *Currency {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CurrencyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CurrencyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CurrencyCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if currency.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized currency.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := currency.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if currency.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized currency.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := currency.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if currency.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized currency.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := currency.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if currency.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized currency.DefaultID (forgotten import ent/runtime?)")
		}
		v := currency.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CurrencyCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Currency.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Currency.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Currency.deleted_at"`)}
	}
	if _, ok := cc.mutation.CoinTypeID(); !ok {
		return &ValidationError{Name: "coin_type_id", err: errors.New(`ent: missing required field "Currency.coin_type_id"`)}
	}
	if _, ok := cc.mutation.PriceVsUsdt(); !ok {
		return &ValidationError{Name: "price_vs_usdt", err: errors.New(`ent: missing required field "Currency.price_vs_usdt"`)}
	}
	return nil
}

func (cc *CurrencyCreate) sqlSave(ctx context.Context) (*Currency, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cc *CurrencyCreate) createSpec() (*Currency, *sqlgraph.CreateSpec) {
	var (
		_node = &Currency{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: currency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currency.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currency.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currency.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currency.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currency.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := cc.mutation.PriceVsUsdt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: currency.FieldPriceVsUsdt,
		})
		_node.PriceVsUsdt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Currency.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CurrencyCreate) OnConflict(opts ...sql.ConflictOption) *CurrencyUpsertOne {
	cc.conflict = opts
	return &CurrencyUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CurrencyCreate) OnConflictColumns(columns ...string) *CurrencyUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CurrencyUpsertOne{
		create: cc,
	}
}

type (
	// CurrencyUpsertOne is the builder for "upsert"-ing
	//  one Currency node.
	CurrencyUpsertOne struct {
		create *CurrencyCreate
	}

	// CurrencyUpsert is the "OnConflict" setter.
	CurrencyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyUpsert) SetCreatedAt(v uint32) *CurrencyUpsert {
	u.Set(currency.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateCreatedAt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyUpsert) AddCreatedAt(v uint32) *CurrencyUpsert {
	u.Add(currency.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyUpsert) SetUpdatedAt(v uint32) *CurrencyUpsert {
	u.Set(currency.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateUpdatedAt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyUpsert) AddUpdatedAt(v uint32) *CurrencyUpsert {
	u.Add(currency.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyUpsert) SetDeletedAt(v uint32) *CurrencyUpsert {
	u.Set(currency.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateDeletedAt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyUpsert) AddDeletedAt(v uint32) *CurrencyUpsert {
	u.Add(currency.FieldDeletedAt, v)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyUpsert) SetCoinTypeID(v uuid.UUID) *CurrencyUpsert {
	u.Set(currency.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateCoinTypeID() *CurrencyUpsert {
	u.SetExcluded(currency.FieldCoinTypeID)
	return u
}

// SetPriceVsUsdt sets the "price_vs_usdt" field.
func (u *CurrencyUpsert) SetPriceVsUsdt(v uint64) *CurrencyUpsert {
	u.Set(currency.FieldPriceVsUsdt, v)
	return u
}

// UpdatePriceVsUsdt sets the "price_vs_usdt" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdatePriceVsUsdt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldPriceVsUsdt)
	return u
}

// AddPriceVsUsdt adds v to the "price_vs_usdt" field.
func (u *CurrencyUpsert) AddPriceVsUsdt(v uint64) *CurrencyUpsert {
	u.Add(currency.FieldPriceVsUsdt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currency.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyUpsertOne) UpdateNewValues() *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(currency.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Currency.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CurrencyUpsertOne) Ignore() *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyUpsertOne) DoNothing() *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyCreate.OnConflict
// documentation for more info.
func (u *CurrencyUpsertOne) Update(set func(*CurrencyUpsert)) *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyUpsertOne) SetCreatedAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyUpsertOne) AddCreatedAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateCreatedAt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyUpsertOne) SetUpdatedAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyUpsertOne) AddUpdatedAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateUpdatedAt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyUpsertOne) SetDeletedAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyUpsertOne) AddDeletedAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateDeletedAt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyUpsertOne) SetCoinTypeID(v uuid.UUID) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateCoinTypeID() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetPriceVsUsdt sets the "price_vs_usdt" field.
func (u *CurrencyUpsertOne) SetPriceVsUsdt(v uint64) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetPriceVsUsdt(v)
	})
}

// AddPriceVsUsdt adds v to the "price_vs_usdt" field.
func (u *CurrencyUpsertOne) AddPriceVsUsdt(v uint64) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddPriceVsUsdt(v)
	})
}

// UpdatePriceVsUsdt sets the "price_vs_usdt" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdatePriceVsUsdt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdatePriceVsUsdt()
	})
}

// Exec executes the query.
func (u *CurrencyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CurrencyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CurrencyUpsertOne.ID is not supported by MySQL driver. Use CurrencyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CurrencyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CurrencyCreateBulk is the builder for creating many Currency entities in bulk.
type CurrencyCreateBulk struct {
	config
	builders []*CurrencyCreate
	conflict []sql.ConflictOption
}

// Save creates the Currency entities in the database.
func (ccb *CurrencyCreateBulk) Save(ctx context.Context) ([]*Currency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Currency, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CurrencyMutation)
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
func (ccb *CurrencyCreateBulk) SaveX(ctx context.Context) []*Currency {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CurrencyCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CurrencyCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Currency.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CurrencyCreateBulk) OnConflict(opts ...sql.ConflictOption) *CurrencyUpsertBulk {
	ccb.conflict = opts
	return &CurrencyUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CurrencyCreateBulk) OnConflictColumns(columns ...string) *CurrencyUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CurrencyUpsertBulk{
		create: ccb,
	}
}

// CurrencyUpsertBulk is the builder for "upsert"-ing
// a bulk of Currency nodes.
type CurrencyUpsertBulk struct {
	create *CurrencyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currency.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyUpsertBulk) UpdateNewValues() *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(currency.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CurrencyUpsertBulk) Ignore() *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyUpsertBulk) DoNothing() *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyCreateBulk.OnConflict
// documentation for more info.
func (u *CurrencyUpsertBulk) Update(set func(*CurrencyUpsert)) *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyUpsertBulk) SetCreatedAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyUpsertBulk) AddCreatedAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateCreatedAt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyUpsertBulk) SetUpdatedAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyUpsertBulk) AddUpdatedAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateUpdatedAt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyUpsertBulk) SetDeletedAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyUpsertBulk) AddDeletedAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateDeletedAt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyUpsertBulk) SetCoinTypeID(v uuid.UUID) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateCoinTypeID() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetPriceVsUsdt sets the "price_vs_usdt" field.
func (u *CurrencyUpsertBulk) SetPriceVsUsdt(v uint64) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetPriceVsUsdt(v)
	})
}

// AddPriceVsUsdt adds v to the "price_vs_usdt" field.
func (u *CurrencyUpsertBulk) AddPriceVsUsdt(v uint64) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.AddPriceVsUsdt(v)
	})
}

// UpdatePriceVsUsdt sets the "price_vs_usdt" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdatePriceVsUsdt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdatePriceVsUsdt()
	})
}

// Exec executes the query.
func (u *CurrencyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CurrencyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
