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

package image

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// OriginalURL applies equality check predicate on the "original_url" field. It's identical to OriginalURLEQ.
func OriginalURL(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOriginalURL), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// OriginalURLEQ applies the EQ predicate on the "original_url" field.
func OriginalURLEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLNEQ applies the NEQ predicate on the "original_url" field.
func OriginalURLNEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLIn applies the In predicate on the "original_url" field.
func OriginalURLIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOriginalURL), v...))
	})
}

// OriginalURLNotIn applies the NotIn predicate on the "original_url" field.
func OriginalURLNotIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOriginalURL), v...))
	})
}

// OriginalURLGT applies the GT predicate on the "original_url" field.
func OriginalURLGT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLGTE applies the GTE predicate on the "original_url" field.
func OriginalURLGTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLLT applies the LT predicate on the "original_url" field.
func OriginalURLLT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLLTE applies the LTE predicate on the "original_url" field.
func OriginalURLLTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLContains applies the Contains predicate on the "original_url" field.
func OriginalURLContains(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLHasPrefix applies the HasPrefix predicate on the "original_url" field.
func OriginalURLHasPrefix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLHasSuffix applies the HasSuffix predicate on the "original_url" field.
func OriginalURLHasSuffix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLEqualFold applies the EqualFold predicate on the "original_url" field.
func OriginalURLEqualFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOriginalURL), v))
	})
}

// OriginalURLContainsFold applies the ContainsFold predicate on the "original_url" field.
func OriginalURLContainsFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOriginalURL), v))
	})
}

// HasGalleryCompany applies the HasEdge predicate on the "gallery_company" edge.
func HasGalleryCompany() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GalleryCompanyTable, GalleryCompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGalleryCompanyWith applies the HasEdge predicate on the "gallery_company" edge with a given conditions (other predicates).
func HasGalleryCompanyWith(preds ...predicate.Company) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GalleryCompanyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GalleryCompanyTable, GalleryCompanyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLogoCompany applies the HasEdge predicate on the "logo_company" edge.
func HasLogoCompany() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, LogoCompanyTable, LogoCompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLogoCompanyWith applies the HasEdge predicate on the "logo_company" edge with a given conditions (other predicates).
func HasLogoCompanyWith(preds ...predicate.Company) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LogoCompanyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, LogoCompanyTable, LogoCompanyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
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
func Not(p predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		p(s.Not())
	})
}
