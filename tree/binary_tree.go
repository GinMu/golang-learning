package main

import "fmt"

// Node define tree
type Node struct {
	Left  *Node
	Data  interface{}
	Right *Node
}

// Initer interface
type Initer interface {
	SetData(data interface{})
}

// Operater interface
type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

// Order interface
type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

// SetData to set node data
func (n *Node) SetData(data interface{}) {
	n.Data = data
}

// PrintBT print node
func (n *Node) PrintBT() {
	printBT(n)
}

// Depth to calculate node's depth
func (n *Node) Depth() int {
	return depth(n)
}

// LeafCount to calculate leaf count
func (n *Node) LeafCount() int {
	return leafCount(n)
}

// PreOrder print node by preorder
func (n *Node) PreOrder() {
	preOrder(n)
}

// InOrder print node by inorder
func (n *Node) InOrder() {
	inOrder(n)
}

// PostOrder print node by postorder
func (n *Node) PostOrder() {
	postOrder(n)
}

// NewNode create new node
func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func printBT(n *Node) {
	if n != nil {
		fmt.Printf("%v", n.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			printBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			printBT(n.Right)
			fmt.Printf(")")
		}
	}
}

func depth(n *Node) int {
	var depleft, depright int
	if n == nil {
		return 0
	}
	depleft = depth(n.Left)
	depright = depth(n.Right)
	if depleft > depright {
		return depleft + 1
	}
	return depright + 1
}

func leafCount(n *Node) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	}
	return leafCount(n.Left) + leafCount(n.Right)
}

func preOrder(n *Node) {
	if n != nil {
		fmt.Printf("%v\n", n.Data)
		preOrder(n.Left)
		preOrder(n.Right)
	}
}

func inOrder(n *Node) {
	if n != nil {
		preOrder(n.Left)
		fmt.Printf("%v\n", n.Data)
		preOrder(n.Right)
	}
}

func postOrder(n *Node) {
	preOrder(n.Left)
	preOrder(n.Right)
	fmt.Printf("%v\n", n.Data)
}

func main() {
	root := NewNode(nil, nil)
	var it Initer
	it = root
	it.SetData("root node")
	a := NewNode(nil, nil)
	a.SetData("left node")
	al := NewNode(nil, nil)
	al.SetData(100)
	ar := NewNode(nil, nil)
	ar.SetData(200)
	a.Left = al
	a.Right = ar
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	root.PrintBT()
	fmt.Println()

	var op Operater = root
	fmt.Println("The depths of the Btree is:", op.Depth())
	fmt.Println("The leaf counts of the Btree is:", op.LeafCount())
	

	var od Order = root
	
	fmt.Println("前序遍历，采用递归算法，按照先访问根节点，再访问左子树，最后访问右子树的次序访问二叉树中的所有节点，且每个节点仅访问一次")
	od.PreOrder()

	fmt.Println("中序遍历，采用递归算法，按照先访问左子树，再访问根节点，最后访问右子树的次序访问二叉树中的所有节点，且每个节点仅访问一次")
	od.InOrder()

	fmt.Println("后序遍历，采用递归算法，按照先访问左子树，再访问右子树，最后访问根节点的次序访问二叉树中的所有节点，且每个节点仅访问一次")
	od.PostOrder()

	fmt.Println()
}
