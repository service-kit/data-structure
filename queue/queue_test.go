package queue

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	q := new(Queue)
	q.Init()
	value := 1
	for i := 0; i < 80; i++ {
		q.Enqueue(value)
		value++
	}
	fmt.Println("enqueue 80 res:")
	fmt.Printf("cap:%v size:%v\n", q.Capacity(), q.Len())
	for _, val := range q.array {
		fmt.Printf("%v ", val)
	}
	fmt.Printf("\n")
	fmt.Println("dequeue 50 res:")
	for i := 0; i < 50; i++ {
		q.Dequeue()
	}
	fmt.Printf("cap:%v size:%v\n", q.Capacity(), q.Len())
	for _, val := range q.array {
		fmt.Printf("%v ", val)
	}
	fmt.Print("\n")
	for i := 0; i < 90; i++ {
		q.Enqueue(value)
		value++
	}
	fmt.Println("enqueue 90 res:")
	fmt.Printf("cap:%v size:%v\n", q.Capacity(), q.Len())
	for _, val := range q.array {
		fmt.Printf("%v ", val)
	}
	fmt.Print("\n")
	for i := 0; i < 70; i++ {
		q.Enqueue(value)
		value++
	}
	fmt.Println("enqueue 70 res:")
	fmt.Printf("cap:%v size:%v\n", q.Capacity(), q.Len())
	for _, val := range q.array {
		fmt.Printf("%v ", val)
	}
	fmt.Printf("\n")
}
