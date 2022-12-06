package main

import "fmt"

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//input2 := []int {1,2,3,4,5,11}
	inMap := make(map[int]struct{})


	if findPair(inMap, input,10) {
		fmt.Println("Pair is exists")
	}else{
		fmt.Println("No Pairs")
	}

}

func findPair(inMap map[int]struct{},input []int, tar int) bool {
	for _, num := range input {

		if _ ,ok:= inMap[tar-num];ok {
			return true
		}else{
			inMap[num] = struct{}{}
		}
	}
	return false
}