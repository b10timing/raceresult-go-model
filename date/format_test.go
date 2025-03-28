// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"testing"
)

func TestDate_String(t *testing.T) {
	cases := []struct {
		value string
	}{
		{"-0001-01-01"},
		{"0000-01-01"},
		{"1000-01-01"},
		{"1970-01-01"},
		{"2000-11-22"},
		{"+10000-01-01"},
	}
	for _, c := range cases {
		d := MustParseISO(c.value)
		value := d.String()
		if value != c.value {
			t.Errorf("String() == %v, want %v", value, c.value)
		}
	}
}

func TestDate_FormatISO(t *testing.T) {
	cases := []struct {
		value string
		n     int
	}{
		{"-5000-02-03", 4},
		{"-05000-02-03", 5},
		{"-005000-02-03", 6},
		{"+0000-01-01", 4},
		{"+00000-01-01", 5},
		{"+1000-01-01", 4},
		{"+01000-01-01", 5},
		{"+1970-01-01", 4},
		{"+001999-12-31", 6},
		{"+999999-12-31", 6},
	}
	for _, c := range cases {
		d := MustParseISO(c.value)
		value := d.FormatISO(c.n)
		if value != c.value {
			t.Errorf("FormatISO(%v) == %v, want %v", c, value, c.value)
		}
	}
}
