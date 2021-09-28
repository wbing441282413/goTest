package utils

func BinarySearch(arr []int, l int, r int, val int) int {
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == val {
			return mid
		}
		if arr[mid] > val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursion(arr []int, l int, r int, val int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if arr[mid] == val {
		return mid
	}
	if arr[mid] < val {
		return BinarySearchRecursion(arr, mid+1, r, val)
	} else {
		return BinarySearchRecursion(arr, l, mid-1, val)
	}

}
