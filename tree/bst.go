package tree

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Key   int
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new TreeNode
func NewTreeNode(key int, value interface{}) *TreeNode {
	return &TreeNode{Key: key, Value: value}
}

// BST
type BinarySearchTree struct {
	Root *TreeNode
}

// Insert inserts a new key-value pair into the BST
func (bst *BinarySearchTree) Insert(key int, value interface{}) {
	bst.Root = insertNode(bst.Root, key, value)
}

func insertNode(node *TreeNode, key int, value interface{}) *TreeNode {
	if node == nil {
		return NewTreeNode(key, value)
	}
	if key < node.Key {
		node.Left = insertNode(node.Left, key, value)
	} else {
		node.Right = insertNode(node.Right, key, value)
	}
	return node
}

// Search searches for a key in the BST
func (bst *BinarySearchTree) Search(key int) (interface{}, bool) {
	node := searchNode(bst.Root, key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

func searchNode(node *TreeNode, key int) *TreeNode {
	if node == nil || node.Key == key {
		return node
	}
	if key < node.Key {
		return searchNode(node.Left, key)
	}
	return searchNode(node.Right, key)
}

// Update updates the value of a key in the BST
func (bst *BinarySearchTree) Update(key int, value interface{}) bool {
	node := searchNode(bst.Root, key)
	if node != nil {
		node.Value = value
		return true
	}
	return false
}

// Delete deletes a key from the BST
func (bst *BinarySearchTree) Delete(key int) {
	bst.Root = deleteNode(bst.Root, key)
}

func deleteNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = deleteNode(node.Left, key)
	} else if key > node.Key {
		node.Right = deleteNode(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minLargerNode := findMin(node.Right)
		node.Key, node.Value = minLargerNode.Key, minLargerNode.Value
		node.Right = deleteNode(node.Right, minLargerNode.Key)
	}
	return node
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
