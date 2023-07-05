package archives

import "github.com/raceresult/go-model/vbdate"

type Participant struct {
	ID             int
	Transponder1   string                 `json:",omitempty"`
	Transponder2   string                 `json:",omitempty"`
	RegNo          string                 `json:",omitempty"`
	Title          string                 `json:",omitempty"`
	Language       string                 `json:",omitempty"`
	Lastname       string                 `json:",omitempty"`
	Firstname      string                 `json:",omitempty"`
	Sex            string                 `json:",omitempty"`
	DateOfBirth    vbdate.VBDate          `json:",omitempty"`
	Street         string                 `json:",omitempty"`
	ZIP            string                 `json:",omitempty"`
	State          string                 `json:",omitempty"`
	City           string                 `json:",omitempty"`
	Country        int                    `json:",omitempty"`
	Nation         int                    `json:",omitempty"`
	Club           string                 `json:",omitempty"`
	License        string                 `json:",omitempty"`
	Phone          string                 `json:",omitempty"`
	CellPhone      string                 `json:",omitempty"`
	Email          string                 `json:",omitempty"`
	AddFields      map[string]interface{} `json:",omitempty"`
	Participations []Participation        `json:",omitempty"`
}

type Participation struct {
	Event   string
	Contest int
	Time    string `json:",omitempty"`
	TotRank int    `json:",omitempty"`
	MFRank  int    `json:",omitempty"`
	AGRank  int    `json:",omitempty"`
	Bib     int
}

type Event struct {
	ID       string
	Name     string
	Date     vbdate.VBDate
	Contests []Contest
}

type Contest struct {
	ID   int
	Name string
}

type ParticipationExt struct {
	EventDate   vbdate.VBDate
	EventName   string
	ContestName string
	FinalTime   string
	TotRank     int
	MFRank      int
	AGRank      int
	Bib         int
}

type Match struct {
	ID        int
	FirstName string
	LastName  string
	Year      int
}
