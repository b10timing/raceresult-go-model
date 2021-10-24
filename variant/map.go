package variant

import (
	"encoding/json"
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
