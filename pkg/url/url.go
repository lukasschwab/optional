// Package url provides utilities for adding optional values to URL parameters
// iff they are non-nil.
package url

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	o "github.com/lukasschwab/optional/pkg/optional"
)

// A BoolFormatter transforms a boolean into its (pre-encoding) value in a URL.
type BoolFormatter = func(b bool) string

// DefaultBoolFormatter returns the bool's name in lowercase: `false` or `true`.
func DefaultBoolFormatter(b bool) string {
	return strconv.FormatBool(b)
}

// AddBoolToQuery formats and adds optional bool B to query Q iff b is non-nil.
func AddBoolToQuery(q *url.Values, name string, b o.Bool, formatter BoolFormatter) {
	if b != nil {
		q.Set(name, formatter(o.ToBool(b)))
	}
}

// A DurationFormatter transforms a duration into its (pre-encoding) value in a
// URL.
type DurationFormatter = func(d time.Duration) string

// DefaultDurationFormatter returns Go's default string representation of a
// duration; unfortunately, it differs from the ISO 8601 duration string format.
//
// Documented here: https://golang.org/pkg/time/#Duration.String
func DefaultDurationFormatter(d time.Duration) string {
	return d.String()
}

// AddDurationToQuery formats and adds optional duration D to query Q iff D is
// non-nil.
func AddDurationToQuery(q *url.Values, name string, d o.Duration, formatter DurationFormatter) {
	if d != nil {
		q.Set(name, formatter(o.ToDuration(d)))
	}
}

// A FloatFormatter transforms a float64 into its (pre-encoding) value in a URL.
type FloatFormatter = func(f float64) string

// DefaultFloatFormatter returns F in scientific notation, e.g. `-1.234456e+78`.
func DefaultFloatFormatter(f float64) string {
	return fmt.Sprintf("%e", f)
}

// AddFloat64ToQuery formats and adds optional float F to query Q iff F is
// non-nil.
func AddFloat64ToQuery(q *url.Values, name string, f o.Float64, formatter FloatFormatter) {
	if f != nil {
		q.Set(name, formatter(o.ToFloat64(f)))
	}
}

// An IntFormatter transforms an int into its (pre-encoding) value in a URL.
type IntFormatter = func(i int) string

// DefaultIntFormatter returns I in a decimal string representation.
func DefaultIntFormatter(i int) string {
	return strconv.Itoa(i)
}

// AddIntToQuery formats and adds optional int I to query Q iff I is non-nil.
func AddIntToQuery(q *url.Values, name string, i o.Int, formatter IntFormatter) {
	if i != nil {
		q.Set(name, formatter(o.ToInt(i)))
	}
}

// AddStringToQuery adds string S to query Q iff S is non-nil.
func AddStringToQuery(q *url.Values, name string, s o.String) {
	if s != nil {
		q.Set(name, o.ToString(s))
	}
}

// A UintFormatter transforms a uint into its (pre-encoding) value in a URL.
type UintFormatter = func(u uint) string

// DefaultUintFormatter returns U in a decimal string representation.
func DefaultUintFormatter(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

// AddUintToQuery formats and adds optional uint U to query Q iff U is non-nil.
func AddUintToQuery(q *url.Values, name string, u o.Uint, formatter UintFormatter) {
	if u != nil {
		q.Set(name, formatter(o.ToUint(u)))
	}
}

// Note: slices are naturally nil-able; while they have no place in package
// optional, I find it convenient to use a similar abstraction.

// A SliceFormatter transforms a slice of strings into its (pre-encoding)
// single-string value in a URL.
type SliceFormatter = func(s []string) string

// SeparatorFormatter returns a SliceFormatter that inserts SEPARATOR between
// each member of the slice.
//
// The following code transforms `[]string{"hello", "world"}` into the single
// string `"hello|world"`:
//
//    SeparatorFormatter("|")([]string{"hello", "world"})
func SeparatorFormatter(separator string) SliceFormatter {
	return func(s []string) string {
		return strings.Join(s, separator)
	}
}

// AddSliceToQuery formats and adds optional string slice S to query Q iff S is
// non-nil.
func AddSliceToQuery(q *url.Values, name string, s []string, formatter SliceFormatter) {
	if s != nil {
		q.Set(name, formatter(s))
	}
}
