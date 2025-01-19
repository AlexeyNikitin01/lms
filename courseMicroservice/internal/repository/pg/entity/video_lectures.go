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

// VideoLecture is an object representing the database table.
type VideoLecture struct {
	ID       int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	CourseID int64  `boil:"course_id" json:"course_id" toml:"course_id" yaml:"course_id"`
	URL      string `boil:"url" json:"url" toml:"url" yaml:"url"`

	R *videoLectureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L videoLectureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VideoLectureColumns = struct {
	ID       string
	CourseID string
	URL      string
}{
	ID:       "id",
	CourseID: "course_id",
	URL:      "url",
}

var VideoLectureTableColumns = struct {
	ID       string
	CourseID string
	URL      string
}{
	ID:       "video_lectures.id",
	CourseID: "video_lectures.course_id",
	URL:      "video_lectures.url",
}

// Generated where

var VideoLectureWhere = struct {
	ID       whereHelperint64
	CourseID whereHelperint64
	URL      whereHelperstring
}{
	ID:       whereHelperint64{field: "\"video_lectures\".\"id\""},
	CourseID: whereHelperint64{field: "\"video_lectures\".\"course_id\""},
	URL:      whereHelperstring{field: "\"video_lectures\".\"url\""},
}

// VideoLectureRels is where relationship names are stored.
var VideoLectureRels = struct {
	Course string
}{
	Course: "Course",
}

// videoLectureR is where relationships are stored.
type videoLectureR struct {
	Course *Course `boil:"Course" json:"Course" toml:"Course" yaml:"Course"`
}

// NewStruct creates a new relationship struct
func (*videoLectureR) NewStruct() *videoLectureR {
	return &videoLectureR{}
}

func (r *videoLectureR) GetCourse() *Course {
	if r == nil {
		return nil
	}
	return r.Course
}

// videoLectureL is where Load methods for each relationship are stored.
type videoLectureL struct{}

var (
	videoLectureAllColumns            = []string{"id", "course_id", "url"}
	videoLectureColumnsWithoutDefault = []string{"course_id", "url"}
	videoLectureColumnsWithDefault    = []string{"id"}
	videoLecturePrimaryKeyColumns     = []string{"id"}
	videoLectureGeneratedColumns      = []string{"id"}
)

type (
	// VideoLectureSlice is an alias for a slice of pointers to VideoLecture.
	// This should almost always be used instead of []VideoLecture.
	VideoLectureSlice []*VideoLecture
	// VideoLectureHook is the signature for custom VideoLecture hook methods
	VideoLectureHook func(context.Context, boil.ContextExecutor, *VideoLecture) error

	videoLectureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	videoLectureType                 = reflect.TypeOf(&VideoLecture{})
	videoLectureMapping              = queries.MakeStructMapping(videoLectureType)
	videoLecturePrimaryKeyMapping, _ = queries.BindMapping(videoLectureType, videoLectureMapping, videoLecturePrimaryKeyColumns)
	videoLectureInsertCacheMut       sync.RWMutex
	videoLectureInsertCache          = make(map[string]insertCache)
	videoLectureUpdateCacheMut       sync.RWMutex
	videoLectureUpdateCache          = make(map[string]updateCache)
	videoLectureUpsertCacheMut       sync.RWMutex
	videoLectureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var videoLectureAfterSelectMu sync.Mutex
var videoLectureAfterSelectHooks []VideoLectureHook

var videoLectureBeforeInsertMu sync.Mutex
var videoLectureBeforeInsertHooks []VideoLectureHook
var videoLectureAfterInsertMu sync.Mutex
var videoLectureAfterInsertHooks []VideoLectureHook

var videoLectureBeforeUpdateMu sync.Mutex
var videoLectureBeforeUpdateHooks []VideoLectureHook
var videoLectureAfterUpdateMu sync.Mutex
var videoLectureAfterUpdateHooks []VideoLectureHook

var videoLectureBeforeDeleteMu sync.Mutex
var videoLectureBeforeDeleteHooks []VideoLectureHook
var videoLectureAfterDeleteMu sync.Mutex
var videoLectureAfterDeleteHooks []VideoLectureHook

var videoLectureBeforeUpsertMu sync.Mutex
var videoLectureBeforeUpsertHooks []VideoLectureHook
var videoLectureAfterUpsertMu sync.Mutex
var videoLectureAfterUpsertHooks []VideoLectureHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VideoLecture) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VideoLecture) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VideoLecture) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VideoLecture) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VideoLecture) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VideoLecture) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VideoLecture) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VideoLecture) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VideoLecture) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range videoLectureAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVideoLectureHook registers your hook function for all future operations.
func AddVideoLectureHook(hookPoint boil.HookPoint, videoLectureHook VideoLectureHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		videoLectureAfterSelectMu.Lock()
		videoLectureAfterSelectHooks = append(videoLectureAfterSelectHooks, videoLectureHook)
		videoLectureAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		videoLectureBeforeInsertMu.Lock()
		videoLectureBeforeInsertHooks = append(videoLectureBeforeInsertHooks, videoLectureHook)
		videoLectureBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		videoLectureAfterInsertMu.Lock()
		videoLectureAfterInsertHooks = append(videoLectureAfterInsertHooks, videoLectureHook)
		videoLectureAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		videoLectureBeforeUpdateMu.Lock()
		videoLectureBeforeUpdateHooks = append(videoLectureBeforeUpdateHooks, videoLectureHook)
		videoLectureBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		videoLectureAfterUpdateMu.Lock()
		videoLectureAfterUpdateHooks = append(videoLectureAfterUpdateHooks, videoLectureHook)
		videoLectureAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		videoLectureBeforeDeleteMu.Lock()
		videoLectureBeforeDeleteHooks = append(videoLectureBeforeDeleteHooks, videoLectureHook)
		videoLectureBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		videoLectureAfterDeleteMu.Lock()
		videoLectureAfterDeleteHooks = append(videoLectureAfterDeleteHooks, videoLectureHook)
		videoLectureAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		videoLectureBeforeUpsertMu.Lock()
		videoLectureBeforeUpsertHooks = append(videoLectureBeforeUpsertHooks, videoLectureHook)
		videoLectureBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		videoLectureAfterUpsertMu.Lock()
		videoLectureAfterUpsertHooks = append(videoLectureAfterUpsertHooks, videoLectureHook)
		videoLectureAfterUpsertMu.Unlock()
	}
}

