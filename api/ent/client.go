// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"api/ent/migrate"

	"api/ent/machine"
	"api/ent/machineaccessories"
	"api/ent/machinespecification"
	"api/ent/session"
	"api/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Machine is the client for interacting with the Machine builders.
	Machine *MachineClient
	// MachineAccessories is the client for interacting with the MachineAccessories builders.
	MachineAccessories *MachineAccessoriesClient
	// MachineSpecification is the client for interacting with the MachineSpecification builders.
	MachineSpecification *MachineSpecificationClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Machine = NewMachineClient(c.config)
	c.MachineAccessories = NewMachineAccessoriesClient(c.config)
	c.MachineSpecification = NewMachineSpecificationClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:                  ctx,
		config:               cfg,
		Machine:              NewMachineClient(cfg),
		MachineAccessories:   NewMachineAccessoriesClient(cfg),
		MachineSpecification: NewMachineSpecificationClient(cfg),
		Session:              NewSessionClient(cfg),
		User:                 NewUserClient(cfg),
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
		ctx:                  ctx,
		config:               cfg,
		Machine:              NewMachineClient(cfg),
		MachineAccessories:   NewMachineAccessoriesClient(cfg),
		MachineSpecification: NewMachineSpecificationClient(cfg),
		Session:              NewSessionClient(cfg),
		User:                 NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Machine.
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
	c.Machine.Use(hooks...)
	c.MachineAccessories.Use(hooks...)
	c.MachineSpecification.Use(hooks...)
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
}

// MachineClient is a client for the Machine schema.
type MachineClient struct {
	config
}

