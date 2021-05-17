package docxlib

import "github.com/gonfva/docxlib/xml"

type Text struct {
	Content string

	pxml *xml.Run
}

// Color allows to set run color
func (r *Text) SetColor(color string) *Text {
	r.pxml.Color(color)
	return r
}

// Size allows to set run size
func (r *Text) SetSize(size int) *Text {
	r.pxml.Size(size)
	return r
}
