// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Permission is an object representing the database table.
type Permission struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updatedAt" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	DeletedAt time.Time `boil:"deletedAt" json:"deletedAt" toml:"deletedAt" yaml:"deletedAt"`

	R *permissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L permissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PermissionColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

var PermissionTableColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "permissions.id",
	Name:      "permissions.name",
	CreatedAt: "permissions.created_at",
	UpdatedAt: "permissions.updated_at",
	DeletedAt: "permissions.deleted_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PermissionWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"permissions\".\"id\""},
	Name:      whereHelperstring{field: "\"permissions\".\"name\""},
	CreatedAt: whereHelpertime_Time{field: "\"permissions\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"permissions\".\"updated_at\""},
	DeletedAt: whereHelpertime_Time{field: "\"permissions\".\"deleted_at\""},
}

// PermissionRels is where relationship names are stored.
var PermissionRels = struct {
	Roles string
}{
	Roles: "Roles",
}

// permissionR is where relationships are stored.
type permissionR struct {
	Roles RoleSlice `boil:"Roles" json:"Roles" toml:"Roles" yaml:"Roles"`
}

// NewStruct creates a new relationship struct
func (*permissionR) NewStruct() *permissionR {
	return &permissionR{}
}

func (r *permissionR) GetRoles() RoleSlice {
	if r == nil {
		return nil
	}
	return r.Roles
}

// permissionL is where Load methods for each relationship are stored.
type permissionL struct{}

var (
	permissionAllColumns            = []string{"id", "name", "created_at", "updated_at", "deleted_at"}
	permissionColumnsWithoutDefault = []string{"name"}
	permissionColumnsWithDefault    = []string{"id", "created_at", "updated_at", "deleted_at"}
	permissionPrimaryKeyColumns     = []string{"id"}
	permissionGeneratedColumns      = []string{"id"}
)

type (
	// PermissionSlice is an alias for a slice of pointers to Permission.
	// This should almost always be used instead of []Permission.
	PermissionSlice []*Permission
	// PermissionHook is the signature for custom Permission hook methods
	PermissionHook func(context.Context, boil.ContextExecutor, *Permission) error

	permissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	permissionType                 = reflect.TypeOf(&Permission{})
	permissionMapping              = queries.MakeStructMapping(permissionType)
	permissionPrimaryKeyMapping, _ = queries.BindMapping(permissionType, permissionMapping, permissionPrimaryKeyColumns)
	permissionInsertCacheMut       sync.RWMutex
	permissionInsertCache          = make(map[string]insertCache)
	permissionUpdateCacheMut       sync.RWMutex
	permissionUpdateCache          = make(map[string]updateCache)
	permissionUpsertCacheMut       sync.RWMutex
	permissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var permissionAfterSelectMu sync.Mutex
var permissionAfterSelectHooks []PermissionHook

var permissionBeforeInsertMu sync.Mutex
var permissionBeforeInsertHooks []PermissionHook
var permissionAfterInsertMu sync.Mutex
var permissionAfterInsertHooks []PermissionHook

var permissionBeforeUpdateMu sync.Mutex
var permissionBeforeUpdateHooks []PermissionHook
var permissionAfterUpdateMu sync.Mutex
var permissionAfterUpdateHooks []PermissionHook

var permissionBeforeDeleteMu sync.Mutex
var permissionBeforeDeleteHooks []PermissionHook
var permissionAfterDeleteMu sync.Mutex
var permissionAfterDeleteHooks []PermissionHook

var permissionBeforeUpsertMu sync.Mutex
var permissionBeforeUpsertHooks []PermissionHook
var permissionAfterUpsertMu sync.Mutex
var permissionAfterUpsertHooks []PermissionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Permission) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Permission) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Permission) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Permission) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Permission) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Permission) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Permission) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Permission) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Permission) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range permissionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPermissionHook registers your hook function for all future operations.
func AddPermissionHook(hookPoint boil.HookPoint, permissionHook PermissionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		permissionAfterSelectMu.Lock()
		permissionAfterSelectHooks = append(permissionAfterSelectHooks, permissionHook)
		permissionAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		permissionBeforeInsertMu.Lock()
		permissionBeforeInsertHooks = append(permissionBeforeInsertHooks, permissionHook)
		permissionBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		permissionAfterInsertMu.Lock()
		permissionAfterInsertHooks = append(permissionAfterInsertHooks, permissionHook)
		permissionAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		permissionBeforeUpdateMu.Lock()
		permissionBeforeUpdateHooks = append(permissionBeforeUpdateHooks, permissionHook)
		permissionBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		permissionAfterUpdateMu.Lock()
		permissionAfterUpdateHooks = append(permissionAfterUpdateHooks, permissionHook)
		permissionAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		permissionBeforeDeleteMu.Lock()
		permissionBeforeDeleteHooks = append(permissionBeforeDeleteHooks, permissionHook)
		permissionBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		permissionAfterDeleteMu.Lock()
		permissionAfterDeleteHooks = append(permissionAfterDeleteHooks, permissionHook)
		permissionAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		permissionBeforeUpsertMu.Lock()
		permissionBeforeUpsertHooks = append(permissionBeforeUpsertHooks, permissionHook)
		permissionBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		permissionAfterUpsertMu.Lock()
		permissionAfterUpsertHooks = append(permissionAfterUpsertHooks, permissionHook)
		permissionAfterUpsertMu.Unlock()
	}
}

