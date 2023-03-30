package main

import "fmt"

func main() {
	var arr = []int{63, 25, 34, 98, 17, 44, 70}
	insertionsort(arr)
	fmt.Println(arr)
}
func insertionsort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] < temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}
