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
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/country"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/website"
	"github.com/google/uuid"
)

// WebsiteUpdate is the builder for updating Website entities.
type WebsiteUpdate struct {
	config
	hooks    []Hook
	mutation *WebsiteMutation
}

// Where appends a list predicates to the WebsiteUpdate builder.
func (wu *WebsiteUpdate) Where(ps ...predicate.Website) *WebsiteUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WebsiteUpdate) SetTitle(s string) *WebsiteUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetDescription sets the "description" field.
func (wu *WebsiteUpdate) SetDescription(s string) *WebsiteUpdate {
	wu.mutation.SetDescription(s)
	return wu
}

// SetURL sets the "url" field.
func (wu *WebsiteUpdate) SetURL(s string) *WebsiteUpdate {
	wu.mutation.SetURL(s)
	return wu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (wu *WebsiteUpdate) SetCompanyID(id uuid.UUID) *WebsiteUpdate {
	wu.mutation.SetCompanyID(id)
	return wu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (wu *WebsiteUpdate) SetNillableCompanyID(id *uuid.UUID) *WebsiteUpdate {
	if id != nil {
		wu = wu.SetCompanyID(*id)
	}
	return wu
}

// SetCompany sets the "company" edge to the Company entity.
func (wu *WebsiteUpdate) SetCompany(c *Company) *WebsiteUpdate {
	return wu.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (wu *WebsiteUpdate) SetCountryID(id uuid.UUID) *WebsiteUpdate {
	wu.mutation.SetCountryID(id)
	return wu
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (wu *WebsiteUpdate) SetNillableCountryID(id *uuid.UUID) *WebsiteUpdate {
	if id != nil {
		wu = wu.SetCountryID(*id)
	}
	return wu
}

// SetCountry sets the "country" edge to the Country entity.
func (wu *WebsiteUpdate) SetCountry(c *Country) *WebsiteUpdate {
	return wu.SetCountryID(c.ID)
}

// Mutation returns the WebsiteMutation object of the builder.
func (wu *WebsiteUpdate) Mutation() *WebsiteMutation {
	return wu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (wu *WebsiteUpdate) ClearCompany() *WebsiteUpdate {
	wu.mutation.ClearCompany()
	return wu
}

// ClearCountry clears the "country" edge to the Country entity.
func (wu *WebsiteUpdate) ClearCountry() *WebsiteUpdate {
	wu.mutation.ClearCountry()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WebsiteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, WebsiteMutation](ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WebsiteUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WebsiteUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WebsiteUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WebsiteUpdate) check() error {
	if v, ok := wu.mutation.Title(); ok {
		if err := website.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Website.title": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Description(); ok {
		if err := website.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Website.description": %w`, err)}
		}
	}
	if v, ok := wu.mutation.URL(); ok {
		if err := website.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Website.url": %w`, err)}
		}
	}
	return nil
}

func (wu *WebsiteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   website.Table,
			Columns: website.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: website.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.SetField(website.FieldTitle, field.TypeString, value)
	}
	if value, ok := wu.mutation.Description(); ok {
		_spec.SetField(website.FieldDescription, field.TypeString, value)
	}
	if value, ok := wu.mutation.URL(); ok {
		_spec.SetField(website.FieldURL, field.TypeString, value)
	}
	if wu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CompanyTable,
			Columns: []string{website.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CompanyTable,
			Columns: []string{website.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CountryTable,
			Columns: []string{website.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CountryTable,
			Columns: []string{website.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{website.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WebsiteUpdateOne is the builder for updating a single Website entity.
type WebsiteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebsiteMutation
}

// SetTitle sets the "title" field.
func (wuo *WebsiteUpdateOne) SetTitle(s string) *WebsiteUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetDescription sets the "description" field.
func (wuo *WebsiteUpdateOne) SetDescription(s string) *WebsiteUpdateOne {
	wuo.mutation.SetDescription(s)
	return wuo
}

// SetURL sets the "url" field.
func (wuo *WebsiteUpdateOne) SetURL(s string) *WebsiteUpdateOne {
	wuo.mutation.SetURL(s)
	return wuo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (wuo *WebsiteUpdateOne) SetCompanyID(id uuid.UUID) *WebsiteUpdateOne {
	wuo.mutation.SetCompanyID(id)
	return wuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (wuo *WebsiteUpdateOne) SetNillableCompanyID(id *uuid.UUID) *WebsiteUpdateOne {
	if id != nil {
		wuo = wuo.SetCompanyID(*id)
	}
	return wuo
}

// SetCompany sets the "company" edge to the Company entity.
func (wuo *WebsiteUpdateOne) SetCompany(c *Company) *WebsiteUpdateOne {
	return wuo.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (wuo *WebsiteUpdateOne) SetCountryID(id uuid.UUID) *WebsiteUpdateOne {
	wuo.mutation.SetCountryID(id)
	return wuo
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (wuo *WebsiteUpdateOne) SetNillableCountryID(id *uuid.UUID) *WebsiteUpdateOne {
	if id != nil {
		wuo = wuo.SetCountryID(*id)
	}
	return wuo
}

// SetCountry sets the "country" edge to the Country entity.
func (wuo *WebsiteUpdateOne) SetCountry(c *Country) *WebsiteUpdateOne {
	return wuo.SetCountryID(c.ID)
}

// Mutation returns the WebsiteMutation object of the builder.
func (wuo *WebsiteUpdateOne) Mutation() *WebsiteMutation {
	return wuo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (wuo *WebsiteUpdateOne) ClearCompany() *WebsiteUpdateOne {
	wuo.mutation.ClearCompany()
	return wuo
}

// ClearCountry clears the "country" edge to the Country entity.
func (wuo *WebsiteUpdateOne) ClearCountry() *WebsiteUpdateOne {
	wuo.mutation.ClearCountry()
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WebsiteUpdateOne) Select(field string, fields ...string) *WebsiteUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Website entity.
func (wuo *WebsiteUpdateOne) Save(ctx context.Context) (*Website, error) {
	return withHooks[*Website, WebsiteMutation](ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WebsiteUpdateOne) SaveX(ctx context.Context) *Website {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WebsiteUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WebsiteUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WebsiteUpdateOne) check() error {
	if v, ok := wuo.mutation.Title(); ok {
		if err := website.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Website.title": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Description(); ok {
		if err := website.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Website.description": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.URL(); ok {
		if err := website.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Website.url": %w`, err)}
		}
	}
	return nil
}

func (wuo *WebsiteUpdateOne) sqlSave(ctx context.Context) (_node *Website, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   website.Table,
			Columns: website.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: website.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Website.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, website.FieldID)
		for _, f := range fields {
			if !website.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != website.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.SetField(website.FieldTitle, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Description(); ok {
		_spec.SetField(website.FieldDescription, field.TypeString, value)
	}
	if value, ok := wuo.mutation.URL(); ok {
		_spec.SetField(website.FieldURL, field.TypeString, value)
	}
	if wuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CompanyTable,
			Columns: []string{website.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CompanyTable,
			Columns: []string{website.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CountryTable,
			Columns: []string{website.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   website.CountryTable,
			Columns: []string{website.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Website{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{website.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}