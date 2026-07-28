package data_structure


// https://pages.cs.wisc.edu/~jinc/
// https://zhuanlan.zhihu.com/p/79980618
// 5 rules in RBT
// 1. Every node has a color either red or black

// 2. The root of the tree is always black

// 3. There are no two adjacent red nodes (Ared node cannot have a red parent or red child)

// 4. Every path from a node (including root) to any of its descendants NULL nodes has the same number of black nodes.

// 5. All leaf nodes are black nodes

// Algorithm	Time Complexity
// Search	          O(log n)
// Insert	          O(log n)
// Delete	          O(log n)

type RedBlackTree struct {
	root *Node
}

type Node struct {
	key    int
	value  string
	color  bool // true for red, false for black
	left   *Node
	right  *Node
	parent *Node
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (rbt *RedBlackTree) Search(key int) *Node {
	// same as a normal BST
	root := rbt.root 
	for root != nil {
		if key == root.key {
			return root
		} else if key < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil
}

func (rbt *RedBlackTree) Insert(key int, value string) *Node{
	// same as a normal BST, mark as red color=true 
	// use insert_fixup(node)

	nodeToInsert := Node{
		key: key,
		value: value,
		color: true,
	}

	if rbt.root == nil {
		rbt.root = &nodeToInsert
	}
	current := rbt.root
	for current != nil {
		if current.key > key && current.left != nil {
			current = current.left
		} else if current.key < key && current.right != nil {
			current = current.right
		} else {
			break
		}
	}

	if current.key > key {
		current.left = &nodeToInsert
	} else {
		current.right = &nodeToInsert
	}
	// insert_fixup
	return rbt.root
}
