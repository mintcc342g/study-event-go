// Code generated by entc, DO NOT EDIT.

package garden

import (
	"study_event_go/ent/predicate"
	"study_event_go/types"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id types.GardenID) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// MentorshipSystemID applies equality check predicate on the "mentorship_system_id" field. It's identical to MentorshipSystemIDEQ.
func MentorshipSystemID(v types.MentorshipSystemID) predicate.Garden {
	vc := uint64(v)
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMentorshipSystemID), vc))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Garden {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Garden {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocation), v))
	})
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Garden {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocation), v...))
	})
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Garden {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocation), v...))
	})
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocation), v))
	})
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocation), v))
	})
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocation), v))
	})
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocation), v))
	})
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocation), v))
	})
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocation), v))
	})
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocation), v))
	})
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocation), v))
	})
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocation), v))
	})
}

// MentorshipSystemIDEQ applies the EQ predicate on the "mentorship_system_id" field.
func MentorshipSystemIDEQ(v types.MentorshipSystemID) predicate.Garden {
	vc := uint64(v)
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMentorshipSystemID), vc))
	})
}

// MentorshipSystemIDNEQ applies the NEQ predicate on the "mentorship_system_id" field.
func MentorshipSystemIDNEQ(v types.MentorshipSystemID) predicate.Garden {
	vc := uint64(v)
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMentorshipSystemID), vc))
	})
}

// MentorshipSystemIDIn applies the In predicate on the "mentorship_system_id" field.
func MentorshipSystemIDIn(vs ...types.MentorshipSystemID) predicate.Garden {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMentorshipSystemID), v...))
	})
}

// MentorshipSystemIDNotIn applies the NotIn predicate on the "mentorship_system_id" field.
func MentorshipSystemIDNotIn(vs ...types.MentorshipSystemID) predicate.Garden {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Garden(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMentorshipSystemID), v...))
	})
}

// MentorshipSystemIDIsNil applies the IsNil predicate on the "mentorship_system_id" field.
func MentorshipSystemIDIsNil() predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMentorshipSystemID)))
	})
}

// MentorshipSystemIDNotNil applies the NotNil predicate on the "mentorship_system_id" field.
func MentorshipSystemIDNotNil() predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMentorshipSystemID)))
	})
}

// HasMentorshipSystem applies the HasEdge predicate on the "mentorship_system" edge.
func HasMentorshipSystem() predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MentorshipSystemTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MentorshipSystemTable, MentorshipSystemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMentorshipSystemWith applies the HasEdge predicate on the "mentorship_system" edge with a given conditions (other predicates).
func HasMentorshipSystemWith(preds ...predicate.MentorshipSystem) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MentorshipSystemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MentorshipSystemTable, MentorshipSystemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Garden) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Garden) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
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
func Not(p predicate.Garden) predicate.Garden {
	return predicate.Garden(func(s *sql.Selector) {
		p(s.Not())
	})
}
