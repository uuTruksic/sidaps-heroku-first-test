// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machineaccessories"
	"api/ent/machinespecification"
	"api/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineQuery is the builder for querying Machine entities.
type MachineQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	predicates               []predicate.Machine
	withMachineSpecification *MachineSpecificationQuery
	withMachineAccessories   *MachineAccessoriesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MachineQuery builder.
func (mq *MachineQuery) Where(ps ...predicate.Machine) *MachineQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MachineQuery) Limit(limit int) *MachineQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MachineQuery) Offset(offset int) *MachineQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MachineQuery) Unique(unique bool) *MachineQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MachineQuery) Order(o ...OrderFunc) *MachineQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryMachineSpecification chains the current query on the "machineSpecification" edge.
func (mq *MachineQuery) QueryMachineSpecification() *MachineSpecificationQuery {
	query := &MachineSpecificationQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(machine.Table, machine.FieldID, selector),
			sqlgraph.To(machinespecification.Table, machinespecification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, machine.MachineSpecificationTable, machine.MachineSpecificationColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMachineAccessories chains the current query on the "machineAccessories" edge.
func (mq *MachineQuery) QueryMachineAccessories() *MachineAccessoriesQuery {
	query := &MachineAccessoriesQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(machine.Table, machine.FieldID, selector),
			sqlgraph.To(machineaccessories.Table, machineaccessories.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, machine.MachineAccessoriesTable, machine.MachineAccessoriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Machine entity from the query.
// Returns a *NotFoundError when no Machine was found.
func (mq *MachineQuery) First(ctx context.Context) (*Machine, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{machine.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MachineQuery) FirstX(ctx context.Context) *Machine {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Machine ID from the query.
// Returns a *NotFoundError when no Machine ID was found.
func (mq *MachineQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{machine.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MachineQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Machine entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Machine entity is found.
// Returns a *NotFoundError when no Machine entities are found.
func (mq *MachineQuery) Only(ctx context.Context) (*Machine, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{machine.Label}
	default:
		return nil, &NotSingularError{machine.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MachineQuery) OnlyX(ctx context.Context) *Machine {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Machine ID in the query.
// Returns a *NotSingularError when more than one Machine ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MachineQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{machine.Label}
	default:
		err = &NotSingularError{machine.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MachineQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Machines.
func (mq *MachineQuery) All(ctx context.Context) ([]*Machine, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MachineQuery) AllX(ctx context.Context) []*Machine {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Machine IDs.
func (mq *MachineQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(machine.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MachineQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MachineQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MachineQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MachineQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MachineQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MachineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MachineQuery) Clone() *MachineQuery {
	if mq == nil {
		return nil
	}
	return &MachineQuery{
		config:                   mq.config,
		limit:                    mq.limit,
		offset:                   mq.offset,
		order:                    append([]OrderFunc{}, mq.order...),
		predicates:               append([]predicate.Machine{}, mq.predicates...),
		withMachineSpecification: mq.withMachineSpecification.Clone(),
		withMachineAccessories:   mq.withMachineAccessories.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithMachineSpecification tells the query-builder to eager-load the nodes that are connected to
// the "machineSpecification" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MachineQuery) WithMachineSpecification(opts ...func(*MachineSpecificationQuery)) *MachineQuery {
	query := &MachineSpecificationQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMachineSpecification = query
	return mq
}

// WithMachineAccessories tells the query-builder to eager-load the nodes that are connected to
// the "machineAccessories" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MachineQuery) WithMachineAccessories(opts ...func(*MachineAccessoriesQuery)) *MachineQuery {
	query := &MachineAccessoriesQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMachineAccessories = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Machine.Query().
//		GroupBy(machine.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MachineQuery) GroupBy(field string, fields ...string) *MachineGroupBy {
	grbuild := &MachineGroupBy{config: mq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	grbuild.label = machine.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Machine.Query().
//		Select(machine.FieldName).
//		Scan(ctx, &v)
func (mq *MachineQuery) Select(fields ...string) *MachineSelect {
	mq.fields = append(mq.fields, fields...)
	selbuild := &MachineSelect{MachineQuery: mq}
	selbuild.label = machine.Label
	selbuild.flds, selbuild.scan = &mq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MachineSelect configured with the given aggregations.
func (mq *MachineQuery) Aggregate(fns ...AggregateFunc) *MachineSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MachineQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !machine.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MachineQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Machine, error) {
	var (
		nodes       = []*Machine{}
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withMachineSpecification != nil,
			mq.withMachineAccessories != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Machine).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Machine{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withMachineSpecification; query != nil {
		if err := mq.loadMachineSpecification(ctx, query, nodes,
			func(n *Machine) { n.Edges.MachineSpecification = []*MachineSpecification{} },
			func(n *Machine, e *MachineSpecification) {
				n.Edges.MachineSpecification = append(n.Edges.MachineSpecification, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := mq.withMachineAccessories; query != nil {
		if err := mq.loadMachineAccessories(ctx, query, nodes,
			func(n *Machine) { n.Edges.MachineAccessories = []*MachineAccessories{} },
			func(n *Machine, e *MachineAccessories) {
				n.Edges.MachineAccessories = append(n.Edges.MachineAccessories, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MachineQuery) loadMachineSpecification(ctx context.Context, query *MachineSpecificationQuery, nodes []*Machine, init func(*Machine), assign func(*Machine, *MachineSpecification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Machine)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.MachineSpecification(func(s *sql.Selector) {
		s.Where(sql.InValues(machine.MachineSpecificationColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.machine_machine_specification
		if fk == nil {
			return fmt.Errorf(`foreign-key "machine_machine_specification" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "machine_machine_specification" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MachineQuery) loadMachineAccessories(ctx context.Context, query *MachineAccessoriesQuery, nodes []*Machine, init func(*Machine), assign func(*Machine, *MachineAccessories)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Machine)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.MachineAccessories(func(s *sql.Selector) {
		s.Where(sql.InValues(machine.MachineAccessoriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.machine_machine_accessories
		if fk == nil {
			return fmt.Errorf(`foreign-key "machine_machine_accessories" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "machine_machine_accessories" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MachineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MachineQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mq *MachineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machine.Table,
			Columns: machine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, machine.FieldID)
		for i := range fields {
			if fields[i] != machine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MachineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(machine.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = machine.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MachineGroupBy is the group-by builder for Machine entities.
type MachineGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MachineGroupBy) Aggregate(fns ...AggregateFunc) *MachineGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MachineGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

func (mgb *MachineGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mgb.fields {
		if !machine.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MachineGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MachineSelect is the builder for selecting fields of Machine entities.
type MachineSelect struct {
	*MachineQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MachineSelect) Aggregate(fns ...AggregateFunc) *MachineSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MachineSelect) Scan(ctx context.Context, v any) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MachineQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

func (ms *MachineSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(ms.sql))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ms.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ms.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}