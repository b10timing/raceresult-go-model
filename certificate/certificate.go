package certificate

import (
	"strconv"
	"strings"

	"github.com/raceresult/go-model/decimal"
)

// Certificate represents a certificate
type Certificate struct {
	Name               string `json:"CertificateName" xml:"CertificateName"`
	PageSize           PageSize
	PageFormat         PageFormat
	PageHeight         int
	PageWidth          int
	SheetHeight        int    `json:",omitempty"`
	SheetWidth         int    `json:",omitempty"`
	MarginTop          int    `json:",omitempty"`
	MarginLeft         int    `json:",omitempty"`
	MarginRight        int    `json:",omitempty"`
	MarginBottom       int    `json:",omitempty"`
	CutLeft            int    `json:",omitempty"`
	CutTop             int    `json:",omitempty"`
	CutBottom          int    `json:",omitempty"`
	CutRight           int    `json:",omitempty"`
	DistanceVertical   int    `json:",omitempty"`
	DistanceHorizontal int    `json:",omitempty"`
	Holes              int    `json:",omitempty"`
	SpecialHoles       string `json:",omitempty"`
	Substrate          string `json:",omitempty"`
	RGBBlackToCMYK     bool   `json:",omitempty"`
	CMYKBlackValue     string `json:",omitempty"`
	PrintNotes         string `json:",omitempty"`
	Copies             int    `json:",omitempty"`
	PrintMode          string `json:",omitempty"`
	Reverse            bool   `json:",omitempty"`
	RoundedCorners     bool   `json:",omitempty"`
	PlotterMarks       bool   `json:",omitempty"`
	Machine            int    `json:",omitempty"`
	BlockSize          int    `json:",omitempty"`
	Version            int
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

// GetWidth returns the width of the certificate in mm
func (q Certificate) GetWidth() decimal.Decimal {
	if q.PageFormat == PFLandscape {
		if q.PageSize == PSUserDefined {
			return decimal.FromInt(q.PageHeight)
		}
		return q.PageSize.Height()
	} else {
		if q.PageSize == PSUserDefined {
			return decimal.FromInt(q.PageWidth)
		}
		return q.PageSize.Width()
	}
}

// GetHeight returns the height of the certificate in mm
func (q Certificate) GetHeight() decimal.Decimal {
	if q.PageFormat == PFLandscape {
		if q.PageSize == PSUserDefined {
			return decimal.FromInt(q.PageWidth)
		}
		return q.PageSize.Width()
	} else {
		if q.PageSize == PSUserDefined {
			return decimal.FromInt(q.PageHeight)
		}
		return q.PageSize.Height()
	}
}

// GetHoles returns the coordinates of holes for this certificates (for bib numbers)
func (q Certificate) GetHoles() [][2]decimal.Decimal {
	var dest [][2]decimal.Decimal
	add := func(x, y decimal.Decimal) {
		dest = append(dest, [2]decimal.Decimal{x, y})
	}

	w := q.GetWidth()
	h := q.GetHeight()
	cl := decimal.FromInt(q.CutLeft)
	ct := decimal.FromInt(q.CutTop)
	cb := decimal.FromInt(q.CutBottom)
	cr := decimal.FromInt(q.CutRight)

	switch q.Holes {
	case 1, 2:
		add(cl+decimal.FromInt(9), ct+decimal.FromInt(9))
		add(w-cr-decimal.FromInt(9), ct+decimal.FromInt(9))
	case 3, 4:
		//  Lower chip
		add(cl+decimal.FromInt(30), h-cb-decimal.FromFloat(9.5))
		add(cl+decimal.FromInt(30), h-cb-decimal.FromInt(22))
		add(w-cr-decimal.FromInt(30), h-cb-decimal.FromFloat(9.5))
		add(w-cr-decimal.FromInt(30), h-cb-decimal.FromInt(22))

		//  upper chip
		if q.Holes == 3 { // duo
			add(cl+decimal.FromInt(30), h-cb-decimal.FromFloat(41.5))
			add(cl+decimal.FromInt(30), h-cb-decimal.FromInt(54))
			add(w-cr-decimal.FromInt(30), h-cb-decimal.FromFloat(41.5))
			add(w-cr-decimal.FromInt(30), h-cb-decimal.FromInt(54))
		}
	}

	switch q.Holes {
	case 1: // normal
		add(cl+decimal.FromInt(9), h-cb-decimal.FromInt(9))
		add(w-cr-decimal.FromInt(9), h-cb-decimal.FromInt(9))

	case 2: // mtb
		add(cl+(w-cl-cr)/2-decimal.FromInt(10), h-cb-decimal.FromInt(9))
		add(cl+(w-cl-cr)/2+decimal.FromInt(10), h-cb-decimal.FromInt(9))

	case 3: // duo
		// bib top holes
		add(cl+decimal.FromInt(9), ct+decimal.FromInt(83+9))
		add(w-cr-decimal.FromInt(9), ct+decimal.FromInt(83+9))

		// bib bottom holes
		add(cl+decimal.FromInt(9), ct+decimal.FromInt(233)-decimal.FromInt(9))
		add(w-cr-decimal.FromInt(9), ct+decimal.FromInt(233)-decimal.FromInt(9))

	case 4: // single
		// bib top holes
		add(cl+decimal.FromInt(9), ct+decimal.FromInt(85+9))
		add(w-cr-decimal.FromInt(9), ct+decimal.FromInt(85+9))

		// bib bottom holes
		add(cl+decimal.FromInt(9), ct+decimal.FromInt(265)-decimal.FromInt(9))
		add(w-cr-decimal.FromInt(9), ct+decimal.FromInt(265)-decimal.FromInt(9))

	case 99: // Special
		if q.SpecialHoles == "" {
			return nil
		}
		for _, s := range strings.Split(q.SpecialHoles, ";") {
			c := strings.Split(s, ":")
			if len(c) < 2 {
				continue
			}

			x, _ := strconv.ParseFloat(strings.TrimSpace(c[0]), 64)
			y, _ := strconv.ParseFloat(strings.TrimSpace(c[1]), 64)
			if x <= 0 || y <= 0 {
				continue
			}
			add(cl+decimal.FromFloat(x), ct+decimal.FromFloat(y))
		}
	}
	return dest
}
