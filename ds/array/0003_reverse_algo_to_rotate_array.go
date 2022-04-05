// reverse algorithm to rotate an array
package main

import "fmt"

func reverse(arr []int, start, end int) {

	for start < end {
		temp := arr[start]

		arr[start] = arr[end]

		arr[end] = temp

		start++
		end--
	}
}

func leftRotate(arr []int, d int) {
	n := len(arr)
	reverse(arr, 0, d-1)
	reverse(arr, d, n-1)
	reverse(arr, 0, n-1)
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // [30 40 50 60 70 80 90 100 10 20]

	leftRotate(arr, 2)

	fmt.Println(arr)
}

// Time Complexity : O(n)

// Auxiliary Space: O(1)
