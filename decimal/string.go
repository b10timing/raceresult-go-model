package decimal

import (
	"errors"
	"strconv"
	"strings"
)

// FromString creates a new decimal from string, returns 0 if parsing fails
func FromString(s string) (Decimal, error) {
	if strings.HasPrefix(s, "-") {
		x, err := FromString(s[1:])
		return -x, err
	}

	a := strings.Split(s, ".")
	switch len(a) {
	case 1:
		x, err := strconv.ParseInt(a[0], 10, 64)
		if err != nil {
			return 0, err
		}
		return FromInt64(x), nil

	case 2:
		var d Decimal
		if a[0] != "" {
			x, err := strconv.ParseInt(a[0], 10, 64)
			if err != nil {
				return 0, err
			}
			d = FromInt64(x)
		}
		x, err := strconv.Atoi(a[1])
		if err != nil {
			return 0, err
		}
		switch len(a[1]) {
		case 1:
			d += Decimal(x * Decimals / 10)
		case 2:
			d += Decimal(x * Decimals / 100)
		case 3:
			d += Decimal(x * Decimals / 1000)
		case 4:
			d += Decimal(x * Decimals / 10000)
		default:
			x, err = strconv.Atoi(a[1][:4])
			d += Decimal(x)
		}
		return d, nil

	default:
		return 0, errors.New("not a valid number")
	}
}

// ToString converts a decimal to a string with max 4 Decimals.
func (s Decimal) ToString() string {
	var p string
	seconds := int64(s) / Decimals
	decs := int(int64(s) - (seconds * Decimals))
	if s < 0 {
		seconds *= -1
		decs *= -1
		p = "-"
	}
	switch {
	case decs == 0:
		return p + strconv.Itoa(int(seconds))
	case decs < 10:
		return p + strconv.Itoa(int(seconds)) + ".000" + strconv.Itoa(decs)
	case decs < 100:
		return p + strconv.Itoa(int(seconds)) + ".00" + strings.TrimRight(strconv.Itoa(decs), "0")
	case decs < 1000:
		return p + strconv.Itoa(int(seconds)) + ".0" + strings.TrimRight(strconv.Itoa(decs), "0")
	default:
		return p + strconv.Itoa(int(seconds)) + "." + strings.TrimRight(strconv.Itoa(decs), "0")
	}
}

// Format formats the number with the given number of decimals
func (x Decimal) Format(decimals int, deciSep, thousandsSep string) string {
	var neg bool
	if x < 0 {
		neg = true
		x = -x
	}

	// convert to string and fill zeros
	s := strconv.Itoa(int(x))
	l := len(strconv.Itoa(Decimals))
	for len(s) < l {
		s = "0" + s
	}

	// integer part
	k := s[:len(s)-4]
	if thousandsSep != "" {
		for i := len(k) - 3; i > 0; i -= 3 {
			k = k[:i] + thousandsSep + k[i:]
		}
	}

	// decimal part
	m := s[len(s)-4 : len(s)-4+decimals]

	// compile
	s = k + deciSep + m

	if neg {
		return "-" + s
	}
	return s
}
