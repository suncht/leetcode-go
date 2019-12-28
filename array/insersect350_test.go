package array

import (
	"fmt"
	"testing"
)

func TestInsersect350(t *testing.T) {
	var array1 = []int{1, 2, 3, 4, 56}
	var array2 = []int{4, 5, 7, 8, 10, 11}

	var res = intersect(array1, array2)
	fmt.Println(res)
}

func TestInsersect350_2(t *testing.T) {
	var array1 = []int{1, 2, 3, 4, 56}
	var array2 = []int{4, 5, 7, 8, 10, 11}

	var res = intersect2(array1, array2)
	fmt.Println(res)
}
