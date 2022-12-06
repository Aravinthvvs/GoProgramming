package main

import "fmt"

func main() {
	ll := &linkedList{}
	ll.addFront(1)
	ll.addFront(2)
	ll.addBack(3)
	fmt.Println("size of linked list :: ",ll.size())
	ll.traverse()
	fmt.Println(ll.front())
	ll.removeFront()
	ll.removeBack()
	fmt.Println(ll.front())
	ll.traverse()
}

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	len  int
}

func (ll *linkedList) addFront(val int) {
	ele := &node{val, nil}

	if ll.head == nil {
		ll.head = ele
	} else {
		ele.next = ll.head
		ll.head = ele
	}
	ll.len++
	return
}

func (ll *linkedList) addBack(val int) {
	ele := &node{val, nil}

	if ll.head == nil {
		ll.head = ele
	} else {
		current := ll.head
		for current.next != nil{
			current = current.next
		}
		current.next = ele
	}

	ll.len++
	return
}

func (ll *linkedList) removeFront()error{

	if ll.head == nil{
		return fmt.Errorf("empty Linkedlist")
	}

	ll.head = ll.head.next
	ll.len --
	return nil
}

func (ll *linkedList) removeBack()error{

	if ll.head == nil{
		return fmt.Errorf("empty Linkedlist")
	}

	current := ll.head
	var pre *node

	for current.next != nil{
		pre = current
		current = current.next
	}
	
	if pre.next != nil{
		pre.next = nil
	}else{
		ll.head = nil
	}

	ll.len --
	return nil
}

func(ll *linkedList)front()(int, error){
	if ll.head == nil{
		return -1,fmt.Errorf("empty Linkedlist")
	}
	return ll.head.value, nil
}

func(ll *linkedList)size()int{
	return ll.len
}

func (ll *linkedList) traverse() {
	current := ll.head

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}