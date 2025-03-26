package emailtemplate

type Preview struct {
	Type        Type
	Bibs        []int
	PIDs        []int
	Sender      string              `json:",omitempty"`
	SenderName  string              `json:",omitempty"`
	ReplyTo     string              `json:",omitempty"`
	CC          string              `json:",omitempty"`
	BCC         string              `json:",omitempty"`
	CellPhone   string              `json:",omitempty"`
	Email       string              `json:",omitempty"`
	Subject     string              `json:",omitempty"`
	Text        string              `json:",omitempty"`
	URL         string              `json:",omitempty"`
	Method      string              `json:",omitempty"`
	Attachments []PreviewAttachment `json:",omitempty"`
	HTTPHeaders []HTTPHeader        `json:",omitempty"`
	Errors      []string            `json:",omitempty"`
}

type PreviewAttachment struct {
	Type  AttachmentType
	Name  string
	Label string
	Bib   int
	PID   int
}
