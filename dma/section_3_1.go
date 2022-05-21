package dma

// This file answers questions of Exercise 3.1 of
// Discrete Mathematics and Its Applications.

// SumList finds the sum of all the integers in a list.
// Question 3.
func SumList(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var sum int

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	return sum
}

// LargestDifference produces as output the largest difference obtained
// by subtracting an integer in the list from the one following it.
// Question 4.
func LargestDifference(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	var max = arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff > max {
			max = diff
		}
	}

	return max
}

// FindDuplicates takes as input a list of n integers
// in non-decreasing order and produces the list of all
// values that occur more than one.
// A list of integers is non-decreasing if each integer
// in the list is at least as large as the previous integer
// in the list.
// Question 5.
func FindDuplicates(arr []int) []int {
	var items = make([]int, 0)

	var i = 0
	var size = len(arr)
	for i < size {
		// Starting from the next one, stops at the first one
		// that is not equal to arr[i]:
		// [2, 2, 3, 4, 6, 6, 6]
		var j = i + 1
		// Count duplicates.
		// This programmed could be easily modified to tell how many
		// duplicates there are for each item.
		var count = 0
		for ; j < size && arr[i] == arr[j]; j++ {
			count++
		}
		// At least one duplicate is found.
		if count > 0 {
			items = append(items, arr[i])
		}
		// Next round loop starts from j.
		i = j
	}

	return items
}
