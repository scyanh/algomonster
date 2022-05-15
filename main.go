package main

import (
	"fmt"
	"github.com/scyanh/algomonster/binarysearch"
)

func main() {

	arr := make([]bool, 6)
	arr = []bool{false, false, false, true, true, true}

	bs := binarysearch.NewFindingTheBoundary()
	idx := bs.FindBoundary(arr)

	fmt.Println("idx=", idx)
}
