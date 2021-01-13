package main

import "errors"

func getMax(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty array")
	}
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}

	return max, nil
}

func countSort(arr []int, exp int) {
	var count [10]int
	output := make([]int, len(arr))

	for _, e := range arr {
		count[(e/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = output[i]
	}
}

func radixSort(arr []int) ([]int, error) {
	if len(arr) == 1 {
		return arr, nil
	}

	max, err := getMax(arr)
	if err != nil {
		return nil, err
	}

	for i := 1; max/i > 0; i *= 10 {
		countSort(arr, i)
	}

	return arr, nil
}
