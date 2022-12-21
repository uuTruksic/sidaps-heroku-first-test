// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/machine"
	"api/ent/machineaccessories"
	"api/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineAccessoriesQuery is the builder for querying MachineAccessories entities.
type MachineAccessoriesQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.MachineAccessories
	withMachine *MachineQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MachineAccessoriesQuery builder.
func (maq *MachineAccessoriesQuery) Where(ps ...predicate.MachineAccessories) *MachineAccessoriesQuery {
	maq.predicates = append(maq.predicates, ps...)
	return maq
}

// Limit adds a limit step to the query.
func (maq *MachineAccessoriesQuery) Limit(limit int) *MachineAccessoriesQuery {
	maq.limit = &limit
	return maq
}

// Offset adds an offset step to the query.
func (maq *MachineAccessoriesQuery) Offset(offset int) *MachineAccessoriesQuery {
	maq.offset = &offset
	return maq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (maq *MachineAccessoriesQuery) Unique(unique bool) *MachineAccessoriesQuery {
	maq.unique = &unique
	return maq
}

// Order adds an order step to the query.
func (maq *MachineAccessoriesQuery) Order(o ...OrderFunc) *MachineAccessoriesQuery {
	maq.order = append(maq.order, o...)
	return maq
}

// QueryMachine chains the current query on the "machine" edge.
func (maq *MachineAccessoriesQuery) QueryMachine() *MachineQuery {
	query := &MachineQuery{config: maq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := maq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := maq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(machineaccessories.Table, machineaccessories.FieldID, selector),
			sqlgraph.To(machine.Table, machine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, machineaccessories.MachineTable, machineaccessories.MachineColumn),
		)
		fromU = sqlgraph.SetNeighbors(maq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MachineAccessories entity from the query.
// Returns a *NotFoundError when no MachineAccessories was found.
func (maq *MachineAccessoriesQuery) First(ctx context.Context) (*MachineAccessories, error) {
	nodes, err := maq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{machineaccessories.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) FirstX(ctx context.Context) *MachineAccessories {
	node, err := maq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MachineAccessories ID from the query.
// Returns a *NotFoundError when no MachineAccessories ID was found.
func (maq *MachineAccessoriesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = maq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{machineaccessories.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) FirstIDX(ctx context.Context) int {
	id, err := maq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MachineAccessories entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MachineAccessories entity is found.
// Returns a *NotFoundError when no MachineAccessories entities are found.
func (maq *MachineAccessoriesQuery) Only(ctx context.Context) (*MachineAccessories, error) {
	nodes, err := maq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{machineaccessories.Label}
	default:
		return nil, &NotSingularError{machineaccessories.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) OnlyX(ctx context.Context) *MachineAccessories {
	node, err := maq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MachineAccessories ID in the query.
// Returns a *NotSingularError when more than one MachineAccessories ID is found.
// Returns a *NotFoundError when no entities are found.
func (maq *MachineAccessoriesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = maq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{machineaccessories.Label}
	default:
		err = &NotSingularError{machineaccessories.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) OnlyIDX(ctx context.Context) int {
	id, err := maq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MachineAccessoriesSlice.
func (maq *MachineAccessoriesQuery) All(ctx context.Context) ([]*MachineAccessories, error) {
	if err := maq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return maq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) AllX(ctx context.Context) []*MachineAccessories {
	nodes, err := maq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MachineAccessories IDs.
func (maq *MachineAccessoriesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := maq.Select(machineaccessories.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) IDsX(ctx context.Context) []int {
	ids, err := maq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (maq *MachineAccessoriesQuery) Count(ctx context.Context) (int, error) {
	if err := maq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return maq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) CountX(ctx context.Context) int {
	count, err := maq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (maq *MachineAccessoriesQuery) Exist(ctx context.Context) (bool, error) {
	if err := maq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return maq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (maq *MachineAccessoriesQuery) ExistX(ctx context.Context) bool {
	exist, err := maq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MachineAccessoriesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (maq *MachineAccessoriesQuery) Clone() *MachineAccessoriesQuery {
	if maq == nil {
		return nil
	}
	return &MachineAccessoriesQuery{
		config:      maq.config,
		limit:       maq.limit,
		offset:      maq.offset,
		order:       append([]OrderFunc{}, maq.order...),
		predicates:  append([]predicate.MachineAccessories{}, maq.predicates...),
		withMachine: maq.withMachine.Clone(),
		// clone intermediate query.
		sql:    maq.sql.Clone(),
		path:   maq.path,
		unique: maq.unique,
	}
}

// WithMachine tells the query-builder to eager-load the nodes that are connected to
// the "machine" edge. The optional arguments are used to configure the query builder of the edge.
func (maq *MachineAccessoriesQuery) WithMachine(opts ...func(*MachineQuery)) *MachineAccessoriesQuery {
	query := &MachineQuery{config: maq.config}
	for _, opt := range opts {
		opt(query)
	}
	maq.withMachine = query
	return maq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name1 string `json:"name1,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MachineAccessories.Query().
//		GroupBy(machineaccessories.FieldName1).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (maq *MachineAccessoriesQuery) GroupBy(field string, fields ...string) *MachineAccessoriesGroupBy {
	grbuild := &MachineAccessoriesGroupBy{config: maq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := maq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return maq.sqlQuery(ctx), nil
	}
	grbuild.label = machineaccessories.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name1 string `json:"name1,omitempty"`
//	}
//
//	client.MachineAccessories.Query().
//		Select(machineaccessories.FieldName1).
//		Scan(ctx, &v)
func (maq *MachineAccessoriesQuery) Select(fields ...string) *MachineAccessoriesSelect {
	maq.fields = append(maq.fields, fields...)
	selbuild := &MachineAccessoriesSelect{MachineAccessoriesQuery: maq}
	selbuild.label = machineaccessories.Label
	selbuild.flds, selbuild.scan = &maq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MachineAccessoriesSelect configured with the given aggregations.
func (maq *MachineAccessoriesQuery) Aggregate(fns ...AggregateFunc) *MachineAccessoriesSelect {
	return maq.Select().Aggregate(fns...)
}

func (maq *MachineAccessoriesQuery) prepareQuery(ctx context.Context) error {
	for _, f := range maq.fields {
		if !machineaccessories.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if maq.path != nil {
		prev, err := maq.path(ctx)
		if err != nil {
			return err
		}
		maq.sql = prev
	}
	return nil
}

func (maq *MachineAccessoriesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MachineAccessories, error) {
	var (
		nodes       = []*MachineAccessories{}
		withFKs     = maq.withFKs
		_spec       = maq.querySpec()
		loadedTypes = [1]bool{
			maq.withMachine != nil,
		}
	)
	if maq.withMachine != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, machineaccessories.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MachineAccessories).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MachineAccessories{config: maq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, maq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := maq.withMachine; query != nil {
		if err := maq.loadMachine(ctx, query, nodes, nil,
			func(n *MachineAccessories, e *Machine) { n.Edges.Machine = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (maq *MachineAccessoriesQuery) loadMachine(ctx context.Context, query *MachineQuery, nodes []*MachineAccessories, init func(*MachineAccessories), assign func(*MachineAccessories, *Machine)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MachineAccessories)
	for i := range nodes {
		if nodes[i].machine_machine_accessories == nil {
			continue
		}
		fk := *nodes[i].machine_machine_accessories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(machine.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "machine_machine_accessories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (maq *MachineAccessoriesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := maq.querySpec()
	_spec.Node.Columns = maq.fields
	if len(maq.fields) > 0 {
		_spec.Unique = maq.unique != nil && *maq.unique
	}
	return sqlgraph.CountNodes(ctx, maq.driver, _spec)
}

func (maq *MachineAccessoriesQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := maq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (maq *MachineAccessoriesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   machineaccessories.Table,
			Columns: machineaccessories.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machineaccessories.FieldID,
			},
		},
		From:   maq.sql,
		Unique: true,
	}
	if unique := maq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := maq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, machineaccessories.FieldID)
		for i := range fields {
			if fields[i] != machineaccessories.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := maq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := maq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := maq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := maq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (maq *MachineAccessoriesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(maq.driver.Dialect())
	t1 := builder.Table(machineaccessories.Table)
	columns := maq.fields
	if len(columns) == 0 {
		columns = machineaccessories.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if maq.sql != nil {
		selector = maq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if maq.unique != nil && *maq.unique {
		selector.Distinct()
	}
	for _, p := range maq.predicates {
		p(selector)
	}
	for _, p := range maq.order {
		p(selector)
	}
	if offset := maq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := maq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MachineAccessoriesGroupBy is the group-by builder for MachineAccessories entities.
type MachineAccessoriesGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (magb *MachineAccessoriesGroupBy) Aggregate(fns ...AggregateFunc) *MachineAccessoriesGroupBy {
	magb.fns = append(magb.fns, fns...)
	return magb
}

// Scan applies the group-by query and scans the result into the given value.
func (magb *MachineAccessoriesGroupBy) Scan(ctx context.Context, v any) error {
	query, err := magb.path(ctx)
	if err != nil {
		return err
	}
	magb.sql = query
	return magb.sqlScan(ctx, v)
}

func (magb *MachineAccessoriesGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range magb.fields {
		if !machineaccessories.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := magb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := magb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (magb *MachineAccessoriesGroupBy) sqlQuery() *sql.Selector {
	selector := magb.sql.Select()
	aggregation := make([]string, 0, len(magb.fns))
	for _, fn := range magb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(magb.fields)+len(magb.fns))
		for _, f := range magb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(magb.fields...)...)
}

// MachineAccessoriesSelect is the builder for selecting fields of MachineAccessories entities.
type MachineAccessoriesSelect struct {
	*MachineAccessoriesQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mas *MachineAccessoriesSelect) Aggregate(fns ...AggregateFunc) *MachineAccessoriesSelect {
	mas.fns = append(mas.fns, fns...)
	return mas
}

// Scan applies the selector query and scans the result into the given value.
func (mas *MachineAccessoriesSelect) Scan(ctx context.Context, v any) error {
	if err := mas.prepareQuery(ctx); err != nil {
		return err
	}
	mas.sql = mas.MachineAccessoriesQuery.sqlQuery(ctx)
	return mas.sqlScan(ctx, v)
}

func (mas *MachineAccessoriesSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(mas.fns))
	for _, fn := range mas.fns {
		aggregation = append(aggregation, fn(mas.sql))
	}
	switch n := len(*mas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		mas.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		mas.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := mas.sql.Query()
	if err := mas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}