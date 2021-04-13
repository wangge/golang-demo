package main

import (
	"fmt"
)

// 二叉树
type Node struct {
	value int
	left *Node
	right *Node
}
func main()  {
	root := Node{value: 3}
	root.left = &Node{}
	root.left.SetValue(0)
	root.left.right = CreateNode(2)
	root.right = &Node{5,nil,nil}
	root.right.left = CreateNode(4)
	fmt.Print("\n前序遍历: ")
	root.PreOrder()
	fmt.Print("\n中序遍历: ")
	root.MiddleOrder()
	fmt.Print("\n后序遍历: ")
	root.PostOrder()
	fmt.Print("\n层次遍历: ")
	root.BreadthFirstSearch()
	fmt.Println("\n层数: ", root.Layers())
	fmt.Println("\n层数: ", root.LayersByQueue())
}

// 打印当前节点
func (node *Node) Print()  {
	fmt.Print(node.value," ")
}

func (node *Node) SetValue(v int)  {
	if node == nil{
		fmt.Println("Setting value to nil node ignored.")
	}
	node.value = v
}
// 前序遍历 左-中-右
func (node *Node) PreOrder()  {
	if(node == nil){
		return
	}
	
	node.Print()
	node.left.PreOrder()
	node.right.PreOrder()
}

// 中序遍历
func (node *Node)MiddleOrder()  {
	if node == nil{
		return
	}
	node.left.MiddleOrder()
	node.Print()
	node.right.MiddleOrder()
}

// 后续遍历
func (node *Node)PostOrder()  {
	if node == nil{
		return
	}
	node.right.PostOrder()
	node.Print()
	node.left.PostOrder()
}
// 层次遍历-广度优先遍历
func (node *Node) BreadthFirstSearch()  {
	if node == nil{
		return
	}

	result := []int{}
	nodes := []*Node{node}
	for len(nodes) >0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result,curNode.value)
		if curNode.left != nil{
			nodes = append(nodes,curNode.left)
		}
		if curNode.right != nil{
			nodes = append(nodes, curNode.right)
		}
	}

	for _,v := range result{
		fmt.Print(v," ")
	}
}

// 层数（递归实现）
func (node *Node)Layers() int {
	if node == nil{
		return 0
	}
	leftLayers := node.left.Layers()
	rightLayers := node.right.Layers()
	if leftLayers > rightLayers{
		return leftLayers + 1
	}
	return rightLayers + 1
}

// 层数（非递归）
func (node *Node)LayersByQueue() int {
	if node == nil {
		return 0
	}
	layers := 0
	nodes :=[]*Node{node}
	for len(nodes) >0{
		layers++
		size := len(nodes)
		count:= 0
		for count < size{
			count++
			curNode := nodes[0]
			nodes = nodes[1:]
			if curNode.left != nil{
				nodes = append(nodes, curNode.left)
			}
			if curNode.right != nil{
				nodes = append(nodes, curNode.right)
			}
		}
	}
	return layers
}

func CreateNode(v int) *Node  {
	return &Node{value: v}
}