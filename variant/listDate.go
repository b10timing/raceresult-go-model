package variant

import (
	"github.com/raceresult/go-model/date"
)

// DateList is a slice of Date
type DateList []date.Date

// NewDateList creates a new DateList
func NewDateList(size int) DateList {
	return make([]date.Date, size)
}

// ToString converts the list into a StringList
func (s DateList) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rDate(v).toString()
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s DateList) ToStringWithDateFormat(df string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rDate(v).toStringWithDateFormat(df)
	}
	return r
}

// ToInt converts the list into an IntList
func (s DateList) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = rDate(v).toInt()
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s DateList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = rDate(v).toFloat64()
	}
	return r
}

// ToDateTime converts the list into a DateTimeList
func (s DateList) ToDateTime() DateTimeList {
	r := NewDateTimeList(len(s))
	for i, v := range s {
		r[i] = rDate(v).toDateTime()
	}
	return r
}

// ToDate converts the list into a DateList
func (s DateList) ToDate() DateList {
	return s
}

// ToBool converts the list into a BoolList
func (s DateList) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = rDate(v).toBool()
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s DateList) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = rDate(v).toDecimal()
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s DateList) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = rDate(v)
	}
	return r
}

// Item returns an item of the list
func (s DateList) Item(index int) Variant {
	return rDate(s[index])
}

// Len returns the length of the list
func (s DateList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s DateList) Abs() RList {
	return s
}

// Val returns a new list having all values converted into numbers
func (s DateList) Val() RList {
	return s.ToFloat64()
}

// Plus adds the values of another list and returns a new list with the sums
func (s DateList) Plus(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] = s[i].AddDate(0, 0, v[i])
		}
		return s
	case DecimalList:
		return s.ToDateTime().Plus(p)
	case BoolList:
		return s.Plus(v.ToInt())
	case DateList:
		return NewVariantList(len(s)) // empty!
	case DateTimeList:
		return NewVariantList(len(s)) // empty!
	case Float64List:
		return s.ToDateTime().Plus(p)
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDate(s[i]).plus(v.Item(i))
		}
		return result
	case VariantList:
		for i := range s {
			v[i] = rDate(s[i]).plus(v.Item(i))
		}
		return v
	default:
		panic("new type not implemented")
	}
}

// Minus substracts the values of another list and returns a new list result
func (s DateList) Minus(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] = s[i].AddDate(0, 0, -v[i])
		}
		return s
	case DecimalList:
		return s.ToDateTime().Minus(p)
	case BoolList:
		return s.Minus(v.ToInt())
	case DateList:
		result := NewIntList(len(s))
		for i := range s {
			result[i] = int(s[i].Sub(v[i]))
		}
		return result
	case DateTimeList:
		return s.ToDateTime().Minus(p)
	case Float64List:
		return s.ToDateTime().Minus(p)
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDate(s[i]).minus(v.Item(i))
		}
		return result
	case VariantList:
		for i := range s {
			v[i] = rDate(s[i]).minus(v.Item(i))
		}
		return v
	default:
		panic("new type not implemented")
	}
}

// Mult multiplies the values of another list and returns a new list with the result
func (s DateList) Mult(p RList) RList {
	return s.ToFloat64().Mult(p)
}

// Div divides the values of another list and returns a new list with the result
func (s DateList) Div(p RList) RList {
	return s.ToFloat64().Div(p)
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s DateList) DivInt(p RList) RList {
	return s.ToFloat64().DivInt(p)
}

// Mod calculates the remainder of an integer division with the values of another list and returns a new list with the result
func (s DateList) Mod(p RList) RList {
	return s.ToFloat64().Mod(p)
}

// Exp calculates a ^ b
func (s DateList) Exp(p RList) RList {
	return s.ToFloat64().Exp(p)
}
