package statistic

import (
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/page"
)

type Aggregation int

const (
	SACount   Aggregation = 1
	SAMinimum Aggregation = 2
	SAMaximum Aggregation = 3
	SAMean    Aggregation = 4
	SASum     Aggregation = 5
)

type Statistics struct {
	Name             string `json:"StatisticName"`
	Type             string
	Row              string
	Col              string
	Filter           string
	OnlyFinishers    bool
	Field            string
	Aggregation      Aggregation
	SortByValue      bool
	SortDesc         bool
	Headline1        string
	Headline2        string
	LineSpacing      decimal.Decimal
	FontName         string
	FontSize         int
	PageFormat       page.Format
	PageMarginBottom decimal.Decimal
	PageMarginLeft   decimal.Decimal
	PageMarginRight  decimal.Decimal
	PageMarginTop    decimal.Decimal
	PageSize         page.Size
	PageHeight       decimal.Decimal
	PageWidth        decimal.Decimal
	TopLeftHeader    string
}
