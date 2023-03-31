package main

import "fmt"

func merge(arr []int, l, mid, r int) {
	leftsize := mid - l + 1
	rightsize := r - mid
	var left, right []int
	for i := 0; i < leftsize; i++ {
		left = append(left, arr[l+i])
	}
	for j := 0; j < rightsize; j++ {
		right = append(right, arr[mid+1+j])
	}
	i, j, k := 0, 0, l
	for i < leftsize && j < rightsize {
		if left[i] >= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < leftsize {
		arr[k] = left[i]
		k++
		i++
	}
	for j < rightsize {
		arr[k] = right[j]
		k++
		j++
	}
}
func mergesort(arr []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2
		mergesort(arr, l, mid)
		mergesort(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}

func main() {
	var arr = []int{63, 57, 10, 12, 22, 10, 5, 57, 10, 44}
	mergesort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
