package registration

import (
	"github.com/raceresult/go-model/variant"
	"github.com/raceresult/go-model/vbdate"
)

type Registration struct {
	Name                                       string
	Key                                        string
	Enabled                                    bool
	EnabledFrom                                vbdate.VBDate
	EnabledTo                                  vbdate.VBDate
	Type                                       string // single, group, group_table
	GroupMin, GroupMax, GroupDefault, GroupInc int
	Steps                                      []Step
	CheckSex                                   bool
	CheckDuplicate                             bool
	OnlinePayment                              bool
	PaymentMethods                             []PaymentMethod
	Confirmation                               Confirmation
	AfterSave                                  []AfterSave
	CSS                                        string
}

type Step struct {
	Name        string
	Label       string
	Enabled     bool
	EnabledFrom vbdate.VBDate
	EnabledTo   vbdate.VBDate
	Elements    []Element
}

type Element struct {
	Type            string // text, box, field, entryfee list, ...
	Label           string
	Enabled         bool
	EnabledFrom     vbdate.VBDate
	EnabledTo       vbdate.VBDate
	Field           *Field
	ShowIf          string
	Color           string
	BackgroundColor string
	Common          int // common in group reg
	ValidationRules []ValidationRule
	Children        []Element
}

type Field struct {
	Name            string // field name
	ControlType     string // checkbox, text field, dropdown, option set, propose boxy
	Mandatory       bool
	DefaultValue    variant.Variant
	Placeholder     string
	Unique          string
	Values          []Value // dropdown: if empty: use default values (contest, ATF), otherwise use these values
	ShowOnCheckPage int
	ForceUpdate     bool
}

type Value struct {
	Value       variant.Variant
	Label       string
	Enabled     bool
	EnabledFrom vbdate.VBDate
	EnabledTo   vbdate.VBDate
	MaxCapacity int
	ShowIf      string
}

type Confirmation struct {
	Expression string
}

type AfterSave struct {
	Type        string
	Value       string
	Destination string
	Filter      string
}

type PaymentMethod struct {
	ID          int
	Label       string
	Enabled     bool
	EnabledFrom vbdate.VBDate
	EnabledTo   vbdate.VBDate
	Filter      string
}

type ValidationRule struct {
	Rule string
	Msg  string
}
