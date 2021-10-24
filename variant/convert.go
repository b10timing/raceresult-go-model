package variant

import (
	"errors"
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/vbdate"
	"strconv"
	"strings"
)

// ToString converts a variant to string.
func ToString(value Variant) string {
	if value == nil {
		return ""
	}
	return value.toString()
}

// ToStringWithDateFormat converts a variant to string with a date format
func ToStringWithDateFormat(value Variant, df string) string {
	if value == nil {
		return ""
	}
	return value.toStringWithDateFormat(df)
}

// ToBool converts a variant to bool.
func ToBool(value Variant) bool {
	if value == nil {
		return false
	}
	return value.toBool()
}

// ToInt converts a variant to int.
func ToInt(value Variant) int {
	if value == nil {
		return 0
	}
	return value.toInt()
}

// ToFloat64 converts a variant to float64
func ToFloat64(value Variant) float64 {
	if value == nil {
		return 0
	}
	return value.toFloat64()
}

// ToDecimal converts a variant to Decimal.
func ToDecimal(value Variant) decimal.Decimal {
	if value == nil {
		return 0
	}
	return value.toDecimal()
}

// ToDate converts a variant to VBDate.
func ToDate(value Variant) vbdate.VBDate {
	if value == nil {
		return vbdate.ZeroDate()
	}
	return value.toDate()
}

// ParseNumber parses a number to int, float or decimal
func ParseNumber(s string) (Variant, error) {
	s2 := strings.ReplaceAll(s, ",", ".")
	a := strings.Split(s2, ".")
	switch len(a) {
	case 1:
		x, err := strconv.Atoi(a[0])
		if err != nil {
			return nil, err
		}
		return RInt(x), nil

	case 2:
		if len(a[1]) <= 4 {
			x, err := decimal.FromString(s2)
			if err != nil {
				return nil, err
			}
			return RDecimal(x), nil
		}
		x, err := strconv.ParseFloat(s2, 64)
		if err != nil {
			return nil, err
		}
		return RFloat(x), nil

	default:
		return nil, errors.New("not a number")
	}
}