// One returns a single videoLecture record from the query.
func (q videoLectureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VideoLecture, error) {
	o := &VideoLecture{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for video_lectures")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all VideoLecture records from the query.
func (q videoLectureQuery) All(ctx context.Context, exec boil.ContextExecutor) (VideoLectureSlice, error) {
	var o []*VideoLecture

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to VideoLecture slice")
	}

	if len(videoLectureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all VideoLecture records in the query.
func (q videoLectureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count video_lectures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q videoLectureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if video_lectures exists")
	}

	return count > 0, nil
}

// Course pointed to by the foreign key.
func (o *VideoLecture) Course(mods ...qm.QueryMod) courseQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CourseID),
	}

	queryMods = append(queryMods, mods...)

	return Courses(queryMods...)
}

// LoadCourse allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (videoLectureL) LoadCourse(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVideoLecture interface{}, mods queries.Applicator) error {
	var slice []*VideoLecture
	var object *VideoLecture

	if singular {
		var ok bool
		object, ok = maybeVideoLecture.(*VideoLecture)
		if !ok {
			object = new(VideoLecture)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeVideoLecture)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeVideoLecture))
			}
		}
	} else {
		s, ok := maybeVideoLecture.(*[]*VideoLecture)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeVideoLecture)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeVideoLecture))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &videoLectureR{}
		}
		args[object.CourseID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &videoLectureR{}
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
		foreign.R.VideoLectures = append(foreign.R.VideoLectures, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CourseID == foreign.ID {
				local.R.Course = foreign
				if foreign.R == nil {
					foreign.R = &courseR{}
				}
				foreign.R.VideoLectures = append(foreign.R.VideoLectures, local)
				break
			}
		}
	}

	return nil
}

// SetCourse of the videoLecture to the related item.
// Sets o.R.Course to related.
// Adds o to related.R.VideoLectures.
func (o *VideoLecture) SetCourse(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Course) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"video_lectures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"course_id"}),
		strmangle.WhereClause("\"", "\"", 2, videoLecturePrimaryKeyColumns),
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
		o.R = &videoLectureR{
			Course: related,
		}
	} else {
		o.R.Course = related
	}

	if related.R == nil {
		related.R = &courseR{
			VideoLectures: VideoLectureSlice{o},
		}
	} else {
		related.R.VideoLectures = append(related.R.VideoLectures, o)
	}

	return nil
}

