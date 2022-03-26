package sorting

import "fmt"

func InsertionSortInt(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		fmt.Printf("Start outer loop %d\n", i)

		var temp = arr[i]
		fmt.Printf("Taken out element at %d:%d\n", i, arr[i])

		var j = i
		fmt.Printf("\tInner loop starts at %d\n", j)
		for j > 0 && arr[j-1] >= temp {
			fmt.Printf("\tMove element at %d to %d\n", j-1, j)

			arr[j] = arr[j-1]
			j--
		}
		fmt.Printf("Put slotted element %d back to %d\n", temp, j)
		arr[j] = temp
		fmt.Printf("End outer loop %d\n\n", i)
	}

	return arr
}
