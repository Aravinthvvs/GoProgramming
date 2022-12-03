package main

import (
	"fmt"
)

type stack []int

var maxStack stack //to store the maxElement

func main() {
	var st stack
	st.push(1)
	st.push(2)
	st.push(4)
	fmt.Println("max element :: ", maxElement())
	fmt.Println(" stack Len :: ", len(st), " isEmpty :: ", st.isStackEmpty())
	err := st.pop()
	st.push(6)
	st.push(5)
	st.push(3)
	fmt.Println("max element :: ", maxElement())
	fmt.Println(" stack Len :: ", len(st), " isEmpty :: ", st.isStackEmpty())
	err = st.pop()
	err = st.pop()
	if err != nil {
		fmt.Println(err)
	}
}

func (st *stack) push(val int) {
	max := val

	if !maxStack.isStackEmpty() {
		top, err := maxStack.top()

		if err != nil {
			fmt.Println("err  in push ", err)
			return
		}

		if max < top {
			max = top
		}
	}
	fmt.Println("pushing the val :: ", val)
	*st = append(*st, val)
	maxStack = append(maxStack, max)
}

func (st *stack) pop() error {
	if st.isStackEmpty() {
		return fmt.Errorf(" len stack is %v , so can't do pop operation", len(*st))
	}
	val, _ := st.top()
	top, _ := maxStack.top()
	//fmt.Println(" val ",val," top ",top)
	if !maxStack.isStackEmpty() && val == top {
		maxSize := len(maxStack) - 1
		maxStack = maxStack[:maxSize]
	}

	size := st.size() - 1
	*st = (*st)[:size]
	return nil
}

func (st *stack) top() (int, error) {
	if st.isStackEmpty() {
		return -1, fmt.Errorf(" len stack is %v , so can't do pop operation", len(*st))
	}
	return (*st)[st.size()-1], nil
}

func (st stack) isStackEmpty() bool {
	return len(st) == 0
}

func (st stack) size() int {
	return len(st)
}

func maxElement() int {
	max, _ := maxStack.top()
	return max
}
