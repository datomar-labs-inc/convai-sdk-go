package convai

import (
	"strings"
)

const (
	EQEquals = iota
	EQExists
	EQBetweenExclusive
	EQBetweenInclusive
	EQHasPrefix
)

type Q map[string]interface{}

func NewExecutionMatcher() *ExecutionMatcher {
	return &ExecutionMatcher{
		Filters: []ExecutionQueryItem{},
		MustNot: []ExecutionQueryItem{},
		Sort:    []ExecutionSort{},
	}
}

type ExecutionMatcher struct {
	EnvID string `json:"envId"`
	Lim   int    `json:"limit"`  // Query Limit
	Off   int    `json:"offset"` // Query Offset

	Filters []ExecutionQueryItem `json:"filters"`
	MustNot []ExecutionQueryItem `json:"mustNot"`
	Sort    []ExecutionSort      `json:"sort"`

	NegateCurrent bool                `json:"negateCurrent"`
	CurrentItem   *ExecutionQueryItem `json:"currentItem"`
}

type ExecutionSort struct {
	Field     string `json:"Field"`
	Ascending bool   `json:"asc"`
}

type ExecutionQueryItem struct {
	Op         int      `json:"op"`
	Field      string   `json:"field"`
	Matcher    []string `json:"matcher"`
	LowerBound *string  `json:"lowerBound"`
	UpperBound *string  `json:"upperBound"`
}

func (e *ExecutionMatcher) Where(field string) *ExecutionMatcher {
	field = strings.TrimSpace(field)

	if len(field) == 0 {
		panic("cannot query blank Field")
	}

	if e.CurrentItem != nil {
		if e.NegateCurrent {
			e.MustNot = append(e.MustNot, *e.CurrentItem)
		} else {
			e.Filters = append(e.Filters, *e.CurrentItem)
		}
	}

	e.CurrentItem = &ExecutionQueryItem{
		Op:    0,
		Field: field,
	}
	e.NegateCurrent = false

	return e
}

func (e *ExecutionMatcher) Not() *ExecutionMatcher {
	e.NegateCurrent = true
	return e
}

func (e *ExecutionMatcher) Limit(limit int) *ExecutionMatcher {
	e.Lim = limit
	return e
}

func (e *ExecutionMatcher) Offset(offset int) *ExecutionMatcher {
	e.Off = offset
	return e
}

func (e *ExecutionMatcher) SortAsc(field string) *ExecutionMatcher {
	e.Sort = append(e.Sort, ExecutionSort{
		Field:     field,
		Ascending: true,
	})
	return e
}

func (e *ExecutionMatcher) SortDesc(field string) *ExecutionMatcher {
	e.Sort = append(e.Sort, ExecutionSort{
		Field:     field,
		Ascending: false,
	})
	return e
}

func (e *ExecutionMatcher) Equals(value ...string) *ExecutionMatcher {
	if e.CurrentItem == nil {
		panic("cannot call equals before calling where")
	}

	e.CurrentItem.Op = EQEquals
	e.CurrentItem.Matcher = append(e.CurrentItem.Matcher, value...)

	return e
}

func (e *ExecutionMatcher) Exists() *ExecutionMatcher {
	if e.CurrentItem == nil {
		panic("cannot call equals before calling where")
	}

	e.CurrentItem.Op = EQExists

	return e
}

func (e *ExecutionMatcher) HasPrefix(prefix string) *ExecutionMatcher {
	if e.CurrentItem == nil {
		panic("cannot call equals before calling where")
	}

	e.CurrentItem.Op = EQHasPrefix
	e.CurrentItem.Matcher = append(e.CurrentItem.Matcher, prefix)

	return e
}

func (e *ExecutionMatcher) Between(low, high string, inclusive bool) *ExecutionMatcher {
	if e.CurrentItem == nil {
		panic("cannot call equals before calling where")
	}

	if inclusive {
		e.CurrentItem.Op = EQBetweenInclusive
	} else {
		e.CurrentItem.Op = EQBetweenExclusive
	}

	e.CurrentItem.LowerBound = &low
	e.CurrentItem.UpperBound = &high

	return e
}
