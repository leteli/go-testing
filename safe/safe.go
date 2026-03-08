package safe

const ErrOutOfRange = "index out of range"

func MustAt[T any](xs []T, i int) T {
	if i < 0 || i >= len(xs) {
		panic(ErrOutOfRange)
	}
	return xs[i]
}
