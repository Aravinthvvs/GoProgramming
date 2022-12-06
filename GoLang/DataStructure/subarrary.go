package main

import (
	"fmt"
	"math")


func main(){
	fmt.Println(minSubArrayLen(7,[]int{2,3,1,2,4,3}))
}

func minSubArrayLen(target int, nums []int) int {
	j, sum, count, min := 0,0,0,math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] == target{
			return 1
		}
		fmt.Println(" index :: ",i," val :: ", nums[i], " sum :: " ,sum)
		if nums[i] + sum >= target {
			if count < min {
				min = count
			}
			sum -= nums[j]
			j++
			i--
			count--
		}else {
			sum += nums[i]
			count ++
		}
	}

	if min == math.MaxInt64 {
		return 0
	}

	return min+1
}