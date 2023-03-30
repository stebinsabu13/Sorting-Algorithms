package main

import "fmt"

func main() {
	var arr = []int{87, 64, 23, 12, 56, 98, 12}
	selectionsort(arr)
	fmt.Println(arr)
}
func selectionsort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
