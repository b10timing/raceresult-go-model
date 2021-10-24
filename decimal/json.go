package decimal

// MarshalJSON converts the number to JSON
func (s Decimal) MarshalJSON() ([]byte, error) {
	return []byte(s.ToString()), nil
}

// UnmarshalJSON parses a Decimal from JSON
func (s *Decimal) UnmarshalJSON(data []byte) error {
	x, err := FromString(string(data))
	if err != nil {
		return err
	}
	*s = x
	return nil
}
