package sorting

// Quicksort Алгоритм быстрой сортировки с использованием рекурсии
func Quicksort(arr []int) []int {
	var supportElm int
	var less []int
	var greater []int

	if len(arr) < 2 {
		return arr
	}

	// сомнительный способ, почитать/поискать лучший вариант
	supportElm = int(float64(len(arr)) / 2)

	root := arr[supportElm]

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
