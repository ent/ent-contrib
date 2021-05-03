// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/dependsonskipped"
	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DependsOnSkippedQuery is the builder for querying DependsOnSkipped entities.
type DependsOnSkippedQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DependsOnSkipped
	// eager-loading edges.
	withSkipped *ImplicitSkippedMessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DependsOnSkippedQuery builder.
func (dosq *DependsOnSkippedQuery) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedQuery {
	dosq.predicates = append(dosq.predicates, ps...)
	return dosq
}

// Limit adds a limit step to the query.
func (dosq *DependsOnSkippedQuery) Limit(limit int) *DependsOnSkippedQuery {
	dosq.limit = &limit
	return dosq
}

// Offset adds an offset step to the query.
func (dosq *DependsOnSkippedQuery) Offset(offset int) *DependsOnSkippedQuery {
	dosq.offset = &offset
	return dosq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dosq *DependsOnSkippedQuery) Unique(unique bool) *DependsOnSkippedQuery {
	dosq.unique = &unique
	return dosq
}

// Order adds an order step to the query.
func (dosq *DependsOnSkippedQuery) Order(o ...OrderFunc) *DependsOnSkippedQuery {
	dosq.order = append(dosq.order, o...)
	return dosq
}

// QuerySkipped chains the current query on the "skipped" edge.
func (dosq *DependsOnSkippedQuery) QuerySkipped() *ImplicitSkippedMessageQuery {
	query := &ImplicitSkippedMessageQuery{config: dosq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dosq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dosq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dependsonskipped.Table, dependsonskipped.FieldID, selector),
			sqlgraph.To(implicitskippedmessage.Table, implicitskippedmessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dependsonskipped.SkippedTable, dependsonskipped.SkippedColumn),
		)
		fromU = sqlgraph.SetNeighbors(dosq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DependsOnSkipped entity from the query.
// Returns a *NotFoundError when no DependsOnSkipped was found.
func (dosq *DependsOnSkippedQuery) First(ctx context.Context) (*DependsOnSkipped, error) {
	nodes, err := dosq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dependsonskipped.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) FirstX(ctx context.Context) *DependsOnSkipped {
	node, err := dosq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DependsOnSkipped ID from the query.
// Returns a *NotFoundError when no DependsOnSkipped ID was found.
func (dosq *DependsOnSkippedQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dosq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dependsonskipped.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) FirstIDX(ctx context.Context) int {
	id, err := dosq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DependsOnSkipped entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DependsOnSkipped entity is not found.
// Returns a *NotFoundError when no DependsOnSkipped entities are found.
func (dosq *DependsOnSkippedQuery) Only(ctx context.Context) (*DependsOnSkipped, error) {
	nodes, err := dosq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dependsonskipped.Label}
	default:
		return nil, &NotSingularError{dependsonskipped.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) OnlyX(ctx context.Context) *DependsOnSkipped {
	node, err := dosq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DependsOnSkipped ID in the query.
// Returns a *NotSingularError when exactly one DependsOnSkipped ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dosq *DependsOnSkippedQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dosq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = &NotSingularError{dependsonskipped.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) OnlyIDX(ctx context.Context) int {
	id, err := dosq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DependsOnSkippeds.
func (dosq *DependsOnSkippedQuery) All(ctx context.Context) ([]*DependsOnSkipped, error) {
	if err := dosq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dosq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) AllX(ctx context.Context) []*DependsOnSkipped {
	nodes, err := dosq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DependsOnSkipped IDs.
func (dosq *DependsOnSkippedQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dosq.Select(dependsonskipped.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) IDsX(ctx context.Context) []int {
	ids, err := dosq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dosq *DependsOnSkippedQuery) Count(ctx context.Context) (int, error) {
	if err := dosq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dosq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) CountX(ctx context.Context) int {
	count, err := dosq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dosq *DependsOnSkippedQuery) Exist(ctx context.Context) (bool, error) {
	if err := dosq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dosq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) ExistX(ctx context.Context) bool {
	exist, err := dosq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DependsOnSkippedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dosq *DependsOnSkippedQuery) Clone() *DependsOnSkippedQuery {
	if dosq == nil {
		return nil
	}
	return &DependsOnSkippedQuery{
		config:      dosq.config,
		limit:       dosq.limit,
		offset:      dosq.offset,
		order:       append([]OrderFunc{}, dosq.order...),
		predicates:  append([]predicate.DependsOnSkipped{}, dosq.predicates...),
		withSkipped: dosq.withSkipped.Clone(),
		// clone intermediate query.
		sql:  dosq.sql.Clone(),
		path: dosq.path,
	}
}

// WithSkipped tells the query-builder to eager-load the nodes that are connected to
// the "skipped" edge. The optional arguments are used to configure the query builder of the edge.
func (dosq *DependsOnSkippedQuery) WithSkipped(opts ...func(*ImplicitSkippedMessageQuery)) *DependsOnSkippedQuery {
	query := &ImplicitSkippedMessageQuery{config: dosq.config}
	for _, opt := range opts {
		opt(query)
	}
	dosq.withSkipped = query
	return dosq
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
//	client.DependsOnSkipped.Query().
//		GroupBy(dependsonskipped.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dosq *DependsOnSkippedQuery) GroupBy(field string, fields ...string) *DependsOnSkippedGroupBy {
	group := &DependsOnSkippedGroupBy{config: dosq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dosq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dosq.sqlQuery(ctx), nil
	}
	return group
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
//	client.DependsOnSkipped.Query().
//		Select(dependsonskipped.FieldName).
//		Scan(ctx, &v)
//
func (dosq *DependsOnSkippedQuery) Select(field string, fields ...string) *DependsOnSkippedSelect {
	dosq.fields = append([]string{field}, fields...)
	return &DependsOnSkippedSelect{DependsOnSkippedQuery: dosq}
}

func (dosq *DependsOnSkippedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dosq.fields {
		if !dependsonskipped.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dosq.path != nil {
		prev, err := dosq.path(ctx)
		if err != nil {
			return err
		}
		dosq.sql = prev
	}
	return nil
}

func (dosq *DependsOnSkippedQuery) sqlAll(ctx context.Context) ([]*DependsOnSkipped, error) {
	var (
		nodes       = []*DependsOnSkipped{}
		_spec       = dosq.querySpec()
		loadedTypes = [1]bool{
			dosq.withSkipped != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DependsOnSkipped{config: dosq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dosq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dosq.withSkipped; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DependsOnSkipped)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Skipped = []*ImplicitSkippedMessage{}
		}
		query.withFKs = true
		query.Where(predicate.ImplicitSkippedMessage(func(s *sql.Selector) {
			s.Where(sql.InValues(dependsonskipped.SkippedColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.depends_on_skipped_skipped
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "depends_on_skipped_skipped" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "depends_on_skipped_skipped" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Skipped = append(node.Edges.Skipped, n)
		}
	}

	return nodes, nil
}

func (dosq *DependsOnSkippedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dosq.querySpec()
	return sqlgraph.CountNodes(ctx, dosq.driver, _spec)
}

func (dosq *DependsOnSkippedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dosq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dosq *DependsOnSkippedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dependsonskipped.Table,
			Columns: dependsonskipped.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dependsonskipped.FieldID,
			},
		},
		From:   dosq.sql,
		Unique: true,
	}
	if unique := dosq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dosq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dependsonskipped.FieldID)
		for i := range fields {
			if fields[i] != dependsonskipped.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dosq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dosq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dosq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dosq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dosq *DependsOnSkippedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dosq.driver.Dialect())
	t1 := builder.Table(dependsonskipped.Table)
	selector := builder.Select(t1.Columns(dependsonskipped.Columns...)...).From(t1)
	if dosq.sql != nil {
		selector = dosq.sql
		selector.Select(selector.Columns(dependsonskipped.Columns...)...)
	}
	for _, p := range dosq.predicates {
		p(selector)
	}
	for _, p := range dosq.order {
		p(selector)
	}
	if offset := dosq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dosq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DependsOnSkippedGroupBy is the group-by builder for DependsOnSkipped entities.
type DependsOnSkippedGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dosgb *DependsOnSkippedGroupBy) Aggregate(fns ...AggregateFunc) *DependsOnSkippedGroupBy {
	dosgb.fns = append(dosgb.fns, fns...)
	return dosgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dosgb *DependsOnSkippedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dosgb.path(ctx)
	if err != nil {
		return err
	}
	dosgb.sql = query
	return dosgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dosgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dosgb.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dosgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) StringsX(ctx context.Context) []string {
	v, err := dosgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dosgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) StringX(ctx context.Context) string {
	v, err := dosgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dosgb.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dosgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) IntsX(ctx context.Context) []int {
	v, err := dosgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dosgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) IntX(ctx context.Context) int {
	v, err := dosgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dosgb.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dosgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dosgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dosgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dosgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dosgb.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dosgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dosgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dosgb *DependsOnSkippedGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dosgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dosgb *DependsOnSkippedGroupBy) BoolX(ctx context.Context) bool {
	v, err := dosgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dosgb *DependsOnSkippedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dosgb.fields {
		if !dependsonskipped.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dosgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dosgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dosgb *DependsOnSkippedGroupBy) sqlQuery() *sql.Selector {
	selector := dosgb.sql
	columns := make([]string, 0, len(dosgb.fields)+len(dosgb.fns))
	columns = append(columns, dosgb.fields...)
	for _, fn := range dosgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dosgb.fields...)
}

// DependsOnSkippedSelect is the builder for selecting fields of DependsOnSkipped entities.
type DependsOnSkippedSelect struct {
	*DependsOnSkippedQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (doss *DependsOnSkippedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := doss.prepareQuery(ctx); err != nil {
		return err
	}
	doss.sql = doss.DependsOnSkippedQuery.sqlQuery(ctx)
	return doss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) ScanX(ctx context.Context, v interface{}) {
	if err := doss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Strings(ctx context.Context) ([]string, error) {
	if len(doss.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := doss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) StringsX(ctx context.Context) []string {
	v, err := doss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = doss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) StringX(ctx context.Context) string {
	v, err := doss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Ints(ctx context.Context) ([]int, error) {
	if len(doss.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := doss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) IntsX(ctx context.Context) []int {
	v, err := doss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = doss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) IntX(ctx context.Context) int {
	v, err := doss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(doss.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := doss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) Float64sX(ctx context.Context) []float64 {
	v, err := doss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = doss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) Float64X(ctx context.Context) float64 {
	v, err := doss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(doss.fields) > 1 {
		return nil, errors.New("ent: DependsOnSkippedSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := doss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) BoolsX(ctx context.Context) []bool {
	v, err := doss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (doss *DependsOnSkippedSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = doss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = fmt.Errorf("ent: DependsOnSkippedSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (doss *DependsOnSkippedSelect) BoolX(ctx context.Context) bool {
	v, err := doss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (doss *DependsOnSkippedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := doss.sqlQuery().Query()
	if err := doss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (doss *DependsOnSkippedSelect) sqlQuery() sql.Querier {
	selector := doss.sql
	selector.Select(selector.Columns(doss.fields...)...)
	return selector
}
