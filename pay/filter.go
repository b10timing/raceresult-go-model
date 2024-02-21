package pay

import (
	"github.com/raceresult/go-model/decimal"
	"time"
)

type RegistrationFilter struct {
	ID        []int
	Event     []int
	Month     []int
	Year      []int
	PaymentID []int
	Search    string
}

type PaymentFilter struct {
	ID             []int
	IDorRetry      []int
	RetryOf        []int
	Event          []int
	Month          []int
	Year           []int
	Method         []int
	Reference      []string
	Email          []string
	Search         []string
	Received       []bool
	ReceivedAmount []decimal.Decimal
	MinCreated     time.Time
	ToPay          decimal.Decimal
}
