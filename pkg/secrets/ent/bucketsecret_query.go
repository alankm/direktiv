// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/pkg/secrets/ent/bucketsecret"
	"github.com/vorteil/direktiv/pkg/secrets/ent/predicate"
)

// BucketSecretQuery is the builder for querying BucketSecret entities.
type BucketSecretQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.BucketSecret
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BucketSecretQuery builder.
func (bsq *BucketSecretQuery) Where(ps ...predicate.BucketSecret) *BucketSecretQuery {
	bsq.predicates = append(bsq.predicates, ps...)
	return bsq
}

// Limit adds a limit step to the query.
func (bsq *BucketSecretQuery) Limit(limit int) *BucketSecretQuery {
	bsq.limit = &limit
	return bsq
}

// Offset adds an offset step to the query.
func (bsq *BucketSecretQuery) Offset(offset int) *BucketSecretQuery {
	bsq.offset = &offset
	return bsq
}

// Order adds an order step to the query.
func (bsq *BucketSecretQuery) Order(o ...OrderFunc) *BucketSecretQuery {
	bsq.order = append(bsq.order, o...)
	return bsq
}

// First returns the first BucketSecret entity from the query.
// Returns a *NotFoundError when no BucketSecret was found.
func (bsq *BucketSecretQuery) First(ctx context.Context) (*BucketSecret, error) {
	nodes, err := bsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bucketsecret.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bsq *BucketSecretQuery) FirstX(ctx context.Context) *BucketSecret {
	node, err := bsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BucketSecret ID from the query.
// Returns a *NotFoundError when no BucketSecret ID was found.
func (bsq *BucketSecretQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bucketsecret.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bsq *BucketSecretQuery) FirstIDX(ctx context.Context) int {
	id, err := bsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BucketSecret entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BucketSecret entity is not found.
// Returns a *NotFoundError when no BucketSecret entities are found.
func (bsq *BucketSecretQuery) Only(ctx context.Context) (*BucketSecret, error) {
	nodes, err := bsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bucketsecret.Label}
	default:
		return nil, &NotSingularError{bucketsecret.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bsq *BucketSecretQuery) OnlyX(ctx context.Context) *BucketSecret {
	node, err := bsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BucketSecret ID in the query.
// Returns a *NotSingularError when exactly one BucketSecret ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bsq *BucketSecretQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = &NotSingularError{bucketsecret.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bsq *BucketSecretQuery) OnlyIDX(ctx context.Context) int {
	id, err := bsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BucketSecrets.
func (bsq *BucketSecretQuery) All(ctx context.Context) ([]*BucketSecret, error) {
	if err := bsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bsq *BucketSecretQuery) AllX(ctx context.Context) []*BucketSecret {
	nodes, err := bsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BucketSecret IDs.
func (bsq *BucketSecretQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bsq.Select(bucketsecret.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bsq *BucketSecretQuery) IDsX(ctx context.Context) []int {
	ids, err := bsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bsq *BucketSecretQuery) Count(ctx context.Context) (int, error) {
	if err := bsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bsq *BucketSecretQuery) CountX(ctx context.Context) int {
	count, err := bsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bsq *BucketSecretQuery) Exist(ctx context.Context) (bool, error) {
	if err := bsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bsq *BucketSecretQuery) ExistX(ctx context.Context) bool {
	exist, err := bsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BucketSecretQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bsq *BucketSecretQuery) Clone() *BucketSecretQuery {
	if bsq == nil {
		return nil
	}
	return &BucketSecretQuery{
		config:     bsq.config,
		limit:      bsq.limit,
		offset:     bsq.offset,
		order:      append([]OrderFunc{}, bsq.order...),
		predicates: append([]predicate.BucketSecret{}, bsq.predicates...),
		// clone intermediate query.
		sql:  bsq.sql.Clone(),
		path: bsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Ns string `json:"ns,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BucketSecret.Query().
//		GroupBy(bucketsecret.FieldNs).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bsq *BucketSecretQuery) GroupBy(field string, fields ...string) *BucketSecretGroupBy {
	group := &BucketSecretGroupBy{config: bsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bsq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Ns string `json:"ns,omitempty"`
//	}
//
//	client.BucketSecret.Query().
//		Select(bucketsecret.FieldNs).
//		Scan(ctx, &v)
//
func (bsq *BucketSecretQuery) Select(field string, fields ...string) *BucketSecretSelect {
	bsq.fields = append([]string{field}, fields...)
	return &BucketSecretSelect{BucketSecretQuery: bsq}
}

func (bsq *BucketSecretQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bsq.fields {
		if !bucketsecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bsq.path != nil {
		prev, err := bsq.path(ctx)
		if err != nil {
			return err
		}
		bsq.sql = prev
	}
	return nil
}

func (bsq *BucketSecretQuery) sqlAll(ctx context.Context) ([]*BucketSecret, error) {
	var (
		nodes = []*BucketSecret{}
		_spec = bsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BucketSecret{config: bsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, bsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bsq *BucketSecretQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bsq.querySpec()
	return sqlgraph.CountNodes(ctx, bsq.driver, _spec)
}

func (bsq *BucketSecretQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bsq *BucketSecretQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bucketsecret.Table,
			Columns: bucketsecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bucketsecret.FieldID,
			},
		},
		From:   bsq.sql,
		Unique: true,
	}
	if fields := bsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bucketsecret.FieldID)
		for i := range fields {
			if fields[i] != bucketsecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, bucketsecret.ValidColumn)
			}
		}
	}
	return _spec
}

func (bsq *BucketSecretQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bsq.driver.Dialect())
	t1 := builder.Table(bucketsecret.Table)
	selector := builder.Select(t1.Columns(bucketsecret.Columns...)...).From(t1)
	if bsq.sql != nil {
		selector = bsq.sql
		selector.Select(selector.Columns(bucketsecret.Columns...)...)
	}
	for _, p := range bsq.predicates {
		p(selector)
	}
	for _, p := range bsq.order {
		p(selector, bucketsecret.ValidColumn)
	}
	if offset := bsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BucketSecretGroupBy is the group-by builder for BucketSecret entities.
type BucketSecretGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bsgb *BucketSecretGroupBy) Aggregate(fns ...AggregateFunc) *BucketSecretGroupBy {
	bsgb.fns = append(bsgb.fns, fns...)
	return bsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bsgb *BucketSecretGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bsgb.path(ctx)
	if err != nil {
		return err
	}
	bsgb.sql = query
	return bsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bsgb.fields) > 1 {
		return nil, errors.New("ent: BucketSecretGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) StringsX(ctx context.Context) []string {
	v, err := bsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) StringX(ctx context.Context) string {
	v, err := bsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bsgb.fields) > 1 {
		return nil, errors.New("ent: BucketSecretGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) IntsX(ctx context.Context) []int {
	v, err := bsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) IntX(ctx context.Context) int {
	v, err := bsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bsgb.fields) > 1 {
		return nil, errors.New("ent: BucketSecretGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bsgb.fields) > 1 {
		return nil, errors.New("ent: BucketSecretGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bsgb *BucketSecretGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bsgb *BucketSecretGroupBy) BoolX(ctx context.Context) bool {
	v, err := bsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bsgb *BucketSecretGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bsgb.fields {
		if !bucketsecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bsgb *BucketSecretGroupBy) sqlQuery() *sql.Selector {
	selector := bsgb.sql
	columns := make([]string, 0, len(bsgb.fields)+len(bsgb.fns))
	columns = append(columns, bsgb.fields...)
	for _, fn := range bsgb.fns {
		columns = append(columns, fn(selector, bucketsecret.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(bsgb.fields...)
}

// BucketSecretSelect is the builder for selecting fields of BucketSecret entities.
type BucketSecretSelect struct {
	*BucketSecretQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bss *BucketSecretSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bss.prepareQuery(ctx); err != nil {
		return err
	}
	bss.sql = bss.BucketSecretQuery.sqlQuery()
	return bss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bss *BucketSecretSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bss.fields) > 1 {
		return nil, errors.New("ent: BucketSecretSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bss *BucketSecretSelect) StringsX(ctx context.Context) []string {
	v, err := bss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bss *BucketSecretSelect) StringX(ctx context.Context) string {
	v, err := bss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bss.fields) > 1 {
		return nil, errors.New("ent: BucketSecretSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bss *BucketSecretSelect) IntsX(ctx context.Context) []int {
	v, err := bss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bss *BucketSecretSelect) IntX(ctx context.Context) int {
	v, err := bss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bss.fields) > 1 {
		return nil, errors.New("ent: BucketSecretSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bss *BucketSecretSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bss *BucketSecretSelect) Float64X(ctx context.Context) float64 {
	v, err := bss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bss.fields) > 1 {
		return nil, errors.New("ent: BucketSecretSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bss *BucketSecretSelect) BoolsX(ctx context.Context) []bool {
	v, err := bss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bss *BucketSecretSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bucketsecret.Label}
	default:
		err = fmt.Errorf("ent: BucketSecretSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bss *BucketSecretSelect) BoolX(ctx context.Context) bool {
	v, err := bss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bss *BucketSecretSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bss.sqlQuery().Query()
	if err := bss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bss *BucketSecretSelect) sqlQuery() sql.Querier {
	selector := bss.sql
	selector.Select(selector.Columns(bss.fields...)...)
	return selector
}
