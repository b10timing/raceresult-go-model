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
	PlotterMarks       bool
	Overprint          bool
	BlockSize          int
	Elements           []Element `json:"Fields" xml:"Element"`
}

// PageCount returns the number of pages the certificate has (max of Page attribute of all elements)
func (q Certificate) PageCount() int {
	c := 1
	for _, item := range q.Elements {
		if c < item.Page {
			c = item.Page
		}
	}
	return c
}
