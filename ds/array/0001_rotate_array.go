package main

import "fmt"

func rotateArray(arr []int, d int) {

	// length of array
	n := len(arr)

	// run a look form 0->d
	for i := 0; i < d; i++ {
		// get first value from array
		val := arr[i]

		// left shift element
		for j := 0; j < n-1; j++ {
			arr[j] = arr[j+1]
		}

		arr[n-1] = val
	}
}

func main() {
	arr := []int{10, 3, -1, 2, 3, 4, 5, 6}

	rotateArray(arr, 3)

	fmt.Println(arr)
}

// output: [2 3 4 5 6 10 -1 3]
// complexity: O(n*d)
// Auxiliary space O(1)
