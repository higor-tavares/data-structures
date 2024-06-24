package trees

import (
	"fmt"
	"testing"
)

func TestInsertNodes(t *testing.T) {
	values := []int{8, 3, 16, 10, 2, 5, 9, 1, 33}
	bst := &BinarySearchTree{}
	for i := 0; i < len(values); i++ {
		bst.Insert(values[i])
	}

	rt := bst.root
	if rt == nil || rt.value != 8 {
		t.Fatalf("The root element is incorrect!")
	}
	if rt.left == nil || rt.left.value != 3 {
		t.Fatalf("The element in left of root is incorrect!")
	}
	if rt.right == nil || rt.right.value != 16 {
		t.Fatalf("The element in right of root is incorrect!")
	}
}

func TestInOrderTransverse(t *testing.T) {
	values := []int{8, 3, 16, 10, 2, 5, 9, 1, 33}
	bst := &BinarySearchTree{}
	for i := 0; i < len(values); i++ {
		bst.Insert(values[i])
	}

	bst.InOrderTransverse(func(node *Node) interface{} {
		fmt.Println(node.value)
		return nil
	})
}
