package variant

// Type defines a type for the different variant implementations.
type Type int

// Constants for the different variant types.
const (
	TypeEmpty Type = iota
	TypeRBool
	TypeRString
	TypeRInt
	TypeRFloat
	TypeRDecimal
	TypeRDate
)

// GetType returns the type of the given variant.
func GetType(v Variant) Type {
	if v == nil {
		return TypeEmpty
	}
	return v.getType()
}

// IsBool checks if given variable is a boolean.
// noinspection GoUnusedExportedFunction
func IsBool(v Variant) bool {
	if v == nil {
		return false
	}
	return v.getType() == TypeRBool
}

// IsString checks if given variable is a string.
func IsString(v Variant) bool {
	if v == nil {
		return false
	}
	return v.getType() == TypeRString
}

// IsInt checks if given variable is an integer.
// noinspection GoUnusedExportedFunction
func IsInt(v Variant) bool {
	if v == nil {
		return false
	}
	return v.getType() == TypeRInt
}

// IsFloat checks if given variable is a float.
// noinspection GoUnusedExportedFunction
func IsFloat(v Variant) bool {
	if v == nil {
		return false
	}
	return v.getType() == TypeRFloat
}

// IsDecimal checks if given variable is a time.
// noinspection GoUnusedExportedFunction
func IsDecimal(v Variant) bool {
	if v == nil {
		return false
	}
	return v.getType() == TypeRDecimal
}

// IsDate checks if given variable is a date.
// noinspection GoUnusedExportedFunction
func IsDate(v Variant) bool {
	if v == nil {
		return false
	}
	return v.getType() == TypeRDate
}

// IsEmpty checks if given variable is empty.
func IsEmpty(v Variant) bool {
	return v == nil
}
