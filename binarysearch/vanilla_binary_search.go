package binarysearch

type vanillaBinarySearch struct {
}

func NewVanillaBinarySearch() vanillaBinarySearch {
	return vanillaBinarySearch{}
}

// BinarySearch
/*
Binary search is an efficient array search algorithm. It works by narrowing down the
search range by half each time. If you have looked up a word in a physical dictionary,
then you've already used binary search in real life. Let's look at a simple example:

Given a sorted array of integers and an integer called target, find the element that equals
to the target and return its index. If the element is not found, return -1.

O(log(N))
*/
func (vanillaBinarySearch) BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] { // then remove right
			right = mid - 1
		} else { // then remove left
			left = mid + 1
		}
	}

	return -1
}
