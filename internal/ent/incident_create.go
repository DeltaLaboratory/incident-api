// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DeltaLaboratory/incident-api/internal/ent/incident"
	"github.com/DeltaLaboratory/incident-api/internal/ent/schema"
	"github.com/google/uuid"
)

// IncidentCreate is the builder for creating a Incident entity.
type IncidentCreate struct {
	config
	mutation *IncidentMutation
	hooks    []Hook
}

// SetIdempotencyKey sets the "idempotency_key" field.
func (ic *IncidentCreate) SetIdempotencyKey(s string) *IncidentCreate {
	ic.mutation.SetIdempotencyKey(s)
	return ic
}

// SetReporter sets the "reporter" field.
func (ic *IncidentCreate) SetReporter(u uuid.UUID) *IncidentCreate {
	ic.mutation.SetReporter(u)
	return ic
}

// SetLocation sets the "location" field.
func (ic *IncidentCreate) SetLocation(sj *schema.GeoJson) *IncidentCreate {
	ic.mutation.SetLocation(sj)
	return ic
}

// SetType sets the "type" field.
func (ic *IncidentCreate) SetType(s string) *IncidentCreate {
	ic.mutation.SetType(s)
	return ic
}

// SetDescription sets the "description" field.
func (ic *IncidentCreate) SetDescription(s string) *IncidentCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetImage sets the "image" field.
func (ic *IncidentCreate) SetImage(b []byte) *IncidentCreate {
	ic.mutation.SetImage(b)
	return ic
}

// SetVote sets the "vote" field.
func (ic *IncidentCreate) SetVote(i int) *IncidentCreate {
	ic.mutation.SetVote(i)
	return ic
}

// SetNillableVote sets the "vote" field if the given value is not nil.
func (ic *IncidentCreate) SetNillableVote(i *int) *IncidentCreate {
	if i != nil {
		ic.SetVote(*i)
	}
	return ic
}

// SetVoteFilter sets the "vote_filter" field.
func (ic *IncidentCreate) SetVoteFilter(b []byte) *IncidentCreate {
	ic.mutation.SetVoteFilter(b)
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *IncidentCreate) SetCreatedAt(t time.Time) *IncidentCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IncidentCreate) SetNillableCreatedAt(t *time.Time) *IncidentCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IncidentCreate) SetID(u uuid.UUID) *IncidentCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *IncidentCreate) SetNillableID(u *uuid.UUID) *IncidentCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// Mutation returns the IncidentMutation object of the builder.
func (ic *IncidentCreate) Mutation() *IncidentMutation {
	return ic.mutation
}

// Save creates the Incident in the database.
func (ic *IncidentCreate) Save(ctx context.Context) (*Incident, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IncidentCreate) SaveX(ctx context.Context) *Incident {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IncidentCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IncidentCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IncidentCreate) defaults() {
	if _, ok := ic.mutation.Vote(); !ok {
		v := incident.DefaultVote
		ic.mutation.SetVote(v)
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := incident.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := incident.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IncidentCreate) check() error {
	if _, ok := ic.mutation.IdempotencyKey(); !ok {
		return &ValidationError{Name: "idempotency_key", err: errors.New(`ent: missing required field "Incident.idempotency_key"`)}
	}
	if _, ok := ic.mutation.Reporter(); !ok {
		return &ValidationError{Name: "reporter", err: errors.New(`ent: missing required field "Incident.reporter"`)}
	}
	if _, ok := ic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Incident.type"`)}
	}
	if _, ok := ic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Incident.description"`)}
	}
	if _, ok := ic.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "Incident.image"`)}
	}
	if _, ok := ic.mutation.Vote(); !ok {
		return &ValidationError{Name: "vote", err: errors.New(`ent: missing required field "Incident.vote"`)}
	}
	if _, ok := ic.mutation.VoteFilter(); !ok {
		return &ValidationError{Name: "vote_filter", err: errors.New(`ent: missing required field "Incident.vote_filter"`)}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Incident.created_at"`)}
	}
	return nil
}

func (ic *IncidentCreate) sqlSave(ctx context.Context) (*Incident, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IncidentCreate) createSpec() (*Incident, *sqlgraph.CreateSpec) {
	var (
		_node = &Incident{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(incident.Table, sqlgraph.NewFieldSpec(incident.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.IdempotencyKey(); ok {
		_spec.SetField(incident.FieldIdempotencyKey, field.TypeString, value)
		_node.IdempotencyKey = value
	}
	if value, ok := ic.mutation.Reporter(); ok {
		_spec.SetField(incident.FieldReporter, field.TypeUUID, value)
		_node.Reporter = value
	}
	if value, ok := ic.mutation.Location(); ok {
		_spec.SetField(incident.FieldLocation, field.TypeOther, value)
		_node.Location = value
	}
	if value, ok := ic.mutation.GetType(); ok {
		_spec.SetField(incident.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(incident.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.Image(); ok {
		_spec.SetField(incident.FieldImage, field.TypeBytes, value)
		_node.Image = &value
	}
	if value, ok := ic.mutation.Vote(); ok {
		_spec.SetField(incident.FieldVote, field.TypeInt, value)
		_node.Vote = value
	}
	if value, ok := ic.mutation.VoteFilter(); ok {
		_spec.SetField(incident.FieldVoteFilter, field.TypeBytes, value)
		_node.VoteFilter = &value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(incident.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// IncidentCreateBulk is the builder for creating many Incident entities in bulk.
type IncidentCreateBulk struct {
	config
	err      error
	builders []*IncidentCreate
}

// Save creates the Incident entities in the database.
func (icb *IncidentCreateBulk) Save(ctx context.Context) ([]*Incident, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Incident, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IncidentMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IncidentCreateBulk) SaveX(ctx context.Context) []*Incident {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IncidentCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IncidentCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
