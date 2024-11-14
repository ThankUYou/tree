package tree

// Tree is an interface for tree operations
type Tree interface {
	Insert(key int, value interface{})
	Delete(key int)
	Search(key int) (interface{}, bool)
	Update(key int, value interface{}) bool
}

// TreeType is an enumeration of tree types
type TreeType int

const (
	BinarySearchTreeType TreeType = iota
	// Add other tree types here
)

// NewTree creates a new tree of the specified type
func NewTree(treeType TreeType) Tree {
	switch treeType {
	case BinarySearchTreeType:
		return &BinarySearchTree{}
	// Add cases for other tree types
	default:
		return nil
	}
}
