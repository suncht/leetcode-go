package model

import (
	"fmt"
	"testing"
)

func TestLinkModel(t *testing.T) {
	var head = GenerateLinkedList(100, 2)
	fmt.Println(StringifyLinkedList(head))
}
