package main

import (
	"fmt"
)

func push(st *[]rune, val rune) {
	//fmt.Println("size of val ", len(*st))
	*st = append(*st, val)
	//fmt.Println(" push size ", len(*st))
}

func pop(st []rune) (res []rune, err string) {
	size := len(st)
	if size == 0 {
		return st, " empty stack"
	}
	st = st[:size-1]
	//fmt.Println("pop stack",st)
	return st, ""
}

func top(st []rune) (res rune, err string) {
	size := len(st)
	if size == 0 {
		return -1, " empty stack"
	}
	return st[size-1], ""
}

func main() {
	exp := string("[]()()()(())")
	res := isBalanced(exp)
	if res{
		fmt.Println("Expression is Balanced")
	}else{
		fmt.Println("Expression is not Balanced")
	}
}

func isBalanced(exp string)bool {

	if len(exp) % 2 != 0{
		return false
	}

	var resultStack []rune

	expMap := map[rune]rune{'(':')','{':'}','[':']'}

	for _, ch := range exp {
		
		if val,ok := expMap[ch];ok{
			resultStack = append(resultStack, val)
			continue
		}

		if topVal,err:=top(resultStack); topVal != ch || err !=""{
			return false
		}
		resultStack, _ = pop(resultStack)
	}
	

	if(len(resultStack) != 0){
		return false
	}
	return true
	
}