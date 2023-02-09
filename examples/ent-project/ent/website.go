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

	"entgo.io/ent/dialect/sql"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/country"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/website"
	"github.com/google/uuid"
)

// Website is the model entity for the Website schema.
type Website struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WebsiteQuery when eager-loading is set.
	Edges            WebsiteEdges `json:"edges"`
	company_websites *uuid.UUID
	country_websites *uuid.UUID
}

// WebsiteEdges holds the relations/edges for other nodes in the graph.
type WebsiteEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebsiteEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[0] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebsiteEdges) CountryOrErr() (*Country, error) {
	if e.loadedTypes[1] {
		if e.Country == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Country, nil
	}
	return nil, &NotLoadedError{edge: "country"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Website) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case website.FieldTitle, website.FieldDescription, website.FieldURL:
			values[i] = new(sql.NullString)
		case website.FieldID:
			values[i] = new(uuid.UUID)
		case website.ForeignKeys[0]: // company_websites
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case website.ForeignKeys[1]: // country_websites
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Website", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Website fields.
func (w *Website) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case website.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				w.ID = *value
			}
		case website.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				w.Title = value.String
			}
		case website.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case website.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				w.URL = value.String
			}
		case website.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field company_websites", values[i])
			} else if value.Valid {
				w.company_websites = new(uuid.UUID)
				*w.company_websites = *value.S.(*uuid.UUID)
			}
		case website.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field country_websites", values[i])
			} else if value.Valid {
				w.country_websites = new(uuid.UUID)
				*w.country_websites = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCompany queries the "company" edge of the Website entity.
func (w *Website) QueryCompany() *CompanyQuery {
	return NewWebsiteClient(w.config).QueryCompany(w)
}

// QueryCountry queries the "country" edge of the Website entity.
func (w *Website) QueryCountry() *CountryQuery {
	return NewWebsiteClient(w.config).QueryCountry(w)
}

// Update returns a builder for updating this Website.
// Note that you need to call Website.Unwrap() before calling this method if this Website
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Website) Update() *WebsiteUpdateOne {
	return NewWebsiteClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Website entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Website) Unwrap() *Website {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Website is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Website) String() string {
	var builder strings.Builder
	builder.WriteString("Website(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("title=")
	builder.WriteString(w.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(w.Description)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(w.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Websites is a parsable slice of Website.
type Websites []*Website

func (w Websites) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}