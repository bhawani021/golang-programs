package main

import "fmt"

func leftRotate(arr []int, d int) {
	for i := 0; i < d; i++ {
		leftRotateByOne(arr)
	}

}

func leftRotateByOne(arr []int) {
	n := len(arr)

	temp := arr[0]
	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[n-1] = temp
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80} // [40 50 60 70 80 10 20 30]

	leftRotate(arr, 3)

	fmt.Println(arr)
}

// Complexity - O(n*d)
