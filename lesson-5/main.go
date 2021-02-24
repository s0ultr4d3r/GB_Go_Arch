package main

import (
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

}
