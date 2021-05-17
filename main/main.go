package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gonfva/docxlib"
)

var fileLocation *string
var readOnly *bool

func init() {
	fileLocation = flag.String("file", "/tmp/new-file.docx", "file location")
	readOnly = flag.Bool("ro", false, "Don't attempt to generate a new file, just read one")
	flag.Parse()
}
func main() {
	if !*readOnly {
		fmt.Printf("Preparing new document to write at %s\n", *fileLocation)

		w := docxlib.New()
		// add new paragraph
		para1 := w.AddParagraph()
		// add text
		para1.AddText("test")

		para1.AddText("test font size").SetSize(22)
		para1.AddText("test color").SetColor("808080")
		para2 := w.AddParagraph()
		para2.AddText("test font size and color").SetSize(22).SetColor("ff0000")

		nextPara := w.AddParagraph()
		nextPara.AddLink("google", `http://google.com`)

		f, err := os.Create(*fileLocation)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		w.Write(f)
		fmt.Println("Document writen. \nNow trying to read it")
	}
	// Now let's try to read the file
	readFile, err := os.Open(*fileLocation)
	if err != nil {
		panic(err)
	}
	fileinfo, err := readFile.Stat()
	if err != nil {
		panic(err)
	}
	size := fileinfo.Size()
	doc, err := docxlib.Parse(readFile, int64(size))
	if err != nil {
		panic(err)
	}
	for _, para := range doc.Paragraphs() {
		for _, child := range para.Children() {
			if child.Kind == "Text" {
				fmt.Printf("\tWe've found a new run with the text ->%s\n", child.Text.Content)
			}
			if child.Kind == "Link" {
				link := child.Link.Target
				text := child.Link.Text
				fmt.Printf("\tWe've found a new hyperlink with ref %s and the text %s\n", link, text)

			}
		}
	}
	fmt.Println("End of main")
}
