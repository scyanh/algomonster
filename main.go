package main

import (
	"fmt"
	"github.com/scyanh/algomonster/binarysearch"
)

func main() {

	arr := make([]int, 3)
	arr = []int{1, 3, 6, 8, 9, 10}

	bs := binarysearch.NewVanillaBinarySearch()
	idx := bs.BinarySearch(arr, 8)

	fmt.Println("idx=", idx)
}
