package stack

import "errors"

type Stack struct {
	Top   int
	Size  int
	Items []int
}

func NewStack(size int) *Stack {
	return &Stack{Top: -1, Items: make([]int, size)}
}

func (stack *Stack) IsEmpty() bool {
	if len(stack.Items) < 0 {
		return true
	}
	return false
}

func (stack *Stack) Push(x int) error {
	if stack.Size == len(stack.Items) {
		return errors.New("StackOverflow Error : ")
	}
	stack.Top++
	stack.Items[stack.Top] = x
	return nil
}

func (stack *Stack) Pop() error {
	if stack.Top == -1 {
		return errors.New("Stack Empty")
	}
	stack.Top--
	return nil
}
func (stack *Stack) Peek() (int, error) {
	if stack.Top == -1 {
		return -1, errors.New("Stack Empty")
	}
	return stack.Items[stack.Top], nil
}
