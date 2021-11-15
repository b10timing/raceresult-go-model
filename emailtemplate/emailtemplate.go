package emailtemplate

type EmailTemplate struct {
	Name                       string
	Type                       Type
	Sender                     string
	SenderName                 string
	CC                         string
	BCC                        string
	ReceiverField              string
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
