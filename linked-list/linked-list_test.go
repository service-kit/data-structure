package linked_list

import (
	"fmt"
	"testing"
)

func Test_SingleLinkedList(t *testing.T) {
	ll := new(SingleLinkedList)
	ll.Init()
	for i := 1; i <= 100; i++ {
		ll.Insert(i, i)
	}
	for i := 1; i <= ll.Len(); i++ {
		fmt.Printf("%v ", ll.GetValue(i))
	}
	fmt.Println()
	fmt.Println("len ", ll.len)
	for i := ll.Len(); i >= 1; i-- {
		ll.Delete(i)
	}
	fmt.Println("len ", ll.len)
	for i := 1; i <= ll.Len(); i++ {
		fmt.Printf("%v ", ll.GetValue(i))
	}
	fmt.Println()
}

func Test_DoubleLinkedList(t *testing.T) {
	d := new(DoubleLinkedList)
	d.Init()
	for i := 0; i < 100; i++ {
		d.PushBack(i)
	}
	for i := 0; i < 100; i++ {
		d.PushFront(i)
	}
	len := d.Len()
	for i := 0; i < len; i++ {
		fmt.Printf("%v ", d.GetNode(i+1).Value())
	}
	fmt.Println()
	d.Delete(99)
	len = d.Len()
	for i := 0; i < len; i++ {
		fmt.Printf("%v ", d.GetNode(i+1).Value())
	}
	fmt.Println()
}
