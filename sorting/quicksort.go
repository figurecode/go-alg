package sorting

type Array []int64

// Quicksort Алгоритм быстрой сортировки с использованием рекурсии
func Quicksort(arr Array) Array {
	var (
		rootIndex     int
		root          int64
		less, greater Array
	)

	if len(arr) < 2 {
		return arr
	}

	// Тут будет возвращаться ближайшее целое, что подходит в данном случае, поэтому нет преобразования типа
	rootIndex = len(arr) / 2
	root = arr[rootIndex]

	for i := 0; i < len(arr); i++ {
		if arr[i] < root {
			less = append(less, arr[i])
		}
		if arr[i] > root {
			greater = append(greater, arr[i])
		}
	}

	return append(append(Quicksort(less), root), Quicksort(greater)...)
}
