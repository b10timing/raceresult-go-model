package label

import "github.com/raceresult/go-model/decimal"

type Label struct {
	Name              string `json:"LabelName"`
	PageFormat        PageFormat
	PageSize          PageSize
	PageMarginTop     decimal.Decimal
	PageMarginLeft    decimal.Decimal
	Width             decimal.Decimal
	Height            decimal.Decimal
	SpacingVertical   decimal.Decimal
	SpacingHorizontal decimal.Decimal
	Direction         LabelDirection
	Expression        string
	FontName          string
	FontSize          int
	FontBold          bool
	FontItalic        bool
	FontUnderline     bool
	Design            string
	BarcodeType       BarcodeType
	Alignment         Alignment
	Filter            string
	Sort1             string
	Sort2             string
	Sort3             string
	SortDesc1         bool
	SortDesc2         bool
	SortDesc3         bool
}

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

type PageFormat int

const (
	PFPortrait  PageFormat = 0
	PFLandscape PageFormat = 1
)

type LabelDirection int

const (
	LDDownThenRight LabelDirection = 0
	LDRightThenDown LabelDirection = 1
)

type BarcodeType int

const (
	BNoBarcode BarcodeType = 0
	BCode39    BarcodeType = 1
	BCodeEAN13 BarcodeType = 2
	BCode128   BarcodeType = 3
)

type Alignment int

const (
	ALeft   Alignment = 1
	ACenter Alignment = 2
	ARight  Alignment = 3
)
