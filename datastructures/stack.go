package datastructures

type Stack[V any] interface {
	Len() int
	IsEmpty() bool
	Push(v V)
	Pop() V
	Peek() *V
}
