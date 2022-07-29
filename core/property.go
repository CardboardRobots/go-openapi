package core

import "strings"

func GetPropertyName(name string) string {
	if len(name) < 1 {
		return ""
	}
	first := name[:1]
	rest := name[1:]
	return strings.ToUpper(first) + rest
}

func GetPropertyType(name string) string {
	switch name {
	case "number":
		return "float32"
	case "integer":
		return "int"
	}
	return name
}
