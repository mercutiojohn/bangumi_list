// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// Bilibili is an object representing the database table.
type Bilibili struct {
	Actors      null.String  `boil:"actors" json:"actors,omitempty" toml:"actors" yaml:"actors,omitempty"`
	Areas       null.String  `boil:"areas" json:"areas,omitempty" toml:"areas" yaml:"areas,omitempty"`
	Alias       null.String  `boil:"alias" json:"alias,omitempty" toml:"alias" yaml:"alias,omitempty"`
	Cover       null.String  `boil:"cover" json:"cover,omitempty" toml:"cover" yaml:"cover,omitempty"`
	Evaluate    null.String  `boil:"evaluate" json:"evaluate,omitempty" toml:"evaluate" yaml:"evaluate,omitempty"`
	JPTitle     null.String  `boil:"jp_title" json:"jp_title,omitempty" toml:"jp_title" yaml:"jp_title,omitempty"`
	MediaID     int64        `boil:"media_id" json:"media_id" toml:"media_id" yaml:"media_id"`
	IsVip       bool         `boil:"is_vip" json:"is_vip" toml:"is_vip" yaml:"is_vip"`
	PubTime     null.Time    `boil:"pub_time" json:"pub_time,omitempty" toml:"pub_time" yaml:"pub_time,omitempty"`
	RatingCount null.Int64   `boil:"rating_count" json:"rating_count,omitempty" toml:"rating_count" yaml:"rating_count,omitempty"`
	RatingScore null.Float64 `boil:"rating_score" json:"rating_score,omitempty" toml:"rating_score" yaml:"rating_score,omitempty"`
	Copyright   null.String  `boil:"copyright" json:"copyright,omitempty" toml:"copyright" yaml:"copyright,omitempty"`
	SeasonID    int64        `boil:"season_id" json:"season_id" toml:"season_id" yaml:"season_id"`
	SeasonTitle null.String  `boil:"season_title" json:"season_title,omitempty" toml:"season_title" yaml:"season_title,omitempty"`
	SeasonType  int64        `boil:"season_type" json:"season_type" toml:"season_type" yaml:"season_type"`
	SeriesTitle null.String  `boil:"series_title" json:"series_title,omitempty" toml:"series_title" yaml:"series_title,omitempty"`
	SquareCover null.String  `boil:"square_cover" json:"square_cover,omitempty" toml:"square_cover" yaml:"square_cover,omitempty"`
	Coins       null.Int64   `boil:"coins" json:"coins,omitempty" toml:"coins" yaml:"coins,omitempty"`
	Danmakus    null.Int64   `boil:"danmakus" json:"danmakus,omitempty" toml:"danmakus" yaml:"danmakus,omitempty"`
	Views       null.Int64   `boil:"views" json:"views,omitempty" toml:"views" yaml:"views,omitempty"`
	Style       null.String  `boil:"style" json:"style,omitempty" toml:"style" yaml:"style,omitempty"`
	Title       null.String  `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	UpMid       null.Int64   `boil:"up_mid" json:"up_mid,omitempty" toml:"up_mid" yaml:"up_mid,omitempty"`
	UpdatedAt   time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt   time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *bilibiliR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bilibiliL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BilibiliColumns = struct {
	Actors      string
	Areas       string
	Alias       string
	Cover       string
	Evaluate    string
	JPTitle     string
	MediaID     string
	IsVip       string
	PubTime     string
	RatingCount string
	RatingScore string
	Copyright   string
	SeasonID    string
	SeasonTitle string
	SeasonType  string
	SeriesTitle string
	SquareCover string
	Coins       string
	Danmakus    string
	Views       string
	Style       string
	Title       string
	UpMid       string
	UpdatedAt   string
	CreatedAt   string
}{
	Actors:      "actors",
	Areas:       "areas",
	Alias:       "alias",
	Cover:       "cover",
	Evaluate:    "evaluate",
	JPTitle:     "jp_title",
	MediaID:     "media_id",
	IsVip:       "is_vip",
	PubTime:     "pub_time",
	RatingCount: "rating_count",
	RatingScore: "rating_score",
	Copyright:   "copyright",
	SeasonID:    "season_id",
	SeasonTitle: "season_title",
	SeasonType:  "season_type",
	SeriesTitle: "series_title",
	SquareCover: "square_cover",
	Coins:       "coins",
	Danmakus:    "danmakus",
	Views:       "views",
	Style:       "style",
	Title:       "title",
	UpMid:       "up_mid",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
}

var BilibiliTableColumns = struct {
	Actors      string
	Areas       string
	Alias       string
	Cover       string
	Evaluate    string
	JPTitle     string
	MediaID     string
	IsVip       string
	PubTime     string
	RatingCount string
	RatingScore string
	Copyright   string
	SeasonID    string
	SeasonTitle string
	SeasonType  string
	SeriesTitle string
	SquareCover string
	Coins       string
	Danmakus    string
	Views       string
	Style       string
	Title       string
	UpMid       string
	UpdatedAt   string
	CreatedAt   string
}{
	Actors:      "bilibili.actors",
	Areas:       "bilibili.areas",
	Alias:       "bilibili.alias",
	Cover:       "bilibili.cover",
	Evaluate:    "bilibili.evaluate",
	JPTitle:     "bilibili.jp_title",
	MediaID:     "bilibili.media_id",
	IsVip:       "bilibili.is_vip",
	PubTime:     "bilibili.pub_time",
	RatingCount: "bilibili.rating_count",
	RatingScore: "bilibili.rating_score",
	Copyright:   "bilibili.copyright",
	SeasonID:    "bilibili.season_id",
	SeasonTitle: "bilibili.season_title",
	SeasonType:  "bilibili.season_type",
	SeriesTitle: "bilibili.series_title",
	SquareCover: "bilibili.square_cover",
	Coins:       "bilibili.coins",
	Danmakus:    "bilibili.danmakus",
	Views:       "bilibili.views",
	Style:       "bilibili.style",
	Title:       "bilibili.title",
	UpMid:       "bilibili.up_mid",
	UpdatedAt:   "bilibili.updated_at",
	CreatedAt:   "bilibili.created_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var BilibiliWhere = struct {
	Actors      whereHelpernull_String
	Areas       whereHelpernull_String
	Alias       whereHelpernull_String
	Cover       whereHelpernull_String
	Evaluate    whereHelpernull_String
	JPTitle     whereHelpernull_String
	MediaID     whereHelperint64
	IsVip       whereHelperbool
	PubTime     whereHelpernull_Time
	RatingCount whereHelpernull_Int64
	RatingScore whereHelpernull_Float64
	Copyright   whereHelpernull_String
	SeasonID    whereHelperint64
	SeasonTitle whereHelpernull_String
	SeasonType  whereHelperint64
	SeriesTitle whereHelpernull_String
	SquareCover whereHelpernull_String
	Coins       whereHelpernull_Int64
	Danmakus    whereHelpernull_Int64
	Views       whereHelpernull_Int64
	Style       whereHelpernull_String
	Title       whereHelpernull_String
	UpMid       whereHelpernull_Int64
	UpdatedAt   whereHelpertime_Time
	CreatedAt   whereHelpertime_Time
}{
	Actors:      whereHelpernull_String{field: "\"bilibili\".\"actors\""},
	Areas:       whereHelpernull_String{field: "\"bilibili\".\"areas\""},
	Alias:       whereHelpernull_String{field: "\"bilibili\".\"alias\""},
	Cover:       whereHelpernull_String{field: "\"bilibili\".\"cover\""},
	Evaluate:    whereHelpernull_String{field: "\"bilibili\".\"evaluate\""},
	JPTitle:     whereHelpernull_String{field: "\"bilibili\".\"jp_title\""},
	MediaID:     whereHelperint64{field: "\"bilibili\".\"media_id\""},
	IsVip:       whereHelperbool{field: "\"bilibili\".\"is_vip\""},
	PubTime:     whereHelpernull_Time{field: "\"bilibili\".\"pub_time\""},
	RatingCount: whereHelpernull_Int64{field: "\"bilibili\".\"rating_count\""},
	RatingScore: whereHelpernull_Float64{field: "\"bilibili\".\"rating_score\""},
	Copyright:   whereHelpernull_String{field: "\"bilibili\".\"copyright\""},
	SeasonID:    whereHelperint64{field: "\"bilibili\".\"season_id\""},
	SeasonTitle: whereHelpernull_String{field: "\"bilibili\".\"season_title\""},
	SeasonType:  whereHelperint64{field: "\"bilibili\".\"season_type\""},
	SeriesTitle: whereHelpernull_String{field: "\"bilibili\".\"series_title\""},
	SquareCover: whereHelpernull_String{field: "\"bilibili\".\"square_cover\""},
	Coins:       whereHelpernull_Int64{field: "\"bilibili\".\"coins\""},
	Danmakus:    whereHelpernull_Int64{field: "\"bilibili\".\"danmakus\""},
	Views:       whereHelpernull_Int64{field: "\"bilibili\".\"views\""},
	Style:       whereHelpernull_String{field: "\"bilibili\".\"style\""},
	Title:       whereHelpernull_String{field: "\"bilibili\".\"title\""},
	UpMid:       whereHelpernull_Int64{field: "\"bilibili\".\"up_mid\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"bilibili\".\"updated_at\""},
	CreatedAt:   whereHelpertime_Time{field: "\"bilibili\".\"created_at\""},
}

// BilibiliRels is where relationship names are stored.
var BilibiliRels = struct {
}{}

// bilibiliR is where relationships are stored.
type bilibiliR struct {
}

// NewStruct creates a new relationship struct
func (*bilibiliR) NewStruct() *bilibiliR {
	return &bilibiliR{}
}

// bilibiliL is where Load methods for each relationship are stored.
type bilibiliL struct{}

var (
	bilibiliAllColumns            = []string{"actors", "areas", "alias", "cover", "evaluate", "jp_title", "media_id", "is_vip", "pub_time", "rating_count", "rating_score", "copyright", "season_id", "season_title", "season_type", "series_title", "square_cover", "coins", "danmakus", "views", "style", "title", "up_mid", "updated_at", "created_at"}
	bilibiliColumnsWithoutDefault = []string{"media_id", "is_vip", "season_type", "updated_at", "created_at"}
	bilibiliColumnsWithDefault    = []string{"actors", "areas", "alias", "cover", "evaluate", "jp_title", "pub_time", "rating_count", "rating_score", "copyright", "season_id", "season_title", "series_title", "square_cover", "coins", "danmakus", "views", "style", "title", "up_mid"}
	bilibiliPrimaryKeyColumns     = []string{"season_id"}
	bilibiliGeneratedColumns      = []string{"season_id"}
)

type (
	// BilibiliSlice is an alias for a slice of pointers to Bilibili.
	// This should almost always be used instead of []Bilibili.
	BilibiliSlice []*Bilibili
	// BilibiliHook is the signature for custom Bilibili hook methods
	BilibiliHook func(context.Context, boil.ContextExecutor, *Bilibili) error

	bilibiliQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bilibiliType                 = reflect.TypeOf(&Bilibili{})
	bilibiliMapping              = queries.MakeStructMapping(bilibiliType)
	bilibiliPrimaryKeyMapping, _ = queries.BindMapping(bilibiliType, bilibiliMapping, bilibiliPrimaryKeyColumns)
	bilibiliInsertCacheMut       sync.RWMutex
	bilibiliInsertCache          = make(map[string]insertCache)
	bilibiliUpdateCacheMut       sync.RWMutex
	bilibiliUpdateCache          = make(map[string]updateCache)
	bilibiliUpsertCacheMut       sync.RWMutex
	bilibiliUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bilibiliAfterSelectHooks []BilibiliHook

var bilibiliBeforeInsertHooks []BilibiliHook
var bilibiliAfterInsertHooks []BilibiliHook

var bilibiliBeforeUpdateHooks []BilibiliHook
var bilibiliAfterUpdateHooks []BilibiliHook

var bilibiliBeforeDeleteHooks []BilibiliHook
var bilibiliAfterDeleteHooks []BilibiliHook

var bilibiliBeforeUpsertHooks []BilibiliHook
var bilibiliAfterUpsertHooks []BilibiliHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Bilibili) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Bilibili) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Bilibili) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Bilibili) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Bilibili) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Bilibili) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Bilibili) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Bilibili) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Bilibili) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bilibiliAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBilibiliHook registers your hook function for all future operations.
func AddBilibiliHook(hookPoint boil.HookPoint, bilibiliHook BilibiliHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bilibiliAfterSelectHooks = append(bilibiliAfterSelectHooks, bilibiliHook)
	case boil.BeforeInsertHook:
		bilibiliBeforeInsertHooks = append(bilibiliBeforeInsertHooks, bilibiliHook)
	case boil.AfterInsertHook:
		bilibiliAfterInsertHooks = append(bilibiliAfterInsertHooks, bilibiliHook)
	case boil.BeforeUpdateHook:
		bilibiliBeforeUpdateHooks = append(bilibiliBeforeUpdateHooks, bilibiliHook)
	case boil.AfterUpdateHook:
		bilibiliAfterUpdateHooks = append(bilibiliAfterUpdateHooks, bilibiliHook)
	case boil.BeforeDeleteHook:
		bilibiliBeforeDeleteHooks = append(bilibiliBeforeDeleteHooks, bilibiliHook)
	case boil.AfterDeleteHook:
		bilibiliAfterDeleteHooks = append(bilibiliAfterDeleteHooks, bilibiliHook)
	case boil.BeforeUpsertHook:
		bilibiliBeforeUpsertHooks = append(bilibiliBeforeUpsertHooks, bilibiliHook)
	case boil.AfterUpsertHook:
		bilibiliAfterUpsertHooks = append(bilibiliAfterUpsertHooks, bilibiliHook)
	}
}

// One returns a single bilibili record from the query.
func (q bilibiliQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Bilibili, error) {
	o := &Bilibili{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bilibili")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Bilibili records from the query.
func (q bilibiliQuery) All(ctx context.Context, exec boil.ContextExecutor) (BilibiliSlice, error) {
	var o []*Bilibili

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Bilibili slice")
	}

	if len(bilibiliAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Bilibili records in the query.
func (q bilibiliQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bilibili rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bilibiliQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bilibili exists")
	}

	return count > 0, nil
}

// Bilibilis retrieves all the records using an executor.
func Bilibilis(mods ...qm.QueryMod) bilibiliQuery {
	mods = append(mods, qm.From("\"bilibili\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bilibili\".*"})
	}

	return bilibiliQuery{q}
}

// FindBilibili retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBilibili(ctx context.Context, exec boil.ContextExecutor, seasonID int64, selectCols ...string) (*Bilibili, error) {
	bilibiliObj := &Bilibili{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bilibili\" where \"season_id\"=?", sel,
	)

	q := queries.Raw(query, seasonID)

	err := q.Bind(ctx, exec, bilibiliObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bilibili")
	}

	if err = bilibiliObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bilibiliObj, err
	}

	return bilibiliObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Bilibili) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bilibili provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bilibiliColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bilibiliInsertCacheMut.RLock()
	cache, cached := bilibiliInsertCache[key]
	bilibiliInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bilibiliAllColumns,
			bilibiliColumnsWithDefault,
			bilibiliColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, bilibiliGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(bilibiliType, bilibiliMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bilibiliType, bilibiliMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bilibili\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bilibili\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into bilibili")
	}

	if !cached {
		bilibiliInsertCacheMut.Lock()
		bilibiliInsertCache[key] = cache
		bilibiliInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Bilibili.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Bilibili) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bilibiliUpdateCacheMut.RLock()
	cache, cached := bilibiliUpdateCache[key]
	bilibiliUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bilibiliAllColumns,
			bilibiliPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, bilibiliGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bilibili, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bilibili\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, bilibiliPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bilibiliType, bilibiliMapping, append(wl, bilibiliPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bilibili row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bilibili")
	}

	if !cached {
		bilibiliUpdateCacheMut.Lock()
		bilibiliUpdateCache[key] = cache
		bilibiliUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bilibiliQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bilibili")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bilibili")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BilibiliSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bilibiliPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bilibili\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bilibiliPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bilibili slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bilibili")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Bilibili) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bilibili provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bilibiliColumnsWithDefault, o)

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

	bilibiliUpsertCacheMut.RLock()
	cache, cached := bilibiliUpsertCache[key]
	bilibiliUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bilibiliAllColumns,
			bilibiliColumnsWithDefault,
			bilibiliColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bilibiliAllColumns,
			bilibiliPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bilibili, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bilibiliPrimaryKeyColumns))
			copy(conflict, bilibiliPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"bilibili\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bilibiliType, bilibiliMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bilibiliType, bilibiliMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert bilibili")
	}

	if !cached {
		bilibiliUpsertCacheMut.Lock()
		bilibiliUpsertCache[key] = cache
		bilibiliUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Bilibili record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Bilibili) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Bilibili provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bilibiliPrimaryKeyMapping)
	sql := "DELETE FROM \"bilibili\" WHERE \"season_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bilibili")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bilibili")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bilibiliQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bilibiliQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bilibili")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bilibili")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BilibiliSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bilibiliBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bilibiliPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bilibili\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bilibiliPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bilibili slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bilibili")
	}

	if len(bilibiliAfterDeleteHooks) != 0 {
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
func (o *Bilibili) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBilibili(ctx, exec, o.SeasonID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BilibiliSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BilibiliSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bilibiliPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bilibili\".* FROM \"bilibili\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bilibiliPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BilibiliSlice")
	}

	*o = slice

	return nil
}

// BilibiliExists checks if the Bilibili row exists.
func BilibiliExists(ctx context.Context, exec boil.ContextExecutor, seasonID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bilibili\" where \"season_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, seasonID)
	}
	row := exec.QueryRowContext(ctx, sql, seasonID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bilibili exists")
	}

	return exists, nil
}
