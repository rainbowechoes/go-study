package parallel

import (
	"fmt"
	"math/rand"
	"sync"
)

type TaskResult struct {
	Url string
	Err error
}

// url and url are individually, can be think as paralleling completely.
func DoParallelTask(urls []string) []TaskResult {
	// 通道数量在一定程度上可以提高并行处理的速度
	taskResults := make(chan TaskResult)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		// worker
		go func(url string) {
			defer wg.Done()
			// do something
			for a := 0; a < rand.Intn(30000); a++ {
				fmt.Printf("a 的值为: %d\n", a)
			}
			taskResults <- TaskResult{url, nil}

		}(url)
	}

	// closer
	go func() {
		wg.Wait()
		close(taskResults)
	}()

	res := make([]TaskResult, len(urls))
	for item := range taskResults {
		res = append(res, item)
	}
	return res
}
