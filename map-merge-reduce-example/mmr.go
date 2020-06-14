package main

import (
	"fmt"

	mergesort "github.com/naumvd95/gotrain/sorting-example/merge-sort"
)

func main() {
	a := []int{2, 1, 3, 4, 50, 78, 32, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(mergesort.MergeSortMultiThread(a))
}
