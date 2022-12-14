package datastructures

type item[V interface{}] struct {
	value V
	next  *item[V]
}

type LinkedStack[V any] struct {
	top  *item[V]
	size int
}

func (stack *LinkedStack[V]) Len() int {
	return stack.size
}

func (stack *LinkedStack[V]) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *LinkedStack[V]) Push(value V) {
	stack.top = &item[V]{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *LinkedStack[V]) Pop() V {
	if stack.Len() <= 0 {
		return *new(V)
	}
	value := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return value
}

func (stack *LinkedStack[V]) Peek() V {
	if stack.Len() <= 0 {
		return *new(V)
	}
	return stack.top.value
}
