// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/predicate"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/undelegates"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
)

// UndelegatesQuery is the builder for querying Undelegates entities.
type UndelegatesQuery struct {
	config
	ctx        *QueryContext
	order      []undelegates.OrderOption
	inters     []Interceptor
	predicates []predicate.Undelegates
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UndelegatesQuery builder.
func (uq *UndelegatesQuery) Where(ps ...predicate.Undelegates) *UndelegatesQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UndelegatesQuery) Limit(limit int) *UndelegatesQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UndelegatesQuery) Offset(offset int) *UndelegatesQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UndelegatesQuery) Unique(unique bool) *UndelegatesQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UndelegatesQuery) Order(o ...undelegates.OrderOption) *UndelegatesQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryOwner chains the current query on the "owner" edge.
func (uq *UndelegatesQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(undelegates.Table, undelegates.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, undelegates.OwnerTable, undelegates.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Undelegates entity from the query.
// Returns a *NotFoundError when no Undelegates was found.
func (uq *UndelegatesQuery) First(ctx context.Context) (*Undelegates, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{undelegates.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UndelegatesQuery) FirstX(ctx context.Context) *Undelegates {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Undelegates ID from the query.
// Returns a *NotFoundError when no Undelegates ID was found.
func (uq *UndelegatesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{undelegates.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UndelegatesQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Undelegates entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Undelegates entity is found.
// Returns a *NotFoundError when no Undelegates entities are found.
func (uq *UndelegatesQuery) Only(ctx context.Context) (*Undelegates, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{undelegates.Label}
	default:
		return nil, &NotSingularError{undelegates.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UndelegatesQuery) OnlyX(ctx context.Context) *Undelegates {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Undelegates ID in the query.
// Returns a *NotSingularError when more than one Undelegates ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UndelegatesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{undelegates.Label}
	default:
		err = &NotSingularError{undelegates.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UndelegatesQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UndelegatesSlice.
func (uq *UndelegatesQuery) All(ctx context.Context) ([]*Undelegates, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Undelegates, *UndelegatesQuery]()
	return withInterceptors[[]*Undelegates](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UndelegatesQuery) AllX(ctx context.Context) []*Undelegates {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Undelegates IDs.
func (uq *UndelegatesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err = uq.Select(undelegates.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UndelegatesQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UndelegatesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UndelegatesQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UndelegatesQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UndelegatesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UndelegatesQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UndelegatesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UndelegatesQuery) Clone() *UndelegatesQuery {
	if uq == nil {
		return nil
	}
	return &UndelegatesQuery{
		config:     uq.config,
		ctx:        uq.ctx.Clone(),
		order:      append([]undelegates.OrderOption{}, uq.order...),
		inters:     append([]Interceptor{}, uq.inters...),
		predicates: append([]predicate.Undelegates{}, uq.predicates...),
		withOwner:  uq.withOwner.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UndelegatesQuery) WithOwner(opts ...func(*UserQuery)) *UndelegatesQuery {
	query := (&UserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withOwner = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Delegator string `json:"delegator,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Undelegates.Query().
//		GroupBy(undelegates.FieldDelegator).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UndelegatesQuery) GroupBy(field string, fields ...string) *UndelegatesGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UndelegatesGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = undelegates.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Delegator string `json:"delegator,omitempty"`
//	}
//
//	client.Undelegates.Query().
//		Select(undelegates.FieldDelegator).
//		Scan(ctx, &v)
func (uq *UndelegatesQuery) Select(fields ...string) *UndelegatesSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UndelegatesSelect{UndelegatesQuery: uq}
	sbuild.label = undelegates.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UndelegatesSelect configured with the given aggregations.
func (uq *UndelegatesQuery) Aggregate(fns ...AggregateFunc) *UndelegatesSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UndelegatesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !undelegates.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UndelegatesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Undelegates, error) {
	var (
		nodes       = []*Undelegates{}
		withFKs     = uq.withFKs
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withOwner != nil,
		}
	)
	if uq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, undelegates.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Undelegates).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Undelegates{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withOwner; query != nil {
		if err := uq.loadOwner(ctx, query, nodes, nil,
			func(n *Undelegates, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UndelegatesQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Undelegates, init func(*Undelegates), assign func(*Undelegates, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Undelegates)
	for i := range nodes {
		if nodes[i].user_undelegates == nil {
			continue
		}
		fk := *nodes[i].user_undelegates
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_undelegates" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uq *UndelegatesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UndelegatesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(undelegates.Table, undelegates.Columns, sqlgraph.NewFieldSpec(undelegates.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, undelegates.FieldID)
		for i := range fields {
			if fields[i] != undelegates.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UndelegatesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(undelegates.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = undelegates.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UndelegatesGroupBy is the group-by builder for Undelegates entities.
type UndelegatesGroupBy struct {
	selector
	build *UndelegatesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UndelegatesGroupBy) Aggregate(fns ...AggregateFunc) *UndelegatesGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UndelegatesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UndelegatesQuery, *UndelegatesGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UndelegatesGroupBy) sqlScan(ctx context.Context, root *UndelegatesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UndelegatesSelect is the builder for selecting fields of Undelegates entities.
type UndelegatesSelect struct {
	*UndelegatesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UndelegatesSelect) Aggregate(fns ...AggregateFunc) *UndelegatesSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UndelegatesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UndelegatesQuery, *UndelegatesSelect](ctx, us.UndelegatesQuery, us, us.inters, v)
}

func (us *UndelegatesSelect) sqlScan(ctx context.Context, root *UndelegatesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
