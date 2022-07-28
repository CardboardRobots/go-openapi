package schemas

import "fmt"

func Tag(tag string) string {
	return fmt.Sprintf("`%v`", tag)
}

func Json(tag string) string {
	return fmt.Sprintf("json:\"%v\"", tag)
}