// One returns a single permission record from the query.
func (q permissionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Permission, error) {
	o := &Permission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for permissions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Permission records from the query.
func (q permissionQuery) All(ctx context.Context, exec boil.ContextExecutor) (PermissionSlice, error) {
	var o []*Permission

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Permission slice")
	}

	if len(permissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Permission records in the query.
func (q permissionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count permissions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q permissionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if permissions exists")
	}

	return count > 0, nil
}

// Roles retrieves all the role's Roles with an executor.
func (o *Permission) Roles(mods ...qm.QueryMod) roleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"roles_permissions\" on \"roles\".\"id\" = \"roles_permissions\".\"role_id\""),
		qm.Where("\"roles_permissions\".\"permission_id\"=?", o.ID),
	)

	return Roles(queryMods...)
}

// LoadRoles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (permissionL) LoadRoles(ctx context.Context, e boil.ContextExecutor, singular bool, maybePermission interface{}, mods queries.Applicator) error {
	var slice []*Permission
	var object *Permission

	if singular {
		var ok bool
		object, ok = maybePermission.(*Permission)
		if !ok {
			object = new(Permission)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePermission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePermission))
			}
		}
	} else {
		s, ok := maybePermission.(*[]*Permission)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePermission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePermission))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &permissionR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &permissionR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.Select("\"roles\".\"id\", \"roles\".\"name\", \"roles\".\"created_at\", \"roles\".\"updated_at\", \"roles\".\"deleted_at\", \"a\".\"permission_id\""),
		qm.From("\"roles\""),
		qm.InnerJoin("\"roles_permissions\" as \"a\" on \"roles\".\"id\" = \"a\".\"role_id\""),
		qm.WhereIn("\"a\".\"permission_id\" in ?", argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load roles")
	}

	var resultSlice []*Role

	var localJoinCols []int64
	for results.Next() {
		one := new(Role)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.Name, &one.CreatedAt, &one.UpdatedAt, &one.DeletedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for roles")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice roles")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for roles")
	}

	if len(roleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Roles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &roleR{}
			}
			foreign.R.Permissions = append(foreign.R.Permissions, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Roles = append(local.R.Roles, foreign)
				if foreign.R == nil {
					foreign.R = &roleR{}
				}
				foreign.R.Permissions = append(foreign.R.Permissions, local)
				break
			}
		}
	}

	return nil
}

// AddRoles adds the given related objects to the existing relationships
// of the permission, optionally inserting them as new records.
// Appends related to o.R.Roles.
// Sets related.R.Permissions appropriately.
func (o *Permission) AddRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Role) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"roles_permissions\" (\"permission_id\", \"role_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &permissionR{
			Roles: related,
		}
	} else {
		o.R.Roles = append(o.R.Roles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &roleR{
				Permissions: PermissionSlice{o},
			}
		} else {
			rel.R.Permissions = append(rel.R.Permissions, o)
		}
	}
	return nil
}

