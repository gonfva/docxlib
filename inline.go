package docxlib

import "github.com/gonfva/docxlib/xml"

type Inline struct {
	childXML *xml.ParagraphChild
	Kind     string
	Link
	Text
}
