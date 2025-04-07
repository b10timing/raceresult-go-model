package datetime

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"
const dateTimeFormat = "2006-01-02 15:04:05"
const timeFormat = "15:04:05"

// DateTime is identical to time.Time with some VB characteristics
type DateTime struct {
	time.Time
	hasZone bool
}

func FromTime(t time.Time, hasZone bool) DateTime {
	return DateTime{Time: t, hasZone: hasZone}
}

// ZeroDate returns the VB zero date 1899-12-30
func ZeroDate() DateTime {
	return New(1899, 12, 30, 0, 0, 0)
}

// Before is equal to time.Time
func (s DateTime) Before(c DateTime) bool {
	switch {
	case s.hasZone && !c.hasZone:
		return s.Time.Before(c.WithTimezone(s.Time.Location()).Time)
	case !s.hasZone && c.hasZone:
		return s.WithTimezone(c.Time.Location()).Time.Before(c.Time)
	case !s.hasZone && !c.hasZone:
		return s.WithTimezone(time.UTC).Time.Before(c.WithTimezone(time.UTC).Time)
	default:
		return s.Time.Before(c.Time)
	}
}

// After is equal to time.Time
func (s DateTime) After(c DateTime) bool {
	switch {
	case s.hasZone && !c.hasZone:
		return s.Time.After(c.WithTimezone(s.Time.Location()).Time)
	case !s.hasZone && c.hasZone:
		return s.WithTimezone(c.Time.Location()).Time.After(c.Time)
	case !s.hasZone && !c.hasZone:
		return s.WithTimezone(time.UTC).Time.After(c.WithTimezone(time.UTC).Time)
	default:
		return s.Time.After(c.Time)
	}
}

// AddDate is equal to time.Time
func (s DateTime) AddDate(year int, month int, day int) DateTime {
	return DateTime{
		Time:    s.Time.AddDate(year, month, day),
		hasZone: s.hasZone,
	}
}

// Add is equal to time.Time
func (s DateTime) Add(d time.Duration) DateTime {
	return DateTime{
		Time:    s.Time.Add(d),
		hasZone: s.hasZone,
	}
}

// Sub is equal to time.Time
func (s DateTime) Sub(d DateTime) time.Duration {
	return s.Time.Sub(d.Time)
}

// ToTime returns the DateTime as time.Time
func (s DateTime) ToTime() time.Time {
	if s.IsZero() {
		return time.Time{}
	}
	return s.Time
}

// MarshalJSON creates a JSON string of the date
func (s DateTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.ToString() + "\""), nil
}

// UnmarshalJSON parses a JSON string to date
func (s *DateTime) UnmarshalJSON(data []byte) error {
	str := string(data)

	switch len(str) {
	case 12:
		d, err := time.ParseInLocation(dateFormat, str[1:len(str)-1], time.UTC)
		if err == nil {
			*s = DateTime{Time: d, hasZone: false}
		}
		return err
	case 21:
		d, err := time.ParseInLocation(dateTimeFormat, str[1:len(str)-1], time.UTC)
		if err == nil {
			*s = DateTime{Time: d, hasZone: false}
		}
		return err
	case 22, 27:
		d, err := time.Parse(time.RFC3339, str[1:len(str)-1])
		if err == nil {
			*s = DateTime{Time: d, hasZone: true}
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
func (s *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (s DateTime) IsZero() bool {
	if s.Time.Equal(ZeroDate().Time) {
		return true
	}
	return s.Time.IsZero()
}

// ToString converts the date to string
func (s DateTime) ToString() string {
	if s.IsZero() {
		return ""
	}

	if s.hasZone {
		return s.Time.Format(time.RFC3339)
	}
	if s.Time.Hour() == 0 && s.Time.Minute() == 0 && s.Time.Second() == 0 {
		return s.Time.Format(dateFormat)
	}
	return s.Time.Format(dateTimeFormat)
}

// ToStringWithDateFormat converts the date to string with the given date format
func (s DateTime) ToStringWithDateFormat(df string) string {
	if s.IsZero() {
		return ""
	}
	if df == "" {
		return s.ToString()
	}
	if s.Time.Unix()%86400 == 0 {
		return s.Time.Format(df)
	}
	return s.Time.Format(df + " " + timeFormat)
}

// WithTimezone returns a new DateTime where the timezone was replaced
func (s DateTime) WithTimezone(tz *time.Location) DateTime {
	if s.IsZero() {
		return s
	}
	if s.hasZone {
		return FromTime(s.Time.In(tz), true)
	}
	return DateTime{
		Time:    time.Date(s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Second(), s.Nanosecond(), tz),
		hasZone: true,
	}
}

// Round is equal to time.Time but returns DateTime
func (s DateTime) Round(d time.Duration) DateTime {
	return DateTime{
		Time:    s.Time.Round(d),
		hasZone: s.hasZone,
	}
}

// Parse parses a date from string
func Parse(str string) (DateTime, bool) {
	if strings.Count(str, ".") == 2 {
		switch len(str) {
		case 10:
			d, err := time.ParseInLocation("02.01.2006", str, time.UTC)
			if err == nil {
				return DateTime{Time: d, hasZone: false}, true
			}
		case 19:
			d, err := time.ParseInLocation("02.01.2006 "+timeFormat, str, time.UTC)
			if err == nil {
				return DateTime{Time: d, hasZone: false}, true
			}
		}
		return ZeroDate(), false
	}
	if strings.Count(str, "/") == 2 {
		switch len(str) {
		case 10:
			d, err := time.ParseInLocation("02/01/2006", str, time.UTC)
			if err == nil {
				return DateTime{Time: d, hasZone: false}, true
			}
		case 19:
			d, err := time.ParseInLocation("02/01/2006 "+timeFormat, str, time.UTC)
			if err == nil {
				return DateTime{Time: d, hasZone: false}, true
			}
		}
		return ZeroDate(), false
	}

	switch len(str) {
	case 10:
		d, err := time.ParseInLocation(dateFormat, str, time.UTC)
		if err == nil {
			return DateTime{Time: d, hasZone: false}, true
		}
	case 19:
		d, err := time.ParseInLocation(dateTimeFormat, str, time.UTC)
		if err == nil {
			return DateTime{Time: d, hasZone: false}, true
		}
	case 20, 25:
		d, err := time.Parse(time.RFC3339, str)
		if err == nil {
			return DateTime{Time: d, hasZone: true}, true
		}
	}
	return ZeroDate(), false
}

// HasZone returns true if the date has a time zone
func (s DateTime) HasZone() bool {
	return s.hasZone
}

// New creates a DateTime with time accuracy
func New(year int, month time.Month, day int, hour, minute, second int) DateTime {
	return DateTime{Time: time.Date(year, month, day, hour, minute, second, 0, time.UTC), hasZone: false}
}

// Now returns a DateTime with the current time without timezone
func Now(withZone bool) DateTime {
	if withZone {
		return DateTime{Time: time.Now(), hasZone: true}
	}

	x := time.Now()
	return DateTime{Time: time.Date(x.Year(), x.Month(), x.Day(), x.Hour(), x.Minute(), x.Second(), x.Nanosecond(), time.UTC), hasZone: false}
}
