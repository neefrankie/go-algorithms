package sorting

func QuickSort[T Number](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T Number](arr []T, p, r int) {
	if p < r {
		q := partition(arr, p, r)
		quickSort(arr, p, q-1)
		quickSort(arr, q+1, r)
	}
}
func partition[T Number](arr []T, p, r int) int {
	var x = arr[r]
	var i = p - 1
	// Loop until one less than r
	for j := p; j < r; j++ {
		if arr[j] <= x {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	q := i + 1
	arr[q], arr[r] = arr[r], arr[q]

	return q
}
