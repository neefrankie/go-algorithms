package dma

import "strings"

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

// CountNegativeInts implements question 6:
// Takes as input a list of n integers and finds the
// number of negative integers in the list.
func CountNegativeInts(ints []int) int {
	var count int

	for _, n := range ints {
		if n < 0 {
			count++
		}
	}

	return count
}

// LastEvenInt implements question 7:
// Takes as input a list of n integers
// and find the location of the last even integer in the list
// or returns -1 if there are no event integers in the list.
func LastEvenInt(arr []int) int {
	var pos int = -1
	for i, n := range arr {
		if n%2 == 0 {
			pos = i
		}
	}

	return pos
}

// IsPalindrome implements question 9 to test if a string is a palindrome.
// A palindrome is a string that reads the same forward and backward. For example:
// A man, a plan, a canal, Panama!
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	var left = 0
	var right = len(s) - 1

	for left < right {
		leftCh := s[left]
		rightCh := s[right]
		if !isLetter(leftCh) {
			left++
			continue
		}
		if !isLetter(rightCh) {
			right--
			continue
		}
		if leftCh != rightCh {
			return false
		}
		left++
		right--
	}

	return true
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

// XPower implements question 10 to compute x^n,
// where x is a real number and n is an integer.
func XPower(x float64, n int) float64 {
	if n >= 0 {
		return xPower(x, n)
	}

	return 1 / xPower(x, -n)
}

func xPower(x float64, n int) float64 {
	if n < 0 {
		panic("n must be non-negative")
	}
	var result float64 = 1

	for i := 1; i <= n; i++ {
		result = result * x
	}

	return result
}
