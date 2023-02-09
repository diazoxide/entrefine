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
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/email"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// EmailUpdate is the builder for updating Email entities.
type EmailUpdate struct {
	config
	hooks    []Hook
	mutation *EmailMutation
}

// Where appends a list predicates to the EmailUpdate builder.
func (eu *EmailUpdate) Where(ps ...predicate.Email) *EmailUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetTitle sets the "title" field.
func (eu *EmailUpdate) SetTitle(s string) *EmailUpdate {
	eu.mutation.SetTitle(s)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EmailUpdate) SetDescription(s string) *EmailUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetAddress sets the "address" field.
func (eu *EmailUpdate) SetAddress(s string) *EmailUpdate {
	eu.mutation.SetAddress(s)
	return eu
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (eu *EmailUpdate) SetCompanyID(id uuid.UUID) *EmailUpdate {
	eu.mutation.SetCompanyID(id)
	return eu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (eu *EmailUpdate) SetNillableCompanyID(id *uuid.UUID) *EmailUpdate {
	if id != nil {
		eu = eu.SetCompanyID(*id)
	}
	return eu
}

// SetCompany sets the "company" edge to the Company entity.
func (eu *EmailUpdate) SetCompany(c *Company) *EmailUpdate {
	return eu.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (eu *EmailUpdate) SetCountryID(id uuid.UUID) *EmailUpdate {
	eu.mutation.SetCountryID(id)
	return eu
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (eu *EmailUpdate) SetNillableCountryID(id *uuid.UUID) *EmailUpdate {
	if id != nil {
		eu = eu.SetCountryID(*id)
	}
	return eu
}

// SetCountry sets the "country" edge to the Country entity.
func (eu *EmailUpdate) SetCountry(c *Country) *EmailUpdate {
	return eu.SetCountryID(c.ID)
}

// Mutation returns the EmailMutation object of the builder.
func (eu *EmailUpdate) Mutation() *EmailMutation {
	return eu.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (eu *EmailUpdate) ClearCompany() *EmailUpdate {
	eu.mutation.ClearCompany()
	return eu
}

// ClearCountry clears the "country" edge to the Country entity.
func (eu *EmailUpdate) ClearCountry() *EmailUpdate {
	eu.mutation.ClearCountry()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmailUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EmailMutation](ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmailUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmailUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmailUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EmailUpdate) check() error {
	if v, ok := eu.mutation.Title(); ok {
		if err := email.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Email.title": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Description(); ok {
		if err := email.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Email.description": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Address(); ok {
		if err := email.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Email.address": %w`, err)}
		}
	}
	return nil
}

func (eu *EmailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   email.Table,
			Columns: email.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: email.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Title(); ok {
		_spec.SetField(email.FieldTitle, field.TypeString, value)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(email.FieldDescription, field.TypeString, value)
	}
	if value, ok := eu.mutation.Address(); ok {
		_spec.SetField(email.FieldAddress, field.TypeString, value)
	}
	if eu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CompanyTable,
			Columns: []string{email.CompanyColumn},
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
	if nodes := eu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CompanyTable,
			Columns: []string{email.CompanyColumn},
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
	if eu.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CountryTable,
			Columns: []string{email.CountryColumn},
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
	if nodes := eu.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CountryTable,
			Columns: []string{email.CountryColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{email.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EmailUpdateOne is the builder for updating a single Email entity.
type EmailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailMutation
}

// SetTitle sets the "title" field.
func (euo *EmailUpdateOne) SetTitle(s string) *EmailUpdateOne {
	euo.mutation.SetTitle(s)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EmailUpdateOne) SetDescription(s string) *EmailUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetAddress sets the "address" field.
func (euo *EmailUpdateOne) SetAddress(s string) *EmailUpdateOne {
	euo.mutation.SetAddress(s)
	return euo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (euo *EmailUpdateOne) SetCompanyID(id uuid.UUID) *EmailUpdateOne {
	euo.mutation.SetCompanyID(id)
	return euo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (euo *EmailUpdateOne) SetNillableCompanyID(id *uuid.UUID) *EmailUpdateOne {
	if id != nil {
		euo = euo.SetCompanyID(*id)
	}
	return euo
}

// SetCompany sets the "company" edge to the Company entity.
func (euo *EmailUpdateOne) SetCompany(c *Company) *EmailUpdateOne {
	return euo.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (euo *EmailUpdateOne) SetCountryID(id uuid.UUID) *EmailUpdateOne {
	euo.mutation.SetCountryID(id)
	return euo
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (euo *EmailUpdateOne) SetNillableCountryID(id *uuid.UUID) *EmailUpdateOne {
	if id != nil {
		euo = euo.SetCountryID(*id)
	}
	return euo
}

// SetCountry sets the "country" edge to the Country entity.
func (euo *EmailUpdateOne) SetCountry(c *Country) *EmailUpdateOne {
	return euo.SetCountryID(c.ID)
}

// Mutation returns the EmailMutation object of the builder.
func (euo *EmailUpdateOne) Mutation() *EmailMutation {
	return euo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (euo *EmailUpdateOne) ClearCompany() *EmailUpdateOne {
	euo.mutation.ClearCompany()
	return euo
}

// ClearCountry clears the "country" edge to the Country entity.
func (euo *EmailUpdateOne) ClearCountry() *EmailUpdateOne {
	euo.mutation.ClearCountry()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmailUpdateOne) Select(field string, fields ...string) *EmailUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Email entity.
func (euo *EmailUpdateOne) Save(ctx context.Context) (*Email, error) {
	return withHooks[*Email, EmailMutation](ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmailUpdateOne) SaveX(ctx context.Context) *Email {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmailUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmailUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EmailUpdateOne) check() error {
	if v, ok := euo.mutation.Title(); ok {
		if err := email.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Email.title": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Description(); ok {
		if err := email.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Email.description": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Address(); ok {
		if err := email.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Email.address": %w`, err)}
		}
	}
	return nil
}

func (euo *EmailUpdateOne) sqlSave(ctx context.Context) (_node *Email, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   email.Table,
			Columns: email.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: email.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Email.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, email.FieldID)
		for _, f := range fields {
			if !email.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != email.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Title(); ok {
		_spec.SetField(email.FieldTitle, field.TypeString, value)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(email.FieldDescription, field.TypeString, value)
	}
	if value, ok := euo.mutation.Address(); ok {
		_spec.SetField(email.FieldAddress, field.TypeString, value)
	}
	if euo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CompanyTable,
			Columns: []string{email.CompanyColumn},
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
	if nodes := euo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CompanyTable,
			Columns: []string{email.CompanyColumn},
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
	if euo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CountryTable,
			Columns: []string{email.CountryColumn},
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
	if nodes := euo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   email.CountryTable,
			Columns: []string{email.CountryColumn},
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
	_node = &Email{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{email.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}