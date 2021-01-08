package main

func merge_sort(arr []int) {
	if len(arr) == 1 {
		return
	}

	mid := len(arr) / 2

	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	merge_sort(left)
	merge_sort(right)
	merge(arr, left, right)
}


func merge(result, left, right []int) {
	var l, r, i int // default to 0

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}
}