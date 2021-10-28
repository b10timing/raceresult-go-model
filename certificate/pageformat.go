package certificate

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

// PageFormat distinguishes between Portrait and Landscape mode
type PageFormat int

const (
	PFPortrait  PageFormat = 0
	PFLandscape PageFormat = 1
)

// UnmarshalJSON reads a PageFormat from JSON
func (q *PageFormat) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	q.fromString(s)
	return nil
}

// UnmarshalXML parses a PageFormat from XML
func (q *PageFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	q.fromString(str)
	return nil
}

func (q *PageFormat) fromString(s string) {
	switch strings.ToLower(s) {
	case "portrait":
		*q = PFPortrait
	case "landscape":
		*q = PFLandscape
	}
}

// MarshalJSON saves a PageFormat in JSON
func (q PageFormat) MarshalJSON() ([]byte, error) {
	var s string
	switch q {
	case PFPortrait:
		s = "Portrait"
	case PFLandscape:
		s = "Landscape"
	}
	return json.Marshal(s)
}
