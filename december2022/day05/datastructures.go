package day05

import "log"

type move struct {
	fromIdx int
	toIdx   int
	count   int
}

type item[V interface{}] struct {
	value V
	next  *item[V]
}

type Stack[V interface{}] struct {
	top  *item[V]
	size int
}

func (stack *Stack[V]) Len() int {
	return stack.size
}

func (stack *Stack[V]) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack[V]) Push(value V) {
	stack.top = &item[V]{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack[V]) Pop() V {
	if stack.Len() <= 0 {
		return *new(V)
	}
	value := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return value
}

func (stack *Stack[V]) Peek() *V {
	if stack.Len() <= 0 {
		return nil
	}
	return &stack.top.value
}

type strategy = string

func (fromStack *Stack[V]) MoveTo(toStack *Stack[V], count int, strategy strategy) {
	switch strategy {
	case "lifo", "LIFO":
		for i := 0; i < count; i++ {
			toStack.Push(fromStack.Pop())
		}
	case "fifo", "FIFO":
		midStack := Stack[V]{top: nil, size: 0}
		for i := 0; i < count; i++ {
			midStack.Push(fromStack.Pop())
		}
		for !midStack.IsEmpty() {
			toStack.Push(midStack.Pop())
		}
	default:
		log.Fatalln("Unknown strategy:", strategy)
	}
}
