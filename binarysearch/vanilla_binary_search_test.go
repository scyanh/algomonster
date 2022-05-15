package binarysearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVanillaBinarySearch(t *testing.T) {
	arr := make([]int, 3)
	arr = []int{1, 3, 6, 8, 9, 10}

	bs := NewVanillaBinarySearch()
	idx := bs.BinarySearch(arr, 8)
	expected := 3

	require.Equal(t, expected, idx)

}
