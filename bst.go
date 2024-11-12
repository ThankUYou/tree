package tree

// 二叉树
type Node struct {
	Value       int   // 节点值
	Left, Right *Node // 左右子树
}

func Insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{
			Value: value,
			Left:  nil,
			Right: nil,
		}
	}
	if value < node.Value {
		node.Left = Insert(node.Left, value)
	} else if value > node.Value {
		node.Right = Insert(node.Right, value)
	}
	return node
}

func Delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = Delete(node.Left, value)
	} else if value > node.Value {
		node.Right = Delete(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minVal := findMinValue(node.Right)
		node.Value = minVal
		node.Right = Delete(node.Right, minVal)
	}
	return node
}

func findMinValue(node *Node) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Value
}

func Search(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	} else if value < node.Value {
		return Search(node.Left, value)
	} else {
		return Search(node.Right, value)
	}
}
