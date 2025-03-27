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
type VBDate struct {
	time.Time
	hasZone bool
}

func FromTime(t time.Time, hasZone bool) VBDate {
	return VBDate{Time: t, hasZone: hasZone}
}

// ZeroDate returns the VB zero date 1899-12-30
func ZeroDate() VBDate {
	return Date(1899, 12, 30)
}

// Before is equal to time.Time
func (s VBDate) Before(c VBDate) bool {
	switch {
	case s.hasZone && !c.hasZone:
		return s.Time.Before(c.WithTimezone(s.Time.Location()).Time)
	case !s.hasZone && c.hasZone:
		return s.WithTimezone(c.Time.Location()).Time.Before(c.Time)
	default:
		return s.Time.Before(c.Time)
	}
}

// After is equal to time.Time
func (s VBDate) After(c VBDate) bool {
	switch {
	case s.hasZone && !c.hasZone:
		return s.Time.After(c.WithTimezone(s.Time.Location()).Time)
	case !s.hasZone && c.hasZone:
		return s.WithTimezone(c.Time.Location()).Time.After(c.Time)
	default:
		return s.Time.After(c.Time)
	}
}

// AddDate is equal to time.Time
func (s VBDate) AddDate(year int, month int, day int) VBDate {
	return VBDate{
		Time:    s.Time.AddDate(year, month, day),
		hasZone: s.hasZone,
	}
}

// Add is equal to time.Time
func (s VBDate) Add(d time.Duration) VBDate {
	return VBDate{
		Time:    s.Time.Add(d),
		hasZone: s.hasZone,
	}
}

// Sub is equal to time.Time
func (s VBDate) Sub(d VBDate) time.Duration {
	return s.Time.Sub(d.Time)
}

// ToTime returns the VBDate as time.Time
func (s VBDate) ToTime() time.Time {
	if s.IsZero() {
		return time.Time{}
	}
	return s.Time
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
			*s = VBDate{Time: d, hasZone: false}
		}
		return err
	case 21:
		d, err := time.ParseInLocation(dateTimeFormat, str[1:len(str)-1], time.UTC)
		if err == nil {
			*s = VBDate{Time: d, hasZone: false}
		}
		return err
	case 22, 27:
		d, err := time.Parse(time.RFC3339, str[1:len(str)-1])
		if err == nil {
			*s = VBDate{Time: d, hasZone: true}
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
	if s.Time.Equal(ZeroDate().Time) {
		return true
	}
	return s.Time.IsZero()
}

// ToString converts the date to string
func (s VBDate) ToString() string {
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
func (s VBDate) ToStringWithDateFormat(df string) string {
	if df == "" {
		return s.ToString()
	}
	if s.IsZero() {
		return ""
	}
	if s.Time.Unix()%86400 == 0 {
		return s.Time.Format(df)
	}
	return s.Time.Format(df + " " + timeFormat)
}

// WithTimezone returns a new VBDate where the timezone was replaced
func (s VBDate) WithTimezone(tz *time.Location) VBDate {
	if s.IsZero() {
		return s
	}
	return VBDate{
		Time:    time.Date(s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Second(), s.Nanosecond(), tz),
		hasZone: true,
	}
}

// Parse parses a date from string
func Parse(str string) (VBDate, bool) {
	if strings.Count(str, ".") == 2 {
		switch len(str) {
		case 10:
			d, err := time.ParseInLocation("02.01.2006", str, time.UTC)
			if err == nil {
				return VBDate{Time: d, hasZone: false}, true
			}
		case 19:
			d, err := time.ParseInLocation("02.01.2006 "+timeFormat, str, time.UTC)
			if err == nil {
				return VBDate{Time: d, hasZone: false}, true
			}
		}
		return ZeroDate(), false
	}
	if strings.Count(str, "/") == 2 {
		switch len(str) {
		case 10:
			d, err := time.ParseInLocation("02/01/2006", str, time.UTC)
			if err == nil {
				return VBDate{Time: d, hasZone: false}, true
			}
		case 19:
			d, err := time.ParseInLocation("02/01/2006 "+timeFormat, str, time.UTC)
			if err == nil {
				return VBDate{Time: d, hasZone: false}, true
			}
		}
		return ZeroDate(), false
	}

	switch len(str) {
	case 10:
		d, err := time.ParseInLocation(dateFormat, str, time.UTC)
		if err == nil {
			return VBDate{Time: d, hasZone: false}, true
		}
	case 19:
		d, err := time.ParseInLocation(dateTimeFormat, str, time.UTC)
		if err == nil {
			return VBDate{Time: d, hasZone: false}, true
		}
	case 20, 25:
		d, err := time.Parse(time.RFC3339, str)
		if err == nil {
			return VBDate{Time: d, hasZone: true}, true
		}
	}
	return ZeroDate(), false
}

// HasZone returns true if the date has a time zone
func (s VBDate) HasZone() bool {
	return s.hasZone
}

// Date creates a VB Date with day accuracy
func Date(year int, month time.Month, day int) VBDate {
	return VBDate{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC), hasZone: false}
}

// DateTime creates a VB Date with time accuracy
func DateTime(year int, month time.Month, day int, hour, minute, second int) VBDate {
	return VBDate{Time: time.Date(year, month, day, hour, minute, second, 0, time.UTC), hasZone: false}
}

// Now returns a VBDate with the current time without timezone
func Now(withZone bool) VBDate {
	if withZone {
		return VBDate{Time: time.Now(), hasZone: true}
	}

	x := time.Now()
	return VBDate{Time: time.Date(x.Year(), x.Month(), x.Day(), x.Hour(), x.Minute(), x.Second(), x.Nanosecond(), time.UTC), hasZone: false}
}
