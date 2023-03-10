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

package warehouse

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the warehouse type in the database.
	Label = "warehouse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLastUpdate holds the string denoting the last_update field in the database.
	FieldLastUpdate = "last_update"
	// FieldOriginalData holds the string denoting the original_data field in the database.
	FieldOriginalData = "original_data"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// FieldFilters holds the string denoting the filters field in the database.
	FieldFilters = "filters"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeVendor holds the string denoting the vendor edge name in mutations.
	EdgeVendor = "vendor"
	// Table holds the table name of the warehouse in the database.
	Table = "warehouses"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "warehouse_products"
	// VendorTable is the table that holds the vendor relation/edge.
	VendorTable = "warehouses"
	// VendorInverseTable is the table name for the Vendor entity.
	// It exists in this package in order to avoid circular dependency with the "vendor" package.
	VendorInverseTable = "vendors"
	// VendorColumn is the table column denoting the vendor relation/edge.
	VendorColumn = "vendor_warehouses"
)

// Columns holds all SQL columns for warehouse fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLastUpdate,
	FieldOriginalData,
	FieldEnabled,
	FieldFilters,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "warehouses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"vendor_warehouses",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
