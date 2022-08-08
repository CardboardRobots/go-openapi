package writer

import (
	"bytes"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/cardboardrobots/go-openapi/entity"
)

func Write(fsys fs.FS, output string, data entity.TemplateData) error {
	rendered, err := RenderTemplate(fsys, output, data)
	if err != nil {
		return err
	}

	formatted, err := Format(rendered)
	if err != nil {
		// Write the unformatted output
		WriteToFile(output, rendered)
		return err
	}

	return WriteToFile(output, formatted)
}

func RenderTemplate(fsys fs.FS, output string, data entity.TemplateData) ([]byte, error) {
	t, err := template.ParseFS(fsys, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("")
	err = t.ExecuteTemplate(buffer, "main.tmpl", data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Format(source []byte) ([]byte, error) {
	bytes, err := format.Source(source)
	return bytes, err
}

func WriteToFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}
