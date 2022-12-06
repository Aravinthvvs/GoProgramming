package main

import "fmt"

func main(){

	var queue []int
	queue = []int{1,3}
	enqueue(&queue,5)
	fmt.Println(queue)

	fmt.Println(peek(queue))
	err := dequeue(&queue)
	enqueue(&queue,2)
	fmt.Println(peek(queue))
	err = dequeue(&queue)
	err = dequeue(&queue)
	err = dequeue(&queue)
	fmt.Println(queue)
	err = dequeue(&queue)
	if err != nil{
		fmt.Println(err)
	}
}

func enqueue(q *[]int, val int){
	*q = append(*q,val)
}

func peek(q []int) (int, error){

	if len(q) == 0{
		return 0, fmt.Errorf(" size of queue is :: %v so can't perform peek",len(q))
	}
	return q[0], nil
}

func dequeue(q *[]int)error{
	if len(*q) == 0{
		return fmt.Errorf(" size of queue is :: %v so can't perform peek",len(*q))
	}
	*q = (*q)[1:]
	return nil
}