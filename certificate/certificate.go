package certificate

// Certificate represents a certificate
type Certificate struct {
	Name               string `json:"CertificateName" xml:"CertificateName"`
	PageSize           PageSize
	PageFormat         PageFormat
	PageHeight         int
	PageWidth          int
	SheetHeight        int
	SheetWidth         int
	MarginTop          int
	MarginLeft         int
	MarginRight        int
	MarginBottom       int
	CutLeft            int
	CutTop             int
	CutBottom          int
	CutRight           int
	DistanceVertical   int
	DistanceHorizontal int
	Holes              int
	SpecialHoles       string
	Substrate          string
	RGBBlackToCMYK     bool
	CMYKBlackValue     string
	PrintNotes         string
	Copies             int
	PrintMode          string
	Reverse            bool
	RoundedCorners     bool
	Overprint          bool
	BlockSize          int
	Elements           []Element `json:"Fields" xml:"Element"`
}
