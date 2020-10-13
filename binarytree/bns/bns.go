/*
二分查找 带边界的二分查找
*/

package bns

import "fmt"

//终止条件left > right
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(mid)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

//终止条件left == right
func LeftBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == target {
			right = mid
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid
		}
	}
	if left == len(arr) || arr[left] != target {
		return -1
	}
	return left
}

//终止条件left == right
func RightBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == target {
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid
		}
	}
	if left == 0 || arr[left-1] != target {
		return -1
	}
	return left - 1
}
