// Copyright 2015 Rick Beton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// These methods allow Date and PeriodOfDays to be fields stored in an
// SQL database by implementing the database/sql/driver interfaces.
// The underlying column type can be an integer (period of days since the epoch),
// a string, or a DATE.

// Scan parses some value. If the value holds an integer, it is treated as the
// period-of-days value that represents a Date. Otherwise, if it holds a string,
// the AutoParse function is used.
//
// This implements sql.Scanner https://golang.org/pkg/database/sql/#Scanner
func (d *Date) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	return d.scanAny(value)
}

func (d *Date) scanAny(value interface{}) (err error) {
	err = nil
	switch v := value.(type) {
	case int64:
		*d = Date{PeriodOfDays(v), true}
	case []byte:
		return d.scanString(string(v))
	case string:
		return d.scanString(v)
	case time.Time:
		*d = NewAt(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful date", value, value)
	}

	return err
}

func (d *Date) scanString(value string) (err error) {
	n, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		*d = Date{PeriodOfDays(n), true}
		return nil
	}
	*d, err = AutoParse(value)
	return err
}

// Value implements driver.Valuer https://golang.org/pkg/database/sql/driver/#Valuer
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}
