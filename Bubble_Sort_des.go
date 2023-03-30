package main

import "fmt"

func main() {
	var arr = []int{63, 25, 56, 98, 12, 37, 80}
	bubblesort(arr)
	fmt.Println(arr)
}
func bubblesort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
