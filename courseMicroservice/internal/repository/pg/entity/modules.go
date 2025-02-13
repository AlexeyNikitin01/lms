// Code generated by SQLBoiler 4.17.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Module is an object representing the database table.
type Module struct {
	ID       int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name     string `boil:"name" json:"name" toml:"name" yaml:"name"`
	CourseID int64  `boil:"course_id" json:"course_id" toml:"course_id" yaml:"course_id"`

	R *moduleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L moduleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ModuleColumns = struct {
	ID       string
	Name     string
	CourseID string
}{
	ID:       "id",
	Name:     "name",
	CourseID: "course_id",
}

var ModuleTableColumns = struct {
	ID       string
	Name     string
	CourseID string
}{
	ID:       "modules.id",
	Name:     "modules.name",
	CourseID: "modules.course_id",
}

// Generated where

var ModuleWhere = struct {
	ID       whereHelperint64
	Name     whereHelperstring
	CourseID whereHelperint64
}{
	ID:       whereHelperint64{field: "\"modules\".\"id\""},
	Name:     whereHelperstring{field: "\"modules\".\"name\""},
	CourseID: whereHelperint64{field: "\"modules\".\"course_id\""},
}

// ModuleRels is where relationship names are stored.
var ModuleRels = struct {
	Course   string
	Lectures string
}{
	Course:   "Course",
	Lectures: "Lectures",
}

// moduleR is where relationships are stored.
type moduleR struct {
	Course   *Course      `boil:"Course" json:"Course" toml:"Course" yaml:"Course"`
	Lectures LectureSlice `boil:"Lectures" json:"Lectures" toml:"Lectures" yaml:"Lectures"`
}

// NewStruct creates a new relationship struct
func (*moduleR) NewStruct() *moduleR {
	return &moduleR{}
}

func (r *moduleR) GetCourse() *Course {
	if r == nil {
		return nil
	}
	return r.Course
}

func (r *moduleR) GetLectures() LectureSlice {
	if r == nil {
		return nil
	}
	return r.Lectures
}

// moduleL is where Load methods for each relationship are stored.
type moduleL struct{}

var (
	moduleAllColumns            = []string{"id", "name", "course_id"}
	moduleColumnsWithoutDefault = []string{"name", "course_id"}
	moduleColumnsWithDefault    = []string{"id"}
	modulePrimaryKeyColumns     = []string{"id"}
	moduleGeneratedColumns      = []string{"id"}
)

type (
	// ModuleSlice is an alias for a slice of pointers to Module.
	// This should almost always be used instead of []Module.
	ModuleSlice []*Module
	// ModuleHook is the signature for custom Module hook methods
	ModuleHook func(context.Context, boil.ContextExecutor, *Module) error

	moduleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	moduleType                 = reflect.TypeOf(&Module{})
	moduleMapping              = queries.MakeStructMapping(moduleType)
	modulePrimaryKeyMapping, _ = queries.BindMapping(moduleType, moduleMapping, modulePrimaryKeyColumns)
	moduleInsertCacheMut       sync.RWMutex
	moduleInsertCache          = make(map[string]insertCache)
	moduleUpdateCacheMut       sync.RWMutex
	moduleUpdateCache          = make(map[string]updateCache)
	moduleUpsertCacheMut       sync.RWMutex
	moduleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var moduleAfterSelectMu sync.Mutex
var moduleAfterSelectHooks []ModuleHook

var moduleBeforeInsertMu sync.Mutex
var moduleBeforeInsertHooks []ModuleHook
var moduleAfterInsertMu sync.Mutex
var moduleAfterInsertHooks []ModuleHook

var moduleBeforeUpdateMu sync.Mutex
var moduleBeforeUpdateHooks []ModuleHook
var moduleAfterUpdateMu sync.Mutex
var moduleAfterUpdateHooks []ModuleHook

var moduleBeforeDeleteMu sync.Mutex
var moduleBeforeDeleteHooks []ModuleHook
var moduleAfterDeleteMu sync.Mutex
var moduleAfterDeleteHooks []ModuleHook

var moduleBeforeUpsertMu sync.Mutex
var moduleBeforeUpsertHooks []ModuleHook
var moduleAfterUpsertMu sync.Mutex
var moduleAfterUpsertHooks []ModuleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Module) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Module) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Module) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Module) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Module) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Module) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Module) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Module) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Module) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moduleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddModuleHook registers your hook function for all future operations.
func AddModuleHook(hookPoint boil.HookPoint, moduleHook ModuleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		moduleAfterSelectMu.Lock()
		moduleAfterSelectHooks = append(moduleAfterSelectHooks, moduleHook)
		moduleAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		moduleBeforeInsertMu.Lock()
		moduleBeforeInsertHooks = append(moduleBeforeInsertHooks, moduleHook)
		moduleBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		moduleAfterInsertMu.Lock()
		moduleAfterInsertHooks = append(moduleAfterInsertHooks, moduleHook)
		moduleAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		moduleBeforeUpdateMu.Lock()
		moduleBeforeUpdateHooks = append(moduleBeforeUpdateHooks, moduleHook)
		moduleBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		moduleAfterUpdateMu.Lock()
		moduleAfterUpdateHooks = append(moduleAfterUpdateHooks, moduleHook)
		moduleAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		moduleBeforeDeleteMu.Lock()
		moduleBeforeDeleteHooks = append(moduleBeforeDeleteHooks, moduleHook)
		moduleBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		moduleAfterDeleteMu.Lock()
		moduleAfterDeleteHooks = append(moduleAfterDeleteHooks, moduleHook)
		moduleAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		moduleBeforeUpsertMu.Lock()
		moduleBeforeUpsertHooks = append(moduleBeforeUpsertHooks, moduleHook)
		moduleBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		moduleAfterUpsertMu.Lock()
		moduleAfterUpsertHooks = append(moduleAfterUpsertHooks, moduleHook)
		moduleAfterUpsertMu.Unlock()
	}
}

