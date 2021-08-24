package defer_func

import "fmt"

func StructRange() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}
}
