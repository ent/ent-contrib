// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithFieldsDelete is the builder for deleting a WithFields entity.
type WithFieldsDelete struct {
	config
	hooks    []Hook
	mutation *WithFieldsMutation
}

// Where adds a new predicate to the WithFieldsDelete builder.
func (wfd *WithFieldsDelete) Where(ps ...predicate.WithFields) *WithFieldsDelete {
	wfd.mutation.predicates = append(wfd.mutation.predicates, ps...)
	return wfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wfd *WithFieldsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wfd.hooks) == 0 {
		affected, err = wfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wfd.mutation = mutation
			affected, err = wfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wfd.hooks) - 1; i >= 0; i-- {
			mut = wfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfd *WithFieldsDelete) ExecX(ctx context.Context) int {
	n, err := wfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wfd *WithFieldsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: withfields.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withfields.FieldID,
			},
		},
	}
	if ps := wfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wfd.driver, _spec)
}

// WithFieldsDeleteOne is the builder for deleting a single WithFields entity.
type WithFieldsDeleteOne struct {
	wfd *WithFieldsDelete
}

// Exec executes the deletion query.
func (wfdo *WithFieldsDeleteOne) Exec(ctx context.Context) error {
	n, err := wfdo.wfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withfields.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wfdo *WithFieldsDeleteOne) ExecX(ctx context.Context) {
	wfdo.wfd.ExecX(ctx)
}
