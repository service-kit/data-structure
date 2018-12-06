package array

import (
	"fmt"
	"testing"
)

func Test_Array(t *testing.T) {
	arr := [10]int{}
	for i, val := range arr {
		fmt.Printf("index: %v value:%v\n", i, val)
	}
	for i, _ := range arr {
		arr[i] = i
	}
	for i, val := range arr {
		fmt.Printf("index: %v value:%v\n", i, val)
	}
}
