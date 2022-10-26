package certificate

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/raceresult/go-model/decimal"
)

// PageSize defines the page size of the certificate
type PageSize int

const (
	PSA1          PageSize = 0
	PSA2          PageSize = 1
	PSA3          PageSize = 2
	PSA4          PageSize = 3
	PSA5          PageSize = 4
	PSLetter      PageSize = 5
	PSLegal       PageSize = 6
	PSA6          PageSize = 7
	PSUserDefined PageSize = 8
)

// UnmarshalJSON parses a PageSize from JSON
func (q *PageSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	q.fromString(s)
	return nil
}

// UnmarshalXML parses a PageSize from XML
func (q *PageSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	q.fromString(str)
	return nil
}

func (q *PageSize) fromString(s string) {
	switch strings.ToUpper(s) {
	case "A1":
		*q = PSA1
	case "A2":
		*q = PSA2
	case "A3":
		*q = PSA3
	case "A4":
		*q = PSA4
	case "A5":
		*q = PSA5
	case "A6":
		*q = PSA6
	case "LEGAL":
		*q = PSLegal
	case "LETTER":
		*q = PSLetter
	default:
		*q = PSUserDefined
	}
}

// MarshalJSON saves a PageSize in JSON
func (q PageSize) MarshalJSON() ([]byte, error) {
	var s string
	switch q {
	case PSA1:
		s = "A1"
	case PSA2:
		s = "A2"
	case PSA3:
		s = "A3"
	case PSA4:
		s = "A4"
	case PSA5:
		s = "A5"
	case PSA6:
		s = "A6"
	case PSLetter:
		s = "Letter"
	case PSLegal:
		s = "Legal"
	default:
		s = "UserDefined"
	}
	return json.Marshal(s)
}

// Height returns the height in mm
func (q PageSize) Height() decimal.Decimal {
	switch q {
	case PSA1:
		return decimal.FromInt(841)
	case PSA2:
		return decimal.FromInt(594)
	case PSA3:
		return decimal.FromInt(420)
	case PSA4:
		return decimal.FromInt(297)
	case PSA5:
		return decimal.FromInt(210)
	case PSA6:
		return decimal.FromInt(148)
	case PSLetter:
		return decimal.FromFloat(11 * 25.4)
	case PSLegal:
		return decimal.FromFloat(14 * 25.4)
	default:
		return 0
	}
}

// Width returns the width in mm
func (q PageSize) Width() decimal.Decimal {
	switch q {
	case PSA1:
		return decimal.FromInt(594)
	case PSA2:
		return decimal.FromInt(420)
	case PSA3:
		return decimal.FromInt(297)
	case PSA4:
		return decimal.FromInt(210)
	case PSA5:
		return decimal.FromInt(148)
	case PSA6:
		return decimal.FromInt(105)
	case PSLetter:
		return decimal.FromFloat(8.5 * 25.4)
	case PSLegal:
		return decimal.FromFloat(8.5 * 25.4)
	default:
		return 0
	}
}
