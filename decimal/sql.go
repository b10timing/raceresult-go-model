package decimal

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

func (q Decimal) Value() (driver.Value, error) {
	return q.ToString(), nil
}

func (q *Decimal) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		var err error
		*q, err = FromString(v)
		return err
	case float64:
		*q = FromFloat(v)
		return nil
	case byte:
		var err error
		*q, err = FromString(string(v))
		return err
	case []uint8:
		var err error
		*q, err = FromString(string(v))
		return err
	default:
		return errors.New("incompatible type for Decimal" + fmt.Sprintf("%T", src))
	}
}
