// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos_learn/internal/data/ent/predicate"
	"kratos_learn/internal/data/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uint32
	delete_time   *time.Time
	uqid          *string
	name          *string
	avatar        *string
	_type         *uint8
	add_type      *int8
	is_enable     *uint8
	addis_enable  *int8
	status        *uint8
	addstatus     *int8
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uint32) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uint32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uint32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uint32, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint32{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDeleteTime sets the "delete_time" field.
func (m *UserMutation) SetDeleteTime(t time.Time) {
	m.delete_time = &t
}

// DeleteTime returns the value of the "delete_time" field in the mutation.
func (m *UserMutation) DeleteTime() (r time.Time, exists bool) {
	v := m.delete_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteTime returns the old "delete_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDeleteTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleteTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleteTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteTime: %w", err)
	}
	return oldValue.DeleteTime, nil
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (m *UserMutation) ClearDeleteTime() {
	m.delete_time = nil
	m.clearedFields[user.FieldDeleteTime] = struct{}{}
}

// DeleteTimeCleared returns if the "delete_time" field was cleared in this mutation.
func (m *UserMutation) DeleteTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldDeleteTime]
	return ok
}

// ResetDeleteTime resets all changes to the "delete_time" field.
func (m *UserMutation) ResetDeleteTime() {
	m.delete_time = nil
	delete(m.clearedFields, user.FieldDeleteTime)
}

// SetUqid sets the "uqid" field.
func (m *UserMutation) SetUqid(s string) {
	m.uqid = &s
}

// Uqid returns the value of the "uqid" field in the mutation.
func (m *UserMutation) Uqid() (r string, exists bool) {
	v := m.uqid
	if v == nil {
		return
	}
	return *v, true
}

// OldUqid returns the old "uqid" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUqid(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUqid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUqid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUqid: %w", err)
	}
	return oldValue.Uqid, nil
}

// ResetUqid resets all changes to the "uqid" field.
func (m *UserMutation) ResetUqid() {
	m.uqid = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetAvatar sets the "avatar" field.
func (m *UserMutation) SetAvatar(s string) {
	m.avatar = &s
}

// Avatar returns the value of the "avatar" field in the mutation.
func (m *UserMutation) Avatar() (r string, exists bool) {
	v := m.avatar
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatar returns the old "avatar" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAvatar(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatar is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatar requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatar: %w", err)
	}
	return oldValue.Avatar, nil
}

// ResetAvatar resets all changes to the "avatar" field.
func (m *UserMutation) ResetAvatar() {
	m.avatar = nil
}

// SetType sets the "type" field.
func (m *UserMutation) SetType(u uint8) {
	m._type = &u
	m.add_type = nil
}

// GetType returns the value of the "type" field in the mutation.
func (m *UserMutation) GetType() (r uint8, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldType(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// AddType adds u to the "type" field.
func (m *UserMutation) AddType(u int8) {
	if m.add_type != nil {
		*m.add_type += u
	} else {
		m.add_type = &u
	}
}

// AddedType returns the value that was added to the "type" field in this mutation.
func (m *UserMutation) AddedType() (r int8, exists bool) {
	v := m.add_type
	if v == nil {
		return
	}
	return *v, true
}

// ResetType resets all changes to the "type" field.
func (m *UserMutation) ResetType() {
	m._type = nil
	m.add_type = nil
}

// SetIsEnable sets the "is_enable" field.
func (m *UserMutation) SetIsEnable(u uint8) {
	m.is_enable = &u
	m.addis_enable = nil
}

// IsEnable returns the value of the "is_enable" field in the mutation.
func (m *UserMutation) IsEnable() (r uint8, exists bool) {
	v := m.is_enable
	if v == nil {
		return
	}
	return *v, true
}

// OldIsEnable returns the old "is_enable" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldIsEnable(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsEnable is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsEnable requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsEnable: %w", err)
	}
	return oldValue.IsEnable, nil
}

// AddIsEnable adds u to the "is_enable" field.
func (m *UserMutation) AddIsEnable(u int8) {
	if m.addis_enable != nil {
		*m.addis_enable += u
	} else {
		m.addis_enable = &u
	}
}

// AddedIsEnable returns the value that was added to the "is_enable" field in this mutation.
func (m *UserMutation) AddedIsEnable() (r int8, exists bool) {
	v := m.addis_enable
	if v == nil {
		return
	}
	return *v, true
}

// ResetIsEnable resets all changes to the "is_enable" field.
func (m *UserMutation) ResetIsEnable() {
	m.is_enable = nil
	m.addis_enable = nil
}

// SetStatus sets the "status" field.
func (m *UserMutation) SetStatus(u uint8) {
	m.status = &u
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *UserMutation) Status() (r uint8, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStatus(ctx context.Context) (v uint8, err error) {
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

// AddStatus adds u to the "status" field.
func (m *UserMutation) AddStatus(u int8) {
	if m.addstatus != nil {
		*m.addstatus += u
	} else {
		m.addstatus = &u
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *UserMutation) AddedStatus() (r int8, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus resets all changes to the "status" field.
func (m *UserMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.delete_time != nil {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.uqid != nil {
		fields = append(fields, user.FieldUqid)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.avatar != nil {
		fields = append(fields, user.FieldAvatar)
	}
	if m._type != nil {
		fields = append(fields, user.FieldType)
	}
	if m.is_enable != nil {
		fields = append(fields, user.FieldIsEnable)
	}
	if m.status != nil {
		fields = append(fields, user.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldDeleteTime:
		return m.DeleteTime()
	case user.FieldUqid:
		return m.Uqid()
	case user.FieldName:
		return m.Name()
	case user.FieldAvatar:
		return m.Avatar()
	case user.FieldType:
		return m.GetType()
	case user.FieldIsEnable:
		return m.IsEnable()
	case user.FieldStatus:
		return m.Status()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldDeleteTime:
		return m.OldDeleteTime(ctx)
	case user.FieldUqid:
		return m.OldUqid(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldAvatar:
		return m.OldAvatar(ctx)
	case user.FieldType:
		return m.OldType(ctx)
	case user.FieldIsEnable:
		return m.OldIsEnable(ctx)
	case user.FieldStatus:
		return m.OldStatus(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldDeleteTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteTime(v)
		return nil
	case user.FieldUqid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUqid(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldAvatar:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatar(v)
		return nil
	case user.FieldType:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case user.FieldIsEnable:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsEnable(v)
		return nil
	case user.FieldStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.add_type != nil {
		fields = append(fields, user.FieldType)
	}
	if m.addis_enable != nil {
		fields = append(fields, user.FieldIsEnable)
	}
	if m.addstatus != nil {
		fields = append(fields, user.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldType:
		return m.AddedType()
	case user.FieldIsEnable:
		return m.AddedIsEnable()
	case user.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldType:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddType(v)
		return nil
	case user.FieldIsEnable:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIsEnable(v)
		return nil
	case user.FieldStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldDeleteTime) {
		fields = append(fields, user.FieldDeleteTime)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldDeleteTime:
		m.ClearDeleteTime()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldDeleteTime:
		m.ResetDeleteTime()
		return nil
	case user.FieldUqid:
		m.ResetUqid()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldAvatar:
		m.ResetAvatar()
		return nil
	case user.FieldType:
		m.ResetType()
		return nil
	case user.FieldIsEnable:
		m.ResetIsEnable()
		return nil
	case user.FieldStatus:
		m.ResetStatus()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
