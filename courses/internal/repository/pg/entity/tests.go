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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Test is an object representing the database table.
type Test struct {
	ID        int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	LectureID int64  `boil:"lecture_id" json:"lecture_id" toml:"lecture_id" yaml:"lecture_id"`

	R *testR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L testL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TestColumns = struct {
	ID        string
	Name      string
	LectureID string
}{
	ID:        "id",
	Name:      "name",
	LectureID: "lecture_id",
}

var TestTableColumns = struct {
	ID        string
	Name      string
	LectureID string
}{
	ID:        "tests.id",
	Name:      "tests.name",
	LectureID: "tests.lecture_id",
}

// Generated where

var TestWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	LectureID whereHelperint64
}{
	ID:        whereHelperint64{field: "\"tests\".\"id\""},
	Name:      whereHelperstring{field: "\"tests\".\"name\""},
	LectureID: whereHelperint64{field: "\"tests\".\"lecture_id\""},
}

// TestRels is where relationship names are stored.
var TestRels = struct {
	Lecture   string
	Questions string
}{
	Lecture:   "Lecture",
	Questions: "Questions",
}

// testR is where relationships are stored.
type testR struct {
	Lecture   *Lecture      `boil:"Lecture" json:"Lecture" toml:"Lecture" yaml:"Lecture"`
	Questions QuestionSlice `boil:"Questions" json:"Questions" toml:"Questions" yaml:"Questions"`
}

// NewStruct creates a new relationship struct
func (*testR) NewStruct() *testR {
	return &testR{}
}

func (r *testR) GetLecture() *Lecture {
	if r == nil {
		return nil
	}
	return r.Lecture
}

func (r *testR) GetQuestions() QuestionSlice {
	if r == nil {
		return nil
	}
	return r.Questions
}

// testL is where Load methods for each relationship are stored.
type testL struct{}

var (
	testAllColumns            = []string{"id", "name", "lecture_id"}
	testColumnsWithoutDefault = []string{"name", "lecture_id"}
	testColumnsWithDefault    = []string{"id"}
	testPrimaryKeyColumns     = []string{"id"}
	testGeneratedColumns      = []string{}
)

type (
	// TestSlice is an alias for a slice of pointers to Test.
	// This should almost always be used instead of []Test.
	TestSlice []*Test
	// TestHook is the signature for custom Test hook methods
	TestHook func(context.Context, boil.ContextExecutor, *Test) error

	testQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	testType                 = reflect.TypeOf(&Test{})
	testMapping              = queries.MakeStructMapping(testType)
	testPrimaryKeyMapping, _ = queries.BindMapping(testType, testMapping, testPrimaryKeyColumns)
	testInsertCacheMut       sync.RWMutex
	testInsertCache          = make(map[string]insertCache)
	testUpdateCacheMut       sync.RWMutex
	testUpdateCache          = make(map[string]updateCache)
	testUpsertCacheMut       sync.RWMutex
	testUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var testAfterSelectMu sync.Mutex
var testAfterSelectHooks []TestHook

var testBeforeInsertMu sync.Mutex
var testBeforeInsertHooks []TestHook
var testAfterInsertMu sync.Mutex
var testAfterInsertHooks []TestHook

var testBeforeUpdateMu sync.Mutex
var testBeforeUpdateHooks []TestHook
var testAfterUpdateMu sync.Mutex
var testAfterUpdateHooks []TestHook

var testBeforeDeleteMu sync.Mutex
var testBeforeDeleteHooks []TestHook
var testAfterDeleteMu sync.Mutex
var testAfterDeleteHooks []TestHook

var testBeforeUpsertMu sync.Mutex
var testBeforeUpsertHooks []TestHook
var testAfterUpsertMu sync.Mutex
var testAfterUpsertHooks []TestHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Test) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Test) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Test) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Test) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Test) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Test) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Test) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Test) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Test) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range testAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTestHook registers your hook function for all future operations.
func AddTestHook(hookPoint boil.HookPoint, testHook TestHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		testAfterSelectMu.Lock()
		testAfterSelectHooks = append(testAfterSelectHooks, testHook)
		testAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		testBeforeInsertMu.Lock()
		testBeforeInsertHooks = append(testBeforeInsertHooks, testHook)
		testBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		testAfterInsertMu.Lock()
		testAfterInsertHooks = append(testAfterInsertHooks, testHook)
		testAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		testBeforeUpdateMu.Lock()
		testBeforeUpdateHooks = append(testBeforeUpdateHooks, testHook)
		testBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		testAfterUpdateMu.Lock()
		testAfterUpdateHooks = append(testAfterUpdateHooks, testHook)
		testAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		testBeforeDeleteMu.Lock()
		testBeforeDeleteHooks = append(testBeforeDeleteHooks, testHook)
		testBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		testAfterDeleteMu.Lock()
		testAfterDeleteHooks = append(testAfterDeleteHooks, testHook)
		testAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		testBeforeUpsertMu.Lock()
		testBeforeUpsertHooks = append(testBeforeUpsertHooks, testHook)
		testBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		testAfterUpsertMu.Lock()
		testAfterUpsertHooks = append(testAfterUpsertHooks, testHook)
		testAfterUpsertMu.Unlock()
	}
}

