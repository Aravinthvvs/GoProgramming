/*How do you synchronize the go-routines so that they print the numbers nicely in order as 1,2,3,4,5 … ?
The program consists of 3 go-routines each of which prints an arithmetic progression (series of numbers) with a common difference of 3.
The first one prints the sequence 1,4,7, …
The second one prints the sequence 2, 5, 8, …
The third one prints the sequence 3, 6, 9, …*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go firstOne(ch1, ch2, &wg)
	go secondOne(ch2, ch3, &wg)
	go thirdOne(ch3, ch1, &wg)

	ch1 <- 1
	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
}

func firstOne(ch1, ch2 chan int, wg *sync.WaitGroup) {
	i := 1
	for val := range ch1 {
		fmt.Println(i)
		//fmt.Println("out of first loop", val)
		i = i + 3
		ch2 <- val
		time.Sleep(time.Duration(1) * time.Second)
	}
	wg.Done()
}

func secondOne(ch2, ch3 chan int, wg *sync.WaitGroup) {
	i := 2
	for val := range ch2 {
		fmt.Println(i)
		//fmt.Println("out of Second loop", val)
		i = i + 3
		ch3 <- val
		time.Sleep(time.Duration(1) * time.Second)
	}
	wg.Done()
}

func thirdOne(ch3, ch1 chan int, wg *sync.WaitGroup) {
	i := 3
	for val := range ch3 {
		fmt.Println(i)
		//fmt.Println("out of Third loop", val)
		i = i + 3
		ch1 <- val
		time.Sleep(time.Duration(1) * time.Second)
	}
	wg.Done()
}