// NewMachineClient returns a client for the Machine from the given config.
func NewMachineClient(c config) *MachineClient {
	return &MachineClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `machine.Hooks(f(g(h())))`.
func (c *MachineClient) Use(hooks ...Hook) {
	c.hooks.Machine = append(c.hooks.Machine, hooks...)
}

// Create returns a builder for creating a Machine entity.
func (c *MachineClient) Create() *MachineCreate {
	mutation := newMachineMutation(c.config, OpCreate)
	return &MachineCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Machine entities.
func (c *MachineClient) CreateBulk(builders ...*MachineCreate) *MachineCreateBulk {
	return &MachineCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Machine.
func (c *MachineClient) Update() *MachineUpdate {
	mutation := newMachineMutation(c.config, OpUpdate)
	return &MachineUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MachineClient) UpdateOne(m *Machine) *MachineUpdateOne {
	mutation := newMachineMutation(c.config, OpUpdateOne, withMachine(m))
	return &MachineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MachineClient) UpdateOneID(id int) *MachineUpdateOne {
	mutation := newMachineMutation(c.config, OpUpdateOne, withMachineID(id))
	return &MachineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Machine.
func (c *MachineClient) Delete() *MachineDelete {
	mutation := newMachineMutation(c.config, OpDelete)
	return &MachineDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MachineClient) DeleteOne(m *Machine) *MachineDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MachineClient) DeleteOneID(id int) *MachineDeleteOne {
	builder := c.Delete().Where(machine.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MachineDeleteOne{builder}
}

// Query returns a query builder for Machine.
func (c *MachineClient) Query() *MachineQuery {
	return &MachineQuery{
		config: c.config,
	}
}

// Get returns a Machine entity by its id.
func (c *MachineClient) Get(ctx context.Context, id int) (*Machine, error) {
	return c.Query().Where(machine.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MachineClient) GetX(ctx context.Context, id int) *Machine {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMachineSpecification queries the machineSpecification edge of a Machine.
func (c *MachineClient) QueryMachineSpecification(m *Machine) *MachineSpecificationQuery {
	query := &MachineSpecificationQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(machine.Table, machine.FieldID, id),
			sqlgraph.To(machinespecification.Table, machinespecification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, machine.MachineSpecificationTable, machine.MachineSpecificationColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMachineAccessories queries the machineAccessories edge of a Machine.
func (c *MachineClient) QueryMachineAccessories(m *Machine) *MachineAccessoriesQuery {
	query := &MachineAccessoriesQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(machine.Table, machine.FieldID, id),
			sqlgraph.To(machineaccessories.Table, machineaccessories.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, machine.MachineAccessoriesTable, machine.MachineAccessoriesColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MachineClient) Hooks() []Hook {
	return c.hooks.Machine
}

// MachineAccessoriesClient is a client for the MachineAccessories schema.
type MachineAccessoriesClient struct {
	config
}

// NewMachineAccessoriesClient returns a client for the MachineAccessories from the given config.
func NewMachineAccessoriesClient(c config) *MachineAccessoriesClient {
	return &MachineAccessoriesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `machineaccessories.Hooks(f(g(h())))`.
func (c *MachineAccessoriesClient) Use(hooks ...Hook) {
	c.hooks.MachineAccessories = append(c.hooks.MachineAccessories, hooks...)
}

// Create returns a builder for creating a MachineAccessories entity.
func (c *MachineAccessoriesClient) Create() *MachineAccessoriesCreate {
	mutation := newMachineAccessoriesMutation(c.config, OpCreate)
	return &MachineAccessoriesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MachineAccessories entities.
func (c *MachineAccessoriesClient) CreateBulk(builders ...*MachineAccessoriesCreate) *MachineAccessoriesCreateBulk {
	return &MachineAccessoriesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MachineAccessories.
func (c *MachineAccessoriesClient) Update() *MachineAccessoriesUpdate {
	mutation := newMachineAccessoriesMutation(c.config, OpUpdate)
	return &MachineAccessoriesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MachineAccessoriesClient) UpdateOne(ma *MachineAccessories) *MachineAccessoriesUpdateOne {
	mutation := newMachineAccessoriesMutation(c.config, OpUpdateOne, withMachineAccessories(ma))
	return &MachineAccessoriesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MachineAccessoriesClient) UpdateOneID(id int) *MachineAccessoriesUpdateOne {
	mutation := newMachineAccessoriesMutation(c.config, OpUpdateOne, withMachineAccessoriesID(id))
	return &MachineAccessoriesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MachineAccessories.
func (c *MachineAccessoriesClient) Delete() *MachineAccessoriesDelete {
	mutation := newMachineAccessoriesMutation(c.config, OpDelete)
	return &MachineAccessoriesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MachineAccessoriesClient) DeleteOne(ma *MachineAccessories) *MachineAccessoriesDeleteOne {
	return c.DeleteOneID(ma.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MachineAccessoriesClient) DeleteOneID(id int) *MachineAccessoriesDeleteOne {
	builder := c.Delete().Where(machineaccessories.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MachineAccessoriesDeleteOne{builder}
}

// Query returns a query builder for MachineAccessories.
func (c *MachineAccessoriesClient) Query() *MachineAccessoriesQuery {
	return &MachineAccessoriesQuery{
		config: c.config,
	}
}

// Get returns a MachineAccessories entity by its id.
func (c *MachineAccessoriesClient) Get(ctx context.Context, id int) (*MachineAccessories, error) {
	return c.Query().Where(machineaccessories.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MachineAccessoriesClient) GetX(ctx context.Context, id int) *MachineAccessories {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMachine queries the machine edge of a MachineAccessories.
func (c *MachineAccessoriesClient) QueryMachine(ma *MachineAccessories) *MachineQuery {
	query := &MachineQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ma.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(machineaccessories.Table, machineaccessories.FieldID, id),
			sqlgraph.To(machine.Table, machine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, machineaccessories.MachineTable, machineaccessories.MachineColumn),
		)
		fromV = sqlgraph.Neighbors(ma.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MachineAccessoriesClient) Hooks() []Hook {
	return c.hooks.MachineAccessories
}

// MachineSpecificationClient is a client for the MachineSpecification schema.
type MachineSpecificationClient struct {
	config
}

// NewMachineSpecificationClient returns a client for the MachineSpecification from the given config.
func NewMachineSpecificationClient(c config) *MachineSpecificationClient {
	return &MachineSpecificationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `machinespecification.Hooks(f(g(h())))`.
func (c *MachineSpecificationClient) Use(hooks ...Hook) {
	c.hooks.MachineSpecification = append(c.hooks.MachineSpecification, hooks...)
}

// Create returns a builder for creating a MachineSpecification entity.
func (c *MachineSpecificationClient) Create() *MachineSpecificationCreate {
	mutation := newMachineSpecificationMutation(c.config, OpCreate)
	return &MachineSpecificationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MachineSpecification entities.
func (c *MachineSpecificationClient) CreateBulk(builders ...*MachineSpecificationCreate) *MachineSpecificationCreateBulk {
	return &MachineSpecificationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MachineSpecification.
func (c *MachineSpecificationClient) Update() *MachineSpecificationUpdate {
	mutation := newMachineSpecificationMutation(c.config, OpUpdate)
	return &MachineSpecificationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MachineSpecificationClient) UpdateOne(ms *MachineSpecification) *MachineSpecificationUpdateOne {
	mutation := newMachineSpecificationMutation(c.config, OpUpdateOne, withMachineSpecification(ms))
	return &MachineSpecificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MachineSpecificationClient) UpdateOneID(id int) *MachineSpecificationUpdateOne {
	mutation := newMachineSpecificationMutation(c.config, OpUpdateOne, withMachineSpecificationID(id))
	return &MachineSpecificationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MachineSpecification.
func (c *MachineSpecificationClient) Delete() *MachineSpecificationDelete {
	mutation := newMachineSpecificationMutation(c.config, OpDelete)
	return &MachineSpecificationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MachineSpecificationClient) DeleteOne(ms *MachineSpecification) *MachineSpecificationDeleteOne {
	return c.DeleteOneID(ms.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MachineSpecificationClient) DeleteOneID(id int) *MachineSpecificationDeleteOne {
	builder := c.Delete().Where(machinespecification.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MachineSpecificationDeleteOne{builder}
}

// Query returns a query builder for MachineSpecification.
func (c *MachineSpecificationClient) Query() *MachineSpecificationQuery {
	return &MachineSpecificationQuery{
		config: c.config,
	}
}

// Get returns a MachineSpecification entity by its id.
func (c *MachineSpecificationClient) Get(ctx context.Context, id int) (*MachineSpecification, error) {
	return c.Query().Where(machinespecification.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MachineSpecificationClient) GetX(ctx context.Context, id int) *MachineSpecification {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMachine queries the machine edge of a MachineSpecification.
func (c *MachineSpecificationClient) QueryMachine(ms *MachineSpecification) *MachineQuery {
	query := &MachineQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ms.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(machinespecification.Table, machinespecification.FieldID, id),
			sqlgraph.To(machine.Table, machine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, machinespecification.MachineTable, machinespecification.MachineColumn),
		)
		fromV = sqlgraph.Neighbors(ms.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MachineSpecificationClient) Hooks() []Hook {
	return c.hooks.MachineSpecification
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a builder for creating a Session entity.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id int) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SessionClient) DeleteOneID(id int) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id int) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id int) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Session.
func (c *SessionClient) QueryUser(s *Session) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, session.UserTable, session.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySessions queries the sessions edge of a User.
func (c *UserClient) QuerySessions(u *User) *SessionQuery {
	query := &SessionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SessionsTable, user.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
