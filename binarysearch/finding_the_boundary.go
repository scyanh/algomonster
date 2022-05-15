package binarysearch

type findingTheBoundary struct {
}

func NewFindingTheBoundary() findingTheBoundary {
	return findingTheBoundary{}
}

// FindBoundary
/*
An array of boolean values is divided into two sections; the left section consists of
all false and the right section consists of all true. Find the boundary of the right section,
i.e. the index of the first true element. If there is no true element, return -1.
*/
func (findingTheBoundary) FindBoundary(arr []bool) int {
	left := 0
	right := len(arr) - 1
	boundary := -1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] {
			boundary = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return boundary
}
