package schemas

func If[T any](test bool, a T, b T) T {
	if test {
		return a
	} else {
		return b
	}
}
