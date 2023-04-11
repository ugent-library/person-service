// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ugent-library/people/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ugent-library/people/ent/organization"
	"github.com/ugent-library/people/ent/organizationperson"
	"github.com/ugent-library/people/ent/person"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Organization is the client for interacting with the Organization builders.
	Organization *OrganizationClient
	// OrganizationPerson is the client for interacting with the OrganizationPerson builders.
	OrganizationPerson *OrganizationPersonClient
	// Person is the client for interacting with the Person builders.
	Person *PersonClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Organization = NewOrganizationClient(c.config)
	c.OrganizationPerson = NewOrganizationPersonClient(c.config)
	c.Person = NewPersonClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Organization:       NewOrganizationClient(cfg),
		OrganizationPerson: NewOrganizationPersonClient(cfg),
		Person:             NewPersonClient(cfg),
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
		ctx:                ctx,
		config:             cfg,
		Organization:       NewOrganizationClient(cfg),
		OrganizationPerson: NewOrganizationPersonClient(cfg),
		Person:             NewPersonClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Organization.
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
	c.Organization.Use(hooks...)
	c.OrganizationPerson.Use(hooks...)
	c.Person.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Organization.Intercept(interceptors...)
	c.OrganizationPerson.Intercept(interceptors...)
	c.Person.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *OrganizationMutation:
		return c.Organization.mutate(ctx, m)
	case *OrganizationPersonMutation:
		return c.OrganizationPerson.mutate(ctx, m)
	case *PersonMutation:
		return c.Person.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// OrganizationClient is a client for the Organization schema.
type OrganizationClient struct {
	config
}

