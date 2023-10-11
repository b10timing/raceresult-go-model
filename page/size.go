package page

import "github.com/raceresult/go-model/decimal"

type Size int

const (
	SizeA1          Size = 0
	SizeA2          Size = 1
	SizeA3          Size = 2
	SizeA4          Size = 3
	SizeA5          Size = 4
	SizeLetter      Size = 5
	SizeLegal       Size = 6
	SizeA6          Size = 7
	SizeUserDefined Size = 8
)

func (s Size) ToMM(landscape bool) (decimal.Decimal, decimal.Decimal) {
	var w, h decimal.Decimal
	switch s {
	case SizeA1:
		w = decimal.FromInt(594)
		h = decimal.FromInt(840)
	case SizeA2:
		w = decimal.FromInt(420)
		h = decimal.FromInt(594)
	case SizeA3:
		w = decimal.FromInt(297)
		h = decimal.FromInt(420)
	case SizeA4:
		w = decimal.FromInt(210)
		h = decimal.FromInt(297)
	case SizeA5:
		w = decimal.FromInt(148)
		h = decimal.FromInt(210)
	case SizeA6:
		w = decimal.FromInt(105)
		h = decimal.FromInt(148)
	case SizeLetter:
		w = decimal.FromInt(216) // 85 * 2.54
		h = decimal.FromInt(279) // 110 * 2.54
	case SizeLegal:
		w = decimal.FromInt(216) // 85 * 2.54
		h = decimal.FromInt(356) // 140 * 2.54
	}

	if landscape {
		return h, w
	}
	return w, h
}
