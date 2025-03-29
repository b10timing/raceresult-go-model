package variant

// BoolList is a slice of bools
type BoolList []bool

// NewBoolList creates a new BoolList
func NewBoolList(size int) BoolList {
	return make([]bool, size)
}

// ToString converts the list into a StringList
func (s BoolList) ToString() StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rBool(v).toString()
	}
	return r
}

// ToStringWithDateFormat converts the list into a StringList using a certain date format
func (s BoolList) ToStringWithDateFormat(string) StringList {
	r := NewStringList(len(s))
	for i, v := range s {
		r[i] = rBool(v).toString()
	}
	return r
}

// ToInt converts the list into an IntList
func (s BoolList) ToInt() IntList {
	r := NewIntList(len(s))
	for i, v := range s {
		r[i] = rBool(v).toInt()
	}
	return r
}

// ToFloat64 converts the list into a Float64List
func (s BoolList) ToFloat64() Float64List {
	r := NewFloat64List(len(s))
	for i, v := range s {
		r[i] = rBool(v).toFloat64()
	}
	return r
}

// ToDateTime converts the list into a DateTimeList
func (s BoolList) ToDateTime() DateTimeList {
	r := NewDateTimeList(len(s))
	for i, v := range s {
		r[i] = rBool(v).toDateTime()
	}
	return r
}

// ToBool converts the list into a BoolList
func (s BoolList) ToBool() BoolList {
	return s
}

// ToDecimal converts the list into a DecimalList
func (s BoolList) ToDecimal() DecimalList {
	r := NewDecimalList(len(s))
	for i, v := range s {
		r[i] = rBool(v).toDecimal()
	}
	return r
}

// ToVariant converts the list into a VariantList
func (s BoolList) ToVariant() VariantList {
	r := NewVariantList(len(s))
	for i, v := range s {
		r[i] = rBool(v)
	}
	return r
}

// Item returns an item of the list
func (s BoolList) Item(index int) Variant {
	return rBool(s[index])
}

// Len returns a the length of the list
func (s BoolList) Len() int {
	return len(s)
}

// Abs returns a new list with the absolute values
func (s BoolList) Abs() RList {
	return s
}

// Val returns a new list having all values converted into numbers
func (s BoolList) Val() RList {
	return NewIntList(len(s))
}

// Plus adds the values of another list and returns a new list with the sums
func (s BoolList) Plus(p RList) RList {
	return s.ToInt().Plus(p)
}

// Minus substracts the values of another list and returns a new list result
func (s BoolList) Minus(p RList) RList {
	return s.ToInt().Minus(p)
}

// Mult multiplies the values of another list and returns a new list with the result
func (s BoolList) Mult(p RList) RList {
	return s.ToInt().Mult(p)
}

// Div divides the values of another list and returns a new list with the result
func (s BoolList) Div(p RList) RList {
	return s.ToInt().Div(p)
}

// DivInt performs integer division with the values of another list and returns a new list with the result
func (s BoolList) DivInt(p RList) RList {
	return s.ToInt().DivInt(p)
}

// Mod calculates the remainer of an integer division with the values of another list and returns a new list with the result
func (s BoolList) Mod(p RList) RList {
	return s.ToInt().Mod(p)
}

// Exp calculates a ^ b
func (s BoolList) Exp(p RList) RList {
	return s.ToInt().Exp(p)
}
