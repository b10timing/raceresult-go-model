package variant

// StringList is a slice of strings
type StringList []string

// NewStringList creates a new StringList
func NewStringList(size int) StringList {
	return make([]string, size)
}

// ToString converts the list into a StringList
func (s StringList) ToString() StringList {
	return s
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s StringList) ToStringWithDateFormat(string) StringList {
	return s
}

// ToInt converts the list into an IntList
func (s StringList) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = rString(v).toInt()
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s StringList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = rString(v).toFloat64()
	}
	return r
}

// ToDateTime converts the list into a DateTimeList
func (s StringList) ToDateTime() DateTimeList {
	r := NewDateTimeList(len(s))
	for i, v := range s {
		r[i] = rString(v).toDateTime()
	}
	return r
}

// ToDate converts the list into a DateList
func (s StringList) ToDate() DateList {
	r := NewDateList(len(s))
	for i, v := range s {
		r[i] = rString(v).toDate()
	}
	return r
}

// ToBool converts the list into a BoolList
func (s StringList) ToBool() BoolList {
	r := NewBoolList(len(s))
	for i, v := range s {
		r[i] = rString(v).toBool()
	}
	return r
}

// ToDecimal converts the list into a DecimalList
func (s StringList) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = rString(v).toDecimal()
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s StringList) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = rString(v)
	}
	return r
}

// Item returns an item of the list
func (s StringList) Item(index int) Variant {
	return rString(s[index])
}

// Len returns a the length of the list
func (s StringList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s StringList) Abs() RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).abs()
	}
	return result
}

// Val returns a new list having all values converted into numbers
func (s StringList) Val() RList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = RString(v).val()
	}
	return r
}

// Plus adds the values of another list and returns a new list with the sums
func (s StringList) Plus(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).plus(p.Item(i))
	}
	return result
}

// Minus substracts the values of another list and returns a new list result
func (s StringList) Minus(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).minus(p.Item(i))
	}
	return result
}

// Mult multiplies the values of another list and returns a new list with the result
func (s StringList) Mult(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).mult(p.Item(i))
	}
	return result
}

// Div divides the values of another list and returns a new list with the result
func (s StringList) Div(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).div(p.Item(i))
	}
	return result
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s StringList) DivInt(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).divInt(p.Item(i))
	}
	return result
}

// Mod calculates the remainer of an integer division with the values of another list and returns a new list with the result
func (s StringList) Mod(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).mod(p.Item(i))
	}
	return result
}

// Exp calculates a ^ b
func (s StringList) Exp(p RList) RList {
	result := NewVariantList(len(s))
	for i := range result {
		result[i] = rString(s[i]).exp(p.Item(i))
	}
	return result
}

// IsAllTheSame returns true if all values in the list are identical
func (s StringList) IsAllTheSame() (string, bool) {
	if len(s) == 0 {
		return "", false
	}
	m := s[0]
	for _, x := range s {
		if x != m {
			return "", false
		}
	}
	return m, true
}
