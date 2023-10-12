package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

type BinarySearchTree struct {
	rootNode *Node
}

func (bst *BinarySearchTree) Add(val int) {

	add(bst.rootNode, val)
}

func add(rootNode *Node, val int) *Node {

	if rootNode == nil {
		rootNode = &Node{Value: val, LeftNode: nil, RightNode: nil}
		return rootNode
	}

	if rootNode.Value > val {
		rootNode.LeftNode = add(rootNode.LeftNode, val)
	} else if rootNode.Value < val {
		rootNode.RightNode = add(rootNode.RightNode, val)
	}
	return rootNode
}

func (bst *BinarySearchTree) Remove(val int) {
	bst.rootNode = remove(bst.rootNode, val)
}

func remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if node.Value > val {
		node.LeftNode = remove(node.LeftNode, val)
	} else if node.Value < val {
		node.RightNode = remove(node.RightNode, val)
	} else {
		if node.LeftNode == nil {
			return node.RightNode
		} else if node.RightNode == nil {
			return node.LeftNode
		} else {

			tempNode := findMin(node.RightNode)

			node.Value = tempNode.Value
			node.RightNode = remove(node.RightNode, tempNode.Value)
		}
	}
	return node
}

func findMin(node *Node) *Node {
	for node.LeftNode != nil {
		node = node.LeftNode
	}
	return node
}

func (bst BinarySearchTree) InOrderTraversal() string {
	sb := strings.Builder{}
	nodetostring(bst.rootNode, &sb)
	return sb.String()
}

func nodetostring(node *Node, str *strings.Builder) {
	if node == nil {
		return
	}
	nodetostring(node.LeftNode, str)
	str.WriteString(fmt.Sprintf("%s ", strconv.Itoa(node.Value)))
	nodetostring(node.RightNode, str)
}

func main() {
	root := &Node{Value: 2, LeftNode: nil, RightNode: nil}
	root.LeftNode = &Node{Value: 1, LeftNode: nil, RightNode: nil}
	root.RightNode = &Node{Value: 3, LeftNode: nil, RightNode: nil}
	bst := BinarySearchTree{rootNode: root}
	bst.Add(14)
	bst.Add(5)
	bst.Add(10)
	bst.Add(12)
	fmt.Println(bst.InOrderTraversal())
	bst.Remove(10)
	fmt.Println(bst.InOrderTraversal())
}
