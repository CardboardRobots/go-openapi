package core

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func ParseDocument(ctx context.Context) {
	loader := openapi3.Loader{Context: ctx}

	fileName := "openapi.yml"
	log.Printf("Loading %v...\n", fileName)
	doc, err := loader.LoadFromFile(fileName)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	err = doc.Validate(ctx)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	// createTemplate("openapi")
	for key, path := range doc.Paths {
		printPath(key, path)
	}

	for key, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value
		fields := []Field{}
		if schema.Type == "object" {
			for key, schemaRef := range schema.Properties {
				schema := schemaRef.Value
				fields = append(fields, Field{
					Name: key,
					Type: Type{
						Name: schema.Type,
					},
				})
			}
			s := Struct{
				Name:   key,
				Fields: fields,
			}
			fmt.Print(s.String())
		}
	}
}

func printPath(key string, path *openapi3.PathItem) {
	fmt.Printf("path: %v\n", key)
	printOperation("Connect", path.Connect)
	printOperation("Delete", path.Delete)
	printOperation("Get", path.Get)
	printOperation("Head", path.Head)
	printOperation("Options", path.Options)
	printOperation("Patch", path.Patch)
	printOperation("Post", path.Post)
	printOperation("Put", path.Put)
	printOperation("Trace", path.Trace)
}

func printOperation(verb string, operation *openapi3.Operation) {
	if operation != nil {
		fmt.Printf("\toperation: %v, %#v\n", verb, operation.Description)
		for key, responseRef := range operation.Responses {
			response := responseRef.Value
			if response != nil {
				description := *response.Description
				fmt.Printf("\t\tresponse: %v, %v\n", key, description)
				for key, content := range response.Content {
					fmt.Printf("\t\t\tcontent: %v, %v\n", key, content.Schema)
				}
			}
		}
	}
}

func WritePackage(writer io.Writer, packageName string) {
	fmt.Fprintf(writer, "package %v\n\n", packageName)
}

type Package struct {
	Name    string
	Structs []Struct
	Funcs   []Func
}

func (p *Package) String() string {
	return Line(0, "package", p.Name) + "\n"
}

type Type struct {
	Name    string
	Pointer bool
}

func (t *Type) String() string {
	if t.Pointer {
		return "*" + t.Name
	} else {
		return t.Name
	}
}

type Field struct {
	Name string
	Type Type
	Tag  string
}

func (f *Field) String() string {
	if f.Tag != "" {
		return Line(1, f.Name, f.Type.String(), Tag(f.Tag))
	} else {
		return Line(1, f.Name, f.Type.String())
	}
}

type Struct struct {
	Name   string
	Fields []Field
}

func (s *Struct) String() string {
	sb := strings.Builder{}
	sb.WriteString(Line(0, "type", s.Name, "struct {"))
	for _, field := range s.Fields {
		sb.WriteString(field.String())
	}
	sb.WriteString("}\n\n")
	return sb.String()
}

type Func struct {
	Name   string
	Params []Type
	Return *Type
}

func Tag(tag string) string {
	return fmt.Sprintf("`%v`", tag)
}

func Json(tag string) string {
	return fmt.Sprintf("json:\"%v\"", tag)
}

func If[T any](test bool, a T, b T) T {
	if test {
		return a
	} else {
		return b
	}
}

func Line(indent int, values ...string) string {
	sb := strings.Builder{}
	for i := 0; i < indent; i++ {
		sb.WriteString("\t")
	}
	sb.WriteString(strings.Join(values, " "))
	sb.WriteString("\n")
	return sb.String()
}
