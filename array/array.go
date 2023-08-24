package array

type Array[T any] struct {
	array []T
}

func New[T any]() *Array[T] {
	return &Array[T]{
		array: []T{},
	}
}

func NewWithEntries[T any](entries ...T) *Array[T] {
	var arr []T

	for _, v := range entries {
		arr = appendValue[T](arr, v)
	}

	return &Array[T]{
		array: arr,
	}
}

func (array *Array[T]) Append(value ...T) {
	var arr []T

	for _, v := range value {
		arr = appendValue[T](arr, v)
	}

	array.array = arr
}

func (array *Array[T]) IndexOf(search_term T) int {
	for i, v := range array.array {
		if v == search_term {
			return i
		}
	}

	return -1
}

func test() {
	t := New[string]()

}
