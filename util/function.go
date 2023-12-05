package util

func Map[T any](array []T, callback func(T) T) []T {
	for i := 0; i < len(array); i++ {
		array[i] = callback(array[i])
	}
	return array
}

func Filter[T any](array []T, callback func(T) bool) []T {
	for i := 0; i < len(array); i++ {
		if !callback(array[i]) {
			array = append(array[:i], array[i+1:]...)
			i--
		}
	}
	return array
}

func Reduce[T any, U any](array []T, callback func(U, T) U, init U) U {
	for _, v := range array {
		init = callback(init, v)
	}
	return init
}
