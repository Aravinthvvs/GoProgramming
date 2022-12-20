package main

import "fmt"

func main() {
	fmt.Println(BinarySearch([]int{7, 6, 5, 4, 3, 2, 1}, 4))
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7}, -3))
}

func BinarySearch(list []int, target int) bool {
	isAssending := true
	if list[0] > list[len(list)-1] {
		isAssending = false
	}

	low := 0
	high := len(list) - 1

	for low <= high {

		mid := low + (high-low)/2
		if list[mid] == target {
			return true
		}

		if isAssending {
			if list[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if list[mid] > target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}
