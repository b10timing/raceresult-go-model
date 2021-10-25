package statistic

import (
	"github.com/raceresult/go-model/decimal"
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
	Headline1        string
	Headline2        string
	LineSpacing      decimal.Decimal
	FontName         string
	FontSize         int
	PageFormat       PageFormat
	PageMarginBottom decimal.Decimal
	PageMarginLeft   decimal.Decimal
	PageMarginRight  decimal.Decimal
	PageMarginTop    decimal.Decimal
	PageSize         PageSize
	TopLeftHeader    string
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
