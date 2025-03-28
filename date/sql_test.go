// Copyright 2015 Rick Beton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"testing"
)

/*
	func TestDate_Scan(t *testing.T) {
		cases := []struct {
			v        interface{}
			expected PeriodOfDays
		}{
			{int64(0), 0},
			{int64(1000), 1000},
			{int64(10000), 10000},
			{int64(0), 0},
			{int64(1000), 1000},
			{int64(10000), 10000},
			{"0", 0},
			{"1000", 1000},
			{"10000", 10000},
			{"2018-12-31", 17896},
			{"31/12/2018", 17896},
			{[]byte("10000"), 10000},
			{PeriodOfDays(10000).Date().Local(), 10000},
		}

		for i, c := range cases {
			r := new(Date)
			e := r.Scan(c.v)
			if e != nil {
				t.Errorf("%d: Got %v for %d", i, e, c.expected)
			}
			if r.DaysSinceEpoch() != c.expected {
				t.Errorf("%d: Got %v, want %d", i, *r, c.expected)
			}

			var d driver.Valuer = *r

			q, e := d.Value()
			if e != nil {
				t.Errorf("%d: Got %v for %d", i, e, c.expected)
			}
			if q.(int64) != int64(c.expected) {
				t.Errorf("%d: Got %v, want %d", i, q, c.expected)
			}
		}
	}
*/
func TestDate_Scan_with_junk(t *testing.T) {
	cases := []struct {
		v        interface{}
		expected string
	}{
		{true, "bool true is not a meaningful date"},
		{true, "bool true is not a meaningful date"},
	}

	for i, c := range cases {
		r := new(Date)
		e := r.Scan(c.v)
		if e.Error() != c.expected {
			t.Errorf("%d: Got %q, want %q", i, e.Error(), c.expected)
		}
	}
}

func TestDate_Scan_with_nil(t *testing.T) {
	var r *Date
	e := r.Scan(nil)
	if e != nil {
		t.Errorf("Got %v", e)
	}
}
