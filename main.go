package main

import "fmt"

func main() {

	test1 := []int{1, 6, 2, 123, 5, 1, 2321, 41, 23}

	test2 := []int{12, 21, 24, 23, 52, 123, 51, 2}

	test3 := []int{64, 1412, 51223, 5312, 4125, 15, 12412}

	test4 := []float64{4, 1, 1, 3, 2, 2, 7, -6, 0}

	//quick_sort
	//fmt.Println(quick_sort(test1))
	fmt.Println(quickSort(test2))
	fmt.Println(quickSort(test3))

	//merge sort
	mergeSort(test1)
	fmt.Println(test1)

	fmt.Println(bucketSort(test4, len(test4)))

	fmt.Println(radixSort(test2))
}
