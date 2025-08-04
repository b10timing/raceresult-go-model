package certificate

import "github.com/raceresult/go-model/decimal"

type Zone struct {
	Top    decimal.Decimal
	Bottom decimal.Decimal
	Type   string
}
