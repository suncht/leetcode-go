package array

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	obj := Constructor(nums)
	param1 := obj.Reset()
	param2 := obj.Shuffle()

	fmt.Println(param1)
	fmt.Println(param2)
}
