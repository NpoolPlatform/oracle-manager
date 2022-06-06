// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CurrenciesColumns holds the columns for the "currencies" table.
	CurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Unique: true},
		{Name: "price_vs_usdt", Type: field.TypeUint64},
	}
	// CurrenciesTable holds the schema information for the "currencies" table.
	CurrenciesTable = &schema.Table{
		Name:       "currencies",
		Columns:    CurrenciesColumns,
		PrimaryKey: []*schema.Column{CurrenciesColumns[0]},
	}
	// RewardsColumns holds the columns for the "rewards" table.
	RewardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Unique: true},
		{Name: "daily_reward", Type: field.TypeUint64},
	}
	// RewardsTable holds the schema information for the "rewards" table.
	RewardsTable = &schema.Table{
		Name:       "rewards",
		Columns:    RewardsColumns,
		PrimaryKey: []*schema.Column{RewardsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CurrenciesTable,
		RewardsTable,
	}
)

func init() {
}
