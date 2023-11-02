package variant

import (
	"github.com/raceresult/go-model/decimal"
)

// DecimalList is a slice of Decimal values
type DecimalList []decimal.Decimal

// NewDecimalList creates a new DecimalList
func NewDecimalList(size int) DecimalList {
	return make([]decimal.Decimal, size)
}

// ToString converts the list into a StringList
func (s DecimalList) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = RDecimal(v).toString()
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s DecimalList) ToStringWithDateFormat(string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = RDecimal(v).toString()
	}
	return r
}

// ToInt converts the list into an IntList
func (s DecimalList) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = RDecimal(v).toInt()
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s DecimalList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = RDecimal(v).toFloat64()
	}
	return r
}

// ToDate converts the list into a DateList
func (s DecimalList) ToDate() DateList {
	r := NewDateList(len(s))
	for i, v := range s {
		r[i] = RDecimal(v).toDate()
	}
	return r
}

// ToBool converts the list into a BoolList
func (s DecimalList) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = RDecimal(v).toBool()
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s DecimalList) ToDecimal() DecimalList {
	return s
}

// ToVariant converts the list into a VariantList
func (s DecimalList) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = RDecimal(v)
	}
	return r
}

// Item returns an item of the list
func (s DecimalList) Item(index int) Variant {
	return RDecimal(s[index])
}

// Len returns a the length of the list
func (s DecimalList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s DecimalList) Abs() RList {
	for i := range s {
		if s[i] < 0 {
			s[i] *= -1
		}
	}
	return s
}

// Val returns a new list having all values converted into numbers
func (s DecimalList) Val() RList {
	return s
}

// Plus adds the values of another list and returns a new list with the sums
func (s DecimalList) Plus(p RList) RList {
	switch v := p.(type) {
	case DecimalList:
		for i := range s {
			s[i] += v[i]
		}
		return s
	case Float64List:
		return s.ToFloat64().Plus(p)
	case BoolList, IntList, DateList:
		return s.Plus(v.ToDecimal())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).plus(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rDecimal(s[i]).plus(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Minus substracts the values of another list and returns a new list result
func (s DecimalList) Minus(p RList) RList {
	switch v := p.(type) {
	case DecimalList:
		for i := range s {
			s[i] -= v[i]
		}
		return s
	case Float64List:
		return s.ToFloat64().Minus(p)
	case IntList, BoolList, DateList:
		return s.Minus(v.ToDecimal())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).minus(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rDecimal(s[i]).minus(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Mult multiplies the values of another list and returns a new list with the result
func (s DecimalList) Mult(p RList) RList {
	switch v := p.(type) {
	case DecimalList:
		for i := range s {
			s[i] = s[i] * v[i] / decimal.Decimals
		}
		return s
	case Float64List:
		return s.ToFloat64().Mult(p)
	case IntList, BoolList, DateList:
		return s.Mult(v.ToDecimal())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).mult(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rDecimal(s[i]).mult(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Div divides the values of another list and returns a new list with the result
func (s DecimalList) Div(p RList) RList {
	switch v := p.(type) {
	case Float64List:
		result := NewFloat64List(len(s))
		for i := range s {
			if v[i] == 0 {
				return s.ToVariant().Div(p)
			}
			result[i] = s[i].Div(v[i])
		}
		return result
	case IntList, BoolList, DateList:
		return s.Div(v.ToFloat64())
	case DecimalList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).div(v.Item(i))
		}
		return result
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).div(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rDecimal(s[i]).div(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s DecimalList) DivInt(p RList) RList {
	switch v := p.(type) {
	case DecimalList:
		for i := range s {
			if v[i] == 0 {
				vl := NewVariantList(len(s))
				for j := range s {
					vl[j] = RDecimal(s[j])
				}
				return vl.DivInt(p)
			}
			x := s[i] / v[i]
			s[i] = decimal.Decimals * x
		}
		return s
	case IntList, BoolList, DateList, Float64List:
		return s.DivInt(v.ToDecimal())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).divInt(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rDecimal(s[i]).divInt(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Mod calculates the remainer of an integer division with the values of another list and returns a new list with the result
func (s DecimalList) Mod(p RList) RList {
	switch v := p.(type) {
	case DecimalList:
		for i := range s {
			if v[i] == 0 {
				vl := NewVariantList(len(s))
				for j := range s {
					vl[j] = RDecimal(s[j])
				}
				return vl.Mod(p)
			}
			s[i] %= v[i]
		}
		return s
	case IntList, BoolList, DateList, Float64List:
		return s.Mod(v.ToDecimal())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rDecimal(s[i]).mod(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rDecimal(s[i]).mod(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Exp calculates a ^ b
func (s DecimalList) Exp(p RList) RList {
	return s.ToFloat64().Exp(p)
}
