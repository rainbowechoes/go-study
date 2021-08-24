package str

import (
	"fmt"
	"testing"
)

func TestConcat(t *testing.T) {
	s := Concat("12", "23")
	fmt.Println(s)
}

func TestReverseStr(t *testing.T) {
	s := "Hello, Go!"
	for _, v := range s {
		fmt.Println("ch is " + string(v))
	}
}