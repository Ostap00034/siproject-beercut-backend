// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent/token"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (tc *TokenCreate) SetToken(s string) *TokenCreate {
	tc.mutation.SetToken(s)
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TokenCreate) SetUserID(s string) *TokenCreate {
	tc.mutation.SetUserID(s)
	return tc
}

// SetExpiresAt sets the "expires_at" field.
func (tc *TokenCreate) SetExpiresAt(t time.Time) *TokenCreate {
	tc.mutation.SetExpiresAt(t)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TokenCreate) SetCreatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := token.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "Token.token"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Token.user_id"`)}
	}
	if _, ok := tc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "Token.expires_at"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Token.created_at"`)}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(token.Table, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Token(); ok {
		_spec.SetField(token.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := tc.mutation.UserID(); ok {
		_spec.SetField(token.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := tc.mutation.ExpiresAt(); ok {
		_spec.SetField(token.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(token.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	err      error
	builders []*TokenCreate
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