// VideoLectures retrieves all the records using an executor.
func VideoLectures(mods ...qm.QueryMod) videoLectureQuery {
	mods = append(mods, qm.From("\"video_lectures\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"video_lectures\".*"})
	}

	return videoLectureQuery{q}
}

// FindVideoLecture retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVideoLecture(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*VideoLecture, error) {
	videoLectureObj := &VideoLecture{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"video_lectures\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, videoLectureObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from video_lectures")
	}

	if err = videoLectureObj.doAfterSelectHooks(ctx, exec); err != nil {
		return videoLectureObj, err
	}

	return videoLectureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VideoLecture) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no video_lectures provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(videoLectureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	videoLectureInsertCacheMut.RLock()
	cache, cached := videoLectureInsertCache[key]
	videoLectureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			videoLectureAllColumns,
			videoLectureColumnsWithDefault,
			videoLectureColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, videoLectureGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(videoLectureType, videoLectureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(videoLectureType, videoLectureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"video_lectures\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"video_lectures\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into video_lectures")
	}

	if !cached {
		videoLectureInsertCacheMut.Lock()
		videoLectureInsertCache[key] = cache
		videoLectureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the VideoLecture.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VideoLecture) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	videoLectureUpdateCacheMut.RLock()
	cache, cached := videoLectureUpdateCache[key]
	videoLectureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			videoLectureAllColumns,
			videoLecturePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, videoLectureGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update video_lectures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"video_lectures\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, videoLecturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(videoLectureType, videoLectureMapping, append(wl, videoLecturePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update video_lectures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for video_lectures")
	}

	if !cached {
		videoLectureUpdateCacheMut.Lock()
		videoLectureUpdateCache[key] = cache
		videoLectureUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q videoLectureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for video_lectures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for video_lectures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VideoLectureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), videoLecturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"video_lectures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, videoLecturePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in videoLecture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all videoLecture")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VideoLecture) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no video_lectures provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(videoLectureColumnsWithDefault, o)

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

	videoLectureUpsertCacheMut.RLock()
	cache, cached := videoLectureUpsertCache[key]
	videoLectureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			videoLectureAllColumns,
			videoLectureColumnsWithDefault,
			videoLectureColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			videoLectureAllColumns,
			videoLecturePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, videoLectureGeneratedColumns)
		update = strmangle.SetComplement(update, videoLectureGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert video_lectures, could not build update column list")
		}

		ret := strmangle.SetComplement(videoLectureAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(videoLecturePrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert video_lectures, could not build conflict column list")
			}

			conflict = make([]string, len(videoLecturePrimaryKeyColumns))
			copy(conflict, videoLecturePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"video_lectures\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(videoLectureType, videoLectureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(videoLectureType, videoLectureMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert video_lectures")
	}

	if !cached {
		videoLectureUpsertCacheMut.Lock()
		videoLectureUpsertCache[key] = cache
		videoLectureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single VideoLecture record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VideoLecture) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no VideoLecture provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), videoLecturePrimaryKeyMapping)
	sql := "DELETE FROM \"video_lectures\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from video_lectures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for video_lectures")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q videoLectureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no videoLectureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from video_lectures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for video_lectures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VideoLectureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(videoLectureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), videoLecturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"video_lectures\" WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, videoLecturePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from videoLecture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for video_lectures")
	}

	if len(videoLectureAfterDeleteHooks) != 0 {
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
func (o *VideoLecture) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVideoLecture(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VideoLectureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VideoLectureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), videoLecturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"video_lectures\".* FROM \"video_lectures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, videoLecturePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in VideoLectureSlice")
	}

	*o = slice

	return nil
}

// VideoLectureExists checks if the VideoLecture row exists.
func VideoLectureExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"video_lectures\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if video_lectures exists")
	}

	return exists, nil
}

// Exists checks if the VideoLecture row exists.
func (o *VideoLecture) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return VideoLectureExists(ctx, exec, o.ID)
}
