package docxlib

import (
	"io"

	"github.com/gonfva/docxlib/xml"
)

// DocxLib is the structure that allow to access the internal represntation
// in memory of the doc (either read or about to be written)
type DocxLib struct {
	lib *xml.LibXML
}

// New generates a new empty docx file that we can manipulate and
// later on, save
func New() *DocxLib {
	lib := xml.New()
	return &DocxLib{lib: lib}
}

// Parse generates a new docx file in memory from a reader
// You can it invoke from a file
//		readFile, err := os.Open(FILE_PATH)
//		if err != nil {
//			panic(err)
//		}
//		fileinfo, err := readFile.Stat()
//		if err != nil {
//			panic(err)
//		}
//		size := fileinfo.Size()
//		doc, err := docxlib.Parse(readFile, int64(size))
// but also you can invoke from a webform (BEWARE of trusting users data!!!)
//
//	func uploadFile(w http.ResponseWriter, r *http.Request) {
//		r.ParseMultipartForm(10 << 20)
//
//		file, handler, err := r.FormFile("file")
//		if err != nil {
//			fmt.Println("Error Retrieving the File")
//			fmt.Println(err)
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		defer file.Close()
//		docxlib.Parse(file, handler.Size)
//	}
func Parse(reader io.ReaderAt, size int64) (doc *DocxLib, err error) {
	libxml, err := xml.Parse(reader, size)
	doc = &DocxLib{lib: libxml}
	return
}

// Write allows to save a docx to a writer
func (f *DocxLib) Write(writer io.Writer) (err error) {
	return f.lib.Write(writer)
}

func (f *DocxLib) Paragraphs() []*Paragraph {
	pars := make([]*Paragraph, 0)
	for _, p := range f.lib.Document.Body.Paragraphs {
		pars = append(pars, &Paragraph{pxml: p, lib: f.lib})
	}

	return pars
}

// AddParagraph adds a new paragraph
func (f *DocxLib) AddParagraph() *Paragraph {
	p := f.lib.AddParagraph()
	para := &Paragraph{pxml: p, lib: f.lib}
	return para
}
