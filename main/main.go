package main

import (
	"fmt"
	"sorting"
)

func main() {
	sortArr := sorting.Array{0, 3, 19, 5, 191, 90, 33, 5, 94, 19, 11, 111, 8, 99, 39, 70, 7, 12, 2, 23, 1, 40, 100, 68, 304, 102, 4, 10, 32, 6}

	fmt.Println("Исходные данные: ", sortArr)
	fmt.Println("Результат: ", sorting.Quicksort(sortArr))
}
