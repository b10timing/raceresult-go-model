package emailtemplate

import (
	"strconv"
	"strings"

	"github.com/raceresult/go-model/variant"
)

type PreviewIDList []int

func (q *PreviewIDList) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "\"")
	*q = PreviewIDList(variant.StringList(strings.Split(s, ",")).ToInt())
	return nil
}
func (q *PreviewIDList) MarshalJSON() ([]byte, error) {
	if len(*q) == 1 {
		return []byte(strconv.Itoa((*q)[0])), nil
	}
	return []byte("\"" + strings.Join(variant.IntList(*q).ToString(), ",") + "\""), nil
}

type Preview struct {
	Bib         PreviewIDList
	PID         PreviewIDList
	CellPhone   string       `json:",omitempty"`
	Email       string       `json:",omitempty"`
	Subject     string       `json:",omitempty"`
	Text        string       `json:",omitempty"`
	Sender      string       `json:",omitempty"`
	URL         string       `json:",omitempty"`
	Attachments []Attachment `json:",omitempty"`
	Auth        string       `json:",omitempty"`
}
