package parallel

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMapReduce(t *testing.T) {
	// 模拟 Java 函数式流编程
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	streamChan := make(chan int)
	outChan := make(chan string)

	go StreamOf(nums, streamChan)
	go MapToStr(streamChan, outChan, func(num int) string {
		return strconv.Itoa(num * num)
	})

	for i := range outChan {
		fmt.Printf("out is %s\n", i)
	}
}
