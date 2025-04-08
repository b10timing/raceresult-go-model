package variant

// IntList is a slice of ints
type IntList []int

// NewIntList creates a new IntList
func NewIntList(size int) IntList {
	return make([]int, size)
}

// ToString converts the list into a StringList
func (s IntList) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rInt(v).toString()
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s IntList) ToStringWithDateFormat(string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rInt(v).toString()
	}
	return r
}

// ToInt converts the list into an IntList
func (s IntList) ToInt() IntList {
	return s
}

// ToFloat64 converts the list into a Float64List
func (s IntList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = rInt(v).toFloat64()
	}
	return r
}

// ToDateTime converts the list into a DateTimeList
func (s IntList) ToDateTime() DateTimeList {
	r := NewDateTimeList(len(s))
	for i, v := range s {
		r[i] = rInt(v).toDateTime()
	}
	return r
}

// ToDate converts the list into a DateList
func (s IntList) ToDate() DateList {
	r := NewDateList(len(s))
	for i, v := range s {
		r[i] = rInt(v).toDate()
	}
	return r
}

// ToBool converts the list into a BoolList
func (s IntList) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = rInt(v).toBool()
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s IntList) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = rInt(v).toDecimal()
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s IntList) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = rInt(v)
	}
	return r
}

// Item returns an item of the list
func (s IntList) Item(index int) Variant {
	return rInt(s[index])
}

// Len returns a the length of the list
func (s IntList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s IntList) Abs() RList {
	for i := range s {
		if s[i] < 0 {
			s[i] *= -1
		}
	}
	return s
}

// Val returns a new list having all values converted into numbers
func (s IntList) Val() RList {
	return s
}

// Plus adds the values of another list and returns a new list with the sums
func (s IntList) Plus(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] += v[i]
		}
		return s
	case DecimalList:
		return s.ToDecimal().Plus(v)
	case BoolList:
		return s.Plus(v.ToInt())
	case DateTimeList, DateList, Float64List, StringList:
		return s.ToFloat64().Plus(v)
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rInt(s[i]).plus(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Minus substracts the values of another list and returns a new list result
func (s IntList) Minus(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] -= v[i]
		}
		return s
	case DecimalList:
		return s.ToDecimal().Minus(v)
	case BoolList:
		return s.Minus(v.ToInt())
	case DateTimeList, DateList, Float64List, StringList:
		return s.ToFloat64().Minus(v)
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rInt(s[i]).minus(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Mult multiplies the values of another list and returns a new list with the result
func (s IntList) Mult(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			s[i] *= v[i]
		}
		return s
	case DecimalList:
		return s.ToDecimal().Mult(v)
	case BoolList:
		return s.Mult(v.ToInt())
	case DateTimeList, DateList, Float64List, StringList:
		return s.ToFloat64().Mult(v)
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rInt(s[i]).mult(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Div divides the values of another list and returns a new list with the result
func (s IntList) Div(p RList) RList {
	switch v := p.(type) {
	case IntList:
		result := NewFloat64List(len(s))
		for i := range s {
			if v[i] == 0 {
				return s.ToVariant().Div(p)
			}
			result[i] = float64(s[i]) / float64(v[i])
		}
		return result
	case DecimalList:
		return s.ToDecimal().Div(v)
	case BoolList:
		return s.Div(v.ToInt())
	case DateTimeList, DateList, Float64List, StringList:
		return s.ToFloat64().Div(v)
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rInt(s[i]).div(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s IntList) DivInt(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			if v[i] == 0 {
				vl := NewVariantList(len(s))
				for j := range s {
					vl[j] = rInt(s[j])
				}
				return vl.DivInt(p)
			}
			s[i] /= v[i]
		}
		return s
	case DecimalList:
		return s.ToDecimal().DivInt(v)
	case BoolList:
		return s.DivInt(v.ToInt())
	case DateTimeList, DateList, Float64List, StringList:
		return s.ToFloat64().DivInt(v.ToFloat64())
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rInt(s[i]).divInt(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Mod calculates the remainer of an integer division with the values of another list and returns a new list with the result
func (s IntList) Mod(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			if v[i] == 0 {
				vl := NewVariantList(len(s))
				for j := range s {
					vl[j] = rInt(s[j])
				}
				return vl.Mod(p)
			}
			s[i] %= v[i]
		}
		return s
	case DecimalList, BoolList, DateTimeList, DateList, Float64List, StringList:
		return s.Mod(v.ToInt())
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rInt(s[i]).mod(v[i])
		}
		return result
	default:
		panic("new type not implemented")
	}
}

// Exp calculates a ^ b
func (s IntList) Exp(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rInt(s[i]).exp(p.Item(i))
	}
	return result
}

// Min returns the minimum value of the list
func (s IntList) Min() int {
	if len(s) == 0 {
		return 0
	}
	m := s[0]
	for _, x := range s {
		if x < m {
			m = x
		}
	}
	return m
}

// Max returns the maximum value of the list
func (s IntList) Max() int {
	if len(s) == 0 {
		return 0
	}
	m := s[0]
	for _, x := range s {
		if x > m {
			m = x
		}
	}
	return m
}

// IsAllTheSame returns true if all values in the list are identical
func (s IntList) IsAllTheSame() (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	m := s[0]
	for _, x := range s {
		if x != m {
			return 0, false
		}
	}
	return m, true
}
