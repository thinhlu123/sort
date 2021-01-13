package main

func quickSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	left := 0
	right := len(a) - 1

	//pivot := right

	//a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
