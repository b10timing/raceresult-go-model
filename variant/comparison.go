package variant

import (
	"golang.org/x/text/collate"
)

// Equals implements v1-equals-v2 for Variant types.
func Equals(v1 Variant, v2 Variant, caseSensitive bool) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 != nil {
		return v1.equals(v2, caseSensitive)
	}
	return v2.equals(v1, caseSensitive)
}

// NotEquals implements not-equal for Variant types.
func NotEquals(v1 Variant, v2 Variant, caseSensitive bool) bool {
	return !Equals(v1, v2, caseSensitive)
}

// Less implements v1-less-v2 for Variant types.
func Less(v1 Variant, v2 Variant, collator *collate.Collator) bool {
	if v1 == nil && v2 == nil {
		return false
	}
	if v1 != nil {
		return v1.less(v2, collator)
	}
	return v2.greater(v1, collator)
}

// Greater implements greater-than for Variant types.
func Greater(v1 Variant, v2 Variant, collator *collate.Collator) bool {
	if v1 == nil && v2 == nil {
		return false
	}
	if v1 != nil {
		return v1.greater(v2, collator)
	}
	return v2.less(v1, collator)
}

// LessOrEquals implements less-or-equal for Variant types.
func LessOrEquals(v1 Variant, v2 Variant, collator *collate.Collator) bool {
	return !Greater(v1, v2, collator)
}

// GreaterOrEquals implements greater-or-equal for Variant types
func GreaterOrEquals(v1 Variant, v2 Variant, collator *collate.Collator) bool {
	return !Less(v1, v2, collator)
}
