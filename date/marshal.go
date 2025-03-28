// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"bytes"
)

// MarshalJSON implements the json.Marshaler interface.
// The date is given in ISO 8601 extended format (e.g. "2006-01-02").
// If the year of the date falls outside the [0,9999] range, this format
// produces an expanded year representation with possibly extra year digits
// beyond the prescribed four-digit minimum and with a + or - sign prefix
// (e.g. , "+12345-06-07", "-0987-06-05").
// Note that the zero value is marshalled as a blank string, which allows
// "omitempty" to work.
func (d Date) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Grow(14)
	buf.WriteByte('"')
	d.WriteTo(buf)
	buf.WriteByte('"')
	return buf.Bytes(), nil
}

// MarshalText implements the encoding.TextMarshaler interface.
// The date is given in ISO 8601 extended format (e.g. "2006-01-02").
// If the year of the date falls outside the [0,9999] range, this format
// produces an expanded year representation with possibly extra year digits
// beyond the prescribed four-digit minimum and with a + or - sign prefix
// (e.g. , "+12345-06-07", "-0987-06-05").
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The date is expected to be in ISO 8601 extended format
// (e.g. "2006-01-02", "+12345-06-07", "-0987-06-05");
// the year must use at least 4 digits and if outside the [0,9999] range
// must be prefixed with a + or - sign.
// Note that the a blank string is unmarshalled as the zero value.
func (d *Date) UnmarshalText(data []byte) (err error) {
	if len(data) == 0 {
		return nil
	}
	u, err := ParseISO(string(data))
	if err == nil {
		d.day = u.day
	}
	return err
}
