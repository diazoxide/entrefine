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
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// LastUpdate applies equality check predicate on the "last_update" field. It's identical to LastUpdateEQ.
func LastUpdate(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdate), v))
	})
}

// OriginalData applies equality check predicate on the "original_data" field. It's identical to OriginalDataEQ.
func OriginalData(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOriginalData), v))
	})
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnabled), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Warehouse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Warehouse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// LastUpdateEQ applies the EQ predicate on the "last_update" field.
func LastUpdateEQ(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateNEQ applies the NEQ predicate on the "last_update" field.
func LastUpdateNEQ(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateIn applies the In predicate on the "last_update" field.
func LastUpdateIn(vs ...time.Time) predicate.Warehouse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastUpdate), v...))
	})
}

// LastUpdateNotIn applies the NotIn predicate on the "last_update" field.
func LastUpdateNotIn(vs ...time.Time) predicate.Warehouse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastUpdate), v...))
	})
}

// LastUpdateGT applies the GT predicate on the "last_update" field.
func LastUpdateGT(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateGTE applies the GTE predicate on the "last_update" field.
func LastUpdateGTE(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateLT applies the LT predicate on the "last_update" field.
func LastUpdateLT(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateLTE applies the LTE predicate on the "last_update" field.
func LastUpdateLTE(v time.Time) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastUpdate), v))
	})
}

// LastUpdateIsNil applies the IsNil predicate on the "last_update" field.
func LastUpdateIsNil() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastUpdate)))
	})
}

// LastUpdateNotNil applies the NotNil predicate on the "last_update" field.
func LastUpdateNotNil() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastUpdate)))
	})
}

// OriginalDataEQ applies the EQ predicate on the "original_data" field.
func OriginalDataEQ(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOriginalData), v))
	})
}

// OriginalDataNEQ applies the NEQ predicate on the "original_data" field.
func OriginalDataNEQ(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOriginalData), v))
	})
}

// OriginalDataIn applies the In predicate on the "original_data" field.
func OriginalDataIn(vs ...string) predicate.Warehouse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOriginalData), v...))
	})
}

// OriginalDataNotIn applies the NotIn predicate on the "original_data" field.
func OriginalDataNotIn(vs ...string) predicate.Warehouse {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOriginalData), v...))
	})
}

// OriginalDataGT applies the GT predicate on the "original_data" field.
func OriginalDataGT(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOriginalData), v))
	})
}

// OriginalDataGTE applies the GTE predicate on the "original_data" field.
func OriginalDataGTE(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOriginalData), v))
	})
}

// OriginalDataLT applies the LT predicate on the "original_data" field.
func OriginalDataLT(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOriginalData), v))
	})
}

// OriginalDataLTE applies the LTE predicate on the "original_data" field.
func OriginalDataLTE(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOriginalData), v))
	})
}

// OriginalDataContains applies the Contains predicate on the "original_data" field.
func OriginalDataContains(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOriginalData), v))
	})
}

// OriginalDataHasPrefix applies the HasPrefix predicate on the "original_data" field.
func OriginalDataHasPrefix(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOriginalData), v))
	})
}

// OriginalDataHasSuffix applies the HasSuffix predicate on the "original_data" field.
func OriginalDataHasSuffix(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOriginalData), v))
	})
}

// OriginalDataIsNil applies the IsNil predicate on the "original_data" field.
func OriginalDataIsNil() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOriginalData)))
	})
}

// OriginalDataNotNil applies the NotNil predicate on the "original_data" field.
func OriginalDataNotNil() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOriginalData)))
	})
}

// OriginalDataEqualFold applies the EqualFold predicate on the "original_data" field.
func OriginalDataEqualFold(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOriginalData), v))
	})
}

// OriginalDataContainsFold applies the ContainsFold predicate on the "original_data" field.
func OriginalDataContainsFold(v string) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOriginalData), v))
	})
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnabled), v))
	})
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnabled), v))
	})
}

// FiltersIsNil applies the IsNil predicate on the "filters" field.
func FiltersIsNil() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFilters)))
	})
}

// FiltersNotNil applies the NotNil predicate on the "filters" field.
func FiltersNotNil() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFilters)))
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVendor applies the HasEdge predicate on the "vendor" edge.
func HasVendor() predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VendorTable, VendorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVendorWith applies the HasEdge predicate on the "vendor" edge with a given conditions (other predicates).
func HasVendorWith(preds ...predicate.Vendor) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VendorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VendorTable, VendorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Warehouse) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Warehouse) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
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
func Not(p predicate.Warehouse) predicate.Warehouse {
	return predicate.Warehouse(func(s *sql.Selector) {
		p(s.Not())
	})
}