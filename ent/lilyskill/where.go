// Code generated by entc, DO NOT EDIT.

package lilyskill

import (
	"study-event-go/ent/predicate"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
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
func IDGT(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LilyID applies equality check predicate on the "lily_id" field. It's identical to LilyIDEQ.
func LilyID(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLilyID), vc))
	})
}

// SkillID applies equality check predicate on the "skill_id" field. It's identical to SkillIDEQ.
func SkillID(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkillID), vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LilyIDEQ applies the EQ predicate on the "lily_id" field.
func LilyIDEQ(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLilyID), vc))
	})
}

// LilyIDNEQ applies the NEQ predicate on the "lily_id" field.
func LilyIDNEQ(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLilyID), vc))
	})
}

// LilyIDIn applies the In predicate on the "lily_id" field.
func LilyIDIn(vs ...types.LilyID) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLilyID), v...))
	})
}

// LilyIDNotIn applies the NotIn predicate on the "lily_id" field.
func LilyIDNotIn(vs ...types.LilyID) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLilyID), v...))
	})
}

// LilyIDGT applies the GT predicate on the "lily_id" field.
func LilyIDGT(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLilyID), vc))
	})
}

// LilyIDGTE applies the GTE predicate on the "lily_id" field.
func LilyIDGTE(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLilyID), vc))
	})
}

// LilyIDLT applies the LT predicate on the "lily_id" field.
func LilyIDLT(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLilyID), vc))
	})
}

// LilyIDLTE applies the LTE predicate on the "lily_id" field.
func LilyIDLTE(v types.LilyID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLilyID), vc))
	})
}

// SkillIDEQ applies the EQ predicate on the "skill_id" field.
func SkillIDEQ(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkillID), vc))
	})
}

// SkillIDNEQ applies the NEQ predicate on the "skill_id" field.
func SkillIDNEQ(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkillID), vc))
	})
}

// SkillIDIn applies the In predicate on the "skill_id" field.
func SkillIDIn(vs ...types.SkillID) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkillID), v...))
	})
}

// SkillIDNotIn applies the NotIn predicate on the "skill_id" field.
func SkillIDNotIn(vs ...types.SkillID) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkillID), v...))
	})
}

// SkillIDGT applies the GT predicate on the "skill_id" field.
func SkillIDGT(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkillID), vc))
	})
}

// SkillIDGTE applies the GTE predicate on the "skill_id" field.
func SkillIDGTE(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkillID), vc))
	})
}

// SkillIDLT applies the LT predicate on the "skill_id" field.
func SkillIDLT(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkillID), vc))
	})
}

// SkillIDLTE applies the LTE predicate on the "skill_id" field.
func SkillIDLTE(v types.SkillID) predicate.LilySkill {
	vc := uint64(v)
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkillID), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LilySkill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LilySkill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LilySkill) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LilySkill) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
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
func Not(p predicate.LilySkill) predicate.LilySkill {
	return predicate.LilySkill(func(s *sql.Selector) {
		p(s.Not())
	})
}