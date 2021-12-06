package decimal

import "testing"

func TestDecimal_Format(t *testing.T) {
	type args struct {
		decimals     int
		deciSep      string
		thousandsSep string
	}
	tests := []struct {
		name string
		x    Decimal
		args args
		want string
	}{
		{
			name: "test1",
			x:    FromFloat(123456789.123),
			args: args{
				decimals:     2,
				deciSep:      ".",
				thousandsSep: ",",
			},
			want: "123,456,789.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.Format(tt.args.decimals, tt.args.deciSep, tt.args.thousandsSep); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
