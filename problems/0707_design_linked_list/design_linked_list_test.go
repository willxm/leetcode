package problem

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	list := Constructor()
	list.Travel()
	list.AddAtIndex(1, 2)
	list.Travel()
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	list.AddAtIndex(0, 1)
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
}
