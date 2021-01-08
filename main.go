package main

import "fmt"

func main(){

	test1 := []int{1,6,2,123,5,1,2321,41,23}

	test2 := []int{12,21,24,23,52,123,51,2}

	test3 := []int{64,1412,51223,5312,4125,15,12412}

	//quick_sort
	//fmt.Println(quick_sort(test1))
	fmt.Println(quick_sort(test2))
	fmt.Println(quick_sort(test3))

	//merge sort
	merge_sort(test1)
	fmt.Println(test1)
}