// One returns a single test record from the query.
func (q testQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Test, error) {
	o := &Test{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for tests")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Test records from the query.
func (q testQuery) All(ctx context.Context, exec boil.ContextExecutor) (TestSlice, error) {
	var o []*Test

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Test slice")
	}

	if len(testAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Test records in the query.
func (q testQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count tests rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q testQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if tests exists")
	}

	return count > 0, nil
}

// Lecture pointed to by the foreign key.
func (o *Test) Lecture(mods ...qm.QueryMod) lectureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LectureID),
	}

	queryMods = append(queryMods, mods...)

	return Lectures(queryMods...)
}

// Questions retrieves all the question's Questions with an executor.
func (o *Test) Questions(mods ...qm.QueryMod) questionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"questions\".\"test_id\"=?", o.ID),
	)

	return Questions(queryMods...)
}

// LoadLecture allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (testL) LoadLecture(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTest interface{}, mods queries.Applicator) error {
	var slice []*Test
	var object *Test

	if singular {
		var ok bool
		object, ok = maybeTest.(*Test)
		if !ok {
			object = new(Test)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTest))
			}
		}
	} else {
		s, ok := maybeTest.(*[]*Test)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTest))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &testR{}
		}
		args[object.LectureID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &testR{}
			}

			args[obj.LectureID] = struct{}{}

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
		qm.WhereIn(`lectures.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Lecture")
	}

	var resultSlice []*Lecture
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Lecture")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for lectures")
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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Lecture = foreign
		if foreign.R == nil {
			foreign.R = &lectureR{}
		}
		foreign.R.Tests = append(foreign.R.Tests, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LectureID == foreign.ID {
				local.R.Lecture = foreign
				if foreign.R == nil {
					foreign.R = &lectureR{}
				}
				foreign.R.Tests = append(foreign.R.Tests, local)
				break
			}
		}
	}

	return nil
}

// LoadQuestions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (testL) LoadQuestions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTest interface{}, mods queries.Applicator) error {
	var slice []*Test
	var object *Test

	if singular {
		var ok bool
		object, ok = maybeTest.(*Test)
		if !ok {
			object = new(Test)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTest))
			}
		}
	} else {
		s, ok := maybeTest.(*[]*Test)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTest))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &testR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &testR{}
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
		qm.From(`questions`),
		qm.WhereIn(`questions.test_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load questions")
	}

	var resultSlice []*Question
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice questions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on questions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for questions")
	}

	if len(questionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Questions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &questionR{}
			}
			foreign.R.Test = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TestID {
				local.R.Questions = append(local.R.Questions, foreign)
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.Test = local
				break
			}
		}
	}

	return nil
}

