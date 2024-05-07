// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Simplewallethq/tg-bot/ent/invoice"
	"github.com/Simplewallethq/tg-bot/ent/invoices_payments"
	"github.com/Simplewallethq/tg-bot/ent/predicate"
)

// InvoicesPaymentsQuery is the builder for querying Invoices_payments entities.
type InvoicesPaymentsQuery struct {
	config
	ctx         *QueryContext
	order       []invoices_payments.OrderOption
	inters      []Interceptor
	predicates  []predicate.Invoices_payments
	withInvoice *InvoiceQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoicesPaymentsQuery builder.
func (ipq *InvoicesPaymentsQuery) Where(ps ...predicate.Invoices_payments) *InvoicesPaymentsQuery {
	ipq.predicates = append(ipq.predicates, ps...)
	return ipq
}

// Limit the number of records to be returned by this query.
func (ipq *InvoicesPaymentsQuery) Limit(limit int) *InvoicesPaymentsQuery {
	ipq.ctx.Limit = &limit
	return ipq
}

// Offset to start from.
func (ipq *InvoicesPaymentsQuery) Offset(offset int) *InvoicesPaymentsQuery {
	ipq.ctx.Offset = &offset
	return ipq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ipq *InvoicesPaymentsQuery) Unique(unique bool) *InvoicesPaymentsQuery {
	ipq.ctx.Unique = &unique
	return ipq
}

// Order specifies how the records should be ordered.
func (ipq *InvoicesPaymentsQuery) Order(o ...invoices_payments.OrderOption) *InvoicesPaymentsQuery {
	ipq.order = append(ipq.order, o...)
	return ipq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (ipq *InvoicesPaymentsQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: ipq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ipq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoices_payments.Table, invoices_payments.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invoices_payments.InvoiceTable, invoices_payments.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ipq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Invoices_payments entity from the query.
// Returns a *NotFoundError when no Invoices_payments was found.
func (ipq *InvoicesPaymentsQuery) First(ctx context.Context) (*Invoices_payments, error) {
	nodes, err := ipq.Limit(1).All(setContextOp(ctx, ipq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoices_payments.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) FirstX(ctx context.Context) *Invoices_payments {
	node, err := ipq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Invoices_payments ID from the query.
// Returns a *NotFoundError when no Invoices_payments ID was found.
func (ipq *InvoicesPaymentsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(1).IDs(setContextOp(ctx, ipq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoices_payments.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) FirstIDX(ctx context.Context) int {
	id, err := ipq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Invoices_payments entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Invoices_payments entity is found.
// Returns a *NotFoundError when no Invoices_payments entities are found.
func (ipq *InvoicesPaymentsQuery) Only(ctx context.Context) (*Invoices_payments, error) {
	nodes, err := ipq.Limit(2).All(setContextOp(ctx, ipq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoices_payments.Label}
	default:
		return nil, &NotSingularError{invoices_payments.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) OnlyX(ctx context.Context) *Invoices_payments {
	node, err := ipq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Invoices_payments ID in the query.
// Returns a *NotSingularError when more than one Invoices_payments ID is found.
// Returns a *NotFoundError when no entities are found.
func (ipq *InvoicesPaymentsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(2).IDs(setContextOp(ctx, ipq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoices_payments.Label}
	default:
		err = &NotSingularError{invoices_payments.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) OnlyIDX(ctx context.Context) int {
	id, err := ipq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Invoices_paymentsSlice.
func (ipq *InvoicesPaymentsQuery) All(ctx context.Context) ([]*Invoices_payments, error) {
	ctx = setContextOp(ctx, ipq.ctx, "All")
	if err := ipq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Invoices_payments, *InvoicesPaymentsQuery]()
	return withInterceptors[[]*Invoices_payments](ctx, ipq, qr, ipq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) AllX(ctx context.Context) []*Invoices_payments {
	nodes, err := ipq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Invoices_payments IDs.
func (ipq *InvoicesPaymentsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ipq.ctx.Unique == nil && ipq.path != nil {
		ipq.Unique(true)
	}
	ctx = setContextOp(ctx, ipq.ctx, "IDs")
	if err = ipq.Select(invoices_payments.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) IDsX(ctx context.Context) []int {
	ids, err := ipq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ipq *InvoicesPaymentsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ipq.ctx, "Count")
	if err := ipq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ipq, querierCount[*InvoicesPaymentsQuery](), ipq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) CountX(ctx context.Context) int {
	count, err := ipq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ipq *InvoicesPaymentsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ipq.ctx, "Exist")
	switch _, err := ipq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ipq *InvoicesPaymentsQuery) ExistX(ctx context.Context) bool {
	exist, err := ipq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoicesPaymentsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ipq *InvoicesPaymentsQuery) Clone() *InvoicesPaymentsQuery {
	if ipq == nil {
		return nil
	}
	return &InvoicesPaymentsQuery{
		config:      ipq.config,
		ctx:         ipq.ctx.Clone(),
		order:       append([]invoices_payments.OrderOption{}, ipq.order...),
		inters:      append([]Interceptor{}, ipq.inters...),
		predicates:  append([]predicate.Invoices_payments{}, ipq.predicates...),
		withInvoice: ipq.withInvoice.Clone(),
		// clone intermediate query.
		sql:  ipq.sql.Clone(),
		path: ipq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (ipq *InvoicesPaymentsQuery) WithInvoice(opts ...func(*InvoiceQuery)) *InvoicesPaymentsQuery {
	query := (&InvoiceClient{config: ipq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ipq.withInvoice = query
	return ipq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		From string `json:"from,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InvoicesPayments.Query().
//		GroupBy(invoices_payments.FieldFrom).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ipq *InvoicesPaymentsQuery) GroupBy(field string, fields ...string) *InvoicesPaymentsGroupBy {
	ipq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoicesPaymentsGroupBy{build: ipq}
	grbuild.flds = &ipq.ctx.Fields
	grbuild.label = invoices_payments.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		From string `json:"from,omitempty"`
//	}
//
//	client.InvoicesPayments.Query().
//		Select(invoices_payments.FieldFrom).
//		Scan(ctx, &v)
func (ipq *InvoicesPaymentsQuery) Select(fields ...string) *InvoicesPaymentsSelect {
	ipq.ctx.Fields = append(ipq.ctx.Fields, fields...)
	sbuild := &InvoicesPaymentsSelect{InvoicesPaymentsQuery: ipq}
	sbuild.label = invoices_payments.Label
	sbuild.flds, sbuild.scan = &ipq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoicesPaymentsSelect configured with the given aggregations.
func (ipq *InvoicesPaymentsQuery) Aggregate(fns ...AggregateFunc) *InvoicesPaymentsSelect {
	return ipq.Select().Aggregate(fns...)
}

func (ipq *InvoicesPaymentsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ipq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ipq); err != nil {
				return err
			}
		}
	}
	for _, f := range ipq.ctx.Fields {
		if !invoices_payments.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ipq.path != nil {
		prev, err := ipq.path(ctx)
		if err != nil {
			return err
		}
		ipq.sql = prev
	}
	return nil
}

func (ipq *InvoicesPaymentsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Invoices_payments, error) {
	var (
		nodes       = []*Invoices_payments{}
		withFKs     = ipq.withFKs
		_spec       = ipq.querySpec()
		loadedTypes = [1]bool{
			ipq.withInvoice != nil,
		}
	)
	if ipq.withInvoice != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, invoices_payments.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Invoices_payments).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Invoices_payments{config: ipq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ipq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ipq.withInvoice; query != nil {
		if err := ipq.loadInvoice(ctx, query, nodes, nil,
			func(n *Invoices_payments, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ipq *InvoicesPaymentsQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*Invoices_payments, init func(*Invoices_payments), assign func(*Invoices_payments, *Invoice)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Invoices_payments)
	for i := range nodes {
		if nodes[i].invoice_payments == nil {
			continue
		}
		fk := *nodes[i].invoice_payments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_payments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ipq *InvoicesPaymentsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ipq.querySpec()
	_spec.Node.Columns = ipq.ctx.Fields
	if len(ipq.ctx.Fields) > 0 {
		_spec.Unique = ipq.ctx.Unique != nil && *ipq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ipq.driver, _spec)
}

func (ipq *InvoicesPaymentsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoices_payments.Table, invoices_payments.Columns, sqlgraph.NewFieldSpec(invoices_payments.FieldID, field.TypeInt))
	_spec.From = ipq.sql
	if unique := ipq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ipq.path != nil {
		_spec.Unique = true
	}
	if fields := ipq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoices_payments.FieldID)
		for i := range fields {
			if fields[i] != invoices_payments.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ipq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ipq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ipq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ipq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ipq *InvoicesPaymentsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ipq.driver.Dialect())
	t1 := builder.Table(invoices_payments.Table)
	columns := ipq.ctx.Fields
	if len(columns) == 0 {
		columns = invoices_payments.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ipq.sql != nil {
		selector = ipq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ipq.ctx.Unique != nil && *ipq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ipq.predicates {
		p(selector)
	}
	for _, p := range ipq.order {
		p(selector)
	}
	if offset := ipq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ipq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvoicesPaymentsGroupBy is the group-by builder for Invoices_payments entities.
type InvoicesPaymentsGroupBy struct {
	selector
	build *InvoicesPaymentsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ipgb *InvoicesPaymentsGroupBy) Aggregate(fns ...AggregateFunc) *InvoicesPaymentsGroupBy {
	ipgb.fns = append(ipgb.fns, fns...)
	return ipgb
}

// Scan applies the selector query and scans the result into the given value.
func (ipgb *InvoicesPaymentsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ipgb.build.ctx, "GroupBy")
	if err := ipgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoicesPaymentsQuery, *InvoicesPaymentsGroupBy](ctx, ipgb.build, ipgb, ipgb.build.inters, v)
}

func (ipgb *InvoicesPaymentsGroupBy) sqlScan(ctx context.Context, root *InvoicesPaymentsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ipgb.fns))
	for _, fn := range ipgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ipgb.flds)+len(ipgb.fns))
		for _, f := range *ipgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ipgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ipgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoicesPaymentsSelect is the builder for selecting fields of InvoicesPayments entities.
type InvoicesPaymentsSelect struct {
	*InvoicesPaymentsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ips *InvoicesPaymentsSelect) Aggregate(fns ...AggregateFunc) *InvoicesPaymentsSelect {
	ips.fns = append(ips.fns, fns...)
	return ips
}

// Scan applies the selector query and scans the result into the given value.
func (ips *InvoicesPaymentsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ips.ctx, "Select")
	if err := ips.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoicesPaymentsQuery, *InvoicesPaymentsSelect](ctx, ips.InvoicesPaymentsQuery, ips, ips.inters, v)
}

func (ips *InvoicesPaymentsSelect) sqlScan(ctx context.Context, root *InvoicesPaymentsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ips.fns))
	for _, fn := range ips.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ips.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ips.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}