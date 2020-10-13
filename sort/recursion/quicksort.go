package recursion

func QuickSort(arr []int, l, r int, key string) {
	if l > r {
		return
	}
	pivotIndex := partion(arr, l, r, key)
	QuickSort(arr, l, pivotIndex-1, key)
	QuickSort(arr, pivotIndex+1, r, key)
}

func partion(arr []int, start, end int, key string) int {
	pivot := arr[start]
	mark := start
	for i := start + 1; i <= end; i++ {
		if key == "desc" {
			if arr[i] > pivot {
				mark++
				arr[i], arr[mark] = arr[mark], arr[i]
			}
		} else if key == "aesc" {
			if arr[i] < pivot {
				mark++
				arr[i], arr[mark] = arr[mark], arr[i]
			}
		}
	}
	arr[start], arr[mark] = arr[mark], arr[start]
	return mark
}