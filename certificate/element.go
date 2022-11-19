package certificate

import (
	"encoding/xml"

	"github.com/raceresult/go-model/decimal"
)

type ElementType int
type ElementPictureStretchMode int
type ElementBarcodeType int

const (
	CETField             ElementType = 0
	CETText              ElementType = 1
	CETPicture           ElementType = 2
	CETPictureName       ElementType = 3
	CETBarcode           ElementType = 4
	CETPictureExpression ElementType = 5
	CETRect              ElementType = 6
	CETLine              ElementType = 7
	CETCircle            ElementType = 8
	CETChipBarcode       ElementType = 11
	CETBibPosition       ElementType = 12

	BNoBarcode ElementBarcodeType = 0
	BCode39    ElementBarcodeType = 1
	BCodeEAN13 ElementBarcodeType = 2
	BCode128   ElementBarcodeType = 3

	PSStretch   ElementPictureStretchMode = 0
	PSNoStretch ElementPictureStretchMode = 1

	ATop    int = 0
	AMiddle int = 1
	ABottom int = 2

	ALeft   int = 1
	ACenter int = 2
	ARight  int = 3
)

// Element represents one element of a certificate
type Element struct {
	XMLName         xml.Name ` json:"-"xml:"Element"`
	Type            ElementType
	Data            string
	Left            decimal.Decimal
	Top             decimal.Decimal
	Width           decimal.Decimal
	Height          decimal.Decimal
	FontName        string
	FontSize        int
	FontBold        bool
	FontItalic      bool
	FontUnderlined  bool
	FontColor       int
	FontColorCMYK   string
	Alignment       int `json:"Alignment" xml:"Alignment"`
	VAlignment      int `json:"vAlignment" xml:"vAlignment"`
	Page            int
	Rotation        int
	DynamicFormat   string                    `json:"DF" xml:"DF"`
	PictureStretch  ElementPictureStretchMode `json:"Stretch" xml:"Stretch"`
	BarcodeType     ElementBarcodeType        `json:"Barcode" xml:"Barcode"`
	Locked          bool
	TextScaling     decimal.Decimal
	TextCharSpacing decimal.Decimal
	OutlineWidth    decimal.Decimal
	OutlineColor    string
	Transparency    decimal.Decimal
}
