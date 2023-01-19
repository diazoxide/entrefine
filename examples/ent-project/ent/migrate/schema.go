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

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 128},
		{Name: "logo", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "description", Type: field.TypeString, Size: 512},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 128},
		{Name: "code", Type: field.TypeString, Size: 2},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:       "countries",
		Columns:    CountriesColumns,
		PrimaryKey: []*schema.Column{CountriesColumns[0]},
	}
	// EmailsColumns holds the columns for the "emails" table.
	EmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 128},
		{Name: "description", Type: field.TypeString, Size: 500},
		{Name: "address", Type: field.TypeString, Size: 128},
		{Name: "company_emails", Type: field.TypeUUID, Nullable: true},
		{Name: "country_emails", Type: field.TypeUUID, Nullable: true},
	}
	// EmailsTable holds the schema information for the "emails" table.
	EmailsTable = &schema.Table{
		Name:       "emails",
		Columns:    EmailsColumns,
		PrimaryKey: []*schema.Column{EmailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emails_companies_emails",
				Columns:    []*schema.Column{EmailsColumns[4]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "emails_countries_emails",
				Columns:    []*schema.Column{EmailsColumns[5]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 128},
		{Name: "original_url", Type: field.TypeString, Size: 128},
		{Name: "company_logo_image", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "company_gallery_images", Type: field.TypeUUID, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_companies_logo_image",
				Columns:    []*schema.Column{ImagesColumns[3]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_companies_gallery_images",
				Columns:    []*schema.Column{ImagesColumns[4]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 128},
		{Name: "description", Type: field.TypeString, Size: 500},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "address", Type: field.TypeString, Size: 128},
		{Name: "postcode", Type: field.TypeString, Size: 8},
		{Name: "type", Type: field.TypeString, Size: 64},
		{Name: "state", Type: field.TypeString, Size: 64},
		{Name: "suburb", Type: field.TypeString, Size: 64},
		{Name: "street_type", Type: field.TypeString, Size: 64},
		{Name: "street_name", Type: field.TypeString, Size: 64},
		{Name: "company_locations", Type: field.TypeUUID, Nullable: true},
		{Name: "country_locations", Type: field.TypeUUID, Nullable: true},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "locations_companies_locations",
				Columns:    []*schema.Column{LocationsColumns[12]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "locations_countries_locations",
				Columns:    []*schema.Column{LocationsColumns[13]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PhonesColumns holds the columns for the "phones" table.
	PhonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 128},
		{Name: "description", Type: field.TypeString, Size: 500},
		{Name: "number", Type: field.TypeString, Size: 24},
		{Name: "type", Type: field.TypeString, Size: 24},
		{Name: "company_phones", Type: field.TypeUUID, Nullable: true},
		{Name: "country_phones", Type: field.TypeUUID, Nullable: true},
	}
	// PhonesTable holds the schema information for the "phones" table.
	PhonesTable = &schema.Table{
		Name:       "phones",
		Columns:    PhonesColumns,
		PrimaryKey: []*schema.Column{PhonesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "phones_companies_phones",
				Columns:    []*schema.Column{PhonesColumns[5]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "phones_countries_phones",
				Columns:    []*schema.Column{PhonesColumns[6]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 128},
		{Name: "description", Type: field.TypeString, Size: 512},
		{Name: "image", Type: field.TypeString, Size: 128},
		{Name: "url", Type: field.TypeString, Size: 128},
		{Name: "last_sell", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"none", "done", "enqueued", "in_progress", "failed"}, Default: "none"},
		{Name: "build_status", Type: field.TypeEnum, Enums: []string{"none", "done", "enqueued", "in_progress", "failed"}, Default: "none"},
		{Name: "vendor_products", Type: field.TypeUUID, Nullable: true},
		{Name: "warehouse_products", Type: field.TypeUUID, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_vendors_products",
				Columns:    []*schema.Column{ProductsColumns[9]},
				RefColumns: []*schema.Column{VendorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "products_warehouses_products",
				Columns:    []*schema.Column{ProductsColumns[10]},
				RefColumns: []*schema.Column{WarehousesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VendorsColumns holds the columns for the "vendors" table.
	VendorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 128},
		{Name: "schema", Type: field.TypeString, Size: 2147483647},
	}
	// VendorsTable holds the schema information for the "vendors" table.
	VendorsTable = &schema.Table{
		Name:       "vendors",
		Columns:    VendorsColumns,
		PrimaryKey: []*schema.Column{VendorsColumns[0]},
	}
	// WarehousesColumns holds the columns for the "warehouses" table.
	WarehousesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 128},
		{Name: "last_update", Type: field.TypeTime, Nullable: true},
		{Name: "original_data", Type: field.TypeString, Nullable: true},
		{Name: "enabled", Type: field.TypeBool, Default: true},
		{Name: "filters", Type: field.TypeJSON, Nullable: true},
		{Name: "vendor_warehouses", Type: field.TypeUUID, Nullable: true},
	}
	// WarehousesTable holds the schema information for the "warehouses" table.
	WarehousesTable = &schema.Table{
		Name:       "warehouses",
		Columns:    WarehousesColumns,
		PrimaryKey: []*schema.Column{WarehousesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "warehouses_vendors_warehouses",
				Columns:    []*schema.Column{WarehousesColumns[6]},
				RefColumns: []*schema.Column{VendorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WebsitesColumns holds the columns for the "websites" table.
	WebsitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString, Size: 128},
		{Name: "description", Type: field.TypeString, Size: 500},
		{Name: "url", Type: field.TypeString, Size: 128},
		{Name: "company_websites", Type: field.TypeUUID, Nullable: true},
		{Name: "country_websites", Type: field.TypeUUID, Nullable: true},
	}
	// WebsitesTable holds the schema information for the "websites" table.
	WebsitesTable = &schema.Table{
		Name:       "websites",
		Columns:    WebsitesColumns,
		PrimaryKey: []*schema.Column{WebsitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "websites_companies_websites",
				Columns:    []*schema.Column{WebsitesColumns[4]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "websites_countries_websites",
				Columns:    []*schema.Column{WebsitesColumns[5]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CompanyCountriesColumns holds the columns for the "company_countries" table.
	CompanyCountriesColumns = []*schema.Column{
		{Name: "company_id", Type: field.TypeUUID},
		{Name: "country_id", Type: field.TypeUUID},
	}
	// CompanyCountriesTable holds the schema information for the "company_countries" table.
	CompanyCountriesTable = &schema.Table{
		Name:       "company_countries",
		Columns:    CompanyCountriesColumns,
		PrimaryKey: []*schema.Column{CompanyCountriesColumns[0], CompanyCountriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_countries_company_id",
				Columns:    []*schema.Column{CompanyCountriesColumns[0]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "company_countries_country_id",
				Columns:    []*schema.Column{CompanyCountriesColumns[1]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompaniesTable,
		CountriesTable,
		EmailsTable,
		ImagesTable,
		LocationsTable,
		PhonesTable,
		ProductsTable,
		VendorsTable,
		WarehousesTable,
		WebsitesTable,
		CompanyCountriesTable,
	}
)

func init() {
	EmailsTable.ForeignKeys[0].RefTable = CompaniesTable
	EmailsTable.ForeignKeys[1].RefTable = CountriesTable
	ImagesTable.ForeignKeys[0].RefTable = CompaniesTable
	ImagesTable.ForeignKeys[1].RefTable = CompaniesTable
	LocationsTable.ForeignKeys[0].RefTable = CompaniesTable
	LocationsTable.ForeignKeys[1].RefTable = CountriesTable
	PhonesTable.ForeignKeys[0].RefTable = CompaniesTable
	PhonesTable.ForeignKeys[1].RefTable = CountriesTable
	ProductsTable.ForeignKeys[0].RefTable = VendorsTable
	ProductsTable.ForeignKeys[1].RefTable = WarehousesTable
	WarehousesTable.ForeignKeys[0].RefTable = VendorsTable
	WebsitesTable.ForeignKeys[0].RefTable = CompaniesTable
	WebsitesTable.ForeignKeys[1].RefTable = CountriesTable
	CompanyCountriesTable.ForeignKeys[0].RefTable = CompaniesTable
	CompanyCountriesTable.ForeignKeys[1].RefTable = CountriesTable
}
