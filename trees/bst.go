package trees

func (bst *BinarySearchTree) Insert(value int) {
	if bst.root == nil {
		bst.root = &Node{value: value}
	} else {
		bst.insertNode(value, bst.root)
	}
}
func (bst *BinarySearchTree) insertNode(value int, node *Node) *Node {
	if node == nil {
		return &Node{value: value}
	}
	if value > node.value {
		if node.right == nil {
			node.right = bst.insertNode(value, node.right)
		} else {
			bst.insertNode(value, node.right)
		}
	} else if value < node.value {
		if node.left == nil {
			node.left = bst.insertNode(value, node.left)
		} else {
			bst.insertNode(value, node.left)
		}
	}
	return node
}

func (bst *BinarySearchTree) InOrderTransverse(callback func(node *Node) interface{}) {
	bst.inOrderTransverseNode(bst.root, callback)
}

func (bst *BinarySearchTree) PreOrderTransverse(callback func(node *Node) interface{}) {
	bst.preOrderTransverseNode(bst.root, callback)
}

func (bst *BinarySearchTree) PostOrderTranverse(callback func(node *Node) interface{}) {
	bst.postOrderTransverseNode(bst.root, callback)
}

func (bst *BinarySearchTree) Remove(value int) {
	bst.root = bst.removeNode(value, bst.root)
}

func (bst *BinarySearchTree) removeNode(value int, node *Node) *Node {
	if node == nil {
		return nil
	}
	if value > node.value {
		node.right = bst.removeNode(value, node.right)
		return node
	} else if value < node.value {
		node.left = bst.removeNode(value, node.left)
		return node
	} else {
		if node.right == nil && node.left == nil {
			node = nil
			return node
		} else if node.left == nil {
			node = node.right
			return node
		} else if node.right == nil {
			node = node.left
			return node
		} else {
			min := bst.minNode(node.right)
			node.value = min.value
			node.right = bst.removeNode(value, node.right)
			return node
		}
	}
}

func (bst *BinarySearchTree) Min() *Node {
	return bst.minNode(bst.root)
}

func (bst *BinarySearchTree) Max() *Node {
	return bst.maxNode(bst.root)
}

func (bst *BinarySearchTree) minNode(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return bst.minNode(node.left)
}

func (bst *BinarySearchTree) maxNode(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return bst.maxNode(node.right)
}

func (bst *BinarySearchTree) inOrderTransverseNode(node *Node, callback func(node *Node) interface{}) {
	if node != nil {
		bst.inOrderTransverseNode(node.left, callback)
		callback(node)
		bst.inOrderTransverseNode(node.right, callback)
	}
}

func (bst *BinarySearchTree) preOrderTransverseNode(node *Node, callback func(node *Node) interface{}) {
	if node != nil {
		callback(node)
		bst.preOrderTransverseNode(node.left, callback)
		bst.preOrderTransverseNode(node.right, callback)
	}
}

func (bst *BinarySearchTree) postOrderTransverseNode(node *Node, callback func(node *Node) interface{}) {
	if node != nil {
		bst.postOrderTransverseNode(node.left, callback)
		bst.postOrderTransverseNode(node.right, callback)
		callback(node)
	}
}
