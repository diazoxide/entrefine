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

package country

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Country {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Country {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Country {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Country {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// HasCompanies applies the HasEdge predicate on the "companies" edge.
func HasCompanies() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CompaniesTable, CompaniesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompaniesWith applies the HasEdge predicate on the "companies" edge with a given conditions (other predicates).
func HasCompaniesWith(preds ...predicate.Company) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CompaniesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CompaniesTable, CompaniesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPhones applies the HasEdge predicate on the "phones" edge.
func HasPhones() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PhonesTable, PhonesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPhonesWith applies the HasEdge predicate on the "phones" edge with a given conditions (other predicates).
func HasPhonesWith(preds ...predicate.Phone) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PhonesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PhonesTable, PhonesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmails applies the HasEdge predicate on the "emails" edge.
func HasEmails() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmailsTable, EmailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmailsWith applies the HasEdge predicate on the "emails" edge with a given conditions (other predicates).
func HasEmailsWith(preds ...predicate.Email) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmailsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmailsTable, EmailsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWebsites applies the HasEdge predicate on the "websites" edge.
func HasWebsites() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebsitesTable, WebsitesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWebsitesWith applies the HasEdge predicate on the "websites" edge with a given conditions (other predicates).
func HasWebsitesWith(preds ...predicate.Website) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WebsitesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebsitesTable, WebsitesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocations applies the HasEdge predicate on the "locations" edge.
func HasLocations() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LocationsTable, LocationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationsWith applies the HasEdge predicate on the "locations" edge with a given conditions (other predicates).
func HasLocationsWith(preds ...predicate.Location) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LocationsTable, LocationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		p(s.Not())
	})
}
