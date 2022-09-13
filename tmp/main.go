package main

import (
	"fmt"
	"github.com/YFR718/go_tools/algorithm/sort"
)

func main() {
	l := []int{1, 5, 3, 6, 0, 2, 9, 2, 0, 1, 4, 6, 2, 8, 3, 5, 9, 0, 5, 3, 2, 8}
	sort.QuickSort(l, 0, len(l)-1)
	fmt.Println(l)

}
