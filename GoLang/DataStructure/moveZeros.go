package main

import "fmt"

func main(){
	fmt.Println(moveZeroes([]int{1,0,0,1,4,3}))
}

func moveZeroes(nums []int) []int {
    for i,j := 0,1;i< len(nums) && j<len(nums);i++{
        if nums[i] == 0{
            if nums[j] != 0{
                nums[i], nums[j] = nums[j], nums[i]
            }else{
                j++
                i--
            }
        }else{
            j++
        }
    }
	return nums
}