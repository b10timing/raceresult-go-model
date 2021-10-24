package variant

import (
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/vbdate"
	"golang.org/x/text/collate"

	"time"
)

// RDate creates a date variant type.
func RDate(v vbdate.VBDate) Variant {
	if v.IsZero() {
		v = vbdate.ZeroDate()
	}
	return rDate(v)
}

// rDate implements a date type.
type rDate vbdate.VBDate

// Type returns the type of the variant.
func (s rDate) getType() Type {
	return TypeRDate
}

func (s rDate) equals(v Variant, _ bool) bool {
	if v == nil {
		return s.isZero()
	}
	switch val := v.(type) {
	case rBool:
		return s.toBool() == bool(val)
	case rString:
		return s.toString() == string(val)
	}
	return vbdate.VBDate(s) == v.toDate()
}

func (s rDate) less(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return vbdate.VBDate(s).Before(vbdate.ZeroDate())
	}
	switch val := v.(type) {
	case rString:
		return s.toString() < string(val)
	case rDate:
		return time.Time(s).Before(time.Time(val))
	}
	return vbdate.VBDate(s).Before(v.toDate())
}

func (s rDate) greater(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return vbdate.ZeroDate().Before(vbdate.VBDate(s))
	}
	switch val := v.(type) {
	case rString:
		return string(val) < s.toString()
	case rDate:
		return time.Time(val).Before(time.Time(s))
	}
	return v.toDate().Before(vbdate.VBDate(s))
}

func (s rDate) toString() string {
	return vbdate.VBDate(s).ToString()
}

func (s rDate) toStringWithDateFormat(df string) string {
	return vbdate.VBDate(s).ToStringWithDateFormat(df)
}

func (s rDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.toString() + "\""), nil
}

// ToFloat64 converts the type to float64.
func (s rDate) toFloat64() float64 {
	return vbdate.VBDate(s).Sub(vbdate.ZeroDate()).Hours() / 24
}

func (s rDate) toDate() vbdate.VBDate {
	return vbdate.VBDate(s)
}

func (s rDate) toBool() bool {
	return !time.Time(s).IsZero()
}

func (s rDate) toInt() int {
	return int(s.toFloat64())
}

func (s rDate) toDecimal() decimal.Decimal {
	return decimal.FromFloat(s.toFloat64())
}

func (s rDate) isZero() bool {
	return vbdate.VBDate(s).IsZero()
}

func (s rDate) abs() Variant {
	return s
}
func (s rDate) val() Variant {
	return RFloat(s.toFloat64())
}

func (s rDate) plus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt:
		return rDate(s.toDate().AddDate(0, 0, p.toInt()))
	case TypeRDecimal:
		return rDate(s.toDate().Add(time.Duration(p.toDecimal() * 100 * 1000)))
	case TypeRBool:
		return rDate(s.toDate().AddDate(0, 0, p.toInt()))
	case TypeRDate:
		return nil
	case TypeRFloat:
		return rDate(s.toDate().Add(time.Duration(p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	case TypeRString:
		if !p.isNumeric() {
			return nil
		}
		return rDate(s.toDate().Add(time.Duration(p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	default:
		return nil
	}
}

func (s rDate) minus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt:
		return rDate(s.toDate().AddDate(0, 0, -p.toInt()))
	case TypeRDecimal:
		return rDate(s.toDate().Add(time.Duration(-p.toDecimal() * 100 * 1000)))
	case TypeRBool:
		return rDate(s.toDate().AddDate(0, 0, -p.toInt()))
	case TypeRDate:
		return rFloat(s.toDate().Sub(p.toDate()).Hours() / 24)
	case TypeRFloat:
		return rDate(s.toDate().Add(time.Duration(-p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	case TypeRString:
		if !p.isNumeric() {
			return nil
		}
		return rDate(s.toDate().Add(time.Duration(-p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	default:
		return nil
	}
}

func (s rDate) mult(p Variant) Variant {
	return RFloat(s.toFloat64()).mult(p)
}

func (s rDate) div(p Variant) Variant {
	return RFloat(s.toFloat64()).div(p)
}

func (s rDate) divInt(p Variant) Variant {
	return RFloat(s.toFloat64()).divInt(p)
}

func (s rDate) mod(p Variant) Variant {
	return RFloat(s.toFloat64()).mod(p)
}

func (s rDate) exp(p Variant) Variant {
	return RFloat(s.toFloat64()).exp(p)
}

func (s rDate) isNumeric() bool {
	return true
}

func (s rDate) toJSON() []byte {
	if s.isZero() {
		return []byte("\"\"")
	}
	return []byte("\"" + time.Time(s).Format(time.RFC3339) + "\"")
}
