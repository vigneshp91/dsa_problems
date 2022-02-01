package main

import (
	"fmt"
	"math"
)

/*****************************

  Write a function that takes in an array of at least two integers and that
  returns an array of the starting and ending indices of the smallest subarray
  in the input array that needs to be sorted in place in order for the entire
  input array to be sorted (in ascending order).
   input = [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19]
   output = [3,9]
******************************/

func SubarraySort(array []int) []int {
	// Write your code here.
	//1,2,4,6,7,7,10,11,12,16,18,19
	result := make([]int, 0)
	if len(array) <= 1 {
		result = append(result, -1)
		result = append(result, -1)
		return result
	} else if len(array) == 2 {
		if array[0] < array[1] {
			result = append(result, -1)
			result = append(result, -1)
			return result
		} else {
			result = append(result, 0)
			result = append(result, 1)
			return result
		}
	}
	tempSubArr := make([]int, 0)
	n := len(array)
	for i := 1; i < n-1; i++ {
		if !(array[i-1] <= array[i] && array[i] <= array[i+1]) {
			tempSubArr = append(tempSubArr, array[i])
		}
	}
	if array[n-1] < array[n-2] {
		tempSubArr = append(tempSubArr, array[n-1])
	}

	if len(tempSubArr) == 0 {
		result = append(result, -1)
		result = append(result, -1)
		return result
	}
	fmt.Println(tempSubArr)
	smallest := getSmallest(tempSubArr)
	largest := getLargest(tempSubArr)
	fmt.Println(smallest)
	fmt.Println(largest)

	for i := 0; i < len(array); i++ {
		if smallest < array[i] {
			result = append(result, i)
			break
		}
	}

	for i := len(array) - 1; i > 0; i-- {
		if largest > array[i] {
			result = append(result, i)
			break
		}
	}
	fmt.Println(result)
	return result
}

func getSmallest(sa []int) int {
	min := math.MaxInt64
	for i := 0; i < len(sa); i++ {
		if min > sa[i] {
			min = sa[i]
		}
	}
	return min
}
func getLargest(sa []int) int {
	max := math.MinInt64
	for i := 0; i < len(sa); i++ {
		if max < sa[i] {
			max = sa[i]
		}
	}
	return max
}
