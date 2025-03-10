// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"primijenjena-informatika-dev/ent/codespace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CodespaceCreate is the builder for creating a Codespace entity.
type CodespaceCreate struct {
	config
	mutation *CodespaceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CodespaceCreate) SetCreatedAt(t time.Time) *CodespaceCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CodespaceCreate) SetNillableCreatedAt(t *time.Time) *CodespaceCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CodespaceCreate) SetUpdatedAt(t time.Time) *CodespaceCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CodespaceCreate) SetNillableUpdatedAt(t *time.Time) *CodespaceCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetMachineType sets the "machine_type" field.
func (cc *CodespaceCreate) SetMachineType(s string) *CodespaceCreate {
	cc.mutation.SetMachineType(s)
	return cc
}

// SetClientIP sets the "client_ip" field.
func (cc *CodespaceCreate) SetClientIP(s string) *CodespaceCreate {
	cc.mutation.SetClientIP(s)
	return cc
}

// SetGithubUserID sets the "github_user_id" field.
func (cc *CodespaceCreate) SetGithubUserID(s string) *CodespaceCreate {
	cc.mutation.SetGithubUserID(s)
	return cc
}

// SetGithubCodespaceID sets the "github_codespace_id" field.
func (cc *CodespaceCreate) SetGithubCodespaceID(s string) *CodespaceCreate {
	cc.mutation.SetGithubCodespaceID(s)
	return cc
}

// SetGithubCodespaceURL sets the "github_codespace_url" field.
func (cc *CodespaceCreate) SetGithubCodespaceURL(s string) *CodespaceCreate {
	cc.mutation.SetGithubCodespaceURL(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CodespaceCreate) SetID(u uuid.UUID) *CodespaceCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CodespaceCreate) SetNillableID(u *uuid.UUID) *CodespaceCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CodespaceMutation object of the builder.
func (cc *CodespaceCreate) Mutation() *CodespaceMutation {
	return cc.mutation
}

// Save creates the Codespace in the database.
func (cc *CodespaceCreate) Save(ctx context.Context) (*Codespace, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CodespaceCreate) SaveX(ctx context.Context) *Codespace {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CodespaceCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CodespaceCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CodespaceCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := codespace.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := codespace.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := codespace.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CodespaceCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Codespace.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Codespace.updated_at"`)}
	}
	if _, ok := cc.mutation.MachineType(); !ok {
		return &ValidationError{Name: "machine_type", err: errors.New(`ent: missing required field "Codespace.machine_type"`)}
	}
	if v, ok := cc.mutation.MachineType(); ok {
		if err := codespace.MachineTypeValidator(v); err != nil {
			return &ValidationError{Name: "machine_type", err: fmt.Errorf(`ent: validator failed for field "Codespace.machine_type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ClientIP(); !ok {
		return &ValidationError{Name: "client_ip", err: errors.New(`ent: missing required field "Codespace.client_ip"`)}
	}
	if v, ok := cc.mutation.ClientIP(); ok {
		if err := codespace.ClientIPValidator(v); err != nil {
			return &ValidationError{Name: "client_ip", err: fmt.Errorf(`ent: validator failed for field "Codespace.client_ip": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GithubUserID(); !ok {
		return &ValidationError{Name: "github_user_id", err: errors.New(`ent: missing required field "Codespace.github_user_id"`)}
	}
	if v, ok := cc.mutation.GithubUserID(); ok {
		if err := codespace.GithubUserIDValidator(v); err != nil {
			return &ValidationError{Name: "github_user_id", err: fmt.Errorf(`ent: validator failed for field "Codespace.github_user_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GithubCodespaceID(); !ok {
		return &ValidationError{Name: "github_codespace_id", err: errors.New(`ent: missing required field "Codespace.github_codespace_id"`)}
	}
	if v, ok := cc.mutation.GithubCodespaceID(); ok {
		if err := codespace.GithubCodespaceIDValidator(v); err != nil {
			return &ValidationError{Name: "github_codespace_id", err: fmt.Errorf(`ent: validator failed for field "Codespace.github_codespace_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GithubCodespaceURL(); !ok {
		return &ValidationError{Name: "github_codespace_url", err: errors.New(`ent: missing required field "Codespace.github_codespace_url"`)}
	}
	if v, ok := cc.mutation.GithubCodespaceURL(); ok {
		if err := codespace.GithubCodespaceURLValidator(v); err != nil {
			return &ValidationError{Name: "github_codespace_url", err: fmt.Errorf(`ent: validator failed for field "Codespace.github_codespace_url": %w`, err)}
		}
	}
	return nil
}

func (cc *CodespaceCreate) sqlSave(ctx context.Context) (*Codespace, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CodespaceCreate) createSpec() (*Codespace, *sqlgraph.CreateSpec) {
	var (
		_node = &Codespace{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(codespace.Table, sqlgraph.NewFieldSpec(codespace.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(codespace.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(codespace.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.MachineType(); ok {
		_spec.SetField(codespace.FieldMachineType, field.TypeString, value)
		_node.MachineType = value
	}
	if value, ok := cc.mutation.ClientIP(); ok {
		_spec.SetField(codespace.FieldClientIP, field.TypeString, value)
		_node.ClientIP = value
	}
	if value, ok := cc.mutation.GithubUserID(); ok {
		_spec.SetField(codespace.FieldGithubUserID, field.TypeString, value)
		_node.GithubUserID = value
	}
	if value, ok := cc.mutation.GithubCodespaceID(); ok {
		_spec.SetField(codespace.FieldGithubCodespaceID, field.TypeString, value)
		_node.GithubCodespaceID = value
	}
	if value, ok := cc.mutation.GithubCodespaceURL(); ok {
		_spec.SetField(codespace.FieldGithubCodespaceURL, field.TypeString, value)
		_node.GithubCodespaceURL = value
	}
	return _node, _spec
}

// CodespaceCreateBulk is the builder for creating many Codespace entities in bulk.
type CodespaceCreateBulk struct {
	config
	err      error
	builders []*CodespaceCreate
}

// Save creates the Codespace entities in the database.
func (ccb *CodespaceCreateBulk) Save(ctx context.Context) ([]*Codespace, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Codespace, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CodespaceMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CodespaceCreateBulk) SaveX(ctx context.Context) []*Codespace {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CodespaceCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CodespaceCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
