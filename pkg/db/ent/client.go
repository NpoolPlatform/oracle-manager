// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/oracle-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/oracle-manager/pkg/db/ent/reward"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Reward is the client for interacting with the Reward builders.
	Reward *RewardClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Reward = NewRewardClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Reward: NewRewardClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Reward: NewRewardClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Reward.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Reward.Use(hooks...)
}

// RewardClient is a client for the Reward schema.
type RewardClient struct {
	config
}

// NewRewardClient returns a client for the Reward from the given config.
func NewRewardClient(c config) *RewardClient {
	return &RewardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `reward.Hooks(f(g(h())))`.
func (c *RewardClient) Use(hooks ...Hook) {
	c.hooks.Reward = append(c.hooks.Reward, hooks...)
}

// Create returns a create builder for Reward.
func (c *RewardClient) Create() *RewardCreate {
	mutation := newRewardMutation(c.config, OpCreate)
	return &RewardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Reward entities.
func (c *RewardClient) CreateBulk(builders ...*RewardCreate) *RewardCreateBulk {
	return &RewardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Reward.
func (c *RewardClient) Update() *RewardUpdate {
	mutation := newRewardMutation(c.config, OpUpdate)
	return &RewardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RewardClient) UpdateOne(r *Reward) *RewardUpdateOne {
	mutation := newRewardMutation(c.config, OpUpdateOne, withReward(r))
	return &RewardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RewardClient) UpdateOneID(id uuid.UUID) *RewardUpdateOne {
	mutation := newRewardMutation(c.config, OpUpdateOne, withRewardID(id))
	return &RewardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Reward.
func (c *RewardClient) Delete() *RewardDelete {
	mutation := newRewardMutation(c.config, OpDelete)
	return &RewardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RewardClient) DeleteOne(r *Reward) *RewardDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RewardClient) DeleteOneID(id uuid.UUID) *RewardDeleteOne {
	builder := c.Delete().Where(reward.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RewardDeleteOne{builder}
}

// Query returns a query builder for Reward.
func (c *RewardClient) Query() *RewardQuery {
	return &RewardQuery{
		config: c.config,
	}
}

// Get returns a Reward entity by its id.
func (c *RewardClient) Get(ctx context.Context, id uuid.UUID) (*Reward, error) {
	return c.Query().Where(reward.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RewardClient) GetX(ctx context.Context, id uuid.UUID) *Reward {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RewardClient) Hooks() []Hook {
	hooks := c.hooks.Reward
	return append(hooks[:len(hooks):len(hooks)], reward.Hooks[:]...)
}
