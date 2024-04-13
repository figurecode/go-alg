package main

import (
	"fmt"
	"sorting"
)

func main() {
	sortArr := []int{3, 5, 33, 19, 39, 99, 70, 12, 23, 46, 2, 6, 7, 5, 1, 4, 10, 32}

	fmt.Println("Исходные данные: ", sortArr)
	fmt.Println("Результат: ", sorting.Quicksort(sortArr))
}
