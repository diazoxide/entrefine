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
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/country"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/email"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/location"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/phone"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/website"
	"github.com/google/uuid"
)

// CountryQuery is the builder for querying Country entities.
type CountryQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	predicates         []predicate.Country
	withCompanies      *CompanyQuery
	withPhones         *PhoneQuery
	withEmails         *EmailQuery
	withWebsites       *WebsiteQuery
	withLocations      *LocationQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Country) error
	withNamedCompanies map[string]*CompanyQuery
	withNamedPhones    map[string]*PhoneQuery
	withNamedEmails    map[string]*EmailQuery
	withNamedWebsites  map[string]*WebsiteQuery
	withNamedLocations map[string]*LocationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CountryQuery builder.
func (cq *CountryQuery) Where(ps ...predicate.Country) *CountryQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CountryQuery) Limit(limit int) *CountryQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CountryQuery) Offset(offset int) *CountryQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CountryQuery) Unique(unique bool) *CountryQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *CountryQuery) Order(o ...OrderFunc) *CountryQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCompanies chains the current query on the "companies" edge.
func (cq *CountryQuery) QueryCompanies() *CompanyQuery {
	query := &CompanyQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, country.CompaniesTable, country.CompaniesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPhones chains the current query on the "phones" edge.
func (cq *CountryQuery) QueryPhones() *PhoneQuery {
	query := &PhoneQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, selector),
			sqlgraph.To(phone.Table, phone.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, country.PhonesTable, country.PhonesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmails chains the current query on the "emails" edge.
func (cq *CountryQuery) QueryEmails() *EmailQuery {
	query := &EmailQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, selector),
			sqlgraph.To(email.Table, email.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, country.EmailsTable, country.EmailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWebsites chains the current query on the "websites" edge.
func (cq *CountryQuery) QueryWebsites() *WebsiteQuery {
	query := &WebsiteQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, selector),
			sqlgraph.To(website.Table, website.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, country.WebsitesTable, country.WebsitesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLocations chains the current query on the "locations" edge.
func (cq *CountryQuery) QueryLocations() *LocationQuery {
	query := &LocationQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, country.LocationsTable, country.LocationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Country entity from the query.
// Returns a *NotFoundError when no Country was found.
func (cq *CountryQuery) First(ctx context.Context) (*Country, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{country.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CountryQuery) FirstX(ctx context.Context) *Country {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Country ID from the query.
// Returns a *NotFoundError when no Country ID was found.
func (cq *CountryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{country.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CountryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Country entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Country entity is found.
// Returns a *NotFoundError when no Country entities are found.
func (cq *CountryQuery) Only(ctx context.Context) (*Country, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{country.Label}
	default:
		return nil, &NotSingularError{country.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CountryQuery) OnlyX(ctx context.Context) *Country {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Country ID in the query.
// Returns a *NotSingularError when more than one Country ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CountryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{country.Label}
	default:
		err = &NotSingularError{country.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CountryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Countries.
func (cq *CountryQuery) All(ctx context.Context) ([]*Country, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CountryQuery) AllX(ctx context.Context) []*Country {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Country IDs.
func (cq *CountryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cq.Select(country.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CountryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CountryQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CountryQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CountryQuery) Exist(ctx context.Context) (bool, error) {
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CountryQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CountryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CountryQuery) Clone() *CountryQuery {
	if cq == nil {
		return nil
	}
	return &CountryQuery{
		config:        cq.config,
		limit:         cq.limit,
		offset:        cq.offset,
		order:         append([]OrderFunc{}, cq.order...),
		predicates:    append([]predicate.Country{}, cq.predicates...),
		withCompanies: cq.withCompanies.Clone(),
		withPhones:    cq.withPhones.Clone(),
		withEmails:    cq.withEmails.Clone(),
		withWebsites:  cq.withWebsites.Clone(),
		withLocations: cq.withLocations.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithCompanies tells the query-builder to eager-load the nodes that are connected to
// the "companies" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithCompanies(opts ...func(*CompanyQuery)) *CountryQuery {
	query := &CompanyQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withCompanies = query
	return cq
}

// WithPhones tells the query-builder to eager-load the nodes that are connected to
// the "phones" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithPhones(opts ...func(*PhoneQuery)) *CountryQuery {
	query := &PhoneQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withPhones = query
	return cq
}

// WithEmails tells the query-builder to eager-load the nodes that are connected to
// the "emails" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithEmails(opts ...func(*EmailQuery)) *CountryQuery {
	query := &EmailQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withEmails = query
	return cq
}

// WithWebsites tells the query-builder to eager-load the nodes that are connected to
// the "websites" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithWebsites(opts ...func(*WebsiteQuery)) *CountryQuery {
	query := &WebsiteQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withWebsites = query
	return cq
}

// WithLocations tells the query-builder to eager-load the nodes that are connected to
// the "locations" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithLocations(opts ...func(*LocationQuery)) *CountryQuery {
	query := &LocationQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withLocations = query
	return cq
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
//	client.Country.Query().
//		GroupBy(country.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CountryQuery) GroupBy(field string, fields ...string) *CountryGroupBy {
	grbuild := &CountryGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = country.Label
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
//	client.Country.Query().
//		Select(country.FieldName).
//		Scan(ctx, &v)
func (cq *CountryQuery) Select(fields ...string) *CountrySelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &CountrySelect{CountryQuery: cq}
	selbuild.label = country.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CountrySelect configured with the given aggregations.
func (cq *CountryQuery) Aggregate(fns ...AggregateFunc) *CountrySelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CountryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !country.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CountryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Country, error) {
	var (
		nodes       = []*Country{}
		_spec       = cq.querySpec()
		loadedTypes = [5]bool{
			cq.withCompanies != nil,
			cq.withPhones != nil,
			cq.withEmails != nil,
			cq.withWebsites != nil,
			cq.withLocations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Country).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Country{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCompanies; query != nil {
		if err := cq.loadCompanies(ctx, query, nodes,
			func(n *Country) { n.Edges.Companies = []*Company{} },
			func(n *Country, e *Company) { n.Edges.Companies = append(n.Edges.Companies, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withPhones; query != nil {
		if err := cq.loadPhones(ctx, query, nodes,
			func(n *Country) { n.Edges.Phones = []*Phone{} },
			func(n *Country, e *Phone) { n.Edges.Phones = append(n.Edges.Phones, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withEmails; query != nil {
		if err := cq.loadEmails(ctx, query, nodes,
			func(n *Country) { n.Edges.Emails = []*Email{} },
			func(n *Country, e *Email) { n.Edges.Emails = append(n.Edges.Emails, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withWebsites; query != nil {
		if err := cq.loadWebsites(ctx, query, nodes,
			func(n *Country) { n.Edges.Websites = []*Website{} },
			func(n *Country, e *Website) { n.Edges.Websites = append(n.Edges.Websites, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withLocations; query != nil {
		if err := cq.loadLocations(ctx, query, nodes,
			func(n *Country) { n.Edges.Locations = []*Location{} },
			func(n *Country, e *Location) { n.Edges.Locations = append(n.Edges.Locations, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedCompanies {
		if err := cq.loadCompanies(ctx, query, nodes,
			func(n *Country) { n.appendNamedCompanies(name) },
			func(n *Country, e *Company) { n.appendNamedCompanies(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedPhones {
		if err := cq.loadPhones(ctx, query, nodes,
			func(n *Country) { n.appendNamedPhones(name) },
			func(n *Country, e *Phone) { n.appendNamedPhones(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedEmails {
		if err := cq.loadEmails(ctx, query, nodes,
			func(n *Country) { n.appendNamedEmails(name) },
			func(n *Country, e *Email) { n.appendNamedEmails(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedWebsites {
		if err := cq.loadWebsites(ctx, query, nodes,
			func(n *Country) { n.appendNamedWebsites(name) },
			func(n *Country, e *Website) { n.appendNamedWebsites(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedLocations {
		if err := cq.loadLocations(ctx, query, nodes,
			func(n *Country) { n.appendNamedLocations(name) },
			func(n *Country, e *Location) { n.appendNamedLocations(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CountryQuery) loadCompanies(ctx context.Context, query *CompanyQuery, nodes []*Country, init func(*Country), assign func(*Country, *Company)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Country)
	nids := make(map[uuid.UUID]map[*Country]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(country.CompaniesTable)
		s.Join(joinT).On(s.C(company.FieldID), joinT.C(country.CompaniesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(country.CompaniesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(country.CompaniesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Country]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "companies" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CountryQuery) loadPhones(ctx context.Context, query *PhoneQuery, nodes []*Country, init func(*Country), assign func(*Country, *Phone)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Country)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Phone(func(s *sql.Selector) {
		s.Where(sql.InValues(country.PhonesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.country_phones
		if fk == nil {
			return fmt.Errorf(`foreign-key "country_phones" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_phones" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CountryQuery) loadEmails(ctx context.Context, query *EmailQuery, nodes []*Country, init func(*Country), assign func(*Country, *Email)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Country)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Email(func(s *sql.Selector) {
		s.Where(sql.InValues(country.EmailsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.country_emails
		if fk == nil {
			return fmt.Errorf(`foreign-key "country_emails" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_emails" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CountryQuery) loadWebsites(ctx context.Context, query *WebsiteQuery, nodes []*Country, init func(*Country), assign func(*Country, *Website)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Country)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Website(func(s *sql.Selector) {
		s.Where(sql.InValues(country.WebsitesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.country_websites
		if fk == nil {
			return fmt.Errorf(`foreign-key "country_websites" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_websites" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CountryQuery) loadLocations(ctx context.Context, query *LocationQuery, nodes []*Country, init func(*Country), assign func(*Country, *Location)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Country)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Location(func(s *sql.Selector) {
		s.Where(sql.InValues(country.LocationsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.country_locations
		if fk == nil {
			return fmt.Errorf(`foreign-key "country_locations" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_locations" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CountryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CountryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   country.Table,
			Columns: country.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: country.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for i := range fields {
			if fields[i] != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CountryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(country.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = country.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedCompanies tells the query-builder to eager-load the nodes that are connected to the "companies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithNamedCompanies(name string, opts ...func(*CompanyQuery)) *CountryQuery {
	query := &CompanyQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedCompanies == nil {
		cq.withNamedCompanies = make(map[string]*CompanyQuery)
	}
	cq.withNamedCompanies[name] = query
	return cq
}

// WithNamedPhones tells the query-builder to eager-load the nodes that are connected to the "phones"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithNamedPhones(name string, opts ...func(*PhoneQuery)) *CountryQuery {
	query := &PhoneQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedPhones == nil {
		cq.withNamedPhones = make(map[string]*PhoneQuery)
	}
	cq.withNamedPhones[name] = query
	return cq
}

// WithNamedEmails tells the query-builder to eager-load the nodes that are connected to the "emails"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithNamedEmails(name string, opts ...func(*EmailQuery)) *CountryQuery {
	query := &EmailQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedEmails == nil {
		cq.withNamedEmails = make(map[string]*EmailQuery)
	}
	cq.withNamedEmails[name] = query
	return cq
}

// WithNamedWebsites tells the query-builder to eager-load the nodes that are connected to the "websites"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithNamedWebsites(name string, opts ...func(*WebsiteQuery)) *CountryQuery {
	query := &WebsiteQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedWebsites == nil {
		cq.withNamedWebsites = make(map[string]*WebsiteQuery)
	}
	cq.withNamedWebsites[name] = query
	return cq
}

// WithNamedLocations tells the query-builder to eager-load the nodes that are connected to the "locations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithNamedLocations(name string, opts ...func(*LocationQuery)) *CountryQuery {
	query := &LocationQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedLocations == nil {
		cq.withNamedLocations = make(map[string]*LocationQuery)
	}
	cq.withNamedLocations[name] = query
	return cq
}

// CountryGroupBy is the group-by builder for Country entities.
type CountryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CountryGroupBy) Aggregate(fns ...AggregateFunc) *CountryGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CountryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *CountryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cgb.fields {
		if !country.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CountryGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// CountrySelect is the builder for selecting fields of Country entities.
type CountrySelect struct {
	*CountryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CountrySelect) Aggregate(fns ...AggregateFunc) *CountrySelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CountrySelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CountryQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *CountrySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(cs.sql))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}