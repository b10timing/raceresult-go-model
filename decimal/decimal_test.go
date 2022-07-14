package decimal

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestDecimal_Round(t *testing.T) {
	x := Decimal(11000)
	r := x.Round(0)
	assert.Equal(t, Decimal(10000), r)

	x = Decimal(15000)
	r = x.Round(0)
	assert.Equal(t, Decimal(20000), r)

	x = Decimal(11230)
	r = x.Round(2)
	assert.Equal(t, Decimal(11200), r)

	x = Decimal(11250)
	r = x.Round(2)
	assert.Equal(t, Decimal(11300), r)
}

func TestDecimal_RoundUp(t *testing.T) {
	x := Decimal(11000)
	r := x.RoundUp(0)
	assert.Equal(t, Decimal(20000), r)

	x = Decimal(10000)
	r = x.RoundUp(0)
	assert.Equal(t, Decimal(10000), r)

	x = Decimal(11230)
	r = x.RoundUp(2)
	assert.Equal(t, Decimal(11300), r)

	x = Decimal(11300)
	r = x.RoundUp(2)
	assert.Equal(t, Decimal(11300), r)
}
