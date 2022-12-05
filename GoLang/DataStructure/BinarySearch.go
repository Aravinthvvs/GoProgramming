package main

import "fmt"

func main() {
	fmt.Println(BinarySearch([]int{1,2,3,4,5,6,7},1))
}

func BinarySearch(list []int, target int) bool {

	low := 0
	high := len(list)-1

	for low <= high{

		mid := low + (high-low)/2
		if list[mid] == target{
			return true
		}else if list[mid] < target{
			low = mid +1
		}else{
			high = mid -1
		}
	}
	return false
}