// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/ent/exhibition"
	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeExhibition = "Exhibition"
)

// ExhibitionMutation represents an operation that mutates the Exhibition nodes in the graph.
type ExhibitionMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	name               *string
	description        *string
	pictures_ids       *[]string
	appendpictures_ids []string
	status             *exhibition.Status
	created_at         *time.Time
	clearedFields      map[string]struct{}
	done               bool
	oldValue           func(context.Context) (*Exhibition, error)
	predicates         []predicate.Exhibition
}

var _ ent.Mutation = (*ExhibitionMutation)(nil)

// exhibitionOption allows management of the mutation configuration using functional options.
type exhibitionOption func(*ExhibitionMutation)

// newExhibitionMutation creates new mutation for the Exhibition entity.
func newExhibitionMutation(c config, op Op, opts ...exhibitionOption) *ExhibitionMutation {
	m := &ExhibitionMutation{
		config:        c,
		op:            op,
		typ:           TypeExhibition,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withExhibitionID sets the ID field of the mutation.
func withExhibitionID(id int) exhibitionOption {
	return func(m *ExhibitionMutation) {
		var (
			err   error
			once  sync.Once
			value *Exhibition
		)
		m.oldValue = func(ctx context.Context) (*Exhibition, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Exhibition.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withExhibition sets the old Exhibition of the mutation.
func withExhibition(node *Exhibition) exhibitionOption {
	return func(m *ExhibitionMutation) {
		m.oldValue = func(context.Context) (*Exhibition, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ExhibitionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ExhibitionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ExhibitionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ExhibitionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Exhibition.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *ExhibitionMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ExhibitionMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Exhibition entity.
// If the Exhibition object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExhibitionMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ExhibitionMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *ExhibitionMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *ExhibitionMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Exhibition entity.
// If the Exhibition object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExhibitionMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *ExhibitionMutation) ResetDescription() {
	m.description = nil
}

// SetPicturesIds sets the "pictures_ids" field.
func (m *ExhibitionMutation) SetPicturesIds(s []string) {
	m.pictures_ids = &s
	m.appendpictures_ids = nil
}

// PicturesIds returns the value of the "pictures_ids" field in the mutation.
func (m *ExhibitionMutation) PicturesIds() (r []string, exists bool) {
	v := m.pictures_ids
	if v == nil {
		return
	}
	return *v, true
}

// OldPicturesIds returns the old "pictures_ids" field's value of the Exhibition entity.
// If the Exhibition object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExhibitionMutation) OldPicturesIds(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPicturesIds is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPicturesIds requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPicturesIds: %w", err)
	}
	return oldValue.PicturesIds, nil
}

// AppendPicturesIds adds s to the "pictures_ids" field.
func (m *ExhibitionMutation) AppendPicturesIds(s []string) {
	m.appendpictures_ids = append(m.appendpictures_ids, s...)
}

// AppendedPicturesIds returns the list of values that were appended to the "pictures_ids" field in this mutation.
func (m *ExhibitionMutation) AppendedPicturesIds() ([]string, bool) {
	if len(m.appendpictures_ids) == 0 {
		return nil, false
	}
	return m.appendpictures_ids, true
}

// ClearPicturesIds clears the value of the "pictures_ids" field.
func (m *ExhibitionMutation) ClearPicturesIds() {
	m.pictures_ids = nil
	m.appendpictures_ids = nil
	m.clearedFields[exhibition.FieldPicturesIds] = struct{}{}
}

// PicturesIdsCleared returns if the "pictures_ids" field was cleared in this mutation.
func (m *ExhibitionMutation) PicturesIdsCleared() bool {
	_, ok := m.clearedFields[exhibition.FieldPicturesIds]
	return ok
}

// ResetPicturesIds resets all changes to the "pictures_ids" field.
func (m *ExhibitionMutation) ResetPicturesIds() {
	m.pictures_ids = nil
	m.appendpictures_ids = nil
	delete(m.clearedFields, exhibition.FieldPicturesIds)
}

// SetStatus sets the "status" field.
func (m *ExhibitionMutation) SetStatus(e exhibition.Status) {
	m.status = &e
}

// Status returns the value of the "status" field in the mutation.
func (m *ExhibitionMutation) Status() (r exhibition.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Exhibition entity.
// If the Exhibition object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExhibitionMutation) OldStatus(ctx context.Context) (v exhibition.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *ExhibitionMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ExhibitionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ExhibitionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Exhibition entity.
// If the Exhibition object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExhibitionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ExhibitionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the ExhibitionMutation builder.
func (m *ExhibitionMutation) Where(ps ...predicate.Exhibition) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ExhibitionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ExhibitionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Exhibition, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ExhibitionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ExhibitionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Exhibition).
func (m *ExhibitionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ExhibitionMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, exhibition.FieldName)
	}
	if m.description != nil {
		fields = append(fields, exhibition.FieldDescription)
	}
	if m.pictures_ids != nil {
		fields = append(fields, exhibition.FieldPicturesIds)
	}
	if m.status != nil {
		fields = append(fields, exhibition.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, exhibition.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ExhibitionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case exhibition.FieldName:
		return m.Name()
	case exhibition.FieldDescription:
		return m.Description()
	case exhibition.FieldPicturesIds:
		return m.PicturesIds()
	case exhibition.FieldStatus:
		return m.Status()
	case exhibition.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ExhibitionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case exhibition.FieldName:
		return m.OldName(ctx)
	case exhibition.FieldDescription:
		return m.OldDescription(ctx)
	case exhibition.FieldPicturesIds:
		return m.OldPicturesIds(ctx)
	case exhibition.FieldStatus:
		return m.OldStatus(ctx)
	case exhibition.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Exhibition field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ExhibitionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case exhibition.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case exhibition.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case exhibition.FieldPicturesIds:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPicturesIds(v)
		return nil
	case exhibition.FieldStatus:
		v, ok := value.(exhibition.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case exhibition.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Exhibition field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ExhibitionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ExhibitionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ExhibitionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Exhibition numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ExhibitionMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(exhibition.FieldPicturesIds) {
		fields = append(fields, exhibition.FieldPicturesIds)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ExhibitionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ExhibitionMutation) ClearField(name string) error {
	switch name {
	case exhibition.FieldPicturesIds:
		m.ClearPicturesIds()
		return nil
	}
	return fmt.Errorf("unknown Exhibition nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ExhibitionMutation) ResetField(name string) error {
	switch name {
	case exhibition.FieldName:
		m.ResetName()
		return nil
	case exhibition.FieldDescription:
		m.ResetDescription()
		return nil
	case exhibition.FieldPicturesIds:
		m.ResetPicturesIds()
		return nil
	case exhibition.FieldStatus:
		m.ResetStatus()
		return nil
	case exhibition.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Exhibition field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ExhibitionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ExhibitionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ExhibitionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ExhibitionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ExhibitionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ExhibitionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ExhibitionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Exhibition unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ExhibitionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Exhibition edge %s", name)
}