// One returns a single module record from the query.
func (q moduleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Module, error) {
	o := &Module{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for modules")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Module records from the query.
func (q moduleQuery) All(ctx context.Context, exec boil.ContextExecutor) (ModuleSlice, error) {
	var o []*Module

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Module slice")
	}

	if len(moduleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Module records in the query.
func (q moduleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count modules rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q moduleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if modules exists")
	}

	return count > 0, nil
}

// Course pointed to by the foreign key.
func (o *Module) Course(mods ...qm.QueryMod) courseQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CourseID),
	}

	queryMods = append(queryMods, mods...)

	return Courses(queryMods...)
}

// Lectures retrieves all the lecture's Lectures with an executor.
func (o *Module) Lectures(mods ...qm.QueryMod) lectureQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"lectures\".\"module_id\"=?", o.ID),
	)

	return Lectures(queryMods...)
}

// LoadCourse allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (moduleL) LoadCourse(ctx context.Context, e boil.ContextExecutor, singular bool, maybeModule interface{}, mods queries.Applicator) error {
	var slice []*Module
	var object *Module

	if singular {
		var ok bool
		object, ok = maybeModule.(*Module)
		if !ok {
			object = new(Module)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeModule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeModule))
			}
		}
	} else {
		s, ok := maybeModule.(*[]*Module)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeModule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeModule))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &moduleR{}
		}
		args[object.CourseID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &moduleR{}
			}

			args[obj.CourseID] = struct{}{}

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
		qm.From(`courses`),
		qm.WhereIn(`courses.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Course")
	}

	var resultSlice []*Course
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Course")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for courses")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for courses")
	}

	if len(courseAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Course = foreign
		if foreign.R == nil {
			foreign.R = &courseR{}
		}
		foreign.R.Modules = append(foreign.R.Modules, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CourseID == foreign.ID {
				local.R.Course = foreign
				if foreign.R == nil {
					foreign.R = &courseR{}
				}
				foreign.R.Modules = append(foreign.R.Modules, local)
				break
			}
		}
	}

	return nil
}

// LoadLectures allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (moduleL) LoadLectures(ctx context.Context, e boil.ContextExecutor, singular bool, maybeModule interface{}, mods queries.Applicator) error {
	var slice []*Module
	var object *Module

	if singular {
		var ok bool
		object, ok = maybeModule.(*Module)
		if !ok {
			object = new(Module)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeModule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeModule))
			}
		}
	} else {
		s, ok := maybeModule.(*[]*Module)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeModule)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeModule))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &moduleR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &moduleR{}
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
		qm.From(`lectures`),
		qm.WhereIn(`lectures.module_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load lectures")
	}

	var resultSlice []*Lecture
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice lectures")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on lectures")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lectures")
	}

	if len(lectureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Lectures = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &lectureR{}
			}
			foreign.R.Module = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ModuleID {
				local.R.Lectures = append(local.R.Lectures, foreign)
				if foreign.R == nil {
					foreign.R = &lectureR{}
				}
				foreign.R.Module = local
				break
			}
		}
	}

	return nil
}

