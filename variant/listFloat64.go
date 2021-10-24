package variant

// Float64List is a slice of float64s
type Float64List []float64

// NewFloat64List creates a new Float64List
func NewFloat64List(size int) Float64List {
	return make([]float64, size)
}

// ToString converts the list into a StringList
func (s Float64List) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rFloat(v).toString()
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s Float64List) ToStringWithDateFormat(string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rFloat(v).toString()
	}
	return r
}

// ToInt converts the list into an IntList
func (s Float64List) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = rFloat(v).toInt()
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s Float64List) ToFloat64() Float64List {
	return s
}

// ToDate converts the list into a DateList
func (s Float64List) ToDate() DateList {
	r := NewDateList(len(s))
	for i, v := range s {
		r[i] = rFloat(v).toDate()
	}
	return r
}

// ToBool converts the list into a BoolList
func (s Float64List) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = rFloat(v).toBool()
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s Float64List) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = rFloat(v).toDecimal()
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s Float64List) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = rFloat(v)
	}
	return r
}

// Item returns an item of the list
func (s Float64List) Item(index int) Variant {
	return rFloat(s[index])
}

// Len returns a the length of the list
func (s Float64List) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s Float64List) Abs() RList {
	for i := range s {
		if s[i] < 0 {
			s[i] *= -1
		}
	}
	return s
}

// Val returns a new list having all values converted into numbers
func (s Float64List) Val() RList {
	return s
}

// Plus adds the values of another list and returns a new list with the sums
func (s Float64List) Plus(p RList) RList {
	switch v := p.(type) {
	case IntList, DecimalList, BoolList, DateList:
		return s.Plus(v.ToFloat64())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rFloat(s[i]).plus(v.Item(i))
		}
		return result
	case Float64List:
		for i := range s {
			s[i] += v[i]
		}
		return s
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rFloat(s[i]).plus(v[i])
		}
		return result
	default:
		panic("New type not implemented")
	}
}

// Minus substracts the values of another list and returns a new list result
func (s Float64List) Minus(p RList) RList {
	switch v := p.(type) {
	case IntList, DecimalList, BoolList, DateList:
		return s.Minus(v.ToFloat64())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rFloat(s[i]).minus(v.Item(i))
		}
		return result
	case Float64List:
		for i := range s {
			s[i] -= v[i]
		}
		return s
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rFloat(s[i]).minus(v[i])
		}
		return result
	default:
		panic("New type not implemented")
	}
}

// Mult multiplies the values of another list and returns a new list with the result
func (s Float64List) Mult(p RList) RList {
	switch v := p.(type) {
	case IntList, DecimalList, BoolList, DateList:
		return s.Mult(v.ToFloat64())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rFloat(s[i]).mult(v.Item(i))
		}
		return result
	case Float64List:
		for i := range s {
			s[i] *= v[i]
		}
		return s
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rFloat(s[i]).mult(v[i])
		}
		return result
	default:
		panic("New type not implemented")
	}
}

// Div divides the values of another list and returns a new list with the result
func (s Float64List) Div(p RList) RList { //nolint:dupl
	switch v := p.(type) {
	case IntList, DecimalList, BoolList, DateList:
		return s.Div(v.ToFloat64())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rFloat(s[i]).div(v.Item(i))
		}
		return result
	case Float64List:
		result := NewFloat64List(len(s))
		for i := range s {
			if v[i] == 0 {
				return s.ToVariant().Div(p)
			}
			result[i] = s[i] / v[i]
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rFloat(s[i]).div(v[i])
		}
		return result
	default:
		panic("New type not implemented")
	}
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s Float64List) DivInt(p RList) RList { //nolint:dupl
	switch v := p.(type) {
	case IntList, DecimalList, BoolList, DateList:
		return s.DivInt(v.ToFloat64())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rFloat(s[i]).divInt(v.Item(i))
		}
		return result
	case Float64List:
		for i := range s {
			if v[i] == 0 {
				vl := NewVariantList(len(s))
				for j := range s {
					vl[j] = RFloat(s[j])
				}
				return vl.DivInt(p)
			}
			s[i] /= v[i]
		}
		return s
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rFloat(s[i]).divInt(v[i])
		}
		return result
	default:
		panic("New type not implemented")
	}
}

// Mod calculates the remainer of an integer division with the values of another list and returns a new list with the result
func (s Float64List) Mod(p RList) RList {
	switch v := p.(type) {
	case IntList:
		for i := range s {
			if v[i] == 0 {
				vl := NewVariantList(len(s))
				for j := range s {
					vl[j] = RFloat(s[j])
				}
				return vl.Mod(p)
			}

			k := int(s[i])
			z := s[i] - float64(k)
			k %= v[i]
			s[i] = float64(k) + z
		}
		return s
	case DecimalList, BoolList, DateList, Float64List:
		return s.Mod(v.ToInt())
	case StringList:
		result := NewVariantList(len(s))
		for i := range s {
			result[i] = rFloat(s[i]).mod(v.Item(i))
		}
		return result
	case VariantList:
		result := NewVariantList(len(s))
		for i := range result {
			result[i] = rFloat(s[i]).mod(v[i])
		}
		return result
	default:
		panic("New type not implemented")
	}
}

// Exp calculates a ^ b
func (s Float64List) Exp(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rFloat(s[i]).exp(p.Item(i))
	}
	return result
}
