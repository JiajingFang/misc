package data_structure

import (
	"fmt"
)

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
	Root *Node
}

type Node struct {
	Key    int
	Value  string
	Color  bool // true for red, false for black
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewRedBlackTree(root *Node) *RedBlackTree {
	return &RedBlackTree{
		Root: &Node{
			Key: root.Key,
			Value: root.Value,
			Color: root.Color,
			Left: root.Left,
			Right: root.Right,
			Parent: root.Parent,
		},
	}
}

func (rbt *RedBlackTree) Search(key int) *Node {
	// same as a normal BST
	root := rbt.Root 
	for root != nil {
		if key == root.Key {
			return root
		} else if key < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

func (rbt *RedBlackTree) Insert(key int, Value string) {
	// same as a normal BST, mark as red color=true 
	// use insert_fixup(node)

	nodeToInsert := Node{
		Key: key,
		Value: Value,
		Color: true,
	}

	if rbt.Root == nil {
		nodeToInsert.Color = false
		rbt.Root = &nodeToInsert
		return 
	}
	current := rbt.Root
	for current != nil {
		if current.Key > key && current.Left != nil {
			current = current.Left
		} else if current.Key < key && current.Right != nil {
			current = current.Right
		} else {
			break
		}
	}

	nodeToInsert.Parent = current
	if current.Key > key {
		current.Left = &nodeToInsert
	} else {
		current.Right = &nodeToInsert
	}
	rbt.insert_fixup(&nodeToInsert)
}

func isRed(n *Node) bool {
	return n != nil && n.Color
}

func (rbt *RedBlackTree) leftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		rbt.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (rbt *RedBlackTree) rightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		rbt.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

func (rbt *RedBlackTree) insert_fixup(z *Node) {
	for z.Parent != nil && isRed(z.Parent) {
		if z.Parent == z.Parent.Parent.Left {
			uncle := z.Parent.Parent.Right
			if isRed(uncle) {
				z.Parent.Color = false
				uncle.Color = false
				z.Parent.Parent.Color = true
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					rbt.leftRotate(z)
				}
				z.Parent.Color = false
				z.Parent.Parent.Color = true
				rbt.rightRotate(z.Parent.Parent)
			}
		} else {
			uncle := z.Parent.Parent.Left
			if isRed(uncle) {
				z.Parent.Color = false
				uncle.Color = false
				z.Parent.Parent.Color = true
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rbt.rightRotate(z)
				}
				z.Parent.Color = false
				z.Parent.Parent.Color = true
				rbt.leftRotate(z.Parent.Parent)
			}
		}
	}
	rbt.Root.Color = false
}


func PrintLevelByLevel (root *Node) {
	var result [][]string 
	if root == nil {
		fmt.Print("empty tree")
		return 
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []string 
		

		for i := 0; i < levelSize; i++ {
			// Dequeue
			// Shadowing! Creates a NEW 'queue' inside this block.
			// The outer 'queue' remains untouched!
			// queue := queue[1:]
			node := queue[0]
			queue = queue[1:]

			// Add node Value to current level list
			currentLevel = append(currentLevel, node.Value)

			// Equeue Left and Right childer 
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// store the completed level into result
		result = append(result, currentLevel)
	}

	fmt.Println("-----print RBT level by level------")
	for _, level := range result{
		fmt.Println("-----------")
		for _, node := range level {
			fmt.Print(node)
		}
		fmt.Println()
	}
}