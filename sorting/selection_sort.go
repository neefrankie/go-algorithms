package sorting

import "goalgorithms/pkg"

func SelectionSort[T pkg.Number](arr []T) {
	for outer := 0; outer < len(arr)-1; outer++ {
		min := outer

		for inner := outer + 1; inner < len(arr); inner++ {
			if arr[inner] < arr[min] {
				arr[min], arr[inner] = arr[inner], arr[min]
			}
		}
	}
}
