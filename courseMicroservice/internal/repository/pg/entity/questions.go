// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Question is an object representing the database table.
type Question struct {
	ID          int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	LectureID   int64             `boil:"lecture_id" json:"lecture_id" toml:"lecture_id" yaml:"lecture_id"`
	Question    string            `boil:"question" json:"question" toml:"question" yaml:"question"`
	Answer      string            `boil:"answer" json:"answer" toml:"answer" yaml:"answer"`
	WrongAnswer types.StringArray `boil:"wrong_answer" json:"wrong_answer" toml:"wrong_answer" yaml:"wrong_answer"`
	CreatedAt   time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt   null.Time         `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *questionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionColumns = struct {
	ID          string
	LectureID   string
	Question    string
	Answer      string
	WrongAnswer string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	LectureID:   "lecture_id",
	Question:    "question",
	Answer:      "answer",
	WrongAnswer: "wrong_answer",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

var QuestionTableColumns = struct {
	ID          string
	LectureID   string
	Question    string
	Answer      string
	WrongAnswer string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "questions.id",
	LectureID:   "questions.lecture_id",
	Question:    "questions.question",
	Answer:      "questions.answer",
	WrongAnswer: "questions.wrong_answer",
	CreatedAt:   "questions.created_at",
	UpdatedAt:   "questions.updated_at",
	DeletedAt:   "questions.deleted_at",
}

// Generated where

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var QuestionWhere = struct {
	ID          whereHelperint64
	LectureID   whereHelperint64
	Question    whereHelperstring
	Answer      whereHelperstring
	WrongAnswer whereHelpertypes_StringArray
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
	DeletedAt   whereHelpernull_Time
}{
	ID:          whereHelperint64{field: "\"questions\".\"id\""},
	LectureID:   whereHelperint64{field: "\"questions\".\"lecture_id\""},
	Question:    whereHelperstring{field: "\"questions\".\"question\""},
	Answer:      whereHelperstring{field: "\"questions\".\"answer\""},
	WrongAnswer: whereHelpertypes_StringArray{field: "\"questions\".\"wrong_answer\""},
	CreatedAt:   whereHelpertime_Time{field: "\"questions\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"questions\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"questions\".\"deleted_at\""},
}

// QuestionRels is where relationship names are stored.
var QuestionRels = struct {
	Lecture string
}{
	Lecture: "Lecture",
}

// questionR is where relationships are stored.
type questionR struct {
	Lecture *Lecture `boil:"Lecture" json:"Lecture" toml:"Lecture" yaml:"Lecture"`
}

// NewStruct creates a new relationship struct
func (*questionR) NewStruct() *questionR {
	return &questionR{}
}

func (r *questionR) GetLecture() *Lecture {
	if r == nil {
		return nil
	}
	return r.Lecture
}

// questionL is where Load methods for each relationship are stored.
type questionL struct{}

var (
	questionAllColumns            = []string{"id", "lecture_id", "question", "answer", "wrong_answer", "created_at", "updated_at", "deleted_at"}
	questionColumnsWithoutDefault = []string{"lecture_id", "question", "answer", "wrong_answer"}
	questionColumnsWithDefault    = []string{"id", "created_at", "updated_at", "deleted_at"}
	questionPrimaryKeyColumns     = []string{"id"}
	questionGeneratedColumns      = []string{"id"}
)

type (
	// QuestionSlice is an alias for a slice of pointers to Question.
	// This should almost always be used instead of []Question.
	QuestionSlice []*Question
	// QuestionHook is the signature for custom Question hook methods
	QuestionHook func(context.Context, boil.ContextExecutor, *Question) error

	questionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionType                 = reflect.TypeOf(&Question{})
	questionMapping              = queries.MakeStructMapping(questionType)
	questionPrimaryKeyMapping, _ = queries.BindMapping(questionType, questionMapping, questionPrimaryKeyColumns)
	questionInsertCacheMut       sync.RWMutex
	questionInsertCache          = make(map[string]insertCache)
	questionUpdateCacheMut       sync.RWMutex
	questionUpdateCache          = make(map[string]updateCache)
	questionUpsertCacheMut       sync.RWMutex
	questionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var questionAfterSelectHooks []QuestionHook

var questionBeforeInsertHooks []QuestionHook
var questionAfterInsertHooks []QuestionHook

var questionBeforeUpdateHooks []QuestionHook
var questionAfterUpdateHooks []QuestionHook

var questionBeforeDeleteHooks []QuestionHook
var questionAfterDeleteHooks []QuestionHook

var questionBeforeUpsertHooks []QuestionHook
var questionAfterUpsertHooks []QuestionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Question) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Question) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Question) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Question) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Question) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Question) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Question) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Question) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Question) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuestionHook registers your hook function for all future operations.
func AddQuestionHook(hookPoint boil.HookPoint, questionHook QuestionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		questionAfterSelectHooks = append(questionAfterSelectHooks, questionHook)
	case boil.BeforeInsertHook:
		questionBeforeInsertHooks = append(questionBeforeInsertHooks, questionHook)
	case boil.AfterInsertHook:
		questionAfterInsertHooks = append(questionAfterInsertHooks, questionHook)
	case boil.BeforeUpdateHook:
		questionBeforeUpdateHooks = append(questionBeforeUpdateHooks, questionHook)
	case boil.AfterUpdateHook:
		questionAfterUpdateHooks = append(questionAfterUpdateHooks, questionHook)
	case boil.BeforeDeleteHook:
		questionBeforeDeleteHooks = append(questionBeforeDeleteHooks, questionHook)
	case boil.AfterDeleteHook:
		questionAfterDeleteHooks = append(questionAfterDeleteHooks, questionHook)
	case boil.BeforeUpsertHook:
		questionBeforeUpsertHooks = append(questionBeforeUpsertHooks, questionHook)
	case boil.AfterUpsertHook:
		questionAfterUpsertHooks = append(questionAfterUpsertHooks, questionHook)
	}
}

// One returns a single question record from the query.
func (q questionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Question, error) {
	o := &Question{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for questions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Question records from the query.
func (q questionQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuestionSlice, error) {
	var o []*Question

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Question slice")
	}

	if len(questionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Question records in the query.
func (q questionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count questions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q questionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if questions exists")
	}

	return count > 0, nil
}

// Lecture pointed to by the foreign key.
func (o *Question) Lecture(mods ...qm.QueryMod) lectureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LectureID),
	}

	queryMods = append(queryMods, mods...)

	return Lectures(queryMods...)
}

// LoadLecture allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionL) LoadLecture(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestion interface{}, mods queries.Applicator) error {
	var slice []*Question
	var object *Question

	if singular {
		var ok bool
		object, ok = maybeQuestion.(*Question)
		if !ok {
			object = new(Question)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeQuestion)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeQuestion))
			}
		}
	} else {
		s, ok := maybeQuestion.(*[]*Question)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeQuestion)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeQuestion))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &questionR{}
		}
		args = append(args, object.LectureID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionR{}
			}

			for _, a := range args {
				if a == obj.LectureID {
					continue Outer
				}
			}

			args = append(args, obj.LectureID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lectures`),
		qm.WhereIn(`lectures.id in ?`, args...),
		qmhelper.WhereIsNull(`lectures.deleted_at`),
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

	if len(questionAfterSelectHooks) != 0 {
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
		foreign.R.Questions = append(foreign.R.Questions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LectureID == foreign.ID {
				local.R.Lecture = foreign
				if foreign.R == nil {
					foreign.R = &lectureR{}
				}
				foreign.R.Questions = append(foreign.R.Questions, local)
				break
			}
		}
	}

	return nil
}

