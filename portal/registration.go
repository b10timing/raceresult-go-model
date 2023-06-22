package portal

import (
	"time"

	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/vbdate"
)

type RegistrationContestEntryFee struct {
	DateStart vbdate.VBDate
	DateEnd   vbdate.VBDate
	RegStart  time.Time
	RegEnd    time.Time
	Fee       decimal.Decimal
	Currency  string
}

type RegistrationContest struct {
	ID          int
	EnabledFrom time.Time
	EnabledTo   time.Time
	Start       time.Time
	Name        string
	Sex         string
	AgeStart    vbdate.VBDate
	AgeEnd      vbdate.VBDate
	EntryFees   []RegistrationContestEntryFee
	SlotsLeft   int
}

type Registration struct {
	Name        string
	Title       string
	Key         string
	TestModeKey string
	EnabledFrom time.Time
	EnabledTo   time.Time
	ButtonText  string
	SlotsLeft   int
	Contests    []RegistrationContest
}
