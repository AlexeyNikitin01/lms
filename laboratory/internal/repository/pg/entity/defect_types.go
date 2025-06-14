// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DefectType is an object representing the database table.
type DefectType struct {
	ID                int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name              string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description       null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	CommonCauses      null.String `boil:"common_causes" json:"common_causes,omitempty" toml:"common_causes" yaml:"common_causes,omitempty"`
	PreventionMethods null.String `boil:"prevention_methods" json:"prevention_methods,omitempty" toml:"prevention_methods" yaml:"prevention_methods,omitempty"`

	R *defectTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L defectTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DefectTypeColumns = struct {
	ID                string
	Name              string
	Description       string
	CommonCauses      string
	PreventionMethods string
}{
	ID:                "id",
	Name:              "name",
	Description:       "description",
	CommonCauses:      "common_causes",
	PreventionMethods: "prevention_methods",
}

var DefectTypeTableColumns = struct {
	ID                string
	Name              string
	Description       string
	CommonCauses      string
	PreventionMethods string
}{
	ID:                "defect_types.id",
	Name:              "defect_types.name",
	Description:       "defect_types.description",
	CommonCauses:      "defect_types.common_causes",
	PreventionMethods: "defect_types.prevention_methods",
}

// Generated where

var DefectTypeWhere = struct {
	ID                whereHelperint
	Name              whereHelperstring
	Description       whereHelpernull_String
	CommonCauses      whereHelpernull_String
	PreventionMethods whereHelpernull_String
}{
	ID:                whereHelperint{field: "\"defect_types\".\"id\""},
	Name:              whereHelperstring{field: "\"defect_types\".\"name\""},
	Description:       whereHelpernull_String{field: "\"defect_types\".\"description\""},
	CommonCauses:      whereHelpernull_String{field: "\"defect_types\".\"common_causes\""},
	PreventionMethods: whereHelpernull_String{field: "\"defect_types\".\"prevention_methods\""},
}

// DefectTypeRels is where relationship names are stored.
var DefectTypeRels = struct {
	DefectFindings string
}{
	DefectFindings: "DefectFindings",
}

// defectTypeR is where relationships are stored.
type defectTypeR struct {
	DefectFindings DefectFindingSlice `boil:"DefectFindings" json:"DefectFindings" toml:"DefectFindings" yaml:"DefectFindings"`
}

// NewStruct creates a new relationship struct
func (*defectTypeR) NewStruct() *defectTypeR {
	return &defectTypeR{}
}

func (r *defectTypeR) GetDefectFindings() DefectFindingSlice {
	if r == nil {
		return nil
	}
	return r.DefectFindings
}

// defectTypeL is where Load methods for each relationship are stored.
type defectTypeL struct{}

var (
	defectTypeAllColumns            = []string{"id", "name", "description", "common_causes", "prevention_methods"}
	defectTypeColumnsWithoutDefault = []string{"name"}
	defectTypeColumnsWithDefault    = []string{"id", "description", "common_causes", "prevention_methods"}
	defectTypePrimaryKeyColumns     = []string{"id"}
	defectTypeGeneratedColumns      = []string{}
)

type (
	// DefectTypeSlice is an alias for a slice of pointers to DefectType.
	// This should almost always be used instead of []DefectType.
	DefectTypeSlice []*DefectType
	// DefectTypeHook is the signature for custom DefectType hook methods
	DefectTypeHook func(context.Context, boil.ContextExecutor, *DefectType) error

	defectTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	defectTypeType                 = reflect.TypeOf(&DefectType{})
	defectTypeMapping              = queries.MakeStructMapping(defectTypeType)
	defectTypePrimaryKeyMapping, _ = queries.BindMapping(defectTypeType, defectTypeMapping, defectTypePrimaryKeyColumns)
	defectTypeInsertCacheMut       sync.RWMutex
	defectTypeInsertCache          = make(map[string]insertCache)
	defectTypeUpdateCacheMut       sync.RWMutex
	defectTypeUpdateCache          = make(map[string]updateCache)
	defectTypeUpsertCacheMut       sync.RWMutex
	defectTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var defectTypeAfterSelectMu sync.Mutex
var defectTypeAfterSelectHooks []DefectTypeHook

var defectTypeBeforeInsertMu sync.Mutex
var defectTypeBeforeInsertHooks []DefectTypeHook
var defectTypeAfterInsertMu sync.Mutex
var defectTypeAfterInsertHooks []DefectTypeHook

var defectTypeBeforeUpdateMu sync.Mutex
var defectTypeBeforeUpdateHooks []DefectTypeHook
var defectTypeAfterUpdateMu sync.Mutex
var defectTypeAfterUpdateHooks []DefectTypeHook

var defectTypeBeforeDeleteMu sync.Mutex
var defectTypeBeforeDeleteHooks []DefectTypeHook
var defectTypeAfterDeleteMu sync.Mutex
var defectTypeAfterDeleteHooks []DefectTypeHook

var defectTypeBeforeUpsertMu sync.Mutex
var defectTypeBeforeUpsertHooks []DefectTypeHook
var defectTypeAfterUpsertMu sync.Mutex
var defectTypeAfterUpsertHooks []DefectTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DefectType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DefectType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DefectType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DefectType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DefectType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DefectType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DefectType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DefectType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DefectType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range defectTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDefectTypeHook registers your hook function for all future operations.
func AddDefectTypeHook(hookPoint boil.HookPoint, defectTypeHook DefectTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		defectTypeAfterSelectMu.Lock()
		defectTypeAfterSelectHooks = append(defectTypeAfterSelectHooks, defectTypeHook)
		defectTypeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		defectTypeBeforeInsertMu.Lock()
		defectTypeBeforeInsertHooks = append(defectTypeBeforeInsertHooks, defectTypeHook)
		defectTypeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		defectTypeAfterInsertMu.Lock()
		defectTypeAfterInsertHooks = append(defectTypeAfterInsertHooks, defectTypeHook)
		defectTypeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		defectTypeBeforeUpdateMu.Lock()
		defectTypeBeforeUpdateHooks = append(defectTypeBeforeUpdateHooks, defectTypeHook)
		defectTypeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		defectTypeAfterUpdateMu.Lock()
		defectTypeAfterUpdateHooks = append(defectTypeAfterUpdateHooks, defectTypeHook)
		defectTypeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		defectTypeBeforeDeleteMu.Lock()
		defectTypeBeforeDeleteHooks = append(defectTypeBeforeDeleteHooks, defectTypeHook)
		defectTypeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		defectTypeAfterDeleteMu.Lock()
		defectTypeAfterDeleteHooks = append(defectTypeAfterDeleteHooks, defectTypeHook)
		defectTypeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		defectTypeBeforeUpsertMu.Lock()
		defectTypeBeforeUpsertHooks = append(defectTypeBeforeUpsertHooks, defectTypeHook)
		defectTypeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		defectTypeAfterUpsertMu.Lock()
		defectTypeAfterUpsertHooks = append(defectTypeAfterUpsertHooks, defectTypeHook)
		defectTypeAfterUpsertMu.Unlock()
	}
}

// One returns a single defectType record from the query.
func (q defectTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DefectType, error) {
	o := &DefectType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for defect_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DefectType records from the query.
func (q defectTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (DefectTypeSlice, error) {
	var o []*DefectType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to DefectType slice")
	}

	if len(defectTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DefectType records in the query.
func (q defectTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count defect_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q defectTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if defect_types exists")
	}

	return count > 0, nil
}

// DefectFindings retrieves all the defect_finding's DefectFindings with an executor.
func (o *DefectType) DefectFindings(mods ...qm.QueryMod) defectFindingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"defect_findings\".\"defect_type_id\"=?", o.ID),
	)

	return DefectFindings(queryMods...)
}

