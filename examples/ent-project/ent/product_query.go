// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/product"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/vendor"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/warehouse"
	"github.com/google/uuid"
)

// ProductQuery is the builder for querying Product entities.
type ProductQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Product
	withWarehouse *WarehouseQuery
	withVendor    *VendorQuery
	withFKs       bool
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*Product) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductQuery builder.
func (pq *ProductQuery) Where(ps ...predicate.Product) *ProductQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *ProductQuery) Limit(limit int) *ProductQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *ProductQuery) Offset(offset int) *ProductQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProductQuery) Unique(unique bool) *ProductQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *ProductQuery) Order(o ...OrderFunc) *ProductQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryWarehouse chains the current query on the "warehouse" edge.
func (pq *ProductQuery) QueryWarehouse() *WarehouseQuery {
	query := &WarehouseQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(warehouse.Table, warehouse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, product.WarehouseTable, product.WarehouseColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVendor chains the current query on the "vendor" edge.
func (pq *ProductQuery) QueryVendor() *VendorQuery {
	query := &VendorQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(vendor.Table, vendor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, product.VendorTable, product.VendorColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Product entity from the query.
// Returns a *NotFoundError when no Product was found.
func (pq *ProductQuery) First(ctx context.Context) (*Product, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{product.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProductQuery) FirstX(ctx context.Context) *Product {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Product ID from the query.
// Returns a *NotFoundError when no Product ID was found.
func (pq *ProductQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{product.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProductQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Product entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Product entity is found.
// Returns a *NotFoundError when no Product entities are found.
func (pq *ProductQuery) Only(ctx context.Context) (*Product, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{product.Label}
	default:
		return nil, &NotSingularError{product.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProductQuery) OnlyX(ctx context.Context) *Product {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Product ID in the query.
// Returns a *NotSingularError when more than one Product ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProductQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = &NotSingularError{product.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProductQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Products.
func (pq *ProductQuery) All(ctx context.Context) ([]*Product, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProductQuery) AllX(ctx context.Context) []*Product {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Product IDs.
func (pq *ProductQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := pq.Select(product.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProductQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProductQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProductQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProductQuery) Exist(ctx context.Context) (bool, error) {
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProductQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProductQuery) Clone() *ProductQuery {
	if pq == nil {
		return nil
	}
	return &ProductQuery{
		config:        pq.config,
		limit:         pq.limit,
		offset:        pq.offset,
		order:         append([]OrderFunc{}, pq.order...),
		predicates:    append([]predicate.Product{}, pq.predicates...),
		withWarehouse: pq.withWarehouse.Clone(),
		withVendor:    pq.withVendor.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithWarehouse tells the query-builder to eager-load the nodes that are connected to
// the "warehouse" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithWarehouse(opts ...func(*WarehouseQuery)) *ProductQuery {
	query := &WarehouseQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withWarehouse = query
	return pq
}

// WithVendor tells the query-builder to eager-load the nodes that are connected to
// the "vendor" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithVendor(opts ...func(*VendorQuery)) *ProductQuery {
	query := &VendorQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withVendor = query
	return pq
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
//	client.Product.Query().
//		GroupBy(product.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProductQuery) GroupBy(field string, fields ...string) *ProductGroupBy {
	grbuild := &ProductGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = product.Label
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
//	client.Product.Query().
//		Select(product.FieldName).
//		Scan(ctx, &v)
func (pq *ProductQuery) Select(fields ...string) *ProductSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &ProductSelect{ProductQuery: pq}
	selbuild.label = product.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ProductSelect configured with the given aggregations.
func (pq *ProductQuery) Aggregate(fns ...AggregateFunc) *ProductSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Product, error) {
	var (
		nodes       = []*Product{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withWarehouse != nil,
			pq.withVendor != nil,
		}
	)
	if pq.withWarehouse != nil || pq.withVendor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, product.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Product).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Product{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withWarehouse; query != nil {
		if err := pq.loadWarehouse(ctx, query, nodes, nil,
			func(n *Product, e *Warehouse) { n.Edges.Warehouse = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withVendor; query != nil {
		if err := pq.loadVendor(ctx, query, nodes, nil,
			func(n *Product, e *Vendor) { n.Edges.Vendor = e }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProductQuery) loadWarehouse(ctx context.Context, query *WarehouseQuery, nodes []*Product, init func(*Product), assign func(*Product, *Warehouse)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Product)
	for i := range nodes {
		if nodes[i].warehouse_products == nil {
			continue
		}
		fk := *nodes[i].warehouse_products
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(warehouse.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "warehouse_products" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProductQuery) loadVendor(ctx context.Context, query *VendorQuery, nodes []*Product, init func(*Product), assign func(*Product, *Vendor)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Product)
	for i := range nodes {
		if nodes[i].vendor_products == nil {
			continue
		}
		fk := *nodes[i].vendor_products
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(vendor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "vendor_products" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *ProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: product.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for i := range fields {
			if fields[i] != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(product.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = product.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductGroupBy is the group-by builder for Product entities.
type ProductGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProductGroupBy) Aggregate(fns ...AggregateFunc) *ProductGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *ProductGroupBy) Scan(ctx context.Context, v any) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *ProductGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range pgb.fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *ProductGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// ProductSelect is the builder for selecting fields of Product entities.
type ProductSelect struct {
	*ProductQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProductSelect) Aggregate(fns ...AggregateFunc) *ProductSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProductSelect) Scan(ctx context.Context, v any) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.ProductQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *ProductSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(ps.sql))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}