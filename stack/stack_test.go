package stack

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := new(Stack)
	stack.Init()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	stack.Push(3)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
