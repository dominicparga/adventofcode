package day05

type move struct {
	fromIdx int
	toIdx   int
	count   int
}

type item[V interface{}] struct {
	value V
	next  *item[V]
}

type stack[V interface{}] struct {
	top  *item[V]
	size int
}

func (stack *stack[V]) Len() int {
	return stack.size
}

func (stack *stack[V]) Push(value V) {
	stack.top = &item[V]{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *stack[V]) Pop() V {
	if stack.Len() <= 0 {
		return *new(V)
	}
	value := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return value
}

func (stack *stack[V]) Peek() *V {
	if stack.Len() <= 0 {
		return nil
	}
	return &stack.top.value
}

func (fromStack *stack[V]) MoveTo(toStack *stack[V], count int) {
	for i := 0; i < count; i++ {
		toStack.Push(fromStack.Pop())
	}
}
