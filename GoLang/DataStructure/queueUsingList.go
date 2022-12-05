package main

import (
	"container/list"
	"fmt"
)

func main() {

	queue := list.New()
	Enqueue(queue, 1)
	Enqueue(queue, 3)
	PrintQueue(queue)
	if err := Dequeue(queue); err != nil{
		fmt.Println(err)
	}
	Enqueue(queue, 2)
	PrintQueue(queue)

}

func Enqueue(queue *list.List, val int) {
	queue.PushBack(val)
}

func Dequeue(queue *list.List) error{
	if queue.Len() == 0{
		return fmt.Errorf(" size of queue is %v so can't perform dequeue", queue.Len())
	}
	element := queue.Front()
	queue.Remove(element)
	return nil
}

func PrintQueue(queue *list.List){
	for queue.Len() > 0{
		ele := queue.Front()
		fmt.Println(ele.Value)
		queue.Remove(ele)
	}
}