package slice

func Find[T any](
	arr []T,
	f func(elem T) bool,
) T {
	for _, v := range arr {
		if f(v) == true {
			return v
		}
	}
	var none T
	return none
}

func Map[T any](
	arr []T,
	f func(elem T, idx int) any,
) []any {
	tmp := make([]any, len(arr))
	for i, v := range arr {
		tmp[i] = f(v, i)
	}
	return tmp
}

func Pop[T any](slice []T) ([]T, T) {
	return slice[:len(slice)-1], slice[len(slice)-1]
}
