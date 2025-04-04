// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"fmt"
	"io"
	"strings"
)

// These are predefined layouts for use in Date.Format and Date.Parse.
// The reference date used in the layouts is the same date used by the
// time package in the standard library:
//
//	Monday, Jan 2, 2006
//
// To define your own format, write down what the reference date would look
// like formatted your way; see the values of the predefined layouts for
// examples. The model is to demonstrate what the reference date looks like
// so that the Parse function and Format method can apply the same
// transformation to a general date value.
const (
	ISO8601  = "2006-01-02" // ISO 8601 extended format
	ISO8601B = "20060102"   // ISO 8601 basic format
	RFC822   = "02-Jan-06"
	RFC822W  = "Mon, 02-Jan-06" // RFC822 with day of the week
	RFC850   = "Monday, 02-Jan-06"
	RFC1123  = "02 Jan 2006"
	RFC1123W = "Mon, 02 Jan 2006" // RFC1123 with day of the week
	RFC3339  = "2006-01-02"
)

// String returns the time formatted in ISO 8601 extended format
// (e.g. "2006-01-02").  If the year of the date falls outside the
// [0,9999] range, this format produces an expanded year representation
// with possibly extra year digits beyond the prescribed four-digit minimum
// and with a + or - sign prefix (e.g. , "+12345-06-07", "-0987-06-05").
func (d Date) String() string {
	buf := &strings.Builder{}
	buf.Grow(12)
	d.WriteTo(buf)
	return buf.String()
}

// WriteTo is as per String, albeit writing to an io.Writer.
func (d Date) WriteTo(w io.Writer) (n64 int64, err error) {
	if d.IsZero() {
		return 0, nil
	}

	var n int
	year, month, day := d.Date()
	if 0 <= year && year < 10000 {
		n, err = fmt.Fprintf(w, "%04d-%02d-%02d", year, month, day)
	} else {
		n, err = fmt.Fprintf(w, "%+05d-%02d-%02d", year, month, day)
	}
	return int64(n), err
}

// Format returns a textual representation of the date value formatted according
// to layout, which defines the format by showing how the reference date,
// defined to be
//
//	Mon, Jan 2, 2006
//
// would be displayed if it were the value; it serves as an example of the
// desired output.
//
// This function actually uses time.Format to format the input and can use any
// layout accepted by time.Format by extending its date to a time at
// 00:00:00.000 UTC.
//
// Additionally, it is able to insert the day-number suffix into the output string.
// This is done by including "nd" in the format string, which will become
//
//	Mon, Jan 2nd, 2006
//
// For example, New Year's Day might be rendered as "Fri, Jan 1st, 2016". To alter
// the suffix strings for a different locale, change DaySuffixes or use FormatWithSuffixes
// instead.
//
// This function cannot currently format Date values according to the expanded
// year variant of ISO 8601; you should use Date.FormatISO to that effect.
func (d Date) Format(layout string) string {
	if d.IsZero() {
		return ""
	}

	return decode(d.day).Format(layout)
}
