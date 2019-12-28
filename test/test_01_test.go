package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	map1 := make(map[string]string)
	fmt.Println(map1)
	for i := 0; i < 10; i++ {
		map1[strconv.Itoa(i)] = strconv.Itoa(i)
	}

	map1["0"] += "asdf"

	fmt.Println(map1)
}
