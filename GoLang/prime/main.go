package main

import ("fmt"
"math")

func main() {
	fmt.Println("Enter the range :")

	var end int64
	fmt.Scanln(&end)

	myChan := make(chan int64)
	endChan := make(chan int64)
	go GetPrimeNums(end, myChan, endChan)

	defer func(){
		close(myChan)
		close(endChan)
	}() 
	for{
		select {
		case val := <-myChan:
			fmt.Println(val)
		case <-endChan:
			return
		}
	}
	

}

func GetPrimeNums(end int64, myChan, endChan chan int64) {

	myChan<-2
	for n := int64(3); n <= end; n += 2 {
		prime := true
		fmt.Println("n", n)
		r := int64(math.Sqrt(float64(n)))
		fmt.Println("r", r)
		for i := int64(3); i < r+1; i += 2 {
			fmt.Println("i", i)
			if n%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			//fmt.Println("inside prime ", n)
			myChan <- n
		}
	}
	endChan <- 0
}
