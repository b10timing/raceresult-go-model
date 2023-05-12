package registration

import (
	"github.com/raceresult/go-model/vbdate"
)

type Registration struct {
	Name                                       string
	Key                                        string
	ChangeKeySalt                              string
	Title                                      string
	Enabled                                    bool
	EnabledFrom                                vbdate.VBDate
	EnabledTo                                  vbdate.VBDate
	TestModeKey                                string
	TestModeValidUntil                         vbdate.VBDate
	Type                                       string
	GroupMin, GroupMax, GroupDefault, GroupInc int
	Steps                                      []Step
	AdditionalValues                           []AdditionalValue
	CheckSex                                   bool
	CheckDuplicate                             bool
	OnlinePayment                              bool
	PaymentMethods                             []PaymentMethod
	Confirmation                               Confirmation
	AfterSave                                  []AfterSave
	CSS                                        string
}

type Step struct {
	Title       string
	Enabled     bool
	EnabledFrom vbdate.VBDate
	EnabledTo   vbdate.VBDate
	Elements    []Element
	ButtonText  string
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
	Format          string
	ClassName       string
	ID              int
	Common          int // common in group reg
	ValidationRules []ValidationRule
	Children        []Element
}

type Field struct {
	Name         string // field name
	ControlType  string // checkbox, text field, dropdown, option set, propose boxy
	Mandatory    int
	DefaultValue string
	Placeholder  string
	Unique       string
	Values       []Value // dropdown: if empty: use default values (contest, ATF), otherwise use these values
}

type Value struct {
	Value       interface{}
	Label       string
	Enabled     bool
	EnabledFrom vbdate.VBDate
	EnabledTo   vbdate.VBDate
	MaxCapacity int
	ShowIf      string
}

type AdditionalValue struct {
	FieldName string
	Source    string
	Value     string
	Filter    string
}

type Confirmation struct {
	Title      string
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
