package main

import "fmt"

func rotateArray(arr []int, d int) {
	n := len(arr)
	for i := 0; i < d; i++ {
		// read first element of array
		temp := arr[0]

		// let rotate array by one
		for j := 0; j < n - 1; j++ {
			arr[j] = arr[j+1]
		}

		// put first element at end of array
		arr[n-1] = temp
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 6}
	rotateArray(arr, 3)

	fmt.Println(arr)
}
