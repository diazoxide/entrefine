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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/image"
	"github.com/google/uuid"
)

// ImageCreate is the builder for creating a Image entity.
type ImageCreate struct {
	config
	mutation *ImageMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (ic *ImageCreate) SetTitle(s string) *ImageCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetOriginalURL sets the "original_url" field.
func (ic *ImageCreate) SetOriginalURL(s string) *ImageCreate {
	ic.mutation.SetOriginalURL(s)
	return ic
}

// SetID sets the "id" field.
func (ic *ImageCreate) SetID(u uuid.UUID) *ImageCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *ImageCreate) SetNillableID(u *uuid.UUID) *ImageCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetGalleryCompanyID sets the "gallery_company" edge to the Company entity by ID.
func (ic *ImageCreate) SetGalleryCompanyID(id uuid.UUID) *ImageCreate {
	ic.mutation.SetGalleryCompanyID(id)
	return ic
}

// SetNillableGalleryCompanyID sets the "gallery_company" edge to the Company entity by ID if the given value is not nil.
func (ic *ImageCreate) SetNillableGalleryCompanyID(id *uuid.UUID) *ImageCreate {
	if id != nil {
		ic = ic.SetGalleryCompanyID(*id)
	}
	return ic
}

// SetGalleryCompany sets the "gallery_company" edge to the Company entity.
func (ic *ImageCreate) SetGalleryCompany(c *Company) *ImageCreate {
	return ic.SetGalleryCompanyID(c.ID)
}

// SetLogoCompanyID sets the "logo_company" edge to the Company entity by ID.
func (ic *ImageCreate) SetLogoCompanyID(id uuid.UUID) *ImageCreate {
	ic.mutation.SetLogoCompanyID(id)
	return ic
}

// SetNillableLogoCompanyID sets the "logo_company" edge to the Company entity by ID if the given value is not nil.
func (ic *ImageCreate) SetNillableLogoCompanyID(id *uuid.UUID) *ImageCreate {
	if id != nil {
		ic = ic.SetLogoCompanyID(*id)
	}
	return ic
}

// SetLogoCompany sets the "logo_company" edge to the Company entity.
func (ic *ImageCreate) SetLogoCompany(c *Company) *ImageCreate {
	return ic.SetLogoCompanyID(c.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (ic *ImageCreate) Mutation() *ImageMutation {
	return ic.mutation
}

// Save creates the Image in the database.
func (ic *ImageCreate) Save(ctx context.Context) (*Image, error) {
	var (
		err  error
		node *Image
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Image)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ImageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImageCreate) SaveX(ctx context.Context) *Image {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImageCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImageCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ImageCreate) defaults() {
	if _, ok := ic.mutation.ID(); !ok {
		v := image.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImageCreate) check() error {
	if _, ok := ic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Image.title"`)}
	}
	if v, ok := ic.mutation.Title(); ok {
		if err := image.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Image.title": %w`, err)}
		}
	}
	if _, ok := ic.mutation.OriginalURL(); !ok {
		return &ValidationError{Name: "original_url", err: errors.New(`ent: missing required field "Image.original_url"`)}
	}
	if v, ok := ic.mutation.OriginalURL(); ok {
		if err := image.OriginalURLValidator(v); err != nil {
			return &ValidationError{Name: "original_url", err: fmt.Errorf(`ent: validator failed for field "Image.original_url": %w`, err)}
		}
	}
	return nil
}

func (ic *ImageCreate) sqlSave(ctx context.Context) (*Image, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ic *ImageCreate) createSpec() (*Image, *sqlgraph.CreateSpec) {
	var (
		_node = &Image{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: image.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: image.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.Title(); ok {
		_spec.SetField(image.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ic.mutation.OriginalURL(); ok {
		_spec.SetField(image.FieldOriginalURL, field.TypeString, value)
		_node.OriginalURL = value
	}
	if nodes := ic.mutation.GalleryCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.GalleryCompanyTable,
			Columns: []string{image.GalleryCompanyColumn},
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
		_node.company_gallery_images = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.LogoCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   image.LogoCompanyTable,
			Columns: []string{image.LogoCompanyColumn},
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
		_node.company_logo_image = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageCreateBulk is the builder for creating many Image entities in bulk.
type ImageCreateBulk struct {
	config
	builders []*ImageCreate
}

// Save creates the Image entities in the database.
func (icb *ImageCreateBulk) Save(ctx context.Context) ([]*Image, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Image, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImageCreateBulk) SaveX(ctx context.Context) []*Image {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImageCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImageCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}