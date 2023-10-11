package list

import (
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/page"
	"github.com/raceresult/go-model/vbdate"
)

// List represents the settings of a list
type List struct {
	Name                           string `json:"ListName" xml:"ListName"`
	BottomPicture                  string
	BottomPictureShow              ShowAt
	ColumnHeadsFontName            string
	ColumnHeadsFontSize            int
	ColumnHeadsFontBold            bool
	ColumnHeadsFontItalic          bool
	ColumnHeadsFontUnderlined      bool
	ColumnHeadsColor               string
	ColumnHeadsShow                ShowAt
	Columns                        int
	ColumnSpacing                  decimal.Decimal
	CoverSheet                     string
	BackSheet                      string
	Design                         string
	DesignShow                     ShowAt
	EveryOtherLineGray             bool
	FontName                       string
	FontSize                       int
	FooterFontName                 string
	FooterFontSize                 int
	FooterFontBold                 bool
	FooterFontItalic               bool
	FooterFontUnderlined           bool
	FooterColor                    string
	GrayLine                       bool
	HeadLine1                      string
	HeadLine1FontName              string
	HeadLine1FontSize              int
	HeadLine1FontBold              bool
	HeadLine1FontItalic            bool
	HeadLine1FontUnderlined        bool
	HeadLine2                      string
	HeadLine1Color                 string
	HeadLine1Show                  ShowAt
	HeadLine2FontName              string
	HeadLine2FontSize              int
	HeadLine2FontBold              bool
	HeadLine2FontItalic            bool
	HeadLine2FontUnderlined        bool
	HeadLine2Color                 string
	HeadLine2Show                  ShowAt
	HeightBottomPicture            decimal.Decimal
	LineColor                      string
	LineBackColor                  string
	LineDynamicFormat              string
	LineSpacing                    decimal.Decimal
	MaxRecords                     int
	MultiplierField                string
	PageFormat                     page.Format
	PageMarginBottom               decimal.Decimal
	PageMarginLeft                 decimal.Decimal
	PageMarginRight                decimal.Decimal
	PageMarginTop                  decimal.Decimal
	PageSize                       page.Size
	PageHeight                     int
	PageWidth                      int
	SepLine                        bool
	TopRightPicture                string
	TopRightPictureShow            ShowAt
	ListHeaderText                 string
	ListFooterText                 string
	ListHeaderFooterFontName       string
	ListHeaderFooterFontSize       int
	ListHeaderFooterFontBold       bool
	ListHeaderFooterFontItalic     bool
	ListHeaderFooterFontUnderlined bool
	ListHeaderFooterAlignment      int
	Remarks                        string
	LastChange                     vbdate.VBDate
	FooterText1                    string
	FooterText2                    string
	FooterText3                    string
	Orders                         []Order  `xml:"Order"`
	Filters                        []Filter `xml:"Filter"`
	Fields                         []Field  `xml:"Field"`
	SelectorResults                []SelectorResult
}

// Order represents one Order/Grouping of a list definition
type Order struct {
	Expression         string `xml:"Exp"`
	Descending         bool   `xml:"D"`
	Grouping2          int    `json:"Grouping" xml:"G"`
	GroupFilterDefault int
	GroupFilterLabel   string
	PageBreak          PageBreak `xml:"P"`
	FontName           string    `xml:"F"`
	FontSize           int       `xml:"S"`
	FontBold           bool      `xml:"B"`
	FontItalic         bool      `xml:"I"`
	FontUnderlined     bool      `xml:"U"`
	Color              string    `xml:"C"`
	BackgroundColor    string    `xml:"BC"`
	Spacing            int       `xml:"SP"`
	Ignore             bool      `json:"-"`
}

// Filter represents one filter of a list definition
type Filter struct {
	OrConjunction bool   `xml:"Or"`
	Expression1   string `xml:"Exp1"`
	Operator      string `xml:"Op"`
	Expression2   string `xml:"Exp2"`
}

// Field represents one field of a list definition
type Field struct {
	Expression     string          `xml:"Exp"`
	Label          string          `xml:"La"`
	Label2         string          `xml:"La2"`
	Alignment      int             `xml:"A"`
	FontBold       bool            `xml:"B"`
	FontItalic     bool            `xml:"I"`
	FontUnderlined bool            `xml:"U"`
	Line           int             `xml:"Li"`
	Color          string          `xml:"C"`
	Link           string          `xml:"Link"`
	ColOffset      int             `xml:"CO"`
	Position       decimal.Decimal `xml:"P"`
	DynamicFormat  string          `xml:"DF"`
	PreviewOnly    bool            `xml:"PO"`
	ResponsiveHide int             `xml:"RH"`
}

// SelectorResult represents one SelectorResult option of a list definition
type SelectorResult struct {
	ResultID  int
	ResultID2 int
	ShowAs    string
}

type ShowAt int

const (
	SANever     ShowAt = 0
	SAFirstPage ShowAt = 1
	SAEveryPage ShowAt = 2
	SALastPage  ShowAt = 3
)

type PageBreak int

const (
	PBNoPageBreak  PageBreak = 0
	PBNewPage      PageBreak = 1
	PBKeepTogether PageBreak = 2
	PBNewColumn    PageBreak = 3
	PBRepeat       PageBreak = 4
)
