package emailtemplate

import (
	"encoding/xml"
	"strings"
)

type EmailTemplate struct {
	Name                       string
	Type                       Type
	Sender                     string
	SenderName                 string
	CC                         string
	BCC                        string
	HTML                       bool
	Subject                    string
	Text                       string
	Header                     string
	Footer                     string
	DefaultFilter              string
	SetCustomFieldAfterSending string
	SaveResultIn               string
	Attachments                []Attachment `xml:"attachment"`
}

type Type int

const (
	TypeSingle     Type = 0
	TypeGroup      Type = 1
	TypeSMS        Type = 2
	TypeWebService Type = 3
)

type Attachment struct {
	Type   AttachmentType `xml:"type"`
	Name   string         `xml:"name"`
	Label  string         `xml:"label"`
	Filter string         `xml:"filter"`
}

type AttachmentType int

const (
	AttachmentTypeFile        AttachmentType = 0
	AttachmentTypeCertificate AttachmentType = 1
	AttachmentTypeURL         AttachmentType = 2
)

// UnmarshalXML parses a AttachmentType from XML
func (q *AttachmentType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	switch strings.ToLower(str) {
	case "file":
		*q = AttachmentTypeFile
	case "url":
		*q = AttachmentTypeURL
	default:
		*q = AttachmentTypeCertificate
	}
	return nil
}
