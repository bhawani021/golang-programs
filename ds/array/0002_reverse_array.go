// reverse an array
package main

import "fmt"

func reverse(arr []int) {
	n := len(arr)

	start := 0
	end := n-1

	for start < end {

		temp := arr[start]

		arr[start] = arr[end]
		arr[end] = temp

		start++
		end--
	}
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	reverse(arr)

	fmt.Println(arr)

}
