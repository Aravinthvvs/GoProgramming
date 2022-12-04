package main

import (
	"fmt"
	"math"
)


func main(){
	fmt.Println(minSubArrayLen(11,[]int{1,1,1,1,1,1,1,1}))
}
func minSubArrayLen(target int, nums []int) int {
    minLen := math.MaxInt64
    sum := 0
    for start, end := 0,0;start < len(nums) && end < len(nums);end++{
        if nums[end] == target{
            return 1
        }
		//fmt.Println("sum :: ",sum )
        if sum+nums[end] >= target{
           
           //fmt.Println("minLen, start, end ",minLen, start, end)
            if minLen > (end-start+1){
                minLen = end - start+1
            }
            sum = sum - nums[start]
            end--
            start++
            
            //fmt.Println("min len end :: ",minLen)
        }else{
            //fmt.Println("sum :: ",sum, "end :: ", end)
            sum = sum + nums[end]
        }
        
    }
    if minLen == math.MaxInt64{
        return 0
    }
    return minLen
}