package certificateset

type Type int

const (
	CSTSingleCertificate    Type = 0
	CSTTeamCertificate      Type = 1
	CSTTeamCertificateMulti Type = 2
)

// CertificateSet represents one certificate set
type CertificateSet struct {
	Name                    string
	Certificate             string `json:"CertificateName" xml:"CertificateName"`
	CertificateSetType      Type
	TeamScore               int
	FilterRankID            int
	FilterRankOperator      string
	FilterRankCompare       int
	FilterAYN               int
	FilterOnlyFinishers     bool
	FilterGeneral           string
	OnlyUnshownCertificates bool
	Sort1                   string
	Sort2                   string
	Sort3                   string
	Sort4                   string
	SortDesc1               bool
	SortDesc2               bool
	SortDesc3               bool
	SortDesc4               bool
}