// SetLecture of the test to the related item.
// Sets o.R.Lecture to related.
// Adds o to related.R.Tests.
func (o *Test) SetLecture(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lecture) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lecture_id"}),
		strmangle.WhereClause("\"", "\"", 2, testPrimaryKeyColumns),
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

	o.LectureID = related.ID
	if o.R == nil {
		o.R = &testR{
			Lecture: related,
		}
	} else {
		o.R.Lecture = related
	}

	if related.R == nil {
		related.R = &lectureR{
			Tests: TestSlice{o},
		}
	} else {
		related.R.Tests = append(related.R.Tests, o)
	}

	return nil
}

// AddQuestions adds the given related objects to the existing relationships
// of the test, optionally inserting them as new records.
// Appends related to o.R.Questions.
// Sets related.R.Test appropriately.
func (o *Test) AddQuestions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Question) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TestID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"questions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"test_id"}),
				strmangle.WhereClause("\"", "\"", 2, questionPrimaryKeyColumns),
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

			rel.TestID = o.ID
		}
	}

	if o.R == nil {
		o.R = &testR{
			Questions: related,
		}
	} else {
		o.R.Questions = append(o.R.Questions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &questionR{
				Test: o,
			}
		} else {
			rel.R.Test = o
		}
	}
	return nil
}

// Tests retrieves all the records using an executor.
func Tests(mods ...qm.QueryMod) testQuery {
	mods = append(mods, qm.From("\"tests\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"tests\".*"})
	}

	return testQuery{q}
}

// FindTest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTest(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Test, error) {
	testObj := &Test{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tests\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, testObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from tests")
	}

	if err = testObj.doAfterSelectHooks(ctx, exec); err != nil {
		return testObj, err
	}

	return testObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Test) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no tests provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(testColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	testInsertCacheMut.RLock()
	cache, cached := testInsertCache[key]
	testInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			testAllColumns,
			testColumnsWithDefault,
			testColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(testType, testMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(testType, testMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tests\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tests\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into tests")
	}

	if !cached {
		testInsertCacheMut.Lock()
		testInsertCache[key] = cache
		testInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Test.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Test) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	testUpdateCacheMut.RLock()
	cache, cached := testUpdateCache[key]
	testUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			testAllColumns,
			testPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update tests, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tests\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, testPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(testType, testMapping, append(wl, testPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update tests row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for tests")
	}

	if !cached {
		testUpdateCacheMut.Lock()
		testUpdateCache[key] = cache
		testUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q testQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for tests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for tests")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), testPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, testPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in test slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all test")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Test) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no tests provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(testColumnsWithDefault, o)

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

	testUpsertCacheMut.RLock()
	cache, cached := testUpsertCache[key]
	testUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			testAllColumns,
			testColumnsWithDefault,
			testColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			testAllColumns,
			testPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert tests, could not build update column list")
		}

		ret := strmangle.SetComplement(testAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(testPrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert tests, could not build conflict column list")
			}

			conflict = make([]string, len(testPrimaryKeyColumns))
			copy(conflict, testPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tests\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(testType, testMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(testType, testMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert tests")
	}

	if !cached {
		testUpsertCacheMut.Lock()
		testUpsertCache[key] = cache
		testUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Test record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Test) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Test provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), testPrimaryKeyMapping)
	sql := "DELETE FROM \"tests\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from tests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for tests")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q testQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no testQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from tests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for tests")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(testBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), testPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, testPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from test slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for tests")
	}

	if len(testAfterDeleteHooks) != 0 {
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
func (o *Test) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), testPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tests\".* FROM \"tests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, testPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in TestSlice")
	}

	*o = slice

	return nil
}

// TestExists checks if the Test row exists.
func TestExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tests\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if tests exists")
	}

	return exists, nil
}

// Exists checks if the Test row exists.
func (o *Test) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TestExists(ctx, exec, o.ID)
}
