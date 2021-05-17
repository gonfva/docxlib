package docxlib

import "github.com/gonfva/docxlib/xml"

type Link struct {
	Target string
	Text   string
	pxml   *xml.Hyperlink
}
