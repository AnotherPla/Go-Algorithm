package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	a := []int{4, 6, 45, 21, 7, 8, 2, 346, 52, 14, 25, 1}
	sort.QuickSort(a, 0, len(a)-1)
	fmt.Printf("a: %v\n", a)

	b := []int{3, 61, 24, 12, 4, 76, 14, 4, 33, 45, 7851, 142, 34}
	sort.Qsort(b, 0, len(b)-1)
	fmt.Printf("b: %v\n", b)
}
