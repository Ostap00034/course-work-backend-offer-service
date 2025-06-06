// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Ostap00034/course-work-backend-offer-service/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Ostap00034/course-work-backend-offer-service/ent/offer"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Offer is the client for interacting with the Offer builders.
	Offer *OfferClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Offer = NewOfferClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
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
		Offer:  NewOfferClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		Offer:  NewOfferClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Offer.
//		Query().
//		Count(ctx)
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
	c.Offer.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Offer.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *OfferMutation:
		return c.Offer.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// OfferClient is a client for the Offer schema.
type OfferClient struct {
	config
}

// NewOfferClient returns a client for the Offer from the given config.
func NewOfferClient(c config) *OfferClient {
	return &OfferClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `offer.Hooks(f(g(h())))`.
func (c *OfferClient) Use(hooks ...Hook) {
	c.hooks.Offer = append(c.hooks.Offer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `offer.Intercept(f(g(h())))`.
func (c *OfferClient) Intercept(interceptors ...Interceptor) {
	c.inters.Offer = append(c.inters.Offer, interceptors...)
}

// Create returns a builder for creating a Offer entity.
func (c *OfferClient) Create() *OfferCreate {
	mutation := newOfferMutation(c.config, OpCreate)
	return &OfferCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Offer entities.
func (c *OfferClient) CreateBulk(builders ...*OfferCreate) *OfferCreateBulk {
	return &OfferCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OfferClient) MapCreateBulk(slice any, setFunc func(*OfferCreate, int)) *OfferCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OfferCreateBulk{err: fmt.Errorf("calling to OfferClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OfferCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OfferCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Offer.
func (c *OfferClient) Update() *OfferUpdate {
	mutation := newOfferMutation(c.config, OpUpdate)
	return &OfferUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OfferClient) UpdateOne(o *Offer) *OfferUpdateOne {
	mutation := newOfferMutation(c.config, OpUpdateOne, withOffer(o))
	return &OfferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OfferClient) UpdateOneID(id uuid.UUID) *OfferUpdateOne {
	mutation := newOfferMutation(c.config, OpUpdateOne, withOfferID(id))
	return &OfferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Offer.
func (c *OfferClient) Delete() *OfferDelete {
	mutation := newOfferMutation(c.config, OpDelete)
	return &OfferDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OfferClient) DeleteOne(o *Offer) *OfferDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OfferClient) DeleteOneID(id uuid.UUID) *OfferDeleteOne {
	builder := c.Delete().Where(offer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OfferDeleteOne{builder}
}

// Query returns a query builder for Offer.
func (c *OfferClient) Query() *OfferQuery {
	return &OfferQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOffer},
		inters: c.Interceptors(),
	}
}

// Get returns a Offer entity by its id.
func (c *OfferClient) Get(ctx context.Context, id uuid.UUID) (*Offer, error) {
	return c.Query().Where(offer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OfferClient) GetX(ctx context.Context, id uuid.UUID) *Offer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OfferClient) Hooks() []Hook {
	return c.hooks.Offer
}

// Interceptors returns the client interceptors.
func (c *OfferClient) Interceptors() []Interceptor {
	return c.inters.Offer
}

func (c *OfferClient) mutate(ctx context.Context, m *OfferMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OfferCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OfferUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OfferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OfferDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Offer mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Offer []ent.Hook
	}
	inters struct {
		Offer []ent.Interceptor
	}
)
