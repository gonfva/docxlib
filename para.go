package docxlib

import "github.com/gonfva/docxlib/xml"

type Paragraph struct {
	lib  *xml.LibXML
	pxml *xml.Paragraph
}

func (p *Paragraph) Children() (ret []*Inline) {
	var inline *Inline
	ret = make([]*Inline, 0)
	for _, child := range p.pxml.Data {
		if child.Link != nil {
			kind := "Link"
			id := child.Link.ID
			target, _ := p.lib.References(id)
			link := Link{Target: target, Text: child.Link.Run.InstrText}
			inline = &Inline{Kind: kind, childXML: &child, Link: link}
		}
		if child.Run != nil {
			kind := "Text"
			text := Text{Content: child.Run.Text.Text, pxml: child.Run}
			inline = &Inline{Kind: kind, childXML: &child, Text: text}
		}
		ret = append(ret, inline)
	}
	return
}

// AddLink adds an hyperlink to paragraph
func (p *Paragraph) AddLink(text string, link string) *Link {
	l := p.pxml.AddLink(text, link)
	hyperlink := Link{Target: link, Text: link, pxml: l}

	return &hyperlink
}

// AddText adds text to paragraph
func (p *Paragraph) AddText(text string) *Text {
	t := p.pxml.AddText(text)
	txt := Text{Content: text, pxml: t}
	return &txt
}
