// Code generated by entc, DO NOT EDIT.

package workflowevents

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/vorteil/direktiv/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
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
func IDGT(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Signature applies equality check predicate on the "signature" field. It's identical to SignatureEQ.
func Signature(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// SignatureEQ applies the EQ predicate on the "signature" field.
func SignatureEQ(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// SignatureNEQ applies the NEQ predicate on the "signature" field.
func SignatureNEQ(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSignature), v))
	})
}

// SignatureIn applies the In predicate on the "signature" field.
func SignatureIn(vs ...[]byte) predicate.WorkflowEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSignature), v...))
	})
}

// SignatureNotIn applies the NotIn predicate on the "signature" field.
func SignatureNotIn(vs ...[]byte) predicate.WorkflowEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSignature), v...))
	})
}

// SignatureGT applies the GT predicate on the "signature" field.
func SignatureGT(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSignature), v))
	})
}

// SignatureGTE applies the GTE predicate on the "signature" field.
func SignatureGTE(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSignature), v))
	})
}

// SignatureLT applies the LT predicate on the "signature" field.
func SignatureLT(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSignature), v))
	})
}

// SignatureLTE applies the LTE predicate on the "signature" field.
func SignatureLTE(v []byte) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSignature), v))
	})
}

// SignatureIsNil applies the IsNil predicate on the "signature" field.
func SignatureIsNil() predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSignature)))
	})
}

// SignatureNotNil applies the NotNil predicate on the "signature" field.
func SignatureNotNil() predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSignature)))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.WorkflowEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.WorkflowEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWfeventswait applies the HasEdge predicate on the "wfeventswait" edge.
func HasWfeventswait() predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventswaitTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventswaitTable, WfeventswaitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWfeventswaitWith applies the HasEdge predicate on the "wfeventswait" edge with a given conditions (other predicates).
func HasWfeventswaitWith(preds ...predicate.WorkflowEventsWait) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventswaitInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventswaitTable, WfeventswaitColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkflowEvents) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkflowEvents) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
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
func Not(p predicate.WorkflowEvents) predicate.WorkflowEvents {
	return predicate.WorkflowEvents(func(s *sql.Selector) {
		p(s.Not())
	})
}
