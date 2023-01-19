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
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/enums"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/product"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/vendor"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/warehouse"
	"github.com/google/uuid"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// LastSell holds the value of the "last_sell" field.
	LastSell *time.Time `json:"last_sell,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status enums.ProcessStatus `json:"status,omitempty"`
	// BuildStatus holds the value of the "build_status" field.
	BuildStatus enums.ProcessStatus `json:"build_status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges              ProductEdges `json:"edges"`
	vendor_products    *uuid.UUID
	warehouse_products *uuid.UUID
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Warehouse holds the value of the warehouse edge.
	Warehouse *Warehouse `json:"warehouse,omitempty"`
	// Vendor holds the value of the vendor edge.
	Vendor *Vendor `json:"vendor,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// WarehouseOrErr returns the Warehouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) WarehouseOrErr() (*Warehouse, error) {
	if e.loadedTypes[0] {
		if e.Warehouse == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: warehouse.Label}
		}
		return e.Warehouse, nil
	}
	return nil, &NotLoadedError{edge: "warehouse"}
}

// VendorOrErr returns the Vendor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) VendorOrErr() (*Vendor, error) {
	if e.loadedTypes[1] {
		if e.Vendor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vendor.Label}
		}
		return e.Vendor, nil
	}
	return nil, &NotLoadedError{edge: "vendor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldName, product.FieldDescription, product.FieldImage, product.FieldURL, product.FieldStatus, product.FieldBuildStatus:
			values[i] = new(sql.NullString)
		case product.FieldLastSell, product.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case product.FieldID:
			values[i] = new(uuid.UUID)
		case product.ForeignKeys[0]: // vendor_products
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case product.ForeignKeys[1]: // warehouse_products
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				pr.Image = value.String
			}
		case product.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pr.URL = value.String
			}
		case product.FieldLastSell:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_sell", values[i])
			} else if value.Valid {
				pr.LastSell = new(time.Time)
				*pr.LastSell = value.Time
			}
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = new(time.Time)
				*pr.CreatedAt = value.Time
			}
		case product.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = enums.ProcessStatus(value.String)
			}
		case product.FieldBuildStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field build_status", values[i])
			} else if value.Valid {
				pr.BuildStatus = enums.ProcessStatus(value.String)
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field vendor_products", values[i])
			} else if value.Valid {
				pr.vendor_products = new(uuid.UUID)
				*pr.vendor_products = *value.S.(*uuid.UUID)
			}
		case product.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field warehouse_products", values[i])
			} else if value.Valid {
				pr.warehouse_products = new(uuid.UUID)
				*pr.warehouse_products = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryWarehouse queries the "warehouse" edge of the Product entity.
func (pr *Product) QueryWarehouse() *WarehouseQuery {
	return (&ProductClient{config: pr.config}).QueryWarehouse(pr)
}

// QueryVendor queries the "vendor" edge of the Product entity.
func (pr *Product) QueryVendor() *VendorQuery {
	return (&ProductClient{config: pr.config}).QueryVendor(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(pr.Image)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(pr.URL)
	builder.WriteString(", ")
	if v := pr.LastSell; v != nil {
		builder.WriteString("last_sell=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pr.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", ")
	builder.WriteString("build_status=")
	builder.WriteString(fmt.Sprintf("%v", pr.BuildStatus))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
