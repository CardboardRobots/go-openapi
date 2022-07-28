package schemas

import "strings"

func Line(indent int, values ...string) string {
	sb := strings.Builder{}
	for i := 0; i < indent; i++ {
		sb.WriteString("\t")
	}
	sb.WriteString(strings.Join(values, " "))
	sb.WriteString("\n")
	return sb.String()
}
