package main

import "fmt"

func main() {
	var arr = []int{87, 64, 23, 12, 56, 98, 16}
	selectionsort(arr)
	fmt.Println(arr)
}
func selectionsort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		max := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		if max != i {
			arr[i], arr[max] = arr[max], arr[i]
		}
	}
}
