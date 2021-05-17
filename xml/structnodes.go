package xml

import (
	"encoding/xml"
	"io"

	"github.com/golang/glog"
)

type ParagraphChild struct {
	Link       *Hyperlink     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	Run        *Run           `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main r,omitempty"`
	Properties *RunProperties `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPr,omitempty"`
}

type Paragraph struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
	Data    []ParagraphChild

	file *LibXML
}

func (p *Paragraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	children := make([]ParagraphChild, 0)
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		switch tt := t.(type) {
		case xml.StartElement:
			var elem ParagraphChild
			if tt.Name.Local == "hyperlink" {
				var value Hyperlink
				d.DecodeElement(&value, &start)
				id := getAtt(tt.Attr, "id")
				anchor := getAtt(tt.Attr, "anchor")
				if id != "" {
					value.ID = id
				}
				if anchor != "" {
					value.ID = anchor
				}
				elem = ParagraphChild{Link: &value}
			} else if tt.Name.Local == "r" {
				var value Run
				d.DecodeElement(&value, &start)
				elem = ParagraphChild{Run: &value}
				if value.InstrText == "" && value.Text == nil {
					glog.V(0).Infof("Empty run, we ignore")
					continue
				}
			} else if tt.Name.Local == "rPr" {
				var value RunProperties
				d.DecodeElement(&value, &start)
				elem = ParagraphChild{Properties: &value}
			} else {
				continue
			}
			children = append(children, elem)
		}

	}
	*p = Paragraph{Data: children}
	return nil

}
