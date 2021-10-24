package variant

import (
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/vbdate"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var compareCases = []struct { //nolint:maligned
	name     string
	v1, v2   Variant
	caseSens bool
	collator *collate.Collator
	equals   bool
	less     bool
	greater  bool
}{
	{
		name:   "both nil",
		equals: true,
	},
	{
		name:   "empty string and nil",
		v1:     RString(""),
		equals: true,
	},
	{
		name:   "nil and empty string",
		v2:     RString(""),
		equals: true,
	},
	{
		name:   "zero int and nil",
		v1:     RInt(0),
		equals: true,
	},
	{
		name:   "zero float and nil",
		v1:     RFloat(0),
		equals: true,
	},
	{
		name:   "zero bool and nil",
		v1:     RBool(false),
		equals: false,
	},
	{
		name:   "empty time and nil",
		v1:     RDecimal(0),
		equals: true,
	},
	{
		name:   "zero date and nil",
		v1:     rDate(vbdate.ZeroDate()),
		equals: true,
	},
	{
		name:   "zero date and nil",
		v1:     rDate(vbdate.ZeroDate()),
		equals: true,
	},
	{
		name: "strings not equal",
		v1:   RString("abc"),
		v2:   RString("cde"),
		less: true,
	},
	{
		name:     "equal strings - case sensitive",
		v1:       RString("abc"),
		v2:       RString("abc"),
		caseSens: true,
		equals:   true,
	},
	{
		name:     "case insensitive",
		v1:       RString("aBc"),
		v2:       RString("Abc"),
		collator: collate.New(language.German, collate.IgnoreCase),
		equals:   true,
		greater:  true,
	},
	{
		name:     "case sensitive sorting",
		v1:       RString("aBc"),
		v2:       RString("Abc"),
		caseSens: true,
		greater:  true,
	},
	{
		name:     "case sensitive sorting with nil",
		v1:       nil,
		v2:       RString("abc"),
		caseSens: true,
		less:     true,
	},
	{
		name:   "string equals int",
		v1:     RString("8"),
		v2:     RInt(8),
		equals: true,
	},
	{
		name:    "string unequals int",
		v1:      RString("abc"),
		v2:      RInt(8),
		greater: true,
	},
	{
		name:   "string equals date",
		v1:     RString("2009-03-05"),
		v2:     rDate(time.Date(2009, 3, 5, 0, 0, 0, 0, time.UTC)),
		equals: true,
	},
	{
		name:   "string equals time",
		v1:     RString("1.6234"),
		v2:     RDecimal(decimal.FromFloat(1.6234)),
		equals: true,
	},
	{
		name:   "string (1) equals bool",
		v1:     RString("1"),
		v2:     RBool(true),
		equals: true,
	},
	{
		name:    "string (true) equals bool",
		v1:      RString("true"),
		v2:      RBool(true),
		equals:  true,
		greater: true,
	},
	{
		name:    "string equals float",
		v1:      RString("42.713859"),
		v2:      RFloat(42.713859),
		equals:  true,
		less:    false,
		greater: false,
	},
}

func TestEquals(t *testing.T) {
	for _, tt := range compareCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Equals(tt.v1, tt.v2, tt.caseSens)
			// dbg.Green(fmt.Sprintf("%+v", tt), got)
			assert.Equal(t, tt.equals, got)
		})
	}
}

func TestLess(t *testing.T) {
	for _, tt := range compareCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Less(tt.v1, tt.v2, tt.collator)
			// dbg.Green(fmt.Sprintf("%+v", tt), got)
			assert.Equal(t, tt.less, got)
		})
	}
}
func TestGreater(t *testing.T) {
	for _, tt := range compareCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Greater(tt.v1, tt.v2, tt.collator)
			// dbg.Green(fmt.Sprintf("%+v", tt), got)
			assert.Equal(t, tt.greater, got)
		})
	}
}

func BenchmarkEquals(b *testing.B) {
	l := len(compareCases)
	for i := 0; i < b.N; i++ {
		tt := compareCases[i%l]
		Equals(tt.v1, tt.v2, tt.caseSens)
	}
}

