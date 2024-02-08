package vbdate

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"
const dateTimeFormat = "2006-01-02 15:04:05"
const timeFormat = "15:04:05"

// VBDate is identical to time.Time with some VB characteristics
type VBDate time.Time

// ZeroDate returns the VB zero date 1899-12-30
func ZeroDate() VBDate {
	return Date(1899, 12, 30)
}

// Before is equal to time.Time
func (s VBDate) Before(c VBDate) bool {
	return time.Time(s).Before(time.Time(c))
}

// After is equal to time.Time
func (s VBDate) After(c VBDate) bool {
	return time.Time(s).After(time.Time(c))
}

// AddDate is equal to time.Time
func (s VBDate) AddDate(year int, month int, day int) VBDate {
	return VBDate(time.Time(s).AddDate(year, month, day))
}

// Add is equal to time.Time
func (s VBDate) Add(d time.Duration) VBDate {
	return VBDate(time.Time(s).Add(d))
}

// Sub is equal to time.Time
func (s VBDate) Sub(d VBDate) time.Duration {
	return time.Time(s).Sub(time.Time(d))
}

// ToTime returns the VBDate as time.Time
func (s VBDate) ToTime() time.Time {
	if s.IsZero() {
		return time.Time{}
	}
	return time.Time(s)
}

// MarshalJSON creates a JSON string of the date
func (s VBDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.ToString() + "\""), nil
}

// UnmarshalJSON parses a JSON string to date
func (s *VBDate) UnmarshalJSON(data []byte) error {
	str := string(data)
	switch len(str) {
	case 12:
		d, err := time.ParseInLocation(dateFormat, str[1:len(str)-1], time.UTC)
		if err == nil {
			*s = VBDate(d)
		}
		return err
	case 21:
		d, err := time.ParseInLocation(dateTimeFormat, str[1:len(str)-1], time.UTC)
		if err == nil {
			*s = VBDate(d)
		}
		return err
	case 2:
		if str == "\"\"" {
			*s = ZeroDate()
			return nil
		}
	case 0:
		*s = ZeroDate()
		return nil
	}
	return errors.New("date time format not supported")
}

// UnmarshalXML parses an XML string
func (s *VBDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	if d, ok := Parse(str); ok {
		*s = d
	}
	return nil
}

// IsZero returns true if the date is either VB ZeroDate or Go ZeroDate
func (s VBDate) IsZero() bool {
	if s == ZeroDate() {
		return true
	}
	return time.Time(s) == time.Time{}
}

// IsAfterVBZero returns true if the date is after the VB ZeroDate
func (s VBDate) IsAfterVBZero() bool {
	return s.After(ZeroDate())
}

// ToString converts the date to string
func (s VBDate) ToString() string {
	if s.IsZero() {
		return ""
	}
	t := time.Time(s)
	if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 {
		return t.Format(dateFormat)
	}
	return t.Format(dateTimeFormat)
}

// ToStringWithDateFormat converts the date to string with the given date format
func (s VBDate) ToStringWithDateFormat(df string) string {
	if df == "" {
		return s.ToString()
	}
	if s.IsZero() {
		return ""
	}
	t := time.Time(s)
	if t.Unix()%86400 == 0 {
		return t.Format(df)
	}
	return t.Format(df + " " + timeFormat)
}

// Year is equal to time.Time
func (s VBDate) Year() int {
	return time.Time(s).Year()
}

// Month is equal to time.Time
func (s VBDate) Month() time.Month {
	return time.Time(s).Month()
}

// Day is equal to time.Time
func (s VBDate) Day() int {
	return time.Time(s).Day()
}

// Hour is equal to time.Time
func (s VBDate) Hour() int {
	return time.Time(s).Hour()
}

// Minute is equal to time.Time
func (s VBDate) Minute() int {
	return time.Time(s).Minute()
}

// Second is equal to time.Time
func (s VBDate) Second() int {
	return time.Time(s).Second()
}

// Nanosecond is equal to time.Time
func (s VBDate) Nanosecond() int {
	return time.Time(s).Nanosecond()
}

// WithTimezone returns a new VBDate where the timezone was replaced
func (s VBDate) WithTimezone(tz *time.Location) VBDate {
	if s.IsZero() {
		return s
	}
	return VBDate(time.Date(s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Second(), s.Nanosecond(), tz))
}

// Parse parses a date from string
func Parse(str string) (VBDate, bool) {
	if strings.Count(str, ".") == 2 {
		switch len(str) {
		case 10:
			d, err := time.ParseInLocation("02.01.2006", str, time.UTC)
			if err == nil {
				return VBDate(d), true
			}
		case 19:
			d, err := time.ParseInLocation("02.01.2006 "+timeFormat, str, time.UTC)
			if err == nil {
				return VBDate(d), true
			}
		}
		return ZeroDate(), false
	}
	if strings.Count(str, "/") == 2 {
		switch len(str) {
		case 10:
			d, err := time.ParseInLocation("02/01/2006", str, time.UTC)
			if err == nil {
				return VBDate(d), true
			}
		case 19:
			d, err := time.ParseInLocation("02/01/2006 "+timeFormat, str, time.UTC)
			if err == nil {
				return VBDate(d), true
			}
		}
		return ZeroDate(), false
	}

	switch len(str) {
	case 10:
		d, err := time.ParseInLocation(dateFormat, str, time.UTC)
		if err == nil {
			return VBDate(d), true
		}
	case 19:
		d, err := time.ParseInLocation(dateTimeFormat, str, time.UTC)
		if err == nil {
			return VBDate(d), true
		}
	case 20, 25:
		d, err := time.Parse(time.RFC3339, str)
		if err == nil {
			return VBDate(d), true
		}
	}
	return ZeroDate(), false
}

// Date creates a VB Date with day accuracy
func Date(year int, month time.Month, day int) VBDate {
	return VBDate(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// DateTime creates a VB Date with time accuracy
func DateTime(year int, month time.Month, day int, hour, minute, second int) VBDate {
	return VBDate(time.Date(year, month, day, hour, minute, second, 0, time.UTC))
}

// Now returns a VBDate with the current time and UTC time zone
func Now() VBDate {
	x := time.Now()
	return VBDate(time.Date(x.Year(), x.Month(), x.Day(), x.Hour(), x.Minute(), x.Second(), x.Nanosecond(), time.UTC))
}
