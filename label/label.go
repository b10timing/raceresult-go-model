package label

import (
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/page"
)

type Label struct {
	Name              string `json:"LabelName"`
	PageFormat        page.Format
	PageSize          page.Size
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
