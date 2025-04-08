package variant

import (
	"github.com/raceresult/go-model/datetime"

	"time"
)

// DateTimeList is a slice of DateTime
type DateTimeList []datetime.DateTime

// NewDateTimeList creates a new DateTimeList
func NewDateTimeList(size int) DateTimeList {
	return make([]datetime.DateTime, size)
}

// ToString converts the list into a StringList
func (s DateTimeList) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toString()
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s DateTimeList) ToStringWithDateFormat(df string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toStringWithDateFormat(df)
	}
	return r
}

// ToInt converts the list into an IntList
func (s DateTimeList) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toInt()
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s DateTimeList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toFloat64()
	}
	return r
}

// ToDate converts the list into a DateList
func (s DateTimeList) ToDate() DateList {
	r := NewDateList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toDate()
	}
	return r
}

// ToDateTime converts the list into a DateTimeList
func (s DateTimeList) ToDateTime() DateTimeList {
	return s
}

// ToBool converts the list into a BoolList
func (s DateTimeList) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toBool()
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s DateTimeList) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v).toDecimal()
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s DateTimeList) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = rDateTime(v)
	}
	return r
}

// Item returns an item of the list
func (s DateTimeList) Item(index int) Variant {
	return rDateTime(s[index])
}

// Len returns the length of the list
func (s DateTimeList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s DateTimeList) Abs() RList {
	return s
}

// Val returns a new list having all values converted into numbers
func (s DateTimeList) Val() RList {
	return s.ToFloat64()
}

// Plus adds the values of another list and returns a new list with the sums
func (s DateTimeList) Plus(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] = s[i].AddDate(0, 0, v[i])
		}
		return s
	case DecimalList:
		for i := range s {
			s[i] = s[i].Add(time.Duration(v[i] * 100 * 1000))
		}
		return s
	case BoolList:
		return s.Plus(v.ToInt())
	case DateTimeList:
		return NewVariantList(len(s)) // empty!
	case DateList:
		return NewVariantList(len(s)) // empty!
	case Float64List:
		for i := range s {
			s[i] = s[i].Add(time.Duration(v[i] * 86400 * 1000 * 1000 * 1000))
		}
		return s
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDateTime(s[i]).plus(v.Item(i))
		}
		return result
	case VariantList:
		for i := range s {
			v[i] = rDateTime(s[i]).plus(v.Item(i))
		}
		return v
	default:
		panic("new type not implemented")
	}
}

// Minus substracts the values of another list and returns a new list result
func (s DateTimeList) Minus(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] = s[i].AddDate(0, 0, -v[i])
		}
		return s
	case DecimalList:
		for i := range s {
			s[i] = s[i].Add(time.Duration(-v[i] * 100 * 1000))
		}
		return s
	case BoolList:
		return s.Minus(v.ToInt())
	case DateTimeList:
		result := NewFloat64List(len(s))
		for i := range s {
			result[i] = s[i].Sub(v[i]).Hours() / 24
		}
		return result
	case DateList:
		result := NewFloat64List(len(s))
		for i := range s {
			result[i] = s[i].Sub(rDate(v[i]).toDateTime()).Hours() / 24
		}
		return result
	case Float64List:
		for i := range s {
			s[i] = s[i].Add(time.Duration(-v[i] * 86400 * 1000 * 1000 * 1000))
		}
		return s
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDateTime(s[i]).minus(v.Item(i))
		}
		return result
	case VariantList:
		for i := range s {
			v[i] = rDateTime(s[i]).minus(v.Item(i))
		}
		return v
	default:
		panic("new type not implemented")
	}
}

// Mult multiplies the values of another list and returns a new list with the result
func (s DateTimeList) Mult(p RList) RList {
	return s.ToFloat64().Mult(p)
}

// Div divides the values of another list and returns a new list with the result
func (s DateTimeList) Div(p RList) RList {
	return s.ToFloat64().Div(p)
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s DateTimeList) DivInt(p RList) RList {
	return s.ToFloat64().DivInt(p)
}

// Mod calculates the remainder of an integer division with the values of another list and returns a new list with the result
func (s DateTimeList) Mod(p RList) RList {
	return s.ToFloat64().Mod(p)
}

// Exp calculates a ^ b
func (s DateTimeList) Exp(p RList) RList {
	return s.ToFloat64().Exp(p)
}
