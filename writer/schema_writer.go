package writer

import (
	"bufio"
	"bytes"
	"go/format"
	"html/template"
	"io/fs"
	"log"
	"os"

	"github.com/cardboardrobots/go-openapi/entity"
)

func WriteTemplate(fsys fs.FS, output string, data entity.TemplateData) {
	t, err := template.ParseFS(fsys, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	buffer := bytes.NewBufferString("")
	err = t.ExecuteTemplate(buffer, "main.tmpl", data)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	bytes, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Print(buffer)
		log.Fatalf("%v", err)
		return
	}

	f, err := os.Create(output)
	if err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.Write(bytes)
	writer.Flush()
}
