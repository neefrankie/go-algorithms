package tree

import "fmt"

type Node struct {
	Key        int
	Data       float64
	LeftChild  *Node
	RightChild *Node
}

func (n *Node) DisplayNode() {
	fmt.Printf("{%d, %f}", n.Key, n.Data)
}
