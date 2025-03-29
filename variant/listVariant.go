package variant

// VariantList is a slice of Variants
//
//goland:noinspection GoNameStartsWithPackageName
type VariantList []Variant

// NewVariantList creates a new VariantList
func NewVariantList(size int) VariantList {
	return make([]Variant, size)
}

// ToString converts the list into a StringList
func (s VariantList) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = ToString(v)
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s VariantList) ToStringWithDateFormat(df string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = ToStringWithDateFormat(v, df)
	}
	return r
}

// ToInt converts the list into an IntList
func (s VariantList) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = ToInt(v)
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s VariantList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = ToFloat64(v)
	}
	return r
}

// ToDateTime converts the list into a DateTimeList
func (s VariantList) ToDateTime() DateTimeList {
	r := NewDateTimeList(len(s))
	for i, v := range s {
		r[i] = ToDateTime(v)
	}
	return r
}

// ToBool converts the list into a BoolList
func (s VariantList) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = ToBool(v)
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s VariantList) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = ToDecimal(v)
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s VariantList) ToVariant() VariantList {
	return s
}

// Item returns an item of the list
func (s VariantList) Item(index int) Variant {
	return s[index]
}

// Len returns a the length of the list
func (s VariantList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s VariantList) Abs() RList {
	result := make(VariantList, len(s))
	for i := range s {
		if s[i] != nil {
			result[i] = s[i].abs()
		}
	}
	return result
}

// Val returns a new list having all values converted into numbers
func (s VariantList) Val() RList {
	result := make(VariantList, len(s))
	for i := range s {
		if s[i] == nil {
			result[i] = RInt(0)
		} else {
			result[i] = s[i].val()
		}
	}
	return result
}

// Plus adds the values of another list and returns a new list with the sums
func (s VariantList) Plus(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].plus(p.Item(i))
	}
	return result
}

// Minus substracts the values of another list and returns a new list result
func (s VariantList) Minus(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].minus(p.Item(i))
	}
	return result
}

// Mult multiplies the values of another list and returns a new list with the result
func (s VariantList) Mult(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].mult(p.Item(i))
	}
	return result
}

// Div divides the values of another list and returns a new list with the result
func (s VariantList) Div(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].div(p.Item(i))
	}
	return result
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s VariantList) DivInt(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].divInt(p.Item(i))
	}
	return result
}

// Mod calculates the remainer of an integer division with the values of another list and returns a new list with the result
func (s VariantList) Mod(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].mod(p.Item(i))
	}
	return result
}

// Exp calculates a ^ b
func (s VariantList) Exp(p RList) RList {
	result := make(VariantList, len(s))
	for i := 0; i < s.Len(); i++ {
		if s[i] == nil {
			continue
		}
		result[i] = s[i].exp(p.Item(i))
	}
	return result
}
