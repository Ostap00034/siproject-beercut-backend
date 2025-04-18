// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/ent/movementhistory"
	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/ent/predicate"
)

// MovementHistoryQuery is the builder for querying MovementHistory entities.
type MovementHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []movementhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.MovementHistory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MovementHistoryQuery builder.
func (mhq *MovementHistoryQuery) Where(ps ...predicate.MovementHistory) *MovementHistoryQuery {
	mhq.predicates = append(mhq.predicates, ps...)
	return mhq
}

// Limit the number of records to be returned by this query.
func (mhq *MovementHistoryQuery) Limit(limit int) *MovementHistoryQuery {
	mhq.ctx.Limit = &limit
	return mhq
}

// Offset to start from.
func (mhq *MovementHistoryQuery) Offset(offset int) *MovementHistoryQuery {
	mhq.ctx.Offset = &offset
	return mhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mhq *MovementHistoryQuery) Unique(unique bool) *MovementHistoryQuery {
	mhq.ctx.Unique = &unique
	return mhq
}

// Order specifies how the records should be ordered.
func (mhq *MovementHistoryQuery) Order(o ...movementhistory.OrderOption) *MovementHistoryQuery {
	mhq.order = append(mhq.order, o...)
	return mhq
}

// First returns the first MovementHistory entity from the query.
// Returns a *NotFoundError when no MovementHistory was found.
func (mhq *MovementHistoryQuery) First(ctx context.Context) (*MovementHistory, error) {
	nodes, err := mhq.Limit(1).All(setContextOp(ctx, mhq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{movementhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mhq *MovementHistoryQuery) FirstX(ctx context.Context) *MovementHistory {
	node, err := mhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MovementHistory ID from the query.
// Returns a *NotFoundError when no MovementHistory ID was found.
func (mhq *MovementHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mhq.Limit(1).IDs(setContextOp(ctx, mhq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{movementhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mhq *MovementHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := mhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MovementHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MovementHistory entity is found.
// Returns a *NotFoundError when no MovementHistory entities are found.
func (mhq *MovementHistoryQuery) Only(ctx context.Context) (*MovementHistory, error) {
	nodes, err := mhq.Limit(2).All(setContextOp(ctx, mhq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{movementhistory.Label}
	default:
		return nil, &NotSingularError{movementhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mhq *MovementHistoryQuery) OnlyX(ctx context.Context) *MovementHistory {
	node, err := mhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MovementHistory ID in the query.
// Returns a *NotSingularError when more than one MovementHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (mhq *MovementHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mhq.Limit(2).IDs(setContextOp(ctx, mhq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{movementhistory.Label}
	default:
		err = &NotSingularError{movementhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mhq *MovementHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := mhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MovementHistories.
func (mhq *MovementHistoryQuery) All(ctx context.Context) ([]*MovementHistory, error) {
	ctx = setContextOp(ctx, mhq.ctx, ent.OpQueryAll)
	if err := mhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MovementHistory, *MovementHistoryQuery]()
	return withInterceptors[[]*MovementHistory](ctx, mhq, qr, mhq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mhq *MovementHistoryQuery) AllX(ctx context.Context) []*MovementHistory {
	nodes, err := mhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MovementHistory IDs.
func (mhq *MovementHistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mhq.ctx.Unique == nil && mhq.path != nil {
		mhq.Unique(true)
	}
	ctx = setContextOp(ctx, mhq.ctx, ent.OpQueryIDs)
	if err = mhq.Select(movementhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mhq *MovementHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := mhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mhq *MovementHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mhq.ctx, ent.OpQueryCount)
	if err := mhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mhq, querierCount[*MovementHistoryQuery](), mhq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mhq *MovementHistoryQuery) CountX(ctx context.Context) int {
	count, err := mhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mhq *MovementHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mhq.ctx, ent.OpQueryExist)
	switch _, err := mhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mhq *MovementHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := mhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MovementHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mhq *MovementHistoryQuery) Clone() *MovementHistoryQuery {
	if mhq == nil {
		return nil
	}
	return &MovementHistoryQuery{
		config:     mhq.config,
		ctx:        mhq.ctx.Clone(),
		order:      append([]movementhistory.OrderOption{}, mhq.order...),
		inters:     append([]Interceptor{}, mhq.inters...),
		predicates: append([]predicate.MovementHistory{}, mhq.predicates...),
		// clone intermediate query.
		sql:  mhq.sql.Clone(),
		path: mhq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PictureID string `json:"picture_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MovementHistory.Query().
//		GroupBy(movementhistory.FieldPictureID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mhq *MovementHistoryQuery) GroupBy(field string, fields ...string) *MovementHistoryGroupBy {
	mhq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MovementHistoryGroupBy{build: mhq}
	grbuild.flds = &mhq.ctx.Fields
	grbuild.label = movementhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PictureID string `json:"picture_id,omitempty"`
//	}
//
//	client.MovementHistory.Query().
//		Select(movementhistory.FieldPictureID).
//		Scan(ctx, &v)
func (mhq *MovementHistoryQuery) Select(fields ...string) *MovementHistorySelect {
	mhq.ctx.Fields = append(mhq.ctx.Fields, fields...)
	sbuild := &MovementHistorySelect{MovementHistoryQuery: mhq}
	sbuild.label = movementhistory.Label
	sbuild.flds, sbuild.scan = &mhq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MovementHistorySelect configured with the given aggregations.
func (mhq *MovementHistoryQuery) Aggregate(fns ...AggregateFunc) *MovementHistorySelect {
	return mhq.Select().Aggregate(fns...)
}

func (mhq *MovementHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mhq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mhq); err != nil {
				return err
			}
		}
	}
	for _, f := range mhq.ctx.Fields {
		if !movementhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mhq.path != nil {
		prev, err := mhq.path(ctx)
		if err != nil {
			return err
		}
		mhq.sql = prev
	}
	return nil
}

func (mhq *MovementHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MovementHistory, error) {
	var (
		nodes = []*MovementHistory{}
		_spec = mhq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MovementHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MovementHistory{config: mhq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mhq *MovementHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mhq.querySpec()
	_spec.Node.Columns = mhq.ctx.Fields
	if len(mhq.ctx.Fields) > 0 {
		_spec.Unique = mhq.ctx.Unique != nil && *mhq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mhq.driver, _spec)
}

func (mhq *MovementHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(movementhistory.Table, movementhistory.Columns, sqlgraph.NewFieldSpec(movementhistory.FieldID, field.TypeInt))
	_spec.From = mhq.sql
	if unique := mhq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mhq.path != nil {
		_spec.Unique = true
	}
	if fields := mhq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, movementhistory.FieldID)
		for i := range fields {
			if fields[i] != movementhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mhq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mhq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mhq *MovementHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mhq.driver.Dialect())
	t1 := builder.Table(movementhistory.Table)
	columns := mhq.ctx.Fields
	if len(columns) == 0 {
		columns = movementhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mhq.sql != nil {
		selector = mhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mhq.ctx.Unique != nil && *mhq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mhq.predicates {
		p(selector)
	}
	for _, p := range mhq.order {
		p(selector)
	}
	if offset := mhq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mhq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MovementHistoryGroupBy is the group-by builder for MovementHistory entities.
type MovementHistoryGroupBy struct {
	selector
	build *MovementHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mhgb *MovementHistoryGroupBy) Aggregate(fns ...AggregateFunc) *MovementHistoryGroupBy {
	mhgb.fns = append(mhgb.fns, fns...)
	return mhgb
}

// Scan applies the selector query and scans the result into the given value.
func (mhgb *MovementHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mhgb.build.ctx, ent.OpQueryGroupBy)
	if err := mhgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MovementHistoryQuery, *MovementHistoryGroupBy](ctx, mhgb.build, mhgb, mhgb.build.inters, v)
}

func (mhgb *MovementHistoryGroupBy) sqlScan(ctx context.Context, root *MovementHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mhgb.fns))
	for _, fn := range mhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mhgb.flds)+len(mhgb.fns))
		for _, f := range *mhgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mhgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mhgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MovementHistorySelect is the builder for selecting fields of MovementHistory entities.
type MovementHistorySelect struct {
	*MovementHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mhs *MovementHistorySelect) Aggregate(fns ...AggregateFunc) *MovementHistorySelect {
	mhs.fns = append(mhs.fns, fns...)
	return mhs
}

// Scan applies the selector query and scans the result into the given value.
func (mhs *MovementHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mhs.ctx, ent.OpQuerySelect)
	if err := mhs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MovementHistoryQuery, *MovementHistorySelect](ctx, mhs.MovementHistoryQuery, mhs, mhs.inters, v)
}

func (mhs *MovementHistorySelect) sqlScan(ctx context.Context, root *MovementHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mhs.fns))
	for _, fn := range mhs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
