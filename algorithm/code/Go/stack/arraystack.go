package stack

import (
	"fmt"
)

type Stack[T any] struct {
	data []T
	top  int
}

type StackInterface[T any] interface {
	Push(T)
	Pop() T
	Top() T
	IsEmpty() bool
	Size() int
	Print()
}

var _ StackInterface[int] = (*Stack[int])(nil)

func NewArrayStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0, 32), top: -1}
}

func (s *Stack[T]) Push(data T) {
	s.top++
	s.data = append(s.data, data)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	value := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return value
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	return s.data[s.top]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack[T]) Size() int {
	return s.top + 1 // 或者 len(s.data)
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0, 32)
	s.top = -1
}

func (s *Stack[T]) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	for i := s.top; i >= 0; i-- {
		fmt.Println(s.data[i], " ")
	}
}
