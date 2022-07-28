package schemas

import (
	"fmt"
	"io"
)

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
