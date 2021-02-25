package main

import (
	"GB/lesson-5/linkedlist"
	"GB/lesson-5/queue"
	"GB/lesson-5/stack"
	"fmt"
)

func main() {
	stack := stack.NewStack(2)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	queue := queue.NewQueue(2)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	list := &linkedlist.List{}
	node := &linkedlist.Node{
		Data: 5,
	}
	list.Append(node)
	node1 := &linkedlist.Node{
		Data: 5,
	}
	list.Add(nil, node1)
	fmt.Println(list.Len())
	node = list.Find(5)
	list.Delete(node) //panic: runtime error: invalid memory address or nil pointer dereference при единственной ноде
	fmt.Println(list.Len())
}
