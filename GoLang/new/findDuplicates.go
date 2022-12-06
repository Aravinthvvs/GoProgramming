package main

import "fmt"

func main() {
	fmt.Println(hasDuplicate([]int{1,2,3,4,5,6,4}, 3))
}

func hasDuplicate(list []int, dist int)bool{

	inputMap := make(map[int]int)

	for i := 0;i <len(list);i++{
		if val, ok := inputMap[list[i]];ok{

			if i - val <= dist{
				return true
			}
		}
		inputMap[list[i]] = i
	}
	return false
}