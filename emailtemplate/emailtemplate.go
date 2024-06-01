package emailtemplate

type HTTPHeader struct {
	Name  string
	Value string
}

type EmailTemplate struct {
	Name                       string
	Type                       Type
	Sender                     string
	SenderName                 string
	ReplyTo                    string
	CC                         string
	BCC                        string
	ReceiverField              string
	HTML                       bool
	Method                     string
	Subject                    string
	Text                       string
	Header                     string
	Footer                     string
	DefaultFilter              string
	SetCustomFieldAfterSending string
	SaveResultIn               string
	Attachments                []Attachment `xml:"attachment"`
	HTTPHeaders                []HTTPHeader
}

type Type int

const (
	TypeSingle     Type = 0
	TypeGroup      Type = 1
	TypeSMS        Type = 2
	TypeWebService Type = 3
	TypeGroupByID  Type = 4
)