func BenchmarkLess(b *testing.B) {
	l := len(compareCases)
	for i := 0; i < b.N; i++ {
		tt := compareCases[i%l]
		Less(tt.v1, tt.v2, tt.collator)
	}
}

func BenchmarkGreater(b *testing.B) {
	l := len(compareCases)
	for i := 0; i < b.N; i++ {
		tt := compareCases[i%l]
		Greater(tt.v1, tt.v2, tt.collator)
	}
}

/*
TODO: 1. Originals

BenchmarkEquals-12     	19116507	        52.7 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	18144996	        63.3 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18047968	        64.1 ns/op	       6 B/op	       0 allocs/op

BenchmarkEquals-12     	22275150	        53.1 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	17960089	        64.5 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18987880	        63.7 ns/op	       6 B/op	       0 allocs/op

BenchmarkEquals-12     	20493670	        53.8 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	18359134	        64.7 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18171920	        65.3 ns/op	       6 B/op	       0 allocs/op

BenchmarkEquals-12     	19362510	        53.1 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	17466604	        64.5 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18878780	        62.9 ns/op	       6 B/op	       0 allocs/op

TODO: using structs (removed again, doesn't improve)

BenchmarkEquals-12     	23028954	        51.1 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	18345182	        65.3 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18000320	        64.2 ns/op	       6 B/op	       0 allocs/op

BenchmarkEquals-12     	23277567	        52.2 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	18253002	        64.3 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18144040	        64.8 ns/op	       6 B/op	       0 allocs/op

BenchmarkEquals-12     	22529976	        52.8 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	17974741	        64.3 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18332013	        65.2 ns/op	       6 B/op	       0 allocs/op

BenchmarkEquals-12     	22615528	        53.0 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	18357307	        64.7 ns/op	       6 B/op	       0 allocs/op
BenchmarkGreater-12    	18269962	        65.6 ns/op	       6 B/op	       0 allocs/op

TODO: (equals only!) rewritten to compare more finegrained

BenchmarkEquals-12     	19880842	        52.2 ns/op	       2 B/op	       0 allocs/op
BenchmarkEquals-12     	23241933	        51.7 ns/op	       2 B/op	       0 allocs/op
BenchmarkEquals-12     	21387001	        52.2 ns/op	       2 B/op	       0 allocs/op
BenchmarkEquals-12     	21848600	        52.6 ns/op	       2 B/op	       0 allocs/op

TODO: (equals only!) rewritten to use type switch instead of constant switch

BenchmarkEquals-12     	23437270	        51.3 ns/op	       2 B/op	       0 allocs/op
BenchmarkEquals-12     	22297714	        51.7 ns/op	       2 B/op	       0 allocs/op
BenchmarkEquals-12     	22960443	        51.9 ns/op	       2 B/op	       0 allocs/op
BenchmarkEquals-12     	23222940	        51.3 ns/op	       2 B/op	       0 allocs/op

TODO: 2. all rewritten to use type switches and more finegrained comparison

BenchmarkEquals-12     	23454974	        49.9 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	20925762	        55.2 ns/op	       2 B/op	       0 allocs/op
BenchmarkGreater-12    	20465773	        56.6 ns/op	       2 B/op	       0 allocs/op

BenchmarkEquals-12     	23010945	        50.5 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	21058872	        55.2 ns/op	       2 B/op	       0 allocs/op
BenchmarkGreater-12    	18900499	        57.2 ns/op	       2 B/op	       0 allocs/op

BenchmarkEquals-12     	23103498	        51.0 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	20159361	        55.2 ns/op	       2 B/op	       0 allocs/op
BenchmarkGreater-12    	18963296	        57.1 ns/op	       2 B/op	       0 allocs/op

BenchmarkEquals-12     	23039710	        51.3 ns/op	       2 B/op	       0 allocs/op
BenchmarkLess-12       	20806618	        55.3 ns/op	       2 B/op	       0 allocs/op
BenchmarkGreater-12    	20003941	        57.0 ns/op	       2 B/op	       0 allocs/op
*/
