// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Ostap00034/siproject-beercut-backend/genre-service/ent/genre"
)

// GenreCreate is the builder for creating a Genre entity.
type GenreCreate struct {
	config
	mutation *GenreMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GenreCreate) SetName(s string) *GenreCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetDescription sets the "description" field.
func (gc *GenreCreate) SetDescription(s string) *GenreCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gc *GenreCreate) SetNillableDescription(s *string) *GenreCreate {
	if s != nil {
		gc.SetDescription(*s)
	}
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GenreCreate) SetCreatedAt(t time.Time) *GenreCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GenreCreate) SetNillableCreatedAt(t *time.Time) *GenreCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// Mutation returns the GenreMutation object of the builder.
func (gc *GenreCreate) Mutation() *GenreMutation {
	return gc.mutation
}

// Save creates the Genre in the database.
func (gc *GenreCreate) Save(ctx context.Context) (*Genre, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenreCreate) SaveX(ctx context.Context) *Genre {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GenreCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GenreCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GenreCreate) defaults() {
	if _, ok := gc.mutation.Description(); !ok {
		v := genre.DefaultDescription
		gc.mutation.SetDescription(v)
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := genre.DefaultCreatedAt
		gc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GenreCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Genre.name"`)}
	}
	if _, ok := gc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Genre.description"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Genre.created_at"`)}
	}
	return nil
}

func (gc *GenreCreate) sqlSave(ctx context.Context) (*Genre, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GenreCreate) createSpec() (*Genre, *sqlgraph.CreateSpec) {
	var (
		_node = &Genre{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(genre.Table, sqlgraph.NewFieldSpec(genre.FieldID, field.TypeInt))
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(genre.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.SetField(genre.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(genre.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// GenreCreateBulk is the builder for creating many Genre entities in bulk.
type GenreCreateBulk struct {
	config
	err      error
	builders []*GenreCreate
}

// Save creates the Genre entities in the database.
func (gcb *GenreCreateBulk) Save(ctx context.Context) ([]*Genre, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Genre, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenreMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GenreCreateBulk) SaveX(ctx context.Context) []*Genre {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GenreCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GenreCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
