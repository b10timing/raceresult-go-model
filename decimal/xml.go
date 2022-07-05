package decimal

import "encoding/xml"

// UnmarshalXML parses an XML string
func (s *Decimal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	x, err := FromString(str)
	if err != nil {
		return err
	}
	*s = x
	return nil
}

// MarshalXML marshall the value into an xml tag
func (s *Decimal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(s.ToString(), start)
}
