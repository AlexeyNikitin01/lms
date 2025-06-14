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

// LectureAnalysisView is an object representing the database table.
type LectureAnalysisView struct {
	LectureTitle  null.String  `boil:"lecture_title" json:"lecture_title,omitempty" toml:"lecture_title" yaml:"lecture_title,omitempty"`
	DefectsFound  null.Int64   `boil:"defects_found" json:"defects_found,omitempty" toml:"defects_found" yaml:"defects_found,omitempty"`
	AvgConfidence null.Float64 `boil:"avg_confidence" json:"avg_confidence,omitempty" toml:"avg_confidence" yaml:"avg_confidence,omitempty"`
	LastAnalysis  null.Time    `boil:"last_analysis" json:"last_analysis,omitempty" toml:"last_analysis" yaml:"last_analysis,omitempty"`
}

var LectureAnalysisViewColumns = struct {
	LectureTitle  string
	DefectsFound  string
	AvgConfidence string
	LastAnalysis  string
}{
	LectureTitle:  "lecture_title",
	DefectsFound:  "defects_found",
	AvgConfidence: "avg_confidence",
	LastAnalysis:  "last_analysis",
}

var LectureAnalysisViewTableColumns = struct {
	LectureTitle  string
	DefectsFound  string
	AvgConfidence string
	LastAnalysis  string
}{
	LectureTitle:  "lecture_analysis_view.lecture_title",
	DefectsFound:  "lecture_analysis_view.defects_found",
	AvgConfidence: "lecture_analysis_view.avg_confidence",
	LastAnalysis:  "lecture_analysis_view.last_analysis",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var LectureAnalysisViewWhere = struct {
	LectureTitle  whereHelpernull_String
	DefectsFound  whereHelpernull_Int64
	AvgConfidence whereHelpernull_Float64
	LastAnalysis  whereHelpernull_Time
}{
	LectureTitle:  whereHelpernull_String{field: "\"lecture_analysis_view\".\"lecture_title\""},
	DefectsFound:  whereHelpernull_Int64{field: "\"lecture_analysis_view\".\"defects_found\""},
	AvgConfidence: whereHelpernull_Float64{field: "\"lecture_analysis_view\".\"avg_confidence\""},
	LastAnalysis:  whereHelpernull_Time{field: "\"lecture_analysis_view\".\"last_analysis\""},
}

var (
	lectureAnalysisViewAllColumns            = []string{"lecture_title", "defects_found", "avg_confidence", "last_analysis"}
	lectureAnalysisViewColumnsWithoutDefault = []string{}
	lectureAnalysisViewColumnsWithDefault    = []string{"lecture_title", "defects_found", "avg_confidence", "last_analysis"}
	lectureAnalysisViewPrimaryKeyColumns     = []string{}
	lectureAnalysisViewGeneratedColumns      = []string{}
)

type (
	// LectureAnalysisViewSlice is an alias for a slice of pointers to LectureAnalysisView.
	// This should almost always be used instead of []LectureAnalysisView.
	LectureAnalysisViewSlice []*LectureAnalysisView
	// LectureAnalysisViewHook is the signature for custom LectureAnalysisView hook methods
	LectureAnalysisViewHook func(context.Context, boil.ContextExecutor, *LectureAnalysisView) error

	lectureAnalysisViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lectureAnalysisViewType           = reflect.TypeOf(&LectureAnalysisView{})
	lectureAnalysisViewMapping        = queries.MakeStructMapping(lectureAnalysisViewType)
	lectureAnalysisViewInsertCacheMut sync.RWMutex
	lectureAnalysisViewInsertCache    = make(map[string]insertCache)
	lectureAnalysisViewUpdateCacheMut sync.RWMutex
	lectureAnalysisViewUpdateCache    = make(map[string]updateCache)
	lectureAnalysisViewUpsertCacheMut sync.RWMutex
	lectureAnalysisViewUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
	// These are used in some views
	_ = fmt.Sprintln("")
	_ = reflect.Int
	_ = strings.Builder{}
	_ = sync.Mutex{}
	_ = strmangle.Plural("")
	_ = strconv.IntSize
)

var lectureAnalysisViewAfterSelectMu sync.Mutex
var lectureAnalysisViewAfterSelectHooks []LectureAnalysisViewHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LectureAnalysisView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lectureAnalysisViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLectureAnalysisViewHook registers your hook function for all future operations.
func AddLectureAnalysisViewHook(hookPoint boil.HookPoint, lectureAnalysisViewHook LectureAnalysisViewHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		lectureAnalysisViewAfterSelectMu.Lock()
		lectureAnalysisViewAfterSelectHooks = append(lectureAnalysisViewAfterSelectHooks, lectureAnalysisViewHook)
		lectureAnalysisViewAfterSelectMu.Unlock()
	}
}

// One returns a single lectureAnalysisView record from the query.
func (q lectureAnalysisViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LectureAnalysisView, error) {
	o := &LectureAnalysisView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for lecture_analysis_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LectureAnalysisView records from the query.
func (q lectureAnalysisViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (LectureAnalysisViewSlice, error) {
	var o []*LectureAnalysisView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to LectureAnalysisView slice")
	}

	if len(lectureAnalysisViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LectureAnalysisView records in the query.
func (q lectureAnalysisViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count lecture_analysis_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lectureAnalysisViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if lecture_analysis_view exists")
	}

	return count > 0, nil
}

// LectureAnalysisViews retrieves all the records using an executor.
func LectureAnalysisViews(mods ...qm.QueryMod) lectureAnalysisViewQuery {
	mods = append(mods, qm.From("\"lecture_analysis_view\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"lecture_analysis_view\".*"})
	}

	return lectureAnalysisViewQuery{q}
}
