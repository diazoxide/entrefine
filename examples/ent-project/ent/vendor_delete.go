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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/predicate"
	"github.com/diazoxide/entrefine/examples/ent-project/ent/vendor"
)

// VendorDelete is the builder for deleting a Vendor entity.
type VendorDelete struct {
	config
	hooks    []Hook
	mutation *VendorMutation
}

// Where appends a list predicates to the VendorDelete builder.
func (vd *VendorDelete) Where(ps ...predicate.Vendor) *VendorDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VendorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, VendorMutation](ctx, vd.sqlExec, vd.mutation, vd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VendorDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VendorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vendor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vendor.FieldID,
			},
		},
	}
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vd.mutation.done = true
	return affected, err
}

// VendorDeleteOne is the builder for deleting a single Vendor entity.
type VendorDeleteOne struct {
	vd *VendorDelete
}

// Exec executes the deletion query.
func (vdo *VendorDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vendor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VendorDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