// LoadDefectFindings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (defectTypeL) LoadDefectFindings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDefectType interface{}, mods queries.Applicator) error {
	var slice []*DefectType
	var object *DefectType

	if singular {
		var ok bool
		object, ok = maybeDefectType.(*DefectType)
		if !ok {
			object = new(DefectType)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDefectType)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDefectType))
			}
		}
	} else {
		s, ok := maybeDefectType.(*[]*DefectType)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDefectType)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDefectType))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &defectTypeR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &defectTypeR{}
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
		qm.From(`defect_findings`),
		qm.WhereIn(`defect_findings.defect_type_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load defect_findings")
	}

	var resultSlice []*DefectFinding
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice defect_findings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on defect_findings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for defect_findings")
	}

	if len(defectFindingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.DefectFindings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &defectFindingR{}
			}
			foreign.R.DefectType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.DefectTypeID {
				local.R.DefectFindings = append(local.R.DefectFindings, foreign)
				if foreign.R == nil {
					foreign.R = &defectFindingR{}
				}
				foreign.R.DefectType = local
				break
			}
		}
	}

	return nil
}

// AddDefectFindings adds the given related objects to the existing relationships
// of the defect_type, optionally inserting them as new records.
// Appends related to o.R.DefectFindings.
// Sets related.R.DefectType appropriately.
func (o *DefectType) AddDefectFindings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*DefectFinding) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.DefectTypeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"defect_findings\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"defect_type_id"}),
				strmangle.WhereClause("\"", "\"", 2, defectFindingPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.AnalysisID, rel.DefectTypeID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.DefectTypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &defectTypeR{
			DefectFindings: related,
		}
	} else {
		o.R.DefectFindings = append(o.R.DefectFindings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &defectFindingR{
				DefectType: o,
			}
		} else {
			rel.R.DefectType = o
		}
	}
	return nil
}

// DefectTypes retrieves all the records using an executor.
func DefectTypes(mods ...qm.QueryMod) defectTypeQuery {
	mods = append(mods, qm.From("\"defect_types\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"defect_types\".*"})
	}

	return defectTypeQuery{q}
}

// FindDefectType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDefectType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DefectType, error) {
	defectTypeObj := &DefectType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"defect_types\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, defectTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from defect_types")
	}

	if err = defectTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return defectTypeObj, err
	}

	return defectTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DefectType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no defect_types provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(defectTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	defectTypeInsertCacheMut.RLock()
	cache, cached := defectTypeInsertCache[key]
	defectTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			defectTypeAllColumns,
			defectTypeColumnsWithDefault,
			defectTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(defectTypeType, defectTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(defectTypeType, defectTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"defect_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"defect_types\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into defect_types")
	}

	if !cached {
		defectTypeInsertCacheMut.Lock()
		defectTypeInsertCache[key] = cache
		defectTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DefectType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DefectType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	defectTypeUpdateCacheMut.RLock()
	cache, cached := defectTypeUpdateCache[key]
	defectTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			defectTypeAllColumns,
			defectTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update defect_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"defect_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, defectTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(defectTypeType, defectTypeMapping, append(wl, defectTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update defect_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for defect_types")
	}

	if !cached {
		defectTypeUpdateCacheMut.Lock()
		defectTypeUpdateCache[key] = cache
		defectTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q defectTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for defect_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for defect_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DefectTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), defectTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"defect_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, defectTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in defectType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all defectType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DefectType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no defect_types provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(defectTypeColumnsWithDefault, o)

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

	defectTypeUpsertCacheMut.RLock()
	cache, cached := defectTypeUpsertCache[key]
	defectTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			defectTypeAllColumns,
			defectTypeColumnsWithDefault,
			defectTypeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			defectTypeAllColumns,
			defectTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert defect_types, could not build update column list")
		}

		ret := strmangle.SetComplement(defectTypeAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(defectTypePrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert defect_types, could not build conflict column list")
			}

			conflict = make([]string, len(defectTypePrimaryKeyColumns))
			copy(conflict, defectTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"defect_types\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(defectTypeType, defectTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(defectTypeType, defectTypeMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert defect_types")
	}

	if !cached {
		defectTypeUpsertCacheMut.Lock()
		defectTypeUpsertCache[key] = cache
		defectTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DefectType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DefectType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no DefectType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), defectTypePrimaryKeyMapping)
	sql := "DELETE FROM \"defect_types\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from defect_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for defect_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q defectTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no defectTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from defect_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for defect_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DefectTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(defectTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), defectTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"defect_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, defectTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from defectType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for defect_types")
	}

	if len(defectTypeAfterDeleteHooks) != 0 {
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
func (o *DefectType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDefectType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DefectTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DefectTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), defectTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"defect_types\".* FROM \"defect_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, defectTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in DefectTypeSlice")
	}

	*o = slice

	return nil
}

// DefectTypeExists checks if the DefectType row exists.
func DefectTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"defect_types\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if defect_types exists")
	}

	return exists, nil
}

// Exists checks if the DefectType row exists.
func (o *DefectType) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DefectTypeExists(ctx, exec, o.ID)
}