// SetCourse of the module to the related item.
// Sets o.R.Course to related.
// Adds o to related.R.Modules.
func (o *Module) SetCourse(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Course) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"modules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"course_id"}),
		strmangle.WhereClause("\"", "\"", 2, modulePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CourseID = related.ID
	if o.R == nil {
		o.R = &moduleR{
			Course: related,
		}
	} else {
		o.R.Course = related
	}

	if related.R == nil {
		related.R = &courseR{
			Modules: ModuleSlice{o},
		}
	} else {
		related.R.Modules = append(related.R.Modules, o)
	}

	return nil
}

// AddLectures adds the given related objects to the existing relationships
// of the module, optionally inserting them as new records.
// Appends related to o.R.Lectures.
// Sets related.R.Module appropriately.
func (o *Module) AddLectures(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Lecture) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ModuleID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"lectures\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"module_id"}),
				strmangle.WhereClause("\"", "\"", 2, lecturePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ModuleID = o.ID
		}
	}

	if o.R == nil {
		o.R = &moduleR{
			Lectures: related,
		}
	} else {
		o.R.Lectures = append(o.R.Lectures, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &lectureR{
				Module: o,
			}
		} else {
			rel.R.Module = o
		}
	}
	return nil
}

// Modules retrieves all the records using an executor.
func Modules(mods ...qm.QueryMod) moduleQuery {
	mods = append(mods, qm.From("\"modules\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"modules\".*"})
	}

	return moduleQuery{q}
}

// FindModule retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindModule(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Module, error) {
	moduleObj := &Module{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"modules\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, moduleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from modules")
	}

	if err = moduleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return moduleObj, err
	}

	return moduleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Module) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no modules provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(moduleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	moduleInsertCacheMut.RLock()
	cache, cached := moduleInsertCache[key]
	moduleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			moduleAllColumns,
			moduleColumnsWithDefault,
			moduleColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, moduleGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(moduleType, moduleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(moduleType, moduleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"modules\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"modules\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into modules")
	}

	if !cached {
		moduleInsertCacheMut.Lock()
		moduleInsertCache[key] = cache
		moduleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Module.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Module) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	moduleUpdateCacheMut.RLock()
	cache, cached := moduleUpdateCache[key]
	moduleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			moduleAllColumns,
			modulePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, moduleGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update modules, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"modules\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, modulePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(moduleType, moduleMapping, append(wl, modulePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update modules row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for modules")
	}

	if !cached {
		moduleUpdateCacheMut.Lock()
		moduleUpdateCache[key] = cache
		moduleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q moduleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for modules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for modules")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ModuleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), modulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"modules\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, modulePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in module slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all module")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Module) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no modules provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(moduleColumnsWithDefault, o)

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

	moduleUpsertCacheMut.RLock()
	cache, cached := moduleUpsertCache[key]
	moduleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			moduleAllColumns,
			moduleColumnsWithDefault,
			moduleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			moduleAllColumns,
			modulePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, moduleGeneratedColumns)
		update = strmangle.SetComplement(update, moduleGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert modules, could not build update column list")
		}

		ret := strmangle.SetComplement(moduleAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(modulePrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert modules, could not build conflict column list")
			}

			conflict = make([]string, len(modulePrimaryKeyColumns))
			copy(conflict, modulePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"modules\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(moduleType, moduleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(moduleType, moduleMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert modules")
	}

	if !cached {
		moduleUpsertCacheMut.Lock()
		moduleUpsertCache[key] = cache
		moduleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Module record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Module) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Module provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), modulePrimaryKeyMapping)
	sql := "DELETE FROM \"modules\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from modules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for modules")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q moduleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no moduleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from modules")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for modules")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ModuleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(moduleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), modulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"modules\" WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, modulePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from module slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for modules")
	}

	if len(moduleAfterDeleteHooks) != 0 {
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
func (o *Module) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindModule(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ModuleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ModuleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), modulePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"modules\".* FROM \"modules\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, modulePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in ModuleSlice")
	}

	*o = slice

	return nil
}

// ModuleExists checks if the Module row exists.
func ModuleExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"modules\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if modules exists")
	}

	return exists, nil
}

// Exists checks if the Module row exists.
func (o *Module) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ModuleExists(ctx, exec, o.ID)
}
