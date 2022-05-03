// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/oracle-manager/pkg/db/ent/reward"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   reward.Table,
			Columns: reward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reward.FieldID,
			},
		},
		Type: "Reward",
		Fields: map[string]*sqlgraph.FieldSpec{
			reward.FieldCreatedAt:   {Type: field.TypeUint32, Column: reward.FieldCreatedAt},
			reward.FieldUpdatedAt:   {Type: field.TypeUint32, Column: reward.FieldUpdatedAt},
			reward.FieldDeletedAt:   {Type: field.TypeUint32, Column: reward.FieldDeletedAt},
			reward.FieldCoinTypeID:  {Type: field.TypeUUID, Column: reward.FieldCoinTypeID},
			reward.FieldDailyReward: {Type: field.TypeUint64, Column: reward.FieldDailyReward},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (rq *RewardQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RewardQuery builder.
func (rq *RewardQuery) Filter() *RewardFilter {
	return &RewardFilter{rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RewardMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RewardMutation builder.
func (m *RewardMutation) Filter() *RewardFilter {
	return &RewardFilter{m}
}

// RewardFilter provides a generic filtering capability at runtime for RewardQuery.
type RewardFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *RewardFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *RewardFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(reward.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *RewardFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(reward.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *RewardFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(reward.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *RewardFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(reward.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *RewardFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(reward.FieldCoinTypeID))
}

// WhereDailyReward applies the entql uint64 predicate on the daily_reward field.
func (f *RewardFilter) WhereDailyReward(p entql.Uint64P) {
	f.Where(p.Field(reward.FieldDailyReward))
}
