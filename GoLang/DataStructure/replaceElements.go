package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(replaceElements([]int{17,18,5,6,4,1}))
}

func replaceElements(nums[]int) []int{
	maxRight := -1
	for index := len(nums) -1; index >= 0; index--{
		nums[index], maxRight = maxRight, int(math.Max(float64(nums[index]),float64(maxRight)))
	}
	return nums
}