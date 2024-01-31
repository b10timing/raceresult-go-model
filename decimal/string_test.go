package decimal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestDecimal_ToString(t *testing.T) {
	tests := []struct {
		name string
		s    Decimal
		want string
	}{
		{
			name: "0",
			s:    0,
			want: "0",
		},
		{
			name: "0.1",
			s:    FromFloat(0.1),
			want: "0.1",
		},
		{
			name: "-0.1",
			s:    FromFloat(-0.1),
			want: "-0.1",
		},
		{
			name: "0.0001",
			s:    FromFloat(0.0001),
			want: "0.0001",
		},
		{
			name: "-0.0001",
			s:    FromFloat(-0.0001),
			want: "-0.0001",
		},
		{
			name: "0.9",
			s:    FromFloat(0.9),
			want: "0.9",
		},
		{
			name: "-0.9",
			s:    FromFloat(-0.9),
			want: "-0.9",
		},
		{
			name: "-1.0001",
			s:    FromFloat(-1.0001),
			want: "-1.0001",
		},
		{
			name: "1.0001",
			s:    FromFloat(1.0001),
			want: "1.0001",
		},
		{
			name: "a",
			s:    FromFloat(-10001),
			want: "-10001",
		},
		{
			name: "10001",
			s:    FromFloat(10001),
			want: "10001",
		},
		{
			name: "max",
			s:    Max,
			want: "922337203685477.5807",
		},
		{
			name: "min",
			s:    Min,
			want: "-922337203685477.5808",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.ToString(), "ToString()")
		})
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Decimal
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "0",
			want: 0,
			args: "0",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "0.1",
			want: FromFloat(0.1),
			args: "0.1",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "-0.1",
			want: FromFloat(-0.1),
			args: "-0.1",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "0.0001",
			want: FromFloat(0.0001),
			args: "0.0001",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "-0.0001",
			want: FromFloat(-0.0001),
			args: "-0.0001",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "0.9",
			want: FromFloat(0.9),
			args: "0.9",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "-0.9",
			want: FromFloat(-0.9),
			args: "-0.9",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "-1.0001",
			want: FromFloat(-1.0001),
			args: "-1.0001",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "1.0001",
			want: FromFloat(1.0001),
			args: "1.0001",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "a",
			want: FromFloat(-10001),
			args: "-10001",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "10001",
			want: FromFloat(10001),
			args: "10001", wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "max",
			want: Max,
			args: "922337203685477.5807",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "min",
			want: Min,
			args: "-922337203685477.5808",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromString(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("FromString(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FromString(%v)", tt.args)
		})
	}
}
