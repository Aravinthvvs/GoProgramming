package main

import "fmt"

var exists = struct{}{}

type set struct {
	m map[int]struct{}
}

func main() {
	
	fmt.Println("distinct values count :: ", distinctValues([]int{1,3,-4,2,5,2,1}))
}

func (s *set) add(val int) {
	s.m[val] = exists
}

func (s *set) remove(val int) {
	delete(s.m, val)
}

func (s *set) contains(val int) bool {
	_, ok := s.m[val]
	return ok
}

func (s *set) size()int{
	return len(s.m)
}

func distinctValues(nums []int)int{
	s := &set{}
	s.m = make(map[int]struct{})
	
	for _,num:= range nums{
		if(num >= 0){
		s.add(num)
		}
	}
	fmt.Println(s.m)
	return s.size()
}