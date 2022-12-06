package main

import (
	"fmt"
	"strconv"
	"sync"
)


func main(){
	
	myChan := make(chan int)
	var wg sync.WaitGroup
	for _, num := range []int{1,2,3,4,5}{
		wg.Add(1)
		n,_ := strconv.Atoi(strconv.Itoa(num))
		go add(n, myChan, &wg)
	}
	
	go func(){
		for {
			select{
				case val,_ := <-myChan:{
				fmt.Println(val)	
			}
		}
	}
	}()
	defer wg.Wait()
}


func add(num int, myChan chan int, wg *sync.WaitGroup){
	num += num
	myChan <- num
	wg.Done()
}