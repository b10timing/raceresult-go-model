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
		w = 594
		h = 840
	case SizeA2:
		w = 420
		h = 594
	case SizeA3:
		w = 297
		h = 420
	case SizeA4:
		w = 210
		h = 297
	case SizeA5:
		w = 148
		h = 210
	case SizeA6:
		w = 105
		h = 148
	case SizeLetter:
		w = 216 // 85 * 2.54
		h = 279 // 110 * 2.54
	case SizeLegal:
		w = 216 // 85 * 2.54
		h = 356 // 140 * 2.54
	}

	if landscape {
		return h, w
	}
	return w, h
}