// SetLecture of the question to the related item.
// Sets o.R.Lecture to related.
// Adds o to related.R.Questions.
func (o *Question) SetLecture(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lecture) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lecture_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionPrimaryKeyColumns),
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
		o.R = &questionR{
			Lecture: related,
		}
	} else {
		o.R.Lecture = related
	}

	if related.R == nil {
		related.R = &lectureR{
			Questions: QuestionSlice{o},
		}
	} else {
		related.R.Questions = append(related.R.Questions, o)
	}

	return nil
}

// Questions retrieves all the records using an executor.
func Questions(mods ...qm.QueryMod) questionQuery {
	mods = append(mods, qm.From("\"questions\""), qmhelper.WhereIsNull("\"questions\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"questions\".*"})
	}

	return questionQuery{q}
}

// FindQuestion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestion(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Question, error) {
	questionObj := &Question{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"questions\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, questionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from questions")
	}

	if err = questionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return questionObj, err
	}

	return questionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Question) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no questions provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(questionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionInsertCacheMut.RLock()
	cache, cached := questionInsertCache[key]
	questionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionAllColumns,
			questionColumnsWithDefault,
			questionColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, questionGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionType, questionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"questions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"questions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into questions")
	}

	if !cached {
		questionInsertCacheMut.Lock()
		questionInsertCache[key] = cache
		questionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Question.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Question) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	questionUpdateCacheMut.RLock()
	cache, cached := questionUpdateCache[key]
	questionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionAllColumns,
			questionPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, questionGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update questions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"questions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, questionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, append(wl, questionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update questions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for questions")
	}

	if !cached {
		questionUpdateCacheMut.Lock()
		questionUpdateCache[key] = cache
		questionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q questionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for questions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, questionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in question slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all question")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Question) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no questions provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(questionColumnsWithDefault, o)

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

	questionUpsertCacheMut.RLock()
	cache, cached := questionUpsertCache[key]
	questionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			questionAllColumns,
			questionColumnsWithDefault,
			questionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			questionAllColumns,
			questionPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, questionGeneratedColumns)
		update = strmangle.SetComplement(update, questionGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert questions, could not build update column list")
		}

		ret := strmangle.SetComplement(questionAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(questionPrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert questions, could not build conflict column list")
			}

			conflict = make([]string, len(questionPrimaryKeyColumns))
			copy(conflict, questionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"questions\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(questionType, questionMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert questions")
	}

	if !cached {
		questionUpsertCacheMut.Lock()
		questionUpsertCache[key] = cache
		questionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Question record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Question) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Question provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionPrimaryKeyMapping)
		sql = "DELETE FROM \"questions\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"questions\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(questionType, questionMapping, append(wl, questionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for questions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q questionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no questionQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for questions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(questionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"questions\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"questions\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, questionPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from question slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for questions")
	}

	if len(questionAfterDeleteHooks) != 0 {
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
func (o *Question) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuestion(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"questions\".* FROM \"questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in QuestionSlice")
	}

	*o = slice

	return nil
}

// QuestionExists checks if the Question row exists.
func QuestionExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"questions\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if questions exists")
	}

	return exists, nil
}