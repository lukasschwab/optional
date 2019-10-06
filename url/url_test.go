package url

import (
	"net/url"
	"testing"
	"time"
)

const key string = "KEY"

var t1 = time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
var t2 = time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
var d = t2.Sub(t1)

func TestDefaultBoolFormatter(t *testing.T) {
	stringTrue := DefaultBoolFormatter(true)
	if stringTrue != "true" {
		t.Errorf("Unexpected default format for true: %s", stringTrue)
	}
	stringFalse := DefaultBoolFormatter(false)
	if stringFalse != "false" {
		t.Errorf("Unexpected default format for false: %s", stringFalse)
	}
}

func TestAddBoolToQuery(t *testing.T) {
	q1 := &url.Values{}
	AddBoolToQuery(q1, key, nil, DefaultBoolFormatter)
	if val := (*q1).Get(key); val != "" {
		t.Errorf("Bool was nil; should be undefined in Values")
	}

	q2 := &url.Values{}
	AddBoolToQuery(q2, key, true, DefaultBoolFormatter)
	if val := (*q2).Get(key); val == "" {
		t.Errorf("Bool was true; shouldn't be undefined in values")
	} else if val != DefaultBoolFormatter(true) {
		t.Errorf("Wrong value for true in values: %s", val)
	}

	q3 := &url.Values{}
	AddBoolToQuery(q3, key, false, DefaultBoolFormatter)
	if val := (*q3).Get(key); val == "" {
		t.Errorf("Bool was false; shouldn't be undefined in values")
	} else if val != DefaultBoolFormatter(false) {
		t.Errorf("Wrong value for false in values: %s", val)
	}
}

func TestDefaultDurationFormatter(t *testing.T) {
	stringD := DefaultDurationFormatter(d)
	if stringD != d.String() {
		t.Errorf("Unexpected default format for duration: %s", stringD)
	}
}

func TestAddDurationToQuery(t *testing.T) {
	q1 := &url.Values{}
	AddDurationToQuery(q1, key, nil, DefaultDurationFormatter)
	if val := (*q1).Get(key); val != "" {
		t.Errorf("Duration was nil; should be undefined in Values")
	}

	q2 := &url.Values{}
	AddDurationToQuery(q2, key, d, DefaultDurationFormatter)
	if val := (*q2).Get(key); val == "" {
		t.Errorf("Duration was non-nil; shouldn't be undefined in values")
	} else if val != d.String() {
		t.Errorf("Wrong value for duration in values: %s", val)
	}
}
