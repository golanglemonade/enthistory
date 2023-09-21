// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory/_examples/basic/ent/character"
	"github.com/flume/enthistory/_examples/basic/ent/friendship"
)

// FriendshipCreate is the builder for creating a Friendship entity.
type FriendshipCreate struct {
	config
	mutation *FriendshipMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FriendshipCreate) SetCreatedAt(t time.Time) *FriendshipCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableCreatedAt(t *time.Time) *FriendshipCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FriendshipCreate) SetUpdatedAt(t time.Time) *FriendshipCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableUpdatedAt(t *time.Time) *FriendshipCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetCharacterID sets the "character_id" field.
func (fc *FriendshipCreate) SetCharacterID(i int) *FriendshipCreate {
	fc.mutation.SetCharacterID(i)
	return fc
}

// SetFriendID sets the "friend_id" field.
func (fc *FriendshipCreate) SetFriendID(i int) *FriendshipCreate {
	fc.mutation.SetFriendID(i)
	return fc
}

// SetID sets the "id" field.
func (fc *FriendshipCreate) SetID(s string) *FriendshipCreate {
	fc.mutation.SetID(s)
	return fc
}

// SetCharacter sets the "character" edge to the Character entity.
func (fc *FriendshipCreate) SetCharacter(c *Character) *FriendshipCreate {
	return fc.SetCharacterID(c.ID)
}

// SetFriend sets the "friend" edge to the Character entity.
func (fc *FriendshipCreate) SetFriend(c *Character) *FriendshipCreate {
	return fc.SetFriendID(c.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fc *FriendshipCreate) Mutation() *FriendshipMutation {
	return fc.mutation
}

// Save creates the Friendship in the database.
func (fc *FriendshipCreate) Save(ctx context.Context) (*Friendship, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FriendshipCreate) SaveX(ctx context.Context) *Friendship {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FriendshipCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FriendshipCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FriendshipCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := friendship.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := friendship.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FriendshipCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Friendship.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Friendship.updated_at"`)}
	}
	if _, ok := fc.mutation.CharacterID(); !ok {
		return &ValidationError{Name: "character_id", err: errors.New(`ent: missing required field "Friendship.character_id"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend_id", err: errors.New(`ent: missing required field "Friendship.friend_id"`)}
	}
	if _, ok := fc.mutation.CharacterID(); !ok {
		return &ValidationError{Name: "character", err: errors.New(`ent: missing required edge "Friendship.character"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend", err: errors.New(`ent: missing required edge "Friendship.friend"`)}
	}
	return nil
}

func (fc *FriendshipCreate) sqlSave(ctx context.Context) (*Friendship, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Friendship.ID type: %T", _spec.ID.Value)
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FriendshipCreate) createSpec() (*Friendship, *sqlgraph.CreateSpec) {
	var (
		_node = &Friendship{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(friendship.Table, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeString))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(friendship.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fc.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.CharacterTable,
			Columns: []string{friendship.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CharacterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FriendID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FriendshipCreateBulk is the builder for creating many Friendship entities in bulk.
type FriendshipCreateBulk struct {
	config
	err      error
	builders []*FriendshipCreate
}

// Save creates the Friendship entities in the database.
func (fcb *FriendshipCreateBulk) Save(ctx context.Context) ([]*Friendship, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Friendship, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) SaveX(ctx context.Context) []*Friendship {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FriendshipCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
