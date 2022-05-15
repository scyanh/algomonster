package binarysearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindingTheBoundary(t *testing.T) {
	arr := make([]bool, 6)
	arr = []bool{false, false, false, true, true, true}

	bs := NewFindingTheBoundary()
	idx := bs.FindBoundary(arr)
	expected := 3

	require.Equal(t, expected, idx)
}
