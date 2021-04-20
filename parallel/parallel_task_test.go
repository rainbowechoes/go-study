package parallel

import (
	"fmt"
	"testing"
)

func TestParallelTasks(t *testing.T) {
	urls := []string{"http://localhost/1", "http://localhost/2", "http://localhost/3", "http://localhost/4", "http://localhost/5", "http://localhost/6", "http://localhost/7"}

	res := DoParallelTask(urls)
	fmt.Println(res)
}
