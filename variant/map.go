package variant

import (
	"encoding/json"
	"strings"
)

type VariantMap map[string]Variant

// UnmarshalJSON parses a VariantMap from JSON
func (s *VariantMap) UnmarshalJSON(data []byte) error {
	// parse into interface
	isettings := make(map[string]interface{})
	if err := json.Unmarshal(data, &isettings); err != nil {
		return err
	}

	// convert to variant
	*s = make(map[string]Variant)
	for k, v := range isettings {
		(*s)[k] = ToVariant(v)
	}
	return nil
}

// GetItem returns the value with the given key with case-insensitive search
func (s *VariantMap) GetItem(key string) (Variant, bool) {
	if v, ok := (*s)[key]; ok {
		return v, true
	}
	for k, v := range *s {
		if strings.EqualFold(k, key) {
			return v, true
		}
	}
	return nil, false
}