// NewOrganizationClient returns a client for the Organization from the given config.
func NewOrganizationClient(c config) *OrganizationClient {
	return &OrganizationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `organization.Hooks(f(g(h())))`.
func (c *OrganizationClient) Use(hooks ...Hook) {
	c.hooks.Organization = append(c.hooks.Organization, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `organization.Intercept(f(g(h())))`.
func (c *OrganizationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Organization = append(c.inters.Organization, interceptors...)
}

// Create returns a builder for creating a Organization entity.
func (c *OrganizationClient) Create() *OrganizationCreate {
	mutation := newOrganizationMutation(c.config, OpCreate)
	return &OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Organization entities.
func (c *OrganizationClient) CreateBulk(builders ...*OrganizationCreate) *OrganizationCreateBulk {
	return &OrganizationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Organization.
func (c *OrganizationClient) Update() *OrganizationUpdate {
	mutation := newOrganizationMutation(c.config, OpUpdate)
	return &OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrganizationClient) UpdateOne(o *Organization) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganization(o))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrganizationClient) UpdateOneID(id int) *OrganizationUpdateOne {
	mutation := newOrganizationMutation(c.config, OpUpdateOne, withOrganizationID(id))
	return &OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Organization.
func (c *OrganizationClient) Delete() *OrganizationDelete {
	mutation := newOrganizationMutation(c.config, OpDelete)
	return &OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrganizationClient) DeleteOne(o *Organization) *OrganizationDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrganizationClient) DeleteOneID(id int) *OrganizationDeleteOne {
	builder := c.Delete().Where(organization.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrganizationDeleteOne{builder}
}

// Query returns a query builder for Organization.
func (c *OrganizationClient) Query() *OrganizationQuery {
	return &OrganizationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrganization},
		inters: c.Interceptors(),
	}
}

// Get returns a Organization entity by its id.
func (c *OrganizationClient) Get(ctx context.Context, id int) (*Organization, error) {
	return c.Query().Where(organization.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrganizationClient) GetX(ctx context.Context, id int) *Organization {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPeople queries the people edge of a Organization.
func (c *OrganizationClient) QueryPeople(o *Organization) *PersonQuery {
	query := (&PersonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, organization.PeopleTable, organization.PeoplePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrganizationPerson queries the organization_person edge of a Organization.
func (c *OrganizationClient) QueryOrganizationPerson(o *Organization) *OrganizationPersonQuery {
	query := (&OrganizationPersonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, id),
			sqlgraph.To(organizationperson.Table, organizationperson.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, organization.OrganizationPersonTable, organization.OrganizationPersonColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrganizationClient) Hooks() []Hook {
	return c.hooks.Organization
}

// Interceptors returns the client interceptors.
func (c *OrganizationClient) Interceptors() []Interceptor {
	return c.inters.Organization
}

func (c *OrganizationClient) mutate(ctx context.Context, m *OrganizationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Organization mutation op: %q", m.Op())
	}
}

// OrganizationPersonClient is a client for the OrganizationPerson schema.
type OrganizationPersonClient struct {
	config
}

// NewOrganizationPersonClient returns a client for the OrganizationPerson from the given config.
func NewOrganizationPersonClient(c config) *OrganizationPersonClient {
	return &OrganizationPersonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `organizationperson.Hooks(f(g(h())))`.
func (c *OrganizationPersonClient) Use(hooks ...Hook) {
	c.hooks.OrganizationPerson = append(c.hooks.OrganizationPerson, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `organizationperson.Intercept(f(g(h())))`.
func (c *OrganizationPersonClient) Intercept(interceptors ...Interceptor) {
	c.inters.OrganizationPerson = append(c.inters.OrganizationPerson, interceptors...)
}

// Create returns a builder for creating a OrganizationPerson entity.
func (c *OrganizationPersonClient) Create() *OrganizationPersonCreate {
	mutation := newOrganizationPersonMutation(c.config, OpCreate)
	return &OrganizationPersonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OrganizationPerson entities.
func (c *OrganizationPersonClient) CreateBulk(builders ...*OrganizationPersonCreate) *OrganizationPersonCreateBulk {
	return &OrganizationPersonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OrganizationPerson.
func (c *OrganizationPersonClient) Update() *OrganizationPersonUpdate {
	mutation := newOrganizationPersonMutation(c.config, OpUpdate)
	return &OrganizationPersonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrganizationPersonClient) UpdateOne(op *OrganizationPerson) *OrganizationPersonUpdateOne {
	mutation := newOrganizationPersonMutation(c.config, OpUpdateOne, withOrganizationPerson(op))
	return &OrganizationPersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrganizationPersonClient) UpdateOneID(id int) *OrganizationPersonUpdateOne {
	mutation := newOrganizationPersonMutation(c.config, OpUpdateOne, withOrganizationPersonID(id))
	return &OrganizationPersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OrganizationPerson.
func (c *OrganizationPersonClient) Delete() *OrganizationPersonDelete {
	mutation := newOrganizationPersonMutation(c.config, OpDelete)
	return &OrganizationPersonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrganizationPersonClient) DeleteOne(op *OrganizationPerson) *OrganizationPersonDeleteOne {
	return c.DeleteOneID(op.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrganizationPersonClient) DeleteOneID(id int) *OrganizationPersonDeleteOne {
	builder := c.Delete().Where(organizationperson.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrganizationPersonDeleteOne{builder}
}

// Query returns a query builder for OrganizationPerson.
func (c *OrganizationPersonClient) Query() *OrganizationPersonQuery {
	return &OrganizationPersonQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrganizationPerson},
		inters: c.Interceptors(),
	}
}

// Get returns a OrganizationPerson entity by its id.
func (c *OrganizationPersonClient) Get(ctx context.Context, id int) (*OrganizationPerson, error) {
	return c.Query().Where(organizationperson.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrganizationPersonClient) GetX(ctx context.Context, id int) *OrganizationPerson {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPeople queries the people edge of a OrganizationPerson.
func (c *OrganizationPersonClient) QueryPeople(op *OrganizationPerson) *PersonQuery {
	query := (&PersonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := op.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationperson.Table, organizationperson.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, organizationperson.PeopleTable, organizationperson.PeopleColumn),
		)
		fromV = sqlgraph.Neighbors(op.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrganizations queries the organizations edge of a OrganizationPerson.
func (c *OrganizationPersonClient) QueryOrganizations(op *OrganizationPerson) *OrganizationQuery {
	query := (&OrganizationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := op.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationperson.Table, organizationperson.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, organizationperson.OrganizationsTable, organizationperson.OrganizationsColumn),
		)
		fromV = sqlgraph.Neighbors(op.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrganizationPersonClient) Hooks() []Hook {
	return c.hooks.OrganizationPerson
}

// Interceptors returns the client interceptors.
func (c *OrganizationPersonClient) Interceptors() []Interceptor {
	return c.inters.OrganizationPerson
}

func (c *OrganizationPersonClient) mutate(ctx context.Context, m *OrganizationPersonMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrganizationPersonCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrganizationPersonUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrganizationPersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrganizationPersonDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OrganizationPerson mutation op: %q", m.Op())
	}
}

// PersonClient is a client for the Person schema.
type PersonClient struct {
	config
}

// NewPersonClient returns a client for the Person from the given config.
func NewPersonClient(c config) *PersonClient {
	return &PersonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `person.Hooks(f(g(h())))`.
func (c *PersonClient) Use(hooks ...Hook) {
	c.hooks.Person = append(c.hooks.Person, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `person.Intercept(f(g(h())))`.
func (c *PersonClient) Intercept(interceptors ...Interceptor) {
	c.inters.Person = append(c.inters.Person, interceptors...)
}

// Create returns a builder for creating a Person entity.
func (c *PersonClient) Create() *PersonCreate {
	mutation := newPersonMutation(c.config, OpCreate)
	return &PersonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Person entities.
func (c *PersonClient) CreateBulk(builders ...*PersonCreate) *PersonCreateBulk {
	return &PersonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Person.
func (c *PersonClient) Update() *PersonUpdate {
	mutation := newPersonMutation(c.config, OpUpdate)
	return &PersonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonClient) UpdateOne(pe *Person) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPerson(pe))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonClient) UpdateOneID(id int) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPersonID(id))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Person.
func (c *PersonClient) Delete() *PersonDelete {
	mutation := newPersonMutation(c.config, OpDelete)
	return &PersonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PersonClient) DeleteOne(pe *Person) *PersonDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PersonClient) DeleteOneID(id int) *PersonDeleteOne {
	builder := c.Delete().Where(person.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonDeleteOne{builder}
}

// Query returns a query builder for Person.
func (c *PersonClient) Query() *PersonQuery {
	return &PersonQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePerson},
		inters: c.Interceptors(),
	}
}

// Get returns a Person entity by its id.
func (c *PersonClient) Get(ctx context.Context, id int) (*Person, error) {
	return c.Query().Where(person.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonClient) GetX(ctx context.Context, id int) *Person {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganizations queries the organizations edge of a Person.
func (c *PersonClient) QueryOrganizations(pe *Person) *OrganizationQuery {
	query := (&OrganizationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.OrganizationsTable, person.OrganizationsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrganizationPerson queries the organization_person edge of a Person.
func (c *PersonClient) QueryOrganizationPerson(pe *Person) *OrganizationPersonQuery {
	query := (&OrganizationPersonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(organizationperson.Table, organizationperson.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, person.OrganizationPersonTable, person.OrganizationPersonColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PersonClient) Hooks() []Hook {
	return c.hooks.Person
}

// Interceptors returns the client interceptors.
func (c *PersonClient) Interceptors() []Interceptor {
	return c.inters.Person
}

func (c *PersonClient) mutate(ctx context.Context, m *PersonMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PersonCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PersonUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PersonDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Person mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Organization, OrganizationPerson, Person []ent.Hook
	}
	inters struct {
		Organization, OrganizationPerson, Person []ent.Interceptor
	}
)
