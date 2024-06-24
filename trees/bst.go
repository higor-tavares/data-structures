package trees

func (tree *BinarySearchTree) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{value: value}
	} else {
		tree.insertNode(value, tree.root)
	}
}
func (tree *BinarySearchTree) insertNode(value int, node *Node) *Node {
	if node == nil {
		return &Node{value: value}
	}
	if value > node.value {
		if node.right == nil {
			node.right = tree.insertNode(value, node.right)
		} else {
			tree.insertNode(value, node.right)
		}
	} else if value < node.value {
		if node.left == nil {
			node.left = tree.insertNode(value, node.left)
		} else {
			tree.insertNode(value, node.left)
		}
	}
	return node
}