// SetRoles removes all previously related items of the
// permission replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Permissions's Roles accordingly.
// Replaces o.R.Roles with related.
// Sets related.R.Permissions's Roles accordingly.
func (o *Permission) SetRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Role) error {
	query := "delete from \"roles_permissions\" where \"permission_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeRolesFromPermissionsSlice(o, related)
	if o.R != nil {
		o.R.Roles = nil
	}

	return o.AddRoles(ctx, exec, insert, related...)
}

// RemoveRoles relationships from objects passed in.
// Removes related items from R.Roles (uses pointer comparison, removal does not keep order)
// Sets related.R.Permissions.
func (o *Permission) RemoveRoles(ctx context.Context, exec boil.ContextExecutor, related ...*Role) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"roles_permissions\" where \"permission_id\" = $1 and \"role_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeRolesFromPermissionsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Roles {
			if rel != ri {
				continue
			}

			ln := len(o.R.Roles)
			if ln > 1 && i < ln-1 {
				o.R.Roles[i] = o.R.Roles[ln-1]
			}
			o.R.Roles = o.R.Roles[:ln-1]
			break
		}
	}

	return nil
}

func removeRolesFromPermissionsSlice(o *Permission, related []*Role) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Permissions {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Permissions)
			if ln > 1 && i < ln-1 {
				rel.R.Permissions[i] = rel.R.Permissions[ln-1]
			}
			rel.R.Permissions = rel.R.Permissions[:ln-1]
			break
		}
	}
}

// Permissions retrieves all the records using an executor.
func Permissions(mods ...qm.QueryMod) permissionQuery {
	mods = append(mods, qm.From("\"permissions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"permissions\".*"})
	}

	return permissionQuery{q}
}

// FindPermission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPermission(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Permission, error) {
	permissionObj := &Permission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"permissions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, permissionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from permissions")
	}

	if err = permissionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return permissionObj, err
	}

	return permissionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Permission) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no permissions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(permissionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	permissionInsertCacheMut.RLock()
	cache, cached := permissionInsertCache[key]
	permissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			permissionAllColumns,
			permissionColumnsWithDefault,
			permissionColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, permissionGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(permissionType, permissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(permissionType, permissionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"permissions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"permissions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into permissions")
	}

	if !cached {
		permissionInsertCacheMut.Lock()
		permissionInsertCache[key] = cache
		permissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Permission.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Permission) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	permissionUpdateCacheMut.RLock()
	cache, cached := permissionUpdateCache[key]
	permissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			permissionAllColumns,
			permissionPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, permissionGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update permissions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"permissions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, permissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(permissionType, permissionMapping, append(wl, permissionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update permissions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for permissions")
	}

	if !cached {
		permissionUpdateCacheMut.Lock()
		permissionUpdateCache[key] = cache
		permissionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q permissionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for permissions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PermissionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), permissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"permissions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, permissionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in permission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all permission")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Permission) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no permissions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(permissionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	permissionUpsertCacheMut.RLock()
	cache, cached := permissionUpsertCache[key]
	permissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			permissionAllColumns,
			permissionColumnsWithDefault,
			permissionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			permissionAllColumns,
			permissionPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, permissionGeneratedColumns)
		update = strmangle.SetComplement(update, permissionGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert permissions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(permissionPrimaryKeyColumns))
			copy(conflict, permissionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"permissions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(permissionType, permissionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(permissionType, permissionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert permissions")
	}

	if !cached {
		permissionUpsertCacheMut.Lock()
		permissionUpsertCache[key] = cache
		permissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Permission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Permission) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Permission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), permissionPrimaryKeyMapping)
	sql := "DELETE FROM \"permissions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for permissions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q permissionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no permissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for permissions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PermissionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(permissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), permissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"permissions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, permissionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from permission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for permissions")
	}

	if len(permissionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Permission) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPermission(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PermissionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PermissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), permissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"permissions\".* FROM \"permissions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, permissionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in PermissionSlice")
	}

	*o = slice

	return nil
}

// PermissionExists checks if the Permission row exists.
func PermissionExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"permissions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if permissions exists")
	}

	return exists, nil
}

// Exists checks if the Permission row exists.
func (o *Permission) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PermissionExists(ctx, exec, o.ID)
}
