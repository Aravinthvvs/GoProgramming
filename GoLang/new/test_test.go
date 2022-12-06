package main

import (
	"fmt"
	"testing"
)

func TestAdd( t *testing.T){
	input :=[]struct{
		testName string
		list []int
		want int
	}{
		{"postivie",
		[]int{1,2,3,4},
		10,
		},
		{
			"negative",[]int{-1,2,-4,-5},
			-8,
		},
	}

	for _, test:= range input{
		fmt.Println("testName :: ", test.testName, "input :: ",test.list)
		t.Run(test.testName,func(t *testing.T) {
			sum := add(test.list)
			if sum != test.want{
				t.Errorf("got %v and want %v",sum,test.want)
			}
		})

	}
}

func TestLinearSearch(t *testing.T){
	testCases := []struct{
		testName string
		list []int
		target int
		output bool
	}{
		{
			"foundValue",
			[]int{1,2,3,4,5,6},
			5,
			true,

		},
		{
			"NotfoundValue",
			[]int{1,2,3,4,5,6},
			7,
			false,

		},
		{
			"notfoundValueEmptyList",
			[]int{},
			5,
			false,

		},
	}

	for _, cases:= range testCases{
		t.Run(cases.testName,func(t *testing.T) {
			out := linearSearch(cases.list,cases.target)

			if out != cases.output{
				t.Errorf(" got %v and want %v",out,cases.output)
			}
		})
	}
}


func BenchmarkAdd(b *testing.B) {
	for i:=0; i< b.N;i++{
		add([]int{1,2,3,4})
	}
}

func ExampleAdd(){
	fmt.Println(add([]int{1,2,3,4}))
	//output:10
}